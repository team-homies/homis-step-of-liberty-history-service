package service

import (
	"main/app/api/comment/resource"
	"main/database/entity"
	"main/database/repository/comment"
)

// 서비스 인터페이스 선언 (메소드와 반환형 기재)
type CommentService interface {
	GetAllComment() (res *[]resource.GetAllCommentResponse, err error)
	CreateComment(id uint, req *resource.CreateCommentRequest) (err error)
	UpdateComment(id uint, req *resource.UpdateCommentRequest) (err error)
	DeleteComment(id uint) (res *resource.DeleteCommentResponse, err error)
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
func (s *commentService) CreateComment(id uint, req *resource.CreateCommentRequest) (err error) {
	for _, newComment := range req {
		comment := resource.CreateCommentRequest{
			Id:      newComment.Id,
			Content: newComment.Content,
		}
		err = comment.comment.Create(req)
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *commentService) UpdateComment(id uint, req *resource.UpdateCommentRequest) (err error) {
	var comment entity.Comment
	comment.Id = id
	if req.Content != "" {
		comment.Content = req.Content
	}
	err = comment.comment.Update(comment)
	if err != nil {
		return nil
	}

	return err
}
func (s *commentService) DeleteComment(id uint) (res *resource.DeleteCommentResponse, err error) {
	// res = new(resource.DeleteCommentResponse)
	// res = &resource.DeleteCommentResponse{
	// 	Comment: resource.CommentResource{
	// 		Id:       1,
	// 		UserId:   1,
	// 		UserName: "논현동",
	// 		Content:  "time to go to bed",
	// 	},
	// }
	res, err = comment.comment.Delete(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
