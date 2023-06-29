// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/atados/v1/health": {
            "get": {
                "description": "healthcheck router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "healthcheck router",
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
        "/atados/v1/voluntary": {
            "get": {
                "description": "get all products by pagination router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voluntary"
                ],
                "summary": "get all products by pagination router",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetAllVoluntariesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "create product router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voluntary"
                ],
                "summary": "create product router",
                "parameters": [
                    {
                        "description": "create product",
                        "name": "createProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateVoluntaryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/CreateVoluntaryVO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/atados/v1/voluntary/{id}": {
            "get": {
                "description": "get product by id router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voluntary"
                ],
                "summary": "get product by id router",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetVoluntaryByIDResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "update product router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voluntary"
                ],
                "summary": "update product router",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update product",
                        "name": "updateProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateVoluntaryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Voluntary updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "delete product router",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voluntary"
                ],
                "summary": "delete product router",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Voluntary deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateVoluntaryRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                }
            }
        },
        "CreateVoluntaryVO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "GetAllVoluntariesResponse": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/VoluntaryPagination"
                },
                "voluntaries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Voluntary"
                    }
                }
            }
        },
        "GetVoluntaryByIDResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "UpdateVoluntaryRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                }
            }
        },
        "Voluntary": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "VoluntaryPagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
