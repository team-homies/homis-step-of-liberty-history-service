package service

import (
	"main/app/api/comment/resource"
	"main/database/repository/comment"
)

// 서비스 인터페이스 선언 (메소드와 반환형 기재)
type CommentService interface {
	GetAllComment() (res *[]resource.GetAllCommentResponse, err error)
	CreateComment(req *resource.CreateCommentRequest) (err error)
	UpdateComment(req *resource.UpdateCommentRequest) (err error)
	DeleteComment(req *resource.DeleteCommentRequest) (res *resource.DeleteCommentResponse, err error)
}

// 함수를 CommentService형으로 선언
func NewCommentService() CommentService {
	// commentService의 주소값 반환
	return &commentService{
		CommentService: &commentService{},
	}
}

// commentService형 선언
type commentService struct {
	// 필요한 리소스 선언
	CommentService
}

func (s *commentService) GetAllComment() (res *[]resource.GetAllCommentResponse, err error) {
	// res = new(resource.GetAllCommentResponse)
	// res = &resource.GetAllCommentResponse{
	// 	Comment: []resource.CommentResource{
	// 		{
	// 			Id:       1,
	// 			UserId:   1,
	// 			UserName: "논현동",
	// 			Content:  "time to go to bed",
	// 		},
	// 		{
	// 			Id:       2,
	// 			UserId:   1,
	// 			UserName: "왱구",
	// 			Content:  "time to go to eat",
	// 		},
	// 		{
	// 			Id:       3,
	// 			UserId:   2,
	// 			UserName: "우림",
	// 			Content:  "time to go to play",
	// 		},
	// 	},
	// }
	//raw 데이터 가져오기
	comments, err := comment.CommentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		res = append(res, resource.GetAllCommentResponse{
			Comment: []resource.CommentResource{
				{
					Id:       comment.Id,
					UserId:   comment.UserId,
					UserName: comment.UserName,
					Content:  comment.Content,
				},
			},
		})
	}

	return res, nil
}
func (s *commentService) CreateComment(req *resource.CreateCommentRequest) (err error) {

	return
}
func (s *commentService) UpdateComment(req *resource.UpdateCommentRequest) (err error) {

	return
}
func (s *commentService) DeleteComment(req *resource.DeleteCommentRequest) (res *resource.DeleteCommentResponse, err error) {
	res = new(resource.DeleteCommentResponse)
	res = &resource.DeleteCommentResponse{
		Comment: resource.CommentResource{
			Id:       1,
			UserId:   1,
			UserName: "논현동",
			Content:  "time to go to bed",
		},
	}

	return
}
