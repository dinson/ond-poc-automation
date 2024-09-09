package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payload struct {
	ID              *primitive.ObjectID `bson::"_id,omitempty"`
	NodeID          string              `bson:"nodeID"`
	InputPayloadIDs []string            `bson:"inputPayloadIDs"`
	Output          string              `bson:"output"`
	CreatedAt       int64               `bson:"createdAt"`
}
