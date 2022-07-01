// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/users": {
            "get": {
                "description": "This api is using for getting users list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "This api is using for creating new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user summary",
                "parameters": [
                    {
                        "description": "user body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/login/{email}/{password}": {
            "post": {
                "description": "This api using for logging registered user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/users/register": {
            "post": {
                "description": "This api is using for registering new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register user summary",
                "parameters": [
                    {
                        "description": "user_body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/verfication": {
            "post": {
                "description": "This api using for verifying registered user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "user body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Emailver"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "This api is using for getting user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "This api is using for updating new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "This api is using for deleting user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.Adress": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "description": "Id                   string   ` + "`" + `protobuf:\"bytes,1,opt,name=id,proto3\" json:\"id\"` + "`" + `\nUserId               string   ` + "`" + `protobuf:\"bytes,2,opt,name=user_id,json=userId,proto3\" json:\"user_id\"` + "`" + `",
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "postal_codes": {
                    "type": "integer"
                }
            }
        },
        "v1.Emailver": {
            "type": "object",
            "properties": {
                "Code": {
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                }
            }
        },
        "v1.Media": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "type": {
                    "description": "Id                   string   ` + "`" + `protobuf:\"bytes,1,opt,name=id,proto3\" json:\"id\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "v1.Post": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "medias": {
                    "description": "UserId               string   ` + "`" + `protobuf:\"bytes,4,opt,name=user_id,json=userId,proto3\" json:\"user_id\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Media"
                    }
                },
                "name": {
                    "description": "Id                   string   ` + "`" + `protobuf:\"bytes,1,opt,name=id,proto3\" json:\"id\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "v1.User": {
            "type": "object",
            "properties": {
                "adress": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Adress"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "post": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Post"
                    }
                },
                "status": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
