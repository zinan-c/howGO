# howGO
learn how to use the go

*For simple project:
project/
├── main.go
├── go.mod
└── *.go

*Midle size prj:
project/
├── cmd/
├── internal/
├── pkg/
└── go.mod

*Large size prj:
project/
├── cmd/
│   ├── user-service/
│   ├── product-service/
│   └── order-service/
├── internal/
├── pkg/
│   ├── common/
│   └── libs/
└── go.mod