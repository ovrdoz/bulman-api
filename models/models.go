package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	Description    *string            `json:"description,omitempty"`
	ScenariosTotal *int64             `json:"scenarios_total,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}

type Scenario struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       *string            `json:"name,omitempty"`
	State      *string            `json:"state,omitempty"`
	URL        *string            `json:"url,omitempty"`
	Method     *string            `json:"method,omitempty"`
	Headers    []Map              `json:"headers,omitempty"`
	Parameters []Map              `json:"parameters,omitempty"`
	Payload    *string            `json:"payload,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	Project    DBRef              `json:"project,omitempty" bson:"project,omitempty"`
}

type Map struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type DBRef struct {
	Ref interface{}        `bson:"$ref"`
	ID  primitive.ObjectID `bson:"$id"`
}

type XMeta struct {
}
