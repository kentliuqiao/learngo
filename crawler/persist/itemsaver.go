package persist

import (
	"context"
	"errors"
	"log"

	"github.com/kentliuqiao/learngo/crawler/engine"
	"github.com/olivere/elastic/v7"
)

// ItemSaver 持久化
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCnt := 0
		for {
			item := <-out
			log.Printf("item saver got item #%d: %v\n", itemCnt, item)
			itemCnt++

			err := save(client, index, item)
			if err != nil {
				log.Printf("item saver error saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.ID != "" {
		indexService.Id(item.ID)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
