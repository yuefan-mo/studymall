package dal

import (
	"github.com/yuefan-mo/studymall/biz/dal/mysql"
	"github.com/yuefan-mo/studymall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
