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
        "/recipes/all": {
            "get": {
                "description": "List All Recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "List all recipes",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/recipes/create": {
            "post": {
                "description": "Create A New Recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Create a new recipe",
                "parameters": [
                    {
                        "description": "Create New Recipe",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipe.Recipe"
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
        "/recipes/delete": {
            "post": {
                "description": "Delete Recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Delete Recipe",
                "parameters": [
                    {
                        "description": "Delete a Recipe",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipe.Recipe"
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
        "/recipes/get": {
            "post": {
                "description": "Get a recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Get a recipe",
                "parameters": [
                    {
                        "description": "Get Recipe",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipe.Recipe"
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
        "/recipes/update": {
            "post": {
                "description": "Update Recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "Update Recipe",
                "parameters": [
                    {
                        "description": "Update a Recipe",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipe.Recipe"
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
        "/search/recipes": {
            "post": {
                "description": "Search Recipes by Title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search Recipes",
                "parameters": [
                    {
                        "description": "Please write search word directly",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
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
        "/shopping/all": {
            "get": {
                "description": "List All Shopping Lists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopping"
                ],
                "summary": "Get shopping lists",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/shopping/create": {
            "post": {
                "description": "Create A New Shopping List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopping"
                ],
                "summary": "Create a new shopping list",
                "parameters": [
                    {
                        "description": "Create New shopping List",
                        "name": "shopping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shopping.Shopping"
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
        "/shopping/delete": {
            "post": {
                "description": "Delete Shopping List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopping"
                ],
                "summary": "Delete Shopping List",
                "parameters": [
                    {
                        "description": "Delete Shopping List",
                        "name": "shopping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shopping.Shopping"
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
        "/shopping/get": {
            "post": {
                "description": "Get a shopping list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopping"
                ],
                "summary": "Get a shopping list",
                "parameters": [
                    {
                        "description": "Get Shopping List",
                        "name": "shopping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shopping.Shopping"
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
        "/shopping/update": {
            "post": {
                "description": "Update Shopping List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopping"
                ],
                "summary": "Update Shopping List",
                "parameters": [
                    {
                        "description": "Update Shopping List",
                        "name": "shopping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shopping.Shopping"
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
        "ingredient.Ingredient": {
            "type": "object",
            "properties": {
                "isexist": {
                    "type": "boolean",
                    "example": false
                },
                "isoptional": {
                    "type": "boolean",
                    "example": true
                },
                "material": {
                    "$ref": "#/definitions/material.Material"
                },
                "measurement": {
                    "description": "gorm.Model  ` + "`" + `json:\"-\" swaggerignore:\"true\"` + "`" + `\nSize        float32",
                    "$ref": "#/definitions/measurement.Measurement"
                }
            }
        },
        "material.Material": {
            "type": "object",
            "properties": {
                "material_photo_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "S3URL1",
                        "S3URL2"
                    ]
                },
                "name": {
                    "type": "string",
                    "example": "banana"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "vegan",
                        "fruit"
                    ]
                }
            }
        },
        "measurement.Measurement": {
            "type": "object",
            "properties": {
                "quantity": {
                    "type": "string",
                    "example": "gram"
                },
                "size": {
                    "description": "gorm.Model ` + "`" + `json:\"-\" swaggerignore:\"true\"` + "`" + `",
                    "type": "number",
                    "example": 2
                }
            }
        },
        "recipe.Recipe": {
            "type": "object",
            "properties": {
                "appropriate_meals": {
                    "description": "Breakfast, Snack , Noon , AfterNoon , Evening , Night",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "breakfast",
                        "snack"
                    ]
                },
                "calori": {
                    "type": "integer",
                    "example": 252
                },
                "cooking_time_minute": {
                    "type": "integer",
                    "example": 10
                },
                "for_how_many_people": {
                    "type": "integer",
                    "example": 2
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ingredient.Ingredient"
                    }
                },
                "photo_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "S3URL1",
                        "S3URL2"
                    ]
                },
                "preperation": {
                    "type": "string",
                    "example": "bla bla bla"
                },
                "preperation_time": {
                    "type": "integer",
                    "example": 15
                },
                "title": {
                    "type": "string",
                    "example": "Sushi With Wassabi"
                },
                "video_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "S3URL1",
                        "S3URL2"
                    ]
                }
            }
        },
        "shopping.Shopping": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string",
                    "example": "1643743448"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ingredient.Ingredient"
                    }
                },
                "start_date": {
                    "type": "string",
                    "example": "1643743444"
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
                "dislikes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "onion"
                    ]
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
                "gender": {
                    "description": "male,female,other",
                    "type": "string",
                    "example": "male"
                },
                "height": {
                    "type": "integer",
                    "example": 170
                },
                "likes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "kebap",
                        "pizza"
                    ]
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
                "preferred_meals": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "breakfast"
                    ]
                },
                "prohibits": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "sugar"
                    ]
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
                "wants": {
                    "description": "gain , lost , protect // (weights)",
                    "type": "string",
                    "example": "gain"
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
