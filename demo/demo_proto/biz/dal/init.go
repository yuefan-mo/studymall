package dal

import (
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
