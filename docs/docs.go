// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/calendar/recipes": {
            "get": {
                "description": "Get recipes of the user weekly",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calendar"
                ],
                "summary": "Get recipes of user weekly",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/calendar/recipes/create": {
            "post": {
                "description": "User creates a recipe in the calendar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calendar"
                ],
                "summary": "Create Recipe In User Calendar",
                "parameters": [
                    {
                        "description": "Create Recipe In Calendar",
                        "name": "calendar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/calendar.Calendar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/calendar/recipes/delete": {
            "post": {
                "description": "Delete Recipe In User Calendar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calendar"
                ],
                "summary": "Delete Recipe In User Calendar",
                "parameters": [
                    {
                        "description": "Delete Recipe In Calendar. Please Use thisfor send request: {'ID':1}",
                        "name": "calendar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/calendar.Calendar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/calendar/recipes/update": {
            "post": {
                "description": "Update Recipe In User Calendar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calendar"
                ],
                "summary": "Update Recipe In User Calendar",
                "parameters": [
                    {
                        "description": "Update Recipe In Calendar",
                        "name": "calendar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/calendar.Calendar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "description": "Logout for the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logout User",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "Sign in with specified user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Signin User",
                "parameters": [
                    {
                        "description": "Sign In",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Signup User",
                "parameters": [
                    {
                        "description": "Create New User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "calendar.Calendar": {
            "type": "object",
            "properties": {
                "date_epoch": {
                    "type": "integer",
                    "example": 1643743444
                },
                "recipe_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "myadress 123121"
                },
                "age": {
                    "type": "integer",
                    "example": 25
                },
                "diet": {
                    "description": "vegaterian , vegan , omnivor , carnivor",
                    "type": "string",
                    "example": "omnivor"
                },
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "favorite_recipes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "height": {
                    "type": "integer",
                    "example": 170
                },
                "name": {
                    "type": "string",
                    "example": "test user"
                },
                "password": {
                    "type": "string",
                    "example": "testpass"
                },
                "phone": {
                    "type": "string",
                    "example": "+905355353535"
                },
                "role": {
                    "description": "root,admin,editor,user,anonymous",
                    "type": "string",
                    "example": "user"
                },
                "username": {
                    "type": "string",
                    "example": "testuser"
                },
                "weight": {
                    "type": "integer",
                    "example": 70
                }
            }
        }
    }
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
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Flow-Diet-Backend API",
	Description: "This is a sample serice for managing orders",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
