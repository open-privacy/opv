package modifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

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
	JSONPointerPath  string `json:"json_pointer_path" validate:"required"`
	ArrayPointerPath string `json:"array_pointer_path" validate:"omitempty,startswith=/,regex=^[a-z-_]+$"`
	FactTypeSlug     string `json:"fact_type_slug" validate:"required"`
	Action           string `json:"action" validate:"required,oneof=tokenize detokenize"`
}

func getValueFromArray(array string, data *gabs.Container) (interface{}, error) {
	keys := strings.Split(array, "/")  // Split the array string by "/"
	value := data.Path(keys[1]).Data() // Get the value corresponding to the key

	return value, nil
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

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := range o.Items {
		wg.Add(1)

		go func(item *OPVBodyModifierItem) {
			defer wg.Done()

			switch item.Action {
			case "tokenize":
				if item.ArrayPointerPath != "" {
					fmt.Println(item.ArrayPointerPath)
					fmt.Println("Array")
					nodes, err := jsonParsed.JSONPointer(item.ArrayPointerPath)
					if err != nil {
						return
					}
					children, _ := nodes.Children()
					for index, child := range children {
						fmt.Println(item.JSONPointerPath)
						fmt.Println(child)
						childNode, err := getValueFromArray(item.JSONPointerPath, child)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						if childNode != "{}" {
							value := ("\"" + string(childNode.(string)) + "\"")
							fmt.Println(value)
							fmt.Println(index)
							factID, err := conn.createFact(item.FactTypeSlug, value)
							if err != nil {
								return
							}
							mu.Lock()
							fmt.Println(factID)
							fmt.Println(item.ArrayPointerPath + "/" + strconv.Itoa(index) + item.JSONPointerPath)
							jsonParsed.SetJSONPointer(factID, item.ArrayPointerPath+"/"+strconv.Itoa(index)+item.JSONPointerPath)
							mu.Unlock()
						}
					}
				} else {
					fmt.Println("Not array")
					node, err := jsonParsed.JSONPointer(item.JSONPointerPath)
					value := node.String()

					factID, err := conn.createFact(item.FactTypeSlug, value)
					if err != nil {
						return
					}

					mu.Lock()
					jsonParsed.SetJSONPointer(factID, item.JSONPointerPath)
					mu.Unlock()
				}

			case "detokenize":
				if item.ArrayPointerPath != "" {
					fmt.Println(item.ArrayPointerPath)
					fmt.Println("Array")
					nodes, err := jsonParsed.JSONPointer(item.ArrayPointerPath)
					fmt.Println(nodes)
					if err != nil {
						fmt.Println(err)
						return
					}
					children, _ := nodes.Children()
					for index, child := range children {
						fmt.Println(child)
						childNode, err := getValueFromArray(item.JSONPointerPath, child)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						if childNode != "{}" {
							fmt.Println(index)
							factID := childNode.(string)
							value, err := conn.getFact(factID)
							fmt.Println(value)
							if err != nil {
								return
							}
							mu.Lock()
							fmt.Println(factID)
							fmt.Println(item.ArrayPointerPath + "/" + strconv.Itoa(index) + item.JSONPointerPath)
							jsonParsed.SetJSONPointer(value, item.ArrayPointerPath+"/"+strconv.Itoa(index)+item.JSONPointerPath)
							mu.Unlock()
						}
					}
				} else {
					fmt.Println("Not array")
					node, err := jsonParsed.JSONPointer(item.JSONPointerPath)
					if err != nil {
						return
					}

					factID := node.Data().(string)

					value, err := conn.getFact(factID)
					if err != nil {
						return
					}

					mu.Lock()
					jsonParsed.SetJSONPointer(value, item.JSONPointerPath)
					mu.Unlock()
				}
			}
		}(&o.Items[i])
	}

	wg.Wait()

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
