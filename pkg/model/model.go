package model

type User struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	Uid       int    `json:"uid" bson:"uid"`
	Following int    `json:"following" bson:"following"`
	Page      []Page `json:"page" bson:"page"`
}
type Page struct {
	Id          int    `json:"id" bson:"id"`
	Subscribers int    `json:"subscribers" bson:"subscribers"`
	Post        []Post `json:"post" bson:"post"`
}
type Post struct {
	Id       int `json:"id" bson:"id"`
	Comments int `json:"comment" bson:"comment"`
	Likes    int `json:"likes" bson:"likes"`
	Dislike  int `json:"dislike" bson:"dislike"`
}
