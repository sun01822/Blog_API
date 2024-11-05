// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user/create": {
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
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user created successfully",
                        "schema": {
                            "$ref": "#/definitions/types.UserResp"
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error creating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error deleting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "description": "Get a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user found successfully",
                        "schema": {
                            "$ref": "#/definitions/types.UserResp"
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error getting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getAll": {
            "get": {
                "description": "Get all users",
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
                        "type": "string",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "users found successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UserResp"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error getting user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Logs in a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "invalid email or password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user updated successfully",
                        "schema": {
                            "$ref": "#/definitions/types.UserResp"
                        }
                    },
                    "400": {
                        "description": "invalid data request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error updating user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.SignUpRequest": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string",
                    "default": "Bangladesh"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "types.UserResp": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string",
                    "default": "Bangladesh"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "job": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "phone": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "types.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string",
                    "default": "Bangladesh"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "job": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "192.168.10.135:8080",
	BasePath:         "/blog_api/v1",
	Schemes:          []string{},
	Title:            "Blog API",
	Description:      "This is a sample server for Blog CRUD Operation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
