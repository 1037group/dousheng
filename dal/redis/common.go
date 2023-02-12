package redis

import (
	"context"
	"fmt"
	"github.com/1037group/dousheng/dal/db"
	"github.com/1037group/dousheng/pkg/configs/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"strings"
	"time"
)

const (
	//涉及到count的模型
	VIDEO    = "video"
	FAVORITE = "favorite"

	//count相关的字段
	VideoFavoriteCount = sql.SQL_VIDEO_VIDEO_FAVORITE_COUNT
	VideoCommentCount  = sql.SQL_VIDEO_VIDEO_COMMENT_COUNT

	//操作类型
	Add = 1  //加
	Sub = -1 //减
)

var (
	//缓存count的key模板，eg：{video:1001}:VideoFavoriteCount
	keyPattern = "{%s:%d}:%s"
	//hash字典key模板，eg：video_1001
	hashKeyPattern = "%s_%d"
	//hash字典域名模板，eg：1001-VideoFavoriteCount
	hashFieldPattern = "%d-%s"
	//redis hash key分片数（避免大key）
	KeySliceSize = 16
)

func GetKeyName(modelName string, modelId int64, fieldName string) string {
	return fmt.Sprintf(keyPattern, modelName, modelId, fieldName)
}

func GetHashKeyName(modelName string, keySlice int) string {
	return fmt.Sprintf(hashKeyPattern, modelName, keySlice)
}

func GetHashFieldName(modelId int64, fieldName string) string {
	return fmt.Sprintf(hashFieldPattern, modelId, fieldName)
}

func updateCachedCount(ctx context.Context, modelName, fieldName string, modelId int64, operatorType int) (int64, error) {
	var key = GetKeyName(modelName, modelId, fieldName)
	var resultCount int64
	var err error
	//1.修改缓存中count的值
	if operatorType == Add {
		resultCount, err = rdb.Incr(ctx, key).Result()
	} else {
		resultCount, err = rdb.Decr(ctx, key).Result()

	}
	if err != nil {
		return resultCount, err
	}
	//2.count修改成功后，修改对应的hash字典，该字典就用来表示发生变化的数据
	//2.1.根据id，找到hash字典的分片key
	var hashKey = GetHashKeyName(modelName, int(modelId)%(KeySliceSize))
	var hashFieldName = GetHashFieldName(modelId, fieldName)
	//eg: hset problems_1 1001-like_count 10
	_, err = rdb.HSet(ctx, hashKey, hashFieldName, resultCount).Result()
	if err != nil {
		//手动回滚(没用redis事务：redis的事务只是将命令打包执行，并不能保证原子性)
		var e1 error
		if operatorType == Add {
			resultCount, e1 = rdb.Decr(ctx, key).Result()
		} else {
			resultCount, e1 = rdb.Incr(ctx, key).Result()
		}
		if e1 != nil {
			return resultCount, err
		}
		return resultCount, err
	}
	return resultCount, err
}

func ScanChangedCountAndUpdateDB(ctx context.Context, modelName string) {
	var hashKey string
	//遍历每个分片
	klog.Infof("Scan redis count cache of [%s] starting...", modelName)
	start := time.Now()
	for i := 0; i < KeySliceSize; i++ {
		hashKey = GetHashKeyName(modelName, i)
		var cursor uint64
		for {
			fields, cur, err := rdb.HScan(ctx, hashKey, cursor, "*-*_count", 10).Result()
			cursor = cur
			if err != nil {
				klog.Errorf("HScan Error：%q", err)
				break
			}
			HandleScannedData(ctx, fields, modelName, hashKey)
			if cursor == 0 {
				break
			}
		}
		//key对应的缓存都scan完毕后，删除key
		size, _ := rdb.HLen(ctx, hashKey).Result()
		if size != 0 {
			_, e := rdb.Del(ctx, hashKey).Result()
			if e != nil {
				klog.Errorf("删除key失败：%s , [error]:%q", hashKey, e)
			}
		}
	}
	cost := time.Now().Sub(start).Seconds()
	klog.Infof("Scan redis count cache of [%s] finished,cost: [%.2f] second", modelName, cost)
}

func HandleScannedData(ctx context.Context, fields []string, modelName string, hashKey string) {
	var id int
	var fieldName string
	var fieldValue int
	var param = make(map[string]interface{})
	for i, key := range fields {
		//所有的k-v会封装为[]string返回：当前index为key index+1就为value
		if strings.Contains(key, "count") {
			split := strings.Split(key, "-")
			id, _ = strconv.Atoi(split[0])
			fieldName = split[1]
			//val, _ := db.HGet(ctx, hashKey, fieldName)
			fieldValue, _ = strconv.Atoi(fields[i+1])
			//封装为Map，更新数据库
			param[fieldName] = fieldValue
			err := updateDB(ctx, modelName, int64(id), param)
			if err != nil {
				klog.Errorf("更新数据库Count失败，表名：%s, ID:%d, 参数：%s , [error]:%q", modelName, id, param, err)
				return
			}
			//如果没有发生异常，删除hash的field
			_, _ = rdb.HDel(ctx, hashKey, fieldName).Result()
		}
	}
}

func updateDB(ctx context.Context, modelName string, id int64, param map[string]interface{}) error {
	var err error
	switch modelName {
	case VIDEO:
		err = db.UpdateFavoriteCount(ctx, db.DB, id, param)
	}
	return err
}

//func CheckRedisExist(ctx context.Context, key string, redisValue interface{}, dbValue int) int {
//	var count int
//	//如果缓存中没有，将数据库中的值set到缓存
//	if redisValue == nil {
//		_, _ = SetStrNxCacheTtl(ctx, key, dbValue, 0)
//		count = dbValue
//	} else {
//		c, _ := strconv.Atoi(redisValue.(string))
//		count = c
//	}
//	return count
//}
