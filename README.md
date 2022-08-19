# adapter
golang service adapter

# how to use

1. get conn from resolver

```golang
srv := "your service name"
conn, err := adapter.NewTarget().Dial(srv)
if err != nil {
    log.Fatal("resolver etcd err", err)
}
defer conn.Close()

// use conn todo anything
```

2. get ipv4 for register

```golang
ip := adaper.NewIP().Local()

// use ip todo anything
```