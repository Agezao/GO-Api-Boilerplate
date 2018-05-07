package productModel

type Product struct {
	ID   uint   `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
