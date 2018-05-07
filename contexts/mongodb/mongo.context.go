package mongocontext

import "github.com/globalsign/mgo"
import . "../../config"

var (
	Session, err = mgo.Dial(Config.Db)
	Db           = Session.DB(Config.DbName)
)

func Disconnect() {
	Session.Close()
}
