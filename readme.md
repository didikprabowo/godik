# Best Practice Go Language
## by Didik Prabowo

### Struktur
```
├── app
│   ├── handlers
│   │   └── example.go
│   └── routes
│       └── web.go
├── config.json
├── Gopkg.lock
├── Gopkg.toml
├── main.go
├── readme.md
├── system
│   ├── http.go
│   └── routes.go
└── vendor
    └── github.com
        └── gorilla
            └── mux
                ├── AUTHORS
                ├── context.go
                ├── doc.go
                ├── go.mod
                ├── ISSUE_TEMPLATE.md
                ├── LICENSE
                ├── middleware.go
                ├── mux.go
                ├── README.md
                ├── regexp.go
                ├── route.go
                └── test_helpers.go
```
#### Description
1. config.json, Mengubah Konfigurasi Host dan Port.
2. App, Workspace mengelola Projek