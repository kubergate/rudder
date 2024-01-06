package datastore

import (
	"github.com/KommodoreX/dp-rudder/pkg/rudder/types"
)

const HTTPRouteBucket = "bkt_HTTPRoute"

// RudderKVStore interface defines the methods that any storage implementation should support.
type RudderKVStore interface {
	Add(key string, value types.Resource[interface{}]) error
	Get(key string) (types.Resource[interface{}], error)
	Update(key, value types.Resource[interface{}]) error
	Delete(key string) error
}
