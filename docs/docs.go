// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/createUser": {
            "post": {
                "description": "Create new user base on the json passed",
                "produces": [
                    "application/json"
                ],
                "summary": "Create-new-user",
                "operationId": "Create-new-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        },
        "/deleteUser/{id}": {
            "get": {
                "description": "Delete the user which have the id passed through the get request",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete user by ID",
                "operationId": "Delete-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        },
        "/getUser/{id}": {
            "get": {
                "description": "Get user by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get user by ID",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        },
        "/getUsers": {
            "get": {
                "description": "Get all user",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all user",
                "operationId": "get-all-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        },
        "/updateUser": {
            "post": {
                "description": "Update user info based on the json passed",
                "produces": [
                    "application/json"
                ],
                "summary": "Update-user",
                "operationId": "Update-user",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
