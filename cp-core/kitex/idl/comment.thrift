include "base.thrift"

struct CompletedComment {
	1: required string ID
	2: required string AuthorID
	3: required string ArticleID
	4: required string ParentID
	5: required string Content
	6: required i64    LikeCount
	7: required i64    CommentCount
	8: required string CreateTime
	9: required string UpdateTime
}

struct CreateCommentRequest {
    1: required string AuthorID
    2: required string ArticleID
    3: optional string ParentID
    4: required string Content
}

struct CreateCommentResponse {
    1: required string ID
}

struct GetCommentRequest {
    1: required string ID
}

struct UpdateCommonRequest {
    1: required string ID
    2: required string Content
}

struct DeleteCommonRequest {
    1: required string ID
}

service CommentService {
    CreateCommentResponse CreateComment(1: CreateCommentRequest req)
    CompletedComment GetComment(1: GetCommentRequest req)
    base.EmptyBody UpdateComment(1: UpdateCommonRequest req)
    base.EmptyBody DeleteComment(1: DeleteCommonRequest req)
}