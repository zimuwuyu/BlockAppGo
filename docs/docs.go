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
        "/v1/blockModel": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "这个接口返回分页的 BlockModel 列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "分页返回 BlockModel 数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.BlockModel"
                            }
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "通过用户名和密码进行登录，并返回 JWT Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回 Token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "用户不存在",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "description": "通过用户名和密码注册新用户(暂时不开放)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册(暂时不开放)",
                "parameters": [
                    {
                        "description": "用户注册信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "model.BlockModel": {
            "type": "object",
            "properties": {
                "bucketName": {
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "hot": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "isRecommend": {
                    "type": "boolean"
                },
                "isShow": {
                    "type": "boolean"
                },
                "modelFile": {
                    "type": "string"
                },
                "modelType": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "totalCount": {
                    "type": "integer"
                },
                "updateTime": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
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
