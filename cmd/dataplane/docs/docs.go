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
        "/fact_types": {
            "post": {
                "description": "create a fact type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fact"
                ],
                "summary": "Create a fact type",
                "operationId": "create-fact-type",
                "parameters": [
                    {
                        "description": "Create Fact Type Parameters",
                        "name": "createFact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.CreateFactType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.FactType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        },
        "/fact_types/{id}": {
            "get": {
                "description": "Show a fact type by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fact"
                ],
                "summary": "Show a fact Type",
                "operationId": "show-fact-type-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Fact Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.FactType"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        },
        "/facts": {
            "post": {
                "description": "create a fact",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fact"
                ],
                "summary": "Create a fact",
                "operationId": "create-fact",
                "parameters": [
                    {
                        "description": "Create Fact Parameters",
                        "name": "createFact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.CreateFact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Fact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        },
        "/facts/{id}": {
            "get": {
                "description": "Show a fact by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fact"
                ],
                "summary": "Show a fact",
                "operationId": "show-fact-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Fact ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Fact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        },
        "/scopes": {
            "post": {
                "description": "Create a scope",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scope"
                ],
                "summary": "Create a scope",
                "operationId": "create-scope",
                "parameters": [
                    {
                        "description": "Create Scope parameters",
                        "name": "createScope",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.CreateScope"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Scope"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        },
        "/scopes/{id}": {
            "get": {
                "description": "Show scope by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Scope"
                ],
                "summary": "Show a scope",
                "operationId": "show-scope-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Scope ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimodel.Scope"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.CreateFact": {
            "type": "object",
            "properties": {
                "fact_type_slug": {
                    "type": "string"
                },
                "scope_custom_id": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "apimodel.CreateFactType": {
            "type": "object",
            "properties": {
                "slug": {
                    "type": "string"
                }
            }
        },
        "apimodel.CreateScope": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "apimodel.Fact": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "fact_type_slug": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "scope_custom_id": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "apimodel.FactType": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "apimodel.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "apimodel.Scope": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "A scope is the unit of encryption isolation unit, usually it represents a person as a scope",
            "name": "Scope"
        },
        {
            "description": "A fact is the unit of PII information, e.g. email, address, phone number, and etc.",
            "name": "Fact"
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
	Title:       "Open Privacy Vault Data Plane API",
	Description: "Open Privacy Vault Data Plane API.",
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
