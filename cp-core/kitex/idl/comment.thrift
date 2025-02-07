struct CompletedComment {
    1: required string ID
    2: required string AuthorID
    3: required string ArticleID
    4: required string ParentID
    5: required string Content
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

service CommentService {
    CreateCommentResponse CreateComment(1: CreateCommentRequest req)
}