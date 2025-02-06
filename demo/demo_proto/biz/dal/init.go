package dal

import (
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal/mysql"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
