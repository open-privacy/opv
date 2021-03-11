// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api_audits": {
            "get": {
                "description": "Query API Audits",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audit"
                ],
                "summary": "Query API Audits",
                "operationId": "query-api-audits",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Plane",
                        "name": "plane",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "HTTP Path",
                        "name": "http_path",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "HTTP Method",
                        "name": "http_method",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Sent HTTP Status",
                        "name": "sent_http_status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Order By",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Order Desc",
                        "name": "order_desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apimodel.APIAudit"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/grants": {
            "post": {
                "description": "Create a grant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Grant"
                ],
                "summary": "Create a grant",
                "operationId": "create-grant",
                "parameters": [
                    {
                        "description": "Create Grant parameters",
                        "name": "createGrant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.CreateGrant"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Grant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "Show the health of the controlplane",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthz"
                ],
                "summary": "Show the health of the controlplane",
                "operationId": "healthz",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Healthz"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.APIAudit": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "hashed_grant_token": {
                    "type": "string"
                },
                "http_method": {
                    "type": "string"
                },
                "http_path": {
                    "type": "string"
                },
                "plane": {
                    "type": "string"
                },
                "sent_http_status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "apimodel.CreateGrant": {
            "type": "object",
            "properties": {
                "allowed_http_methods": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "domain": {
                    "type": "string"
                }
            }
        },
        "apimodel.Grant": {
            "type": "object",
            "properties": {
                "allowed_http_methods": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "domain": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "apimodel.Healthz": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "object"
                }
            }
        }
    },
    "tags": [
        {
            "description": "A grant is the unit of grant to access certain resource in OPV",
            "name": "Grant"
        }
    ]
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Open Privacy Vault Control Plane API",
	Description: "Open Privacy Vault Control Plane API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
