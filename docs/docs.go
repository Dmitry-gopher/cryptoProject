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
        "/getavg": {
            "get": {
                "description": "Returns a list of coins with their average rates in USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coins"
                ],
                "summary": "Get average coin rates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List of coin titles (separated by commas, for example: BTC,ETH)",
                        "name": "titles",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coin list",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CoinDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/getlast": {
            "get": {
                "description": "Returns a list of coins with their latest rates in USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coins"
                ],
                "summary": "Get last coin rates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List of coin titles (separated by commas, for example: BTC,ETH)",
                        "name": "titles",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coin list",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CoinDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/getmax": {
            "get": {
                "description": "Returns a list of coins with their max rates in USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coins"
                ],
                "summary": "Get max coin rates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List of coin titles (separated by commas, for example: BTC,ETH)",
                        "name": "titles",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coin list",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CoinDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/getmin": {
            "get": {
                "description": "Returns a list of coins with their min rates in USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "coins"
                ],
                "summary": "Get min coin rates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List of coin titles (separated by commas, for example: BTC,ETH)",
                        "name": "titles",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coin list",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CoinDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CoinDTO": {
            "type": "object",
            "properties": {
                "current_rate": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Crypto API",
	Description:      "This is a sample server Crypto server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
