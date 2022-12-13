package cf_pubsub_callback

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

var ctx = context.Background()

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func EntryPoint(ctx context.Context, m PubSubMessage) (err error) {
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			log.Println(panicErr)
			err = errors.New(fmt.Sprintf("recovery: %v", panicErr))
			return
		}
	}()

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Event ID: %v\n", meta.EventID)
	log.Printf("Event type: %v\n", meta.EventType)
	// log.Printf("Bucket: %v\n", e.Bucket)
	// log.Printf("File: %v\n", e.Name)
	// log.Printf("Metageneration: %v\n", e.Metageneration)
	// log.Printf("Created: %v\n", e.TimeCreated)
	// log.Printf("Updated: %v\n", e.Updated)

	return
}
