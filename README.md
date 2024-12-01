# Project
Golang Logging Example on Elasticsearch


## Dependencies
- [go-elasticsearch](https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv)
- [zap](https://pkg.go.dev/go.uber.org/zap)


## Project Structure
```
📦 
├─ app
│  ├─ Dockerfile
│  └─ main.go
├─ config
│  ├─ elastic.go
│  └─ logger.go
│  └─ logger_es_writer.go
├─ src
│  ├─ services
│  │  └─ my_service.go
├─ .env
├─ .gitignore
├─ docker-compose.yml
├─ go.mod
├─ go.sum
└─ README.md
```