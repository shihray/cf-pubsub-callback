package cf_pubsub_callback

import (
	"context"
	"errors"
	"fmt"
	"github.com/shihray/cf-pubsub-callback/internal/helper"
	"github.com/shihray/cf-pubsub-callback/internal/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"cloud.google.com/go/functions/metadata"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	pbDirect "github.com/lctech-tw/gin-proto/dist/go/star/direct"
	"github.com/lctech-tw/go-pkg/sysvar"
)

var (
	token = "b5g5PgZXUw6drUcoLBNBMgXXY7DxfabE"
	ctx   = context.Background()
	cli   pbDirect.DirectServiceClient
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	address := sysvar.Get("STAR_ADDRESS", "gapi.jkf-star.jkf.io")
	port := sysvar.Get("STAR_PORT", ":18380")

	conn, err := grpc.DialContext(ctx, address+port, grpc.WithInsecure(), grpc.WithAuthority(address))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cli = pbDirect.NewDirectServiceClient(conn)
}

func EntryPoint(ctx context.Context, m PubSubMessage) (err error) {
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			err = errors.New(fmt.Sprintf("recovery: %v", panicErr))
			log.WithError(err).Error("recovery:", panicErr)
			return
		}
	}()

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Event ID: %v", meta.EventID)
	log.Printf("Event type: %v", meta.EventType)

	var payload model.PubSubAction
	if err = helper.Unmarshal(m.Data, &payload); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("EventName: %v", payload.EventName)
	log.Printf("UID: %v", payload.Uid)
	log.Printf("ObjectId: %v", payload.ObjectId)

	switch payload.EventName {
	case "AvatarUpdated":
		_, err = cli.CFCallback(ctx, &pbDirect.CFCallbackReq{
			Token:        token,
			StarObjectId: payload.ObjectId,
			Action:       "avatar_updated",
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	case "AccountDeleted":
		_, err = cli.CFCallback(ctx, &pbDirect.CFCallbackReq{
			Token:        token,
			StarObjectId: payload.ObjectId,
			Action:       "account_deleted",
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	case "DisplayNameUpdated":
		_, err = cli.CFCallback(ctx, &pbDirect.CFCallbackReq{
			Token:        token,
			StarObjectId: payload.ObjectId,
			Action:       "display_name_updated",
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}
