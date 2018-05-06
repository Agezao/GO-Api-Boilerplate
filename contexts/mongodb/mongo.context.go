package mongocontext

import "github.com/globalsign/mgo"
import . "../../config"

var Session mgo.Session

func init() {
	Session, err := mgo.Dial(Config.Db)
	if nil != err {
		panic(err)
	}
	defer Session.Close()

	Session.SetMode(mgo.Monotonic, true)
}
