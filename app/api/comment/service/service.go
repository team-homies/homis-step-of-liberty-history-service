package service

import (
	"main/app/api/comment/resource"
	"main/database/entity"
	"main/database/repository/comment"

	"gorm.io/gorm"
)

// 서비스 인터페이스 선언 (메소드와 반환형 기재)
type CommentService interface {
	GetAllComment() (res *[]resource.GetAllCommentResponse, err error)
	CreateComment(id int, req *resource.CreateCommentRequest) (err error)
	UpdateComment(id int, req *resource.UpdateCommentRequest) (err error)
	DeleteComment(id int) (res *resource.DeleteCommentResponse, err error)
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
	db *gorm.DB
}

func (s *commentService) GetAllComment() (res *[]resource.GetAllCommentResponse, err error) {
	//raw 데이터 가져오기
	comments, err := comment.NewCommentRepository(s.db).GetAll()
	if err != nil {
		return nil, err
	}
	var commentRes []resource.GetAllCommentResponse

	for _, comment := range comments {
		commentRes = append(commentRes, resource.GetAllCommentResponse{
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
func (s *commentService) CreateComment(id int, req []*resource.CreateCommentRequest) (err error) {
	for _, newComment := range req {
		commentRes := resource.CreateCommentRequest{
			Id:      newComment.Id,
			Content: newComment.Content,
		}
		err = comment.NewCommentRepository(s.db).Create()
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *commentService) UpdateComment(id int, req *resource.UpdateCommentRequest) (err error) {
	var commentT entity.Comment
	commentT.ID = id
	if req.Content != "" {
		commentT.Content = req.Content
	}
	err = comment.NewCommentRepository(s.db).Update(&req)
	if err != nil {
		return nil
	}

	return err
}
func (s *commentService) DeleteComment(id int) (res *resource.DeleteCommentResponse, err error) {
	// res = new(resource.DeleteCommentResponse)
	// res = &resource.DeleteCommentResponse{
	// 	Comment: resource.CommentResource{
	// 		Id:       1,
	// 		UserId:   1,
	// 		UserName: "논현동",
	// 		Content:  "time to go to bed",
	// 	},
	// }
	res, err = comment.NewCommentRepository(s.db).Delete()
	if err != nil {
		return nil, err
	}

	return res, nil
}
