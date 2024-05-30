// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Comment  CommentPath
	DexEvent DexEventPath
}

type CommentPath struct {
	FindAllComment string
	CreateComment  string
	UpdateComment  string
	DeleteComment  string
}

type DexEventPath struct {
	FindDexEvent   string
	CreateDexEvent string
}
