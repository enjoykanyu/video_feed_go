
```
douyin
├─ 📁gateway
│  ├─ 📁cmd
│  │  └─ 📄main.go
│  ├─ 📁internal
│  │  ├─ 📁controller
│  │  │  └─ 📄user.go
│  │  ├─ 📁grpc_client
│  │  │  └─ 📄grpc_client.go
│  │  └─ 📁router
│  │     └─ 📄router.go
│  ├─ 📄go.mod
│  └─ 📄go.sum
├─ 📁user_service
│  ├─ 📁internal
│  │  ├─ 📁config
│  │  │  ├─ 📄application.yaml
│  │  │  └─ 📄config.go
│  │  ├─ 📁handler
│  │  │  └─ 📄user.go
│  │  ├─ 📁repostry
│  │  │  └─ 📄user.go
│  │  ├─ 📁service
│  │  │  └─ 📄user.go
│  │  └─ 📁store
│  │     └─ 📁model
│  │        ├─ 📄model.go
│  │        └─ 📄user.go
│  ├─ 📁proto
│  │  ├─ 📁pb
│  │  │  └─ 📄user.proto
│  │  ├─ 📄user.pb.go
│  │  └─ 📄user_grpc.pb.go
│  ├─ 📄go.mod
│  ├─ 📄go.sum
│  └─ 📄main.go
├─ 📄go.mod
└─ 📄go.sum
```