package redis

import (
	"context"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/cperror"
	"github.com/redis/go-redis/v9"
	"reflect"
	"strconv"
	"strings"
)

type CommentCache struct {
	CommentID    string
	UserID       string
	ArticleID    string
	ParentID     string
	LikeCount    int
	CommentCount int
	Content      string
	CreateTime   string
}

type RangeOption struct {
	SortOrder string
	SortBy    string
}

func convertOrderValue2Score(scoreVal reflect.Value) (score float64, err error) {
	switch scoreVal.Kind() {
	case reflect.String:
		score, err = strconv.ParseFloat(scoreVal.String(), 64)
		if err != nil {
			return 0.0, err
		}
	case reflect.Float64:
		score = scoreVal.Float()
	case reflect.Int:
		score = float64(scoreVal.Int())
	default:
		return 0.0, fmt.Errorf("unsupported type %s", scoreVal.Kind().String())
	}
	return score, nil
}

// ZAddComment add single comment into an exist zset
func (rdb *RedisClient) ZAddComment(ctx context.Context, cc *CommentCache, order string) error {
	key := fmt.Sprintf(constants.RedisKeyCommentIndex, cc.ArticleID, strings.ToUpper(order))
	existFlag, err := rdb.Expire(ctx, key, 60).Result()
	if err != nil {
		return err
	}

	member := fmt.Sprintf("%s:%s:%s:%s:%d:%d:%s",
		cc.CommentID,
		cc.UserID,
		cc.ArticleID,
		cc.ParentID,
		cc.LikeCount,
		cc.CommentCount,
		cc.CreateTime,
	)

	scoreVal := reflect.ValueOf(cc).FieldByName(order)
	score, err := convertOrderValue2Score(scoreVal)
	if err != nil {
		return err
	}

	if !existFlag {
		return cperror.ErrKeyNotExistsRedis
	} else {
		rdb.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: member,
		})
	}
	return nil
}

// ZAddComments add comments of article into a new zset
//func ZAddComments(ctx context.Context, rdb *RedisClient, ccs []*CommentCache) error {
//
//}

// ZRangeComments get comments of article from zset
func (rdb *RedisClient) ZRangeComments(ctx context.Context, articleID string) ([]*CommentCache, error) {
	return nil, nil
}
