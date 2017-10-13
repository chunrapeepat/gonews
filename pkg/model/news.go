package model

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// News structure type
type News struct {
	ID          string
	Topic       string
	Image       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	newsStorage []*News
	mutexNews   sync.Mutex
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

func CreateNews(news *News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	mutexNews.Lock()
	defer mutexNews.Unlock()
	newsStorage = append(newsStorage, news)
}

func ListNews() []*News {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	r := make([]*News, len(newsStorage))
	for i := range newsStorage {
		r[i] = newsStorage[i]
	}
	return r
}

func GetNews(id string) *News {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	for _, news := range newsStorage {
		if news.ID == id {
			return news
		}
	}
	return nil
}

func DeleteNews(id string) {
	mutexNews.Lock()
	defer mutexNews.Unlock()
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
		}
	}
}
