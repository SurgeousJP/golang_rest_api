package models

type Book struct{
	Name string `json:"name" bson:"book_name"`
	ImageURL string `json:"book_img_url" bson:"book_img_url"`
	Author string `json:"author" bson:"book_author"`
	Price float64 `json:"price" bson:"book_price"`
	Supplier string `json:"supplier" bson:"book_supplier"`
	Publisher string `json:"publisher" bson:"book_publisher"`
	Layout string `json:"book_layout" bson:"book_layout"`
	Series string `json:"series" bson:"book_series"`
}