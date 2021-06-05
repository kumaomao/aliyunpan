package global

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type _redisDb struct {
	//保存redis连接
	conn  *redis.Pool
}

var RedisDb = new (_redisDb)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			conn,e = redis.Dial("tcp", addr)
			return conn,e
		},
	}
}

//获取
func (r *_redisDb) Get(name string) string {
	ret, _ := redis.String(r.conn.Get().Do("Get",name))

	return ret
}

//设置
func (r *_redisDb) Set(name string,value string)  {
	_, err := r.conn.Get().Do("Set",name,value)
	if err != nil{
		fmt.Println(err)
	}
}

//设置缓存 带过期时间
func (r *_redisDb) SetEx(name string,value string,time int)  {
	_, err := r.conn.Get().Do("SETEX",name,time,value)
	if err != nil{
		fmt.Println(err)
	}
}

func init() {
	RedisDb.conn = newPool("127.0.0.1:6379")
}


