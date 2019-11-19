# 博客 REST API 设计

## 总体设计

+ 博客不开放用户注册接口，只能由管理员登录。
+ 博客的评论系统对接 Disqus 平台。

## 获取图形验证码

管理员在登录博客之前，需要获取图形验证码，**该接口无需鉴权**。

### Request

```
GET /api/v1/user/captcha
```

### Response

>Status: 200 OK
>
>Location: /api/v1/user/captcha

```json
{
   	"captcha_id": "UUID of captcha image",
    "base64": "Base64 encoding of captcha image"
}
```

## 管理员登录

管理员通过访问此接口获取 `Access Token` ，**该接口无需鉴权**。

### Request

```
POST /api/v1/users/login
```

| 字段         | 类型   | 描述           |
| ------------ | ------ | -------------- |
| username     | string | 管理员用户名   |
| password     | string | 管理员密码     |
| captcha_id   | string | 验证码 ID      |
| captcha_code | string | 验证码识别结果 |

### Response

>Status: 200 OK
>
>Location: /api/v1/user/login

```json
{
    "access_token": "string",
    "expires_at": "date"
}
```

## 获取分类列表

我们可访问此接口获取博客的分类列表，**该接口无需鉴权**。

### Request

```
GET /api/v1/categories
```

| 字段     | 类型   | 描述   |
| -------- | ------ | ------ |
| page     | number | 页码   |
| per_page | number | 页大小 |

### Response

> Status: 200 OK
>
> Location: /api/v1/categories

```json
{
    "categories": [
        {
        "category_id": 0,
        "category_name": "Programming",
        "count": 63
    	},
    	{
        "category_id": 1,
        "category_name": "Finance",
        "count": 31
    	}
    ]
}
```

## 创建分类列表

管理员可访问此接口实现分类列表的创建，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized` 。

### Request

```
POST /api/v1/categories
```

| 字段          | 类型   | 描述         |
| ------------- | ------ | ------------ |
| category_name | string | 分类列表名称 |

### Response

> Status: 201 Created
>
> Location: /api/v1/categories

```json
{
    "category_id": 0,
    "category_name": "程序设计"
}
```

## 更新分类列表

管理员可访问此接口实现分类列表的更新，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized`。

### Request

```
PUT /api/v1/categories/:category_id
```

| 字段          | 类型   | 描述         |
| ------------- | ------ | ------------ |
| category_name | string | 分类列表名称 |

### Response

> Status: 200 OK
>
> Location: /api/v1/categories/0

## 删除分类列表

管理员可访问此接口实现分类列表的删除，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized`。

### Request

```
DELETE /api/v1/categories:category_id
```

### Response

> Status: 204 No Content
>
> Location: /api/v1/categories/0

## 获取分类内的文章

我们可通过此接口获取分类内的文章，**该接口无需鉴权**。

### Request

```
GET /api/v1/categories/:category_id
```

| 字段     | 类型   | 描述   |
| -------- | ------ | ------ |
| page     | number | 页码   |
| per_page | number | 页大小 |

### Response

> Status: 200 OK
>
> Location: /api/v1/categories/0

```json
{
    "category_name": "Programming",
    "articles": [
        {
            "article_id": 0,
            "article_title": "Golang Lua",
            "created_at": "2019-03-05"
        },
        {
            "article_id": 1,
            "article_title": "Golang 热重启",
            "created_at": "2019-07-17"
        }
    ]
}
```

## 获取文章

我们可通过此接口获取文章，**该接口无需鉴权**。

### Request

```
GET /api/v1/articles:article_id
```

### Response

> Status: 200 OK
>
> Location: /api/v1/articles/0

```json
{
    "article_id": 0,
    "article_title": "Golang Lua",
    "article_content": "xxxxxx"
}
```

## 创建文章

管理员可通过此接口创建文章，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized`。

### Request

```
POST /api/v1/articles
```

| 字段            | 类型   | 描述        |
| --------------- | ------ | ----------- |
| category_id     | number | 分类列表 ID |
| article_title   | string | 文章标题    |
| article_content | string | 文章内容    |

### Response

> Status: 201 Created
>
> Location: /api/v1/articles

```json
{
    "article_id": 0,
    "category_id": 0,
    "article_title": "Golang Lua",
    "article_content": "xxxxxx"
}
```

## 更新文章

管理员可通过此接口实现对文章的更新，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized`。

### Request

```
PUT /api/v1/articles/0
```

| 字段            | 类型   | 描述        |
| --------------- | ------ | ----------- |
| category_id     | number | 分类列表 ID |
| article_title   | string | 文章标题    |
| article_content | string | 文章内容    |

### Response

> Status: 200 OK
>
> Location: /api/v1/articles/0

```json
{
    "article_id": 0,
    "category_id": 0,
    "article_title": "Golang Lua",
    "article_content": "xxxxxx"
}
```

## 删除文章

管理员可通过此接口删除文章，**该接口需要鉴权**，若无鉴权返回 `401 Unauthorized`。

### Request

```
DELETE /api/v1/articles/:article_id
```

### Response

> Status: 204 No Content
>
> Location: /api/v1/articles/0