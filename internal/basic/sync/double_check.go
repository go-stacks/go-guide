package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrStore(key K,
	newVale V) (val V, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	s.m[key] = newVale
	return newVale, false
}

type valProvider[V any] func() V

func (s *SafeMap[K, V]) LoadOrStoreHeavy(key K, p valProvider[V]) (interface{}, bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}

	val = p()
	s.m[key] = val
	return val, false
}
