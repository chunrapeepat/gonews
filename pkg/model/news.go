package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News structure type
type News struct {
	ID          bson.ObjectId `bson:"_id"`
	Topic       string
	Image       string
	Description string
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
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

// CreateNews create a news and Insert into database
func CreateNews(news *News) error {
	news.ID = bson.NewObjectId()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil
}

//ListNews list all news from database
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

//GetNews get news from database with id
func GetNews(id string) (*News, error) {
	objectID := bson.ObjectId(id)
	if !objectID.Valid() {
		return nil, fmt.Errorf("Invalid ID")
	}
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

//DeleteNews delete new from database by id
func DeleteNews(id string) error {
	objectID := bson.ObjectId(id)
	if !objectID.Valid() {
		return fmt.Errorf("Invalid ID")
	}
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}
