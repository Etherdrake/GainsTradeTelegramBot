package concurrentmaps

import (
	"github.com/Etherdrake/telegram-bot-api/v7"
	"sync"
)

type BoardMessagesCache struct {
	mx    sync.Mutex
	cache map[int64]*tgbotapi.Message
}

func New() *BoardMessagesCache {
	return &BoardMessagesCache{
		cache: make(map[int64]*tgbotapi.Message),
	}
}

func (bm *BoardMessagesCache) Get(key int64) (value *tgbotapi.Message, ok bool) {
	bm.mx.Lock()
	defer bm.mx.Unlock()
	value, ok = bm.cache[key]
	return
}

func (bm *BoardMessagesCache) Set(key int64, value *tgbotapi.Message) {
	bm.mx.Lock()
	defer bm.mx.Unlock()
	bm.cache[key] = value
}

func (bm *BoardMessagesCache) Delete(key int64, value *tgbotapi.Message) {
	bm.mx.Lock()
	defer bm.mx.Unlock()
	delete(bm.cache, key)
}
