typedef string ArticleStatus
const ArticleStatus ArticleStatusDraft     = "draft"
const ArticleStatus ArticleStatusPublished = "published"
const ArticleStatus ArticleStatusPrivate   = "private"


struct CreateArticleRequest {
	1: required string        Author
	2: required string        Title
	3: required string        Content
	4: optional ArticleStatus Status   (go.tag = "default:published")
}



struct CreateArticleResponse {
	1: required string        ID
	2: required ArticleStatus Status
}

service ArticleService {
    CreateArticleResponse CreateArticle(1: CreateArticleRequest req)
}