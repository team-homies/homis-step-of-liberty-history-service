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
	UpdateComment(eventId int, req *resource.UpdateCommentRequest) (err error)
	DeleteComment(eventId int, req *resource.DeleteCommentRequest) (res *resource.DeleteCommentResponse, err error)
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
	db *gorm.DB
}

// [혈서 목록 조회] query : EventId
func (s *commentService) FindAllComment(eventId int) (res []resource.FindAllCommentResponse, err error) {
	commentRepository := repository.NewRepository()
	res = []resource.FindAllCommentResponse{}

	// 1. 만들어진 레포지토리를 사용해서 데이터를 가져온다
	commentFind, err := commentRepository.FindAll(eventId)
	if err != nil {
		return nil, err
	}
	// 2. 리턴한다

	for _, commentAll := range commentFind {
		res = append(res, resource.FindAllCommentResponse{
			UserId:  commentAll.UserId,
			Content: commentAll.Content,
		})
	}

	return res, nil
}

// [혈서 등록] query : EventId, body : UserId, Content
func (s *commentService) CreateComment(req *resource.CreateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 1. 만들어진 레포지토리를 사용해서 데이터를 입력한다
	comments := entity.Comment{
		Content: req.Content,
	}

	// 2. 리턴
	err = commentRepository.Create(req.EventId, &comments)
	if err != nil {
		return err
	}

	return nil
}

// [혈서 수정] query : EventId, body : UserId, Content
func (s *commentService) UpdateComment(eventId int, req *resource.UpdateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 어떤 유저인지 조건이 없음
	err = commentRepository.Update(eventId, req)

	// commentUpdate.ID = id
	// if req.Content != "" {
	// commentT.Content = req.Content
	// }
	// commentRepository.Update()
	if err != nil {
		return err
	}

	return
}

// [혈서 삭제] query : EventId, body : UserId
func (s *commentService) DeleteComment(eventId int, req *resource.DeleteCommentRequest) (res *resource.DeleteCommentResponse, err error) {
	commentRepository := repository.NewRepository()
	req = new(resource.DeleteCommentRequest)
	res = new(resource.DeleteCommentResponse)

	// 1. 만들어진 레포지토리를 사용해서 데이터를 삭제한다
	commentDel, err := commentRepository.Delete(eventId, req)
	if err != nil {
		return nil, err
	}

	// 2. 지운 데이터를 res에 반영시킨다
	res.UserId = commentDel.UserId
	res.Content = commentDel.Content
	// 3. 리턴
	return res, nil
}
