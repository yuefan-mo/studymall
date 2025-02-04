package dal

import (
	"github.com/yuefan-mo/studymall/demo/demo_thrift/biz/dal/mysql"
	"github.com/yuefan-mo/studymall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
