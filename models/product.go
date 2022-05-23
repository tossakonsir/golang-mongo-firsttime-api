package models

type Product struct {
	Itemname    string `json:"itemname" bson:"itemname"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Selltime    string `json:"selltime" bson:"selltime"`
	Qty         int    `json:"qty" bson:"qty"`
	Token       string `json:"token" bson:"token"`
}
