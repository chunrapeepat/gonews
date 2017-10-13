package model

import (
	"crypto/rand"
	"encoding/base64"
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
	newsStorage = append(newsStorage, news)
}

func ListNews() []*News {
	return newsStorage
}

func GetNews(id string) *News {
	for _, news := range newsStorage {
		if news.ID == id {
			return news
		}
	}
	return nil
}

func DeleteNews(id string) {
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
		}
	}
}
