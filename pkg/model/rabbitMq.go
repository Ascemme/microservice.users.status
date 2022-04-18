package model

type Massage struct {
	Uid   int    `json:"uid" bson:"uid"`
	Page  int    `json:"page" bson:"page"`
	Post  int    `json:"post" bson:"post"`
	Value string `json:"value"bson:"value"`
}
