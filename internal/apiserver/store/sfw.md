
store.go 定义
```
type Factory interface {
	Users() UserStore
	Secrets() SecretStore
	Policies() PolicyStore
	Close() error
}
```

抽象层：
- user.go 定义 UserStore interface
- secret.go 定义 SecretStore interface
- ploicy.go 定义 PolicyStore interface

实现层：
- mysql/user.go 实现 UserStore interface
- mysql/secret.go 实现 SecretStore interface
- mysql/ploicy.go 实现 PolicyStore interface

