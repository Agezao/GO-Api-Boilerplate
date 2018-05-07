package productsservice

import "github.com/globalsign/mgo/bson"
import . "../../models/product"
import ctx "../../contexts/mongodb"

var coll = ctx.Db.C("products")

func Get(query bson.M) []Product {
	var result []Product
	coll.Find(query).All(&result)
	return result
}
