package main

import "sync"

/**
  @author: CodeWater
  @since: 2023/11/24
  @desc: 手动实现一个线程安全的map
**/

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁,保护下面的map字段
	m            map[int]int
}

// NewRWMap 创建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

// Get 获取key对应的value
func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

// Set 设置key对应的value
func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

// Delete 删除key对应的value
func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

// Len 返回map的长度
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
