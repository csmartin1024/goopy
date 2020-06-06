package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Goop is
type Goop struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Color     string             `json:"color" bson:"color"`
	Opacity   string             `json:"opacity" bson:"opacity"`
	Speed     int64              `json:"speed" bson:"speed"`
	Texture   string             `json:"texture" bson:"texture"`
	Viscosity string             `json:"viscosity" bson:"viscosity"`
}
