// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Dex DexPath
}

type DexPath struct {
	GetTags string
}
