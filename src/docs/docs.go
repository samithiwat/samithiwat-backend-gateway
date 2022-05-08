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
        "contact": {
            "name": "Samithiwat",
            "url": "https://samithiwat.dev",
            "email": "admin@samithiwat.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "get": {
                "description": "Return the arrays of user dto if successfully",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.User"
                        }
                    },
                    "400": {
                        "description": "Invalid query param",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "503": {
                        "description": "Service is down",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Return the user dto if successfully",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create the user",
                "parameters": [
                    {
                        "description": "user dto",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/proto.User"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "404": {
                        "description": "Not found user",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "503": {
                        "description": "Service is down",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Return the user dto if successfully",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get specific user with id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.User"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "404": {
                        "description": "Not found user",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "503": {
                        "description": "Service is down",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "Return the user dto if successfully",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete the user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.User"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "404": {
                        "description": "Not found user",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "503": {
                        "description": "Service is down",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    }
                }
            },
            "patch": {
                "description": "Return the user dto if successfully",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update the existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user dto",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.User"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "404": {
                        "description": "Not found user",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    },
                    "503": {
                        "description": "Service is down",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserDto": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserDto": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "model.ResponseErr": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "proto.Contact": {
            "type": "object",
            "properties": {
                "facebook": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instagram": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "twitter": {
                    "type": "string"
                }
            }
        },
        "proto.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "province": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "proto.Log": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/proto.User"
                }
            }
        },
        "proto.Organization": {
            "type": "object",
            "properties": {
                "contact": {
                    "$ref": "#/definitions/proto.Contact"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/proto.Location"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Log"
                    }
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.User"
                    }
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Role"
                    }
                },
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Team"
                    }
                }
            }
        },
        "proto.Permission": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Role"
                    }
                }
            }
        },
        "proto.Role": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Permission"
                    }
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.User"
                    }
                }
            }
        },
        "proto.Team": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Log"
                    }
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.User"
                    }
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "$ref": "#/definitions/proto.Organization"
                },
                "subTeams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Team"
                    }
                }
            }
        },
        "proto.User": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/proto.Location"
                },
                "contact": {
                    "$ref": "#/definitions/proto.Contact"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Log"
                    }
                },
                "organizations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Organization"
                    }
                },
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.Team"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "Auth Token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "# User Tag API Documentation\r\n**User** functions goes here",
            "name": "user"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "Samithiwat Backend",
	Description:      "# Samithiwat's API\r\nThis is the documentation for https://samithiwat.dev",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
