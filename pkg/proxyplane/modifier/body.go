package modifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/go-playground/validator/v10"
	"github.com/google/martian/parse"
	"github.com/open-privacy/opv/pkg/config"
)

func init() {
	parse.Register("opv.body.Modifier", NewOPVBodyModifierFromJSON)
}

type OPVBodyModifier struct {
	Scope                         []parse.ModifierType  `json:"scope" validate:"gt=0,dive,oneof=request response"`
	Items                         []OPVBodyModifierItem `json:"items"`
	OPVDataplaneGrantToken        string                `json:"-" validate:"required"`
	OPVDataplaneGrantTokenFromEnv string                `json:"opv_dataplane_grant_token_from_env"`
	OPVDataplaneBaseURL           string                `json:"opv_dataplane_base_url"`
}

type OPVBodyModifierItem struct {
	JSONPointerPath string `json:"json_pointer_path" validate:"required"`
	FactTypeSlug    string `json:"fact_type_slug" validate:"required"`
	Action          string `json:"action" validate:"required,oneof=tokenize detokenize"`
}

func (o *OPVBodyModifier) Render(contentType string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(contentType, "application/json") {
		return nil, fmt.Errorf("Content-Type %s not supported", contentType)
	}

	jsonParsed, err := gabs.ParseJSONBuffer(body)
	if err != nil {
		return nil, err
	}

	conn := newConn(o.OPVDataplaneGrantToken, o.OPVDataplaneBaseURL)

	for _, item := range o.Items {
		switch item.Action {
		case "tokenize":
			node, err := jsonParsed.JSONPointer(item.JSONPointerPath)
			if err != nil {
				return nil, err
			}
			value := node.String()
			factID, err := conn.createFact(item.FactTypeSlug, value)
			if err != nil {
				return nil, err
			}
			jsonParsed.SetJSONPointer(factID, item.JSONPointerPath)
		case "detokenize":
			node, err := jsonParsed.JSONPointer(item.JSONPointerPath)
			if err != nil {
				return nil, err
			}
			factID := node.Data().(string)
			value, err := conn.getFact(factID)
			if err != nil {
				return nil, err
			}
			jsonParsed.SetJSONPointer(value, item.JSONPointerPath)
		}
	}

	return jsonParsed.Bytes(), nil
}

func (o *OPVBodyModifier) ModifyRequest(req *http.Request) error {
	body, err := o.Render(req.Header.Get("Content-Type"), req.Body)
	if err != nil {
		return err
	}
	req.ContentLength = int64(len(body))
	req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	return nil
}

func (o *OPVBodyModifier) ModifyResponse(res *http.Response) error {
	body, err := o.Render(res.Header.Get("Content-Type"), res.Body)
	if err != nil {
		return err
	}
	res.ContentLength = int64(len(body))
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewReader(body))
	return nil
}

func NewOPVBodyModifierFromJSON(b []byte) (*parse.Result, error) {
	mod := &OPVBodyModifier{}
	if err := json.Unmarshal(b, mod); err != nil {
		return nil, err
	}

	if mod.OPVDataplaneBaseURL == "" {
		mod.OPVDataplaneBaseURL = config.ENV.ProxyPlaneDefaultDPBaseURL
	}
	if mod.OPVDataplaneGrantTokenFromEnv == "" {
		mod.OPVDataplaneGrantToken = config.ENV.ProxyPlaneDefaultDPGrantToken
	} else {
		mod.OPVDataplaneGrantToken = os.Getenv(mod.OPVDataplaneGrantTokenFromEnv)
	}

	validate := validator.New()
	if err := validate.Struct(mod); err != nil {
		return nil, err
	}

	return parse.NewResult(mod, mod.Scope)
}
