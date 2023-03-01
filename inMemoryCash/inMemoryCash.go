package inMemoryCash

import (
	"sync"
	"time"
)

type item struct {
	value     string
	lastAcces int64
}

type TTLCache struct {
	m                 map[string]*item
	l                 sync.Mutex
	expireAtTimestamp int64
}

func NewCache(ln int, maxTTL int) *TTLCache {
	m := &TTLCache{
		m: make(map[string]*item, ln),
	}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAcces > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return m

}

func (m *TTLCache) Len() int {
	return len(m.m)

}
func (m *TTLCache) Set(k string, v string) {
	m.l.Lock()

	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.lastAcces = time.Now().Unix()
	m.l.Unlock()
	return

}

func (m *TTLCache) Get(k string) (v string) {
	m.l.Lock()
	if it, ok := m.m[k]; ok {
		v = it.value
		it.lastAcces = time.Now().Unix()
	}
	m.l.Unlock()

	return
}

func (m *TTLCache) Delete(k string) {
	m.l.Lock()

	if _, ok := m.m[k]; ok {
		delete(m.m, k)

	}

	m.l.Unlock()

}
