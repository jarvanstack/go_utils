package httpu

import (
	"sync"
	"time"

	"github.com/golang/groupcache/lru"
)

type LRUTokenBucket struct {
	lru      *lru.Cache
	rate     int64
	capacity int64
}

func NewLRUTokenBucket(maxEntity, rate, capacity int64) *LRUTokenBucket {
	return &LRUTokenBucket{
		lru:      lru.New(int(maxEntity)),
		rate:     rate,
		capacity: capacity,
	}
}

func (l *LRUTokenBucket) Allow(ip string) bool {
	v, ok := l.lru.Get(ip)
	if !ok {
		//没有 cache
		l.lru.Add(ip, NewTokenBucket(l.rate, l.capacity))
		return true
	}
	//cache
	tb := v.(*TokenBucket)
	return tb.Allow()
}

//令牌桶限流算法
type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //桶上次放token的时间戳 s
	lock         sync.Mutex
}

func NewTokenBucket(rate, capacity int64) *TokenBucket {
	return &TokenBucket{
		rate:         rate,
		capacity:     capacity,
		tokens:       0,
		lastTokenSec: 0,
		lock:         sync.Mutex{},
	}
}
func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}
