// Code generated by swaggo/swag. DO NOT EDIT
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
        "/ads": {
            "post": {
                "description": "Receive an ad based on input paramenters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ads"
                ],
                "summary": "Receive Ad",
                "parameters": [
                    {
                        "description": "Request Ad",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestAd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AdResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "videoURL": {
                    "type": "string"
                }
            }
        },
        "models.RequestAd": {
            "type": "object",
            "properties": {
                "countryCode": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "userID": {
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
