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
        "/api/v1/witchcraft": {
            "get": {
                "description": "Get All Requirments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Witchcraft"
                ],
                "summary": "Get All Requirments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Witchcraft"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create New Requirment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Witchcraft"
                ],
                "summary": "Create New Requirment",
                "parameters": [
                    {
                        "description": "Create Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateWitchcraft"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/witchcraft/{id}": {
            "get": {
                "description": "Get Requirment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Witchcraft"
                ],
                "summary": "Get  Requirment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Witchcraft"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Requirment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Witchcraft"
                ],
                "summary": "Update Requirment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateWitchcraft"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Witchcraft"
                ],
                "summary": "Delete Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateWitchcraft": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "balanceAmount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "payment": {
                    "type": "integer"
                },
                "place": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "visitorsDuration": {
                    "type": "string"
                },
                "witchcraftId": {
                    "type": "string"
                }
            }
        },
        "model.UpdateWitchcraft": {
            "type": "object",
            "properties": {
                "VisitorsDuration": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "balanceAmount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "payment": {
                    "type": "integer"
                },
                "place": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                }
            }
        },
        "model.Witchcraft": {
            "type": "object",
            "properties": {
                "VisitorsDuration": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "balanceAmount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "payment": {
                    "type": "integer"
                },
                "place": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "witchcraftId": {
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
