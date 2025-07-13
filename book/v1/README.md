# 演示

## 启动

```sh
➜ go run v1/main.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/books                --> main.(*BookApiHandler).ListBook-fm (3 handlers)
[GIN-debug] POST   /api/books                --> main.(*BookApiHandler).CreateBook-fm (3 handlers)
[GIN-debug] GET    /api/books/:bn            --> main.(*BookApiHandler).GetBook-fm (3 handlers)
[GIN-debug] PUT    /api/books/:bn            --> main.(*BookApiHandler).UpdateBook-fm (3 handlers)
[GIN-debug] DELETE /api/books/:bn            --> main.(*BookApiHandler).DeleteBook-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

## CRUD 示例

- 创建书籍：
```sh
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d '{"title": "Go 语言", "author": "张三", "price": 39.99}'
```

- 获取所有书籍：
```sh
curl http://localhost:8080/api/books
```


- 根据 ID 获取书籍：
```sh
curl http://localhost:8080/api/books/1
```

- 更新书籍：
```sh
curl -X PUT http://localhost:8080/api/books/1 -H "Content-Type: application/json" -d '{"title": "Go 语言进阶", "author": "张三", "price": 49.99}'
```

- 删除书籍：
```sh
curl -X DELETE http://localhost:8080/api/books/1
```

这样就完成了一个简单的使用 MySQL 的 Book CRUD 示例。你可以根据需要进一步扩展功能。