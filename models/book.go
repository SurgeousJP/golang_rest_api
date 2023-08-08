package models

type Book struct{
	Name string `json:"name" bson:"name"`
	ImageURL string `json:"book_img_url" bson:"book_img_url"`
	Author string `json:"author" bson:"author"`
	Price float64 `json:"price" bson:"price"`
	Supplier string `json:"supplier" bson:"supplier"`
	Publisher string `json:"publisher" bson:"publisher"`
	Layout string `json:"book_layout" bson:"book_layout"`
	Series string `json:"series" bson:"series"`
}