package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFromKey(key string, size, page int64) ([]string, error) {
	//2.确定查询索引起点
	start := (page - 1) * size
	end := start + size - 1
	// 3.ZREVRANGE 查询  返回的是post_id的结果集由大到小
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取ID
	//1.根据用户请求中的order参数获取redis中的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Size, p.Page)
}

//func GetPostVoteData(pids []string) (data []int64, err error) {
//	//data = make([]int64, 0, len(ids))
//	//for _, id := range ids {
//	//	key := getRedisKey(KeyPostVotedZSetPF + id)
//	//	//查找key中分数是1元素的数量-》统计每篇帖子的赞成票数量
//	//	v := client.ZCount(key, "1", "1").Val()
//	//	data = append(data, v)
//	//}
//	//使用pipeline一次发送多条命令，减少RTT
//	data = make([]int64, 0, len(pids))
//	pipeline := client.Pipeline()
//	for _, id := range pids {
//		key := getRedisKey(KeyPostVotedZSetPF + id)
//		pipeline.ZCount(key, "1", "1")
//	}
//	cmders, err := pipeline.Exec()
//	if err != nil {
//		zap.L().Error("pipeline.Exec() fail ", zap.Error(err))
//		return
//	}
//	data = make([]int64, 0, len(pids))
//	for _, cmder := range cmders {
//		v := cmder.(*redis.IntCmd).Val()
//		data = append(data, v)
//	}
//	return data, nil
//}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	data = make([]int64, 0, len(ids))
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
		v := client.ZCount(key, "1", "1").Val()
		data = append(data, v)
	}
	return

	//这个写法有问题
	// 使用pipeline一次发送多条命令,减少RTT
	//pipeline := client.Pipeline()
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	pipeline.ZCount(key, "1", "1")
	//}
	//cmders, err := pipeline.Exec()
	//if err != nil {
	//	return
	//}
	//data = make([]int64, 0, len(cmders))
	//for _, cmder := range cmders {
	//	v := cmder.(*redis.IntCmd).Val()
	//	data = append(data, v)
	//}
	//return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//使用interstore把分区的帖子set与帖子的zset生成一个新的zset

	//针对新的zset按之前的逻辑取数据

	//社区的key
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	orderKey := p.Order
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//利用缓存key减zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		//不存在需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return []string{}, err
		}
	}
	//存在的话根据key查询ids
	//2.确定查询索引起点
	return getIDsFromKey(key, p.Size, p.Page)
}
