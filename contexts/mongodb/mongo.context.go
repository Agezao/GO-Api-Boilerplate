package mongocontext

import "github.com/globalsign/mgo"
import . "../../config"

var Session mgo.Session
var Db mgo.Database

func init() {
	Connect()
}

func Disconnect() {
	Session.Close()
}

func Connect() {
	Session, err := mgo.Dial(Config.Db)
	if nil != err {
		panic(err)
	}
	defer Session.Close()

	//Session.SetMode(mgo.Monotonic, true)
}
