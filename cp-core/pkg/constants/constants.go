package constants

const (
	SIDPrefixArticle = "art"
	SIDPrefixComment = "cmt"
)

const (
	SortOrderAsc               = "Asc"
	SortOrderDesc              = "Desc"
	SortOrderCommentCreateTime = "CreateTime"
	SortOrderCommentLikeCount  = "LikeCount"
)

const (
	RedisKeyCommentIndex = "CP:COMMENT:%s:%s" // "cp:service_name:article_id:order_index"
)
