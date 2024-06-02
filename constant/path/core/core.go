// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Comment  CommentPath
	Dex DexPath
}

type CommentPath struct {
	FindAllComment string
	CreateComment  string
	UpdateComment  string
	DeleteComment  string
}

type DexPath struct {
	FindDexEvent   string
	CreateDexEvent string
	GetTags string
}
