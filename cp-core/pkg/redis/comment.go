package redis

import (
	"context"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
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
		score = float64(int64(scoreVal.Int()))
	default:
		return 0.0, fmt.Errorf("unsupported type %s", scoreVal.Kind().String())
	}
	return score, nil
}

func AddComment(ctx context.Context, rdb *RedisClient, cc *CommentCache) error {

	err := ZAddComment(ctx, rdb, cc, constants.SortOrderCommentCreateTime)
	if err != nil {
		return err
	}

	err = ZAddComment(ctx, rdb, cc, constants.SortOrderCommentLikeCount)
	if err != nil {
		return err
	}
	return nil
}

func ZAddComment(ctx context.Context, rdb *RedisClient, cc *CommentCache, order string) error {
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

	if existFlag {
		rdb.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: member,
		})
	} else {
		// ...
	}
	return nil
}
