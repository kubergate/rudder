package datastore

import (
	"errors"
	"k8s.io/apimachinery/pkg/types"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	"sync"
)

type HttpRouteKVStore struct {
	mu    sync.RWMutex
	store map[string]*gwapiv1b1.HTTPRoute
}

func NewHttpRouteKVStore() *HttpRouteKVStore {
	return &HttpRouteKVStore{
		store: make(map[string]*gwapiv1b1.HTTPRoute),
	}
}

func (httpRouteKVStore *HttpRouteKVStore) Get(key types.NamespacedName) (gwapiv1b1.HTTPRoute, bool) {
	httpRouteKVStore.mu.RLock()
	defer httpRouteKVStore.mu.RUnlock()

	value, exists := httpRouteKVStore.store[key.String()]
	if exists {
		return *value, exists
	}
	return gwapiv1b1.HTTPRoute{}, exists
}

func (httpRouteKVStore *HttpRouteKVStore) Add(key types.NamespacedName, value gwapiv1b1.HTTPRoute) (*gwapiv1b1.HTTPRoute, error) {
	httpRouteKVStore.mu.Lock()
	defer httpRouteKVStore.mu.Unlock()

	httpRouteKVStore.store[key.String()] = &value
	return &value, nil
}

func (httpRouteKVStore *HttpRouteKVStore) Delete(key types.NamespacedName) error {
	httpRouteKVStore.mu.Lock()
	defer httpRouteKVStore.mu.Unlock()

	if _, exists := httpRouteKVStore.store[key.String()]; !exists {
		return errors.New("key not found")
	}
	delete(httpRouteKVStore.store, key.String())
	return nil
}

func (httpRouteKVStore *HttpRouteKVStore) Contains(key types.NamespacedName) bool {
	httpRouteKVStore.mu.RLock()
	defer httpRouteKVStore.mu.RUnlock()

	_, exists := httpRouteKVStore.store[key.String()]
	return exists
}
