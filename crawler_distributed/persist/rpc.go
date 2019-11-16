package persist

import (
	"github.com/olivere/elastic/v7"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/persist"
	"log"
)

// ItemSaver的RPC service

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		log.Printf("item %v saved.", item)
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}