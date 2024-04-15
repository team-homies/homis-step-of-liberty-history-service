// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Comment CommentPath
	Dex     DexPath
}

type CommentPath struct {
	GetAllComment string
	CreateComment string
	UpdateComment string
	DeleteComment string
}

type DexPath struct {
	GetDex    string
	CreateDex string
}
