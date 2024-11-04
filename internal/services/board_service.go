package services

import (
	"log"

	"gmsprt-golang/internal/models"
	"gmsprt-golang/pkg/gorm_scopes"
	"gmsprt-golang/pkg/string_utils"

	"gorm.io/gorm"
)

type BoardService struct {
	logger *log.Logger
	db     *gorm.DB
}

func New(logger *log.Logger, db *gorm.DB) *BoardService {
	boardService := BoardService{
		logger: logger,
		db:     db,
	}
	return &boardService
}

// [GET] /boards
func (boardService *BoardService) findBoards(listBoardSummary *[]BoardSummary, pageable *gorm_scopes.Pageable, query map[string]interface{}) error {

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
		boardSummary := BoardSummary{
			BoardCommon: BoardCommon{
				ID:        info.ID,
				CreatedAt: info.CreatedAt,
				UpdatedAt: info.UpdatedAt,
				DeletedAt: info.DeletedAt.Time,
				Title:     info.Title,
				YnUse:     info.YnUse,
				Name:      info.Name,
			},
			ContentSummary: string_utils.Substr(info.PlainText, 255),
		}
		*listBoardSummary = append(*listBoardSummary, boardSummary)
	}

	return nil
}

// [POST] /boards
func (boardService *BoardService) addBoard(boardDetails *BoardDetails, boardAdd *BoardAdd) error {

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

	boardDetails.BoardCommon.ID = info.ID
	boardDetails.BoardCommon.CreatedAt = info.CreatedAt
	boardDetails.BoardCommon.UpdatedAt = info.UpdatedAt
	boardDetails.BoardCommon.DeletedAt = info.DeletedAt.Time
	boardDetails.BoardCommon.Title = info.Title
	boardDetails.BoardCommon.YnUse = info.YnUse
	boardDetails.BoardCommon.Name = info.Name
	boardDetails.Content = info.Content
	boardDetails.PlainText = info.PlainText

	return nil
}

// [GET] /boards/{ID}
func (boardService *BoardService) findBoard(boardDetails *BoardDetails, ID uint) error {

	var info models.Board

	db := boardService.db

	session := db.Session(&gorm.Session{})

	if err := session.First(&info, ID).Error; err != nil {
		return err
	}

	// TODO struct to struct 속성 복사: struct 임베디드 고려
	boardDetails.BoardCommon.ID = info.ID
	boardDetails.BoardCommon.CreatedAt = info.CreatedAt
	boardDetails.BoardCommon.UpdatedAt = info.UpdatedAt
	boardDetails.BoardCommon.DeletedAt = info.DeletedAt.Time
	boardDetails.BoardCommon.Title = info.Title
	boardDetails.BoardCommon.YnUse = info.YnUse
	boardDetails.BoardCommon.Name = info.Name
	boardDetails.Content = info.Content
	boardDetails.PlainText = info.PlainText

	return nil
}

// [PATCH] /boards/{ID}
func (boardService *BoardService) modifyBoard(boardDetails *BoardDetails, ID uint, boardModify *BoardModify) error {

	var info models.Board

	db := boardService.db

	session := db.Session(&gorm.Session{})

	if err := session.First(&info, ID).Error; err != nil {
		return err
	}

	if err := session.Save(&info).Error; err != nil {
		return err
	}

	// TODO struct to struct 속성 복사: struct 임베디드 고려
	boardDetails.BoardCommon.ID = info.ID
	boardDetails.BoardCommon.CreatedAt = info.CreatedAt
	boardDetails.BoardCommon.UpdatedAt = info.UpdatedAt
	boardDetails.BoardCommon.DeletedAt = info.DeletedAt.Time
	boardDetails.BoardCommon.Title = info.Title
	boardDetails.BoardCommon.YnUse = info.YnUse
	boardDetails.BoardCommon.Name = info.Name
	boardDetails.Content = info.Content
	boardDetails.PlainText = info.PlainText

	return nil
}

// [DELETE] /boards/{ID}
func (boardService *BoardService) deleteBoardsById(ID uint) error {

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
