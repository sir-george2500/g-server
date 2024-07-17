// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "George S Mulbah",
            "url": "https://github.com/sir-george2500"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/feed_follows": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve all feed follows for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed_follows"
                ],
                "summary": "Get Feed Follows",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.FeedsFollow"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new feed follow for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed_follows"
                ],
                "summary": "Create Feed Follow",
                "parameters": [
                    {
                        "description": "Feed Follow Parameters",
                        "name": "feed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.paramenters"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/database.FeedsFollow"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/feed_follows/{feedFollowID}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a feed follow for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed_follows"
                ],
                "summary": "Delete Feed Follow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Feed Follow ID",
                        "name": "feedFollowID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/feeds": {
            "get": {
                "description": "Get the list of feeds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feeds"
                ],
                "summary": "Get Feeds",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Feed"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new feed with the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feeds"
                ],
                "summary": "Create Feed",
                "parameters": [
                    {
                        "description": "Feed Parameters",
                        "name": "feed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.paramenters_feed_swag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.CreateFeed_swag"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve posts associated with the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get Posts for User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve details of the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "User Parameters",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.CreateUser"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.FeedsFollow": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "feedID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "main.CreateFeed_swag": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "main.CreateUser": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "main.Feed": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "main.Post": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "feed_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "publishe_at": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "main.User": {
            "type": "object",
            "properties": {
                "api_key": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "main.paramenters": {
            "type": "object",
            "properties": {
                "feed_id": {
                    "type": "string"
                }
            }
        },
        "main.paramenters_feed_swag": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
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
	Version:          "1.0",
	Host:             "g-server-7fg9.onrender.com",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "g-server API",
	Description:      "g-server - a full-blown RSS feed aggregator.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
