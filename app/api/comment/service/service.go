package service

import (
	"main/app/api/comment/resource"
	"main/database/repository"

	"gorm.io/gorm"
)

// 서비스 인터페이스 선언 (메소드와 반환형 기재)
type CommentService interface {
	FindAllComment() (res []resource.GetAllCommentResponse, err error)
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

// 혈서목록조회 서비스
func (s *commentService) FindAllComment() (res []resource.GetAllCommentResponse, err error) {
	commentRepository := repository.NewRepository()
	res = []resource.GetAllCommentResponse{}

	// 1. 만들어진 레포지토리를 사용해서 데이터를 가져온다
	commentFind, err := commentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// 2. 리턴한다
	for _, commentAll := range commentFind {
		res = append(res, resource.GetAllCommentResponse{
			ID:      int(commentAll.ID),
			Content: commentAll.Content,
		})
	}

	return res, nil
}

// 혈서등록 서비스
func (s *commentService) CreateComment(userId int, req *resource.CreateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()
	// 1. 만들어진 레포지토리를 사용해서 데이터를 입력한다
	// commentCreate := entity.Comment{
	// 	Content: req.Content,
	// }
	// 2. 리턴
	err = commentRepository.Create(uint64(req.Id), req.Content)
	if err != nil {
		return err
	}

	return nil
}

// 혈서수정 서비스
func (s *commentService) UpdateComment(id int, req *resource.UpdateCommentRequest) (err error) {
	commentRepository := repository.NewRepository()

	// 어떤 유저인지 조건이 없음
	err = commentRepository.Update(
		uint64(req.Id),
		req.Content,
	)

	// commentUpdate.ID = id
	// if req.Content != "" {
	// commentT.Content = req.Content
	// }
	// commentRepository.Update()
	if err != nil {
		return nil
	}

	return err
}

// 혈서삭제 서비스
func (s *commentService) DeleteComment(id int) (res *resource.DeleteCommentResponse, err error) {
	commentRepository := repository.NewRepository()
	// 1. 만들어진 레포지토리를 사용해서 데이터를 삭제한다
	commentRepository.Delete(id)

	// 2. 리턴
	if err != nil {
		return nil, err
	}

	return res, nil
}
