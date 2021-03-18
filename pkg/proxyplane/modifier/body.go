package modifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
	"github.com/google/martian/parse"
)

type OPVBodyModifier struct {
	Items []OPVBodyModifierItem `json:"items"`
	Scope []parse.ModifierType  `json:"scope"`

	conn *conn
}

type OPVBodyModifierItem struct {
	JSONPointerPath string `json:"json_pointer_path"`
	FactTypeSlug    string `json:"fact_type_slug"`
	Action          string `json:"action"`
}

func (o *OPVBodyModifier) Render(contentType string, body io.Reader) ([]byte, error) {
	if contentType != "application/json" {
		return nil, fmt.Errorf("Content-Type %s not supported", contentType)
	}

	jsonParsed, err := gabs.ParseJSONBuffer(body)
	if err != nil {
		return nil, err
	}

	for _, item := range o.Items {
		switch item.Action {
		case "tokenize":
			node, err := jsonParsed.JSONPointer(item.JSONPointerPath)
			if err != nil {
				return nil, err
			}
			value := node.String()
			factID, err := o.conn.createFact(item.FactTypeSlug, value)
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
			value, err := o.conn.getFact(factID)
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
	mod.conn = mustNewConn()
	return parse.NewResult(mod, mod.Scope)
}
