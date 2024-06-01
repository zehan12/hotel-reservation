package types

type User struct {
	Id        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
