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
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         *string            `json:"name,omitempty"`
	ProjectID    primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty"`
	State        *string            `json:"state,omitempty"`
	URL          *string            `json:"url,omitempty"`
	Host         *string            `json:"host,omitempty"`
	Method       *string            `json:"method,omitempty"`
	XMetaPayload *XMeta             `json:"x-meta-payload,omitempty"`
	XMetaHeaders *XMeta             `json:"x-meta-headers,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

type XMeta struct {
}
