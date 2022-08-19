package adapter

import (
	"github.com/tiantour/protector"
	"google.golang.org/grpc/resolver"
)

func init() {
	b := protector.NewResolver()
	resolver.Register(b)
}
