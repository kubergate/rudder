package datastore

import (
	"github.com/kubergate/rudder/pkg/rudder/types"
)

// RudderKVStore interface defines the methods that any storage implementation should support.
type RudderKVStore[Key comparable] interface {
	Add(key Key, value types.Resource[any]) error
	Get(key Key) (types.Resource[any], error)
	Delete(key Key) error
}
