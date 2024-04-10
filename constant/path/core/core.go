// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Comment CommentPath
	Dex     DexPath
}

type CommentPath struct {
	GetComment    string
	GetComments   string
	CreateComment string
	UpdateComment string
	DeleteComment string
}

type DexPath struct {
	GetDex    string
	GetDexs   string
	CreateDex string
	UpdateDex string
	DeleteDex string
}
