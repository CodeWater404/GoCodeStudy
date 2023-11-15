package redis

/**
  @author: CodeWater
  @since: 2023/11/15
  @desc:  redis key注意使用命名空间的方式,方便查询和拆分,常见的有：，/，_进行分割
**/

const (
	Prefix             = "bluebell:"   //项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  //zset; 帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" //zset; 记录用户及投票类型; 参数是post_id

	KeyCommunitySetPF = "community:" //set; 保存每个分区下帖子的id
)

// getRedisKey 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
