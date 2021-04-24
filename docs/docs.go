// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/evelive3/go-gin-api/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/admin": {
            "get": {
                "description": "管理员列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "管理员列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "第几页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页显示条数",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "mobile",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.listResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            },
            "post": {
                "description": "新增管理员",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "新增管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.createResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/info": {
            "get": {
                "description": "管理员详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "管理员详情",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.detailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/login": {
            "post": {
                "description": "管理员登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "管理员登出",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.logoutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/modify_password": {
            "patch": {
                "description": "修改密码",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "旧密码",
                        "name": "old_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "new_password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.modifyPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/reset_password/{id}": {
            "patch": {
                "description": "重置密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hashId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.resetPasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/used": {
            "patch": {
                "description": "更新管理员为启用/禁用",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "更新管理员为启用/禁用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hashid",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "是否启用 1:是 -1:否",
                        "name": "used",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.updateUsedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/admin/{id}": {
            "delete": {
                "description": "删除管理员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.admin"
                ],
                "summary": "删除管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hashId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin_handler.deleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/authorized": {
            "get": {
                "description": "调用方列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "调用方列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "第几页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页显示条数",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "调用方key",
                        "name": "business_key",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "调用方secret",
                        "name": "business_secret",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "调用方对接人",
                        "name": "business_developer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.listResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            },
            "post": {
                "description": "新增调用方",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "新增调用方",
                "parameters": [
                    {
                        "type": "string",
                        "description": "调用方key",
                        "name": "business_key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "调用方对接人",
                        "name": "business_developer",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.createResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/authorized/used": {
            "patch": {
                "description": "更新调用方为启用/禁用",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "更新调用方为启用/禁用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hashid",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "是否启用 1:是 -1:否",
                        "name": "used",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.updateUsedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/authorized/{id}": {
            "delete": {
                "description": "删除调用方",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "删除调用方",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hashId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.deleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/authorized_api": {
            "get": {
                "description": "调用方接口地址列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "调用方接口地址列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "调用方key",
                        "name": "business_key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.listAPIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            },
            "post": {
                "description": "授权调用方接口地址",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "授权调用方接口地址",
                "parameters": [
                    {
                        "type": "string",
                        "description": "HashID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求方法",
                        "name": "method",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求地址",
                        "name": "api",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.createAPIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/authorized_api/{id}": {
            "delete": {
                "description": "删除调用方接口地址",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.authorized"
                ],
                "summary": "删除调用方接口地址",
                "parameters": [
                    {
                        "type": "string",
                        "description": "主键ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorized_handler.deleteAPIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/config/email": {
            "patch": {
                "description": "修改邮件配置",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.config"
                ],
                "summary": "修改邮件配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱服务器",
                        "name": "host",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "端口",
                        "name": "port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "发件人邮箱",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "发件人密码",
                        "name": "pass",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "收件人邮箱地址，多个用,分割",
                        "name": "to",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config_handler.emailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/tool/hashids/decode/{id}": {
            "get": {
                "description": "HashIds 解密",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.tool"
                ],
                "summary": "HashIds 解密",
                "parameters": [
                    {
                        "type": "string",
                        "description": "需解密的密文",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool_handler.hashIdsDecodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/api/tool/hashids/encode/{id}": {
            "get": {
                "description": "HashIds 加密",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API.tool"
                ],
                "summary": "HashIds 加密",
                "parameters": [
                    {
                        "type": "string",
                        "description": "需加密的数字",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool_handler.hashIdsEncodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin_handler.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "admin_handler.deleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "admin_handler.detailResponse": {
            "type": "object",
            "properties": {
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "admin_handler.listData": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "created_user": {
                    "description": "创建人",
                    "type": "string"
                },
                "hashid": {
                    "description": "hashid",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "is_used": {
                    "description": "是否启用 1:是 -1:否",
                    "type": "integer"
                },
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                },
                "updated_user": {
                    "description": "更新人",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "admin_handler.listResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin_handler.listData"
                    }
                },
                "pagination": {
                    "type": "object",
                    "properties": {
                        "current_page": {
                            "type": "integer"
                        },
                        "pre_page_count": {
                            "type": "integer"
                        },
                        "total": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "admin_handler.loginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "用户身份标识",
                    "type": "string"
                }
            }
        },
        "admin_handler.logoutResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "description": "用户账号",
                    "type": "string"
                }
            }
        },
        "admin_handler.modifyPasswordResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "description": "用户账号",
                    "type": "string"
                }
            }
        },
        "admin_handler.modifyPersonalInfoResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "description": "用户账号",
                    "type": "string"
                }
            }
        },
        "admin_handler.resetPasswordResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "admin_handler.updateUsedResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "authorized_handler.createAPIResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "authorized_handler.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "authorized_handler.deleteAPIResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "authorized_handler.deleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "authorized_handler.listAPIData": {
            "type": "object",
            "properties": {
                "api": {
                    "description": "调用方对接人",
                    "type": "string"
                },
                "business_key": {
                    "description": "调用方key",
                    "type": "string"
                },
                "hash_id": {
                    "description": "hashID",
                    "type": "string"
                },
                "method": {
                    "description": "调用方secret",
                    "type": "string"
                }
            }
        },
        "authorized_handler.listAPIResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/authorized_handler.listAPIData"
                    }
                }
            }
        },
        "authorized_handler.listData": {
            "type": "object",
            "properties": {
                "business_developer": {
                    "description": "调用方对接人",
                    "type": "string"
                },
                "business_key": {
                    "description": "调用方key",
                    "type": "string"
                },
                "business_secret": {
                    "description": "调用方secret",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "created_user": {
                    "description": "创建人",
                    "type": "string"
                },
                "hashid": {
                    "description": "hashid",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "is_used": {
                    "description": "是否启用 1:是 -1:否",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                },
                "updated_user": {
                    "description": "更新人",
                    "type": "string"
                }
            }
        },
        "authorized_handler.listResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/authorized_handler.listData"
                    }
                },
                "pagination": {
                    "type": "object",
                    "properties": {
                        "current_page": {
                            "type": "integer"
                        },
                        "pre_page_count": {
                            "type": "integer"
                        },
                        "total": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "authorized_handler.updateUsedResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "code.Failure": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务码",
                    "type": "integer"
                },
                "message": {
                    "description": "描述信息",
                    "type": "string"
                }
            }
        },
        "config_handler.emailResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱地址",
                    "type": "string"
                }
            }
        },
        "tool_handler.hashIdsDecodeResponse": {
            "type": "object",
            "properties": {
                "val": {
                    "description": "解密后的值",
                    "type": "integer"
                }
            }
        },
        "tool_handler.hashIdsEncodeResponse": {
            "type": "object",
            "properties": {
                "val": {
                    "description": "加密后的值",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "127.0.0.1:9999",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "swagger 接口文档",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
