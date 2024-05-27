package service

import (
	"main/app/api/comment/resource"
	"main/database/entity"
	"main/database/repository"

	"gorm.io/gorm"
)

// 서비스 인터페이스 선언 (메소드와 반환형 기재)
type CommentService interface {
	FindAllComment(eventId int) (res []resource.FindAllCommentResponse, err error)
	CreateComment(req *resource.CreateCommentRequest) (err error)
	UpdateComment(req *resource.UpdateCommentRequest) (err error)
	DeleteComment(req *resource.DeleteCommentRequest) (err error)
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
	CommentService
}

// [혈서 목록 조회] Param : Id (eventId값)
func (s *commentService) FindAllComment(eventId int) (res []resource.FindAllCommentResponse, err error) {
	commentRepository := repository.NewRepository()
	res = []resource.FindAllCommentResponse{}

	// 1. 만들어진 레포지토리를 사용해서 데이터를 가져온다
	commentFind, err := commentRepository.FindAll(eventId)
	if err != nil {
		return nil, err
	}

	for _, commentAll := range commentFind {
		res = append(res, resource.FindAllCommentResponse{
			UserId:  commentAll.UserId,
			Content: commentAll.Content,
		})
	}

	// 2. 리턴한다
	return
}

// [혈서 등록] Param : Id, body : UserId, Content
func (s *commentService) CreateComment(req *resource.CreateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 1. 받아온 parameter들을 변수에 저장한다
	commentCreate := entity.Comment{
		Model: gorm.Model{
			ID: uint(req.Id),
		},
		UserId:  req.UserId,
		Content: req.Content,
	}

	// 2. 만들어진 레포지토리를 사용해서 데이터를 입력한다
	err = commentRepository.Create(&commentCreate)
	if err != nil {
		return err
	}

	// 3. 리턴
	return
}

// [혈서 수정] Param : Id, body : Id, Content
func (s *commentService) UpdateComment(req *resource.UpdateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 1. 받아온 parameter들을 변수에 저장한다
	var comment entity.Comment
	comment.Model.ID = uint(req.Id)
	if req.Id != 0 {
		comment.Model.ID = uint(req.Id)
	}
	if req.Content != "" {
		comment.Content = req.Content
	}

	// 2. 만들어진 레포지토리를 사용해서 데이터를 수정
	err = commentRepository.Update(&comment)
	if err != nil {
		return
	}

	// 3. 리턴
	return
}

// [혈서 삭제] Param : Id, body : Id
func (s *commentService) DeleteComment(req *resource.DeleteCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 1. 받아온 parameter들을 변수에 저장한다
	commentDelete := entity.Comment{
		Model: gorm.Model{
			ID: uint(req.Id),
		},
	}

	// 2. 만들어진 레포지토리를 사용해서 데이터를 삭제한다
	err = commentRepository.Delete(&commentDelete)
	if err != nil {
		return err
	}

	// 3. 리턴
	return
}
