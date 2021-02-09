package CacheHelper

import (
	logHelper "GoMVC/utils/LogHelper"
	"time"

	appconf "github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
)

const defaultTTL = 60

var redisdb *redis.Client

func initClient() (err error) {
	conf, err1 := appconf.NewConfig("ini", "conf/db.conf")
	if err1 != nil {
		return
	}
	redis_Addr := conf.String("client_Addr")
	redisdb = redis.NewClient(&redis.Options{
		Addr:     redis_Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

/*****string******/
func Add(key string, val string) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.Set(key, val, 0).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

// 不存在才插入数据并设置默认过期时间
func AddNX(key string, val string) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.SetNX(key, val, defaultTTL).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

// 设置带有效期的string
func Add_Expire(key string, val string, expTime time.Time) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.Set(key, val, time.Duration(expTime.Second())).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

// 获取值
func Get(key string) string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.Get(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return ""
	}
	return val
}

// 获取过期时间
func GetTTL(key string) time.Duration {
	if redisdb == nil {
		initClient()
	}
	tm, err := redisdb.TTL(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return time.Duration(-1)
	}
	return tm
}

// 自增
func Incr(key string) int64 {
	if redisdb == nil {
		initClient()
	}
	result, err := redisdb.Incr(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return -1
	}
	return result
}

/************hash***************/
//设置hash 对象
func Hash_Set(key string, val map[string]interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.HMSet(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

// 设置hash对象的某个属性
func Hash_SetField(key string, field string, val interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.HSet(key, field, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

func Hash_SetFieldNX(key string, field string, val interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.HSetNX(key, field, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

///获取hash对象
func Hash_Get(key string) interface{} {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.HMGet(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

// 获取hash对象的某个属性值
func Hash_GetField(key string, field string) string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.HGet(key, field).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return ""
	}
	return val
}

/*************list*****************/
func List_LPush(key string, val interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.LPush(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}
func List_RPush(key string, val interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.RPush(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}
func List_Set(key string, index int64, val interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.LSet(key, index, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}
func List_Len(key string) int64 {
	if redisdb == nil {
		initClient()
	}
	len, err := redisdb.LLen(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return -1
	}
	return len
}
func List_BLPop(key string) []string {
	if redisdb == nil {
		initClient()
	}
	arr, err := redisdb.BLPop(1*time.Second, key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return arr
}

/**************set*****************/
//添加
func Set_Add(key string, val []interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.SAdd(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}
func Set_Rem(key string, val []interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.SRem(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

///数量
func Set_Card(key string) int64 {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.SCard(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return -1
	}
	return val
}

//差集
func Set_Diff(key string, key2 string) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.SDiff(key, key2).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

//交集
func Set_Inter(key string, key2 string) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.SInter(key, key2).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

//并集
func Set_Union(key string, key2 string) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.SUnion(key, key2).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

//成员
func Set_Members(key string) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.SMembers(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

/**************sortset*****************/
//添加
func ZSet_Add(key string, val redis.Z) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.ZAdd(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}
func ZSet_Rem(key string, val ...interface{}) bool {
	if redisdb == nil {
		initClient()
	}
	err := redisdb.ZRem(key, val).Err()
	if err != nil {
		logHelper.LogError_e(err)
		return false
	}
	return true
}

///数量
func ZSet_Card(key string) int64 {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.ZCard(key).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return -1
	}
	return val
}

//返回有序集合指定区间内的成员
func ZSet_Range(key string, start int64, stop int64) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.ZRange(key, start, stop).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

//返回有序集合指定区间内的成员分数从高到低
func ZSet_RevRange(key string, start int64, stop int64) []string {
	if redisdb == nil {
		initClient()
	}
	val, err := redisdb.ZRevRange(key, start, stop).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}

//指定分数区间的成员列表
func ZSet_RevRangeByScore(key string, max string, min string, limit int64, offect int64) []string {
	if redisdb == nil {
		initClient()
	}
	opt := redis.ZRangeBy{Min: "(" + min, Max: "(" + max, Offset: offect, Count: limit}
	val, err := redisdb.ZRangeByScore(key, opt).Result()
	if err != nil {
		logHelper.LogError_e(err)
		return nil
	}
	return val
}
