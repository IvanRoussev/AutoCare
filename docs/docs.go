// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import (
	"github.com/swaggo/swag"
)

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "version": "1.0",
        "title": "Auto Care",
        "contact": {
            "email": "ivan.roussev12@gmail.com",
            "name": "Ivan Roussev"
        }
    },
    "host": "localhost:8080",
    "consumes": ["application/json"],
    "produces": ["application/json"],
    "paths": {
        "/users": {
            "post": {
                "tags": ["user"],
                "summary": "Creates a user",
                "description": "Create yourself a user",
                "parameters": [
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "properties": {
                                "username": {
                                    "type": "string",
                                    "example": "AmazingLama23"
                                },
                                "password": {
                                    "type": "string",
                                    "example": "Password"
                                },
                                "full_name": {
                                    "type": "string",
                                    "example": "John Smith"
                                },
                                "email": {
                                    "type": "string",
                                    "example": "example@gmail.com"
                                },
                            },
                            "required": ["username", "password", "full_name", "email"]
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created User"
                    },
                    "400": {
                        "description": "Bad request - Invalid input"
                    },
                    "500": {
                        "description": "Internal server error"
                    },
                    "403": {
                        "description": "Status Forbidden Username is already taken"
                    }
                }
            }
        }
    }
}
`

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
