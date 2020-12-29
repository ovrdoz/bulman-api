package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	Description    *string            `json:"description,omitempty"`
	ScenariosTotal *int64             `json:"scenarios_total,omitempty"`
}

type Scenarios struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         *string            `json:"name,omitempty"`
	Project      *string            `json:"project,omitempty"`
	State        *string            `json:"state,omitempty"`
	URL          *string            `json:"url,omitempty"`
	Host         *string            `json:"host,omitempty"`
	Method       *string            `json:"method,omitempty"`
	XMetaPayload *XMeta             `json:"x-meta-payload,omitempty"`
	XMetaHeaders *XMeta             `json:"x-meta-headers,omitempty"`
}

type XMeta struct {
}
