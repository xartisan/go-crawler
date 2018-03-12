package persist

import (
	"log"

	"context"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"github.com/xartisan/go-crawler/engine"
)

func SaveWorker(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return out, nil
	}
	go func() {
		log.Println("fdasfads;lfjdsa;klfjdas;lkfkjdas;lkfjdas;lkfdajsl;k")
		itemCount := 0
		for {
			item := <-out
			log.Printf("Persist Worker: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Save worker error saving err: %v", err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}
	indexService := client.Index().Index(index).
		Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("%+v\n", resp)
	return nil
}
