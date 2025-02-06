struct CreateCommentRequest {
    1: required string AuthorID
    2: required string ArticleID
    3: optional string ParentID
    4: required string Content
}

struct CreateCommentResponse {
    1: required string ID
}

service CommonService {
    CreateCommentResponse CreateComment(1: CreateCommentRequest req)
}