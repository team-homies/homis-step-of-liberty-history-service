package service

import "main/app/api/comment/resource"

// 서비스 인터페이스
type CommentService interface {
	GetComment(req *resource.GetCommentRequest) (res *resource.GetCommentResponse, err error)
	GetComments() (resource.GetCommentsResponse, error)
	CreateComment(req *resource.CreateCommentRequest) error
	UpdateComment(req *resource.UpdateCommentRequest) (res *resource.UpdateCommentResponse, err error)
	DeleteComment(id int) error
}

// comment 서비스 함수
func NewCommentService() CommentService {
	return &commentService{
		// 위의 서비스 인터페이스 Comment서비스
		CommentService: &commentService{},
	}
}

type commentService struct {
	CommentService
}

func (s *commentService) GetComment(req *resource.GetCommentRequest) (res *resource.GetCommentResponse, err error) {
	res = new(resource.GetCommentResponse)
	res = &resource.GetCommentResponse{
		Comment: resource.CommentResource{
			Id:       1,
			UserId:   1,
			UserName: "논현동",
			Content:  "time to go to bed",
		},
	}

	return
}
