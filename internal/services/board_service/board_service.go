package board_service

import (
	"gmsprt-golang/internal/models"
	"gmsprt-golang/internal/repository/gorm_scopes"
	"gmsprt-golang/internal/utils"

	"gorm.io/gorm"
)

type BoardService struct {
	db *gorm.DB
}

func NewBoardService(db *gorm.DB) *BoardService {
	boardService := BoardService{
		db: db,
	}
	return &boardService
}

/*
reflect 를 통한 struct 간에 attribute 복사 방법:
copy := reflect.New(reflect.TypeOf(original)).Elem().Interface().(Car)
일단 사용하지 않도록 한다.
*/
func BoardModelToBoardDetails(info *models.Board, boardDetails *BoardDetails) {
	boardDetails.BoardCommon.ID = info.ID
	boardDetails.BoardCommon.CreatedAt = info.CreatedAt
	boardDetails.BoardCommon.UpdatedAt = info.UpdatedAt
	boardDetails.BoardCommon.DeletedAt = info.DeletedAt.Time
	boardDetails.BoardCommon.Title = info.Title
	boardDetails.BoardCommon.YnUse = info.YnUse
	boardDetails.BoardCommon.Name = info.Name

	boardDetails.Content = info.Content
	boardDetails.PlainText = info.PlainText
}

func BoardModelToBoardSummary(info *models.Board, boardSummary *BoardSummary) {
	boardSummary.BoardCommon.ID = info.ID
	boardSummary.BoardCommon.CreatedAt = info.CreatedAt
	boardSummary.BoardCommon.UpdatedAt = info.UpdatedAt
	boardSummary.BoardCommon.DeletedAt = info.DeletedAt.Time
	boardSummary.BoardCommon.Title = info.Title
	boardSummary.BoardCommon.YnUse = info.YnUse
	boardSummary.BoardCommon.Name = info.Name

	boardSummary.ContentSummary = utils.Substr(info.PlainText, 255)
}

// [GET] /boards
func (boardService *BoardService) FindBoards(listBoardSummary *[]BoardSummary, pageable *gorm_scopes.Pageable, query map[string]interface{}) error {

	// logger := boardService.logger
	db := boardService.db

	session := db.Session(&gorm.Session{})

	// Paging
	if pageable != nil {
		session = session.Scopes(gorm_scopes.Paginate(pageable))
	}

	for column, value := range query {
		// TODO column-value 유효성 확인?
		session = session.Scopes(gorm_scopes.WhereEqual(column, value))
	}

	var boards []models.Board
	if err := session.Find(&boards).Error; err != nil {
		return err
	}
	for _, info := range boards {
		// TODO struct to struct 속성 복사: struct 임베디드 고려
		boardSummary := BoardSummary{}
		BoardModelToBoardSummary(&info, &boardSummary)
		*listBoardSummary = append(*listBoardSummary, boardSummary)
	}

	return nil
}

// [POST] /boards
func (boardService *BoardService) AddBoard(boardDetails *BoardDetails, boardAdd *BoardAdd) error {

	db := boardService.db

	session := db.Session(&gorm.Session{})

	info := models.Board{
		Title:     boardAdd.Title,
		Content:   boardAdd.Content,
		PlainText: boardAdd.PlainText,
		Name:      boardAdd.Name,
		Pwd:       boardAdd.Pwd,
	}
	if err := session.Create(&info).Error; err != nil {
		return err
	}

	BoardModelToBoardDetails(&info, boardDetails)
	return nil
}

// [GET] /boards/{ID}
func (boardService *BoardService) FindBoard(boardDetails *BoardDetails, ID uint) error {

	var info models.Board

	db := boardService.db

	session := db.Session(&gorm.Session{})

	if err := session.First(&info, ID).Error; err != nil {
		return err
	}

	BoardModelToBoardDetails(&info, boardDetails)
	return nil
}

// [PATCH] /boards/{ID}
func (boardService *BoardService) ModifyBoard(boardDetails *BoardDetails, ID uint, boardModify *BoardModify) error {

	var info models.Board

	db := boardService.db

	session := db.Session(&gorm.Session{})

	if err := session.First(&info, ID).Error; err != nil {
		return err
	}

	if err := session.Save(&info).Error; err != nil {
		return err
	}

	BoardModelToBoardDetails(&info, boardDetails)
	return nil
}

// [DELETE] /boards/{ID}
func (boardService *BoardService) DeleteBoardsById(ID uint) error {

	var info models.Board

	db := boardService.db

	session := db.Session(&gorm.Session{})

	if err := session.First(&info, ID).Error; err != nil {
		return err
	}

	if err := session.Delete(&info).Error; err != nil {
		return err
	}

	return nil
}
