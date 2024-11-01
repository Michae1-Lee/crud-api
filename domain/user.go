package domain

type User struct {
	Id       int    `json:"id" bson:"id"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}
