include "base.thrift"

typedef string ArticleStatus
const ArticleStatus ArticleStatusDraft     = "draft"
const ArticleStatus ArticleStatusPublished = "published"
const ArticleStatus ArticleStatusPrivate   = "private"

struct Article {
	1:  required string        ID
	2:  required string        AuthorID
	3:  required string        Title
	4:  required string        Content
	5:  required ArticleStatus Status
	6:  required i64           ViewCount
	7:  required i64           LikeCount
	8:  required i64           CommentCount
	9:  required string        CreateTime
	10: required string        UpdateTime
}

struct CreateArticleRequest {
	1: required string        AuthorID
	2: required string        Title
	3: required string        Content
	4: required ArticleStatus Status    (go.tag = "default:published")
}

struct CreateArticleResponse {
	1: required string        ID
	2: required ArticleStatus Status
}

struct GetArticleRequest {
	1: required string ID
}

struct UpdateArticleRequest {
	1: required string        ID
	2: optional string        Title
	3: optional string        Content
	4: optional ArticleStatus Status
}

struct ListArticlesRequest {
	1: optional i32                 PageNumber  (go.tag = "default:1")
	2: optional i32                 PageSize    (go.tag = "default:10")
	3: optional string              SortOrder   (go.tag = "default:desc")
	4: optional string              SortBy      (go.tag = "default:create_time")
	5: optional list<ArticleStatus> Statuses
	6: optional list<string>        AuthorIDs
}

struct ListArticlesResponse {
	1: required i64           TotalCount
	2: required i32           PageNumber
	3: required i32           PageSize
	4: required list<Article> Articles
}

struct DeleteArticleRequest {
	1: required string ID
}

service ArticleService {
    CreateArticleResponse CreateArticle(1: CreateArticleRequest req)
    Article GetArticle(1: GetArticleRequest req)
    ListArticlesResponse ListArticles(1: ListArticlesRequest req)
    base.EmptyBody UpdateArticle(1: UpdateArticleRequest req)
    base.EmptyBody DeleteArticle(1: DeleteArticleRequest req)
}