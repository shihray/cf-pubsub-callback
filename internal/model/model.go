package model

//{
//	"event_name":"PasswordUpdated",
//	"updated_time":"2022-04-21T08:48:03.864248091Z",
//	"object_id":"79485370-33fb-40c4-acff-1a2fba667ec7",
//	"uid":7000007
//}

type PubSubAction struct {
	EventName   string `json:"event_name"`
	ObjectId    string `json:"object_id"`
	Uid         int64  `json:"uid"`
	UpdatedTime string `json:"updated_time"`
}
