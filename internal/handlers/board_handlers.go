package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gmsprt-golang/internal/services"
	"gmsprt-golang/pkg/gorm_scopes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BoardHandlers struct {
	boardService *services.BoardService
}

func NewBoardHandlers(db *gorm.DB) *BoardHandlers {
	boardService := services.NewBoardService(db)
	boardHandlers := BoardHandlers{
		boardService: boardService,
	}
	return &boardHandlers
}

// [GET] /boards
func (boardHandlers *BoardHandlers) GetBoards(c *gin.Context) {

	listBoardSummary := []services.BoardSummary{}
	pageable := gorm_scopes.Pageable{
		Page:     1,
		PageSize: 10,
	}

	query := make(map[string]interface{})
	for key, value := range c.Request.URL.Query() {
		switch key {
		case "page":
			pageable.Page, _ = strconv.Atoi(value[0])
		case "page_size":
			pageable.PageSize, _ = strconv.Atoi(value[0])
		default:
			// TODO board table 에 존재하는 컬럼만 걸러내야함
			query[key] = value[0]
		}
	}

	if err := boardHandlers.boardService.FindBoards(&listBoardSummary, &pageable, query); err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"pageable": pageable,
		"data":     listBoardSummary,
	})
}

// [POST] /boards
func (boardHandlers *BoardHandlers) PostBoard(c *gin.Context) {

	var boardDetails services.BoardDetails

	var boardAdd services.BoardAdd
	if err := c.BindJSON(&boardAdd); err != nil {
		log.Println(err.Error())
	}

	if err := boardHandlers.boardService.AddBoard(&boardDetails, &boardAdd); err != nil {
		log.Println(err.Error())
	}

	c.Redirect(http.StatusCreated, fmt.Sprintf("/boards/%d", boardDetails.ID))
}

// [GET] /boards/{ID}
func (boardHandlers *BoardHandlers) GetBoard(c *gin.Context) {

	var boardDetails services.BoardDetails

	paramID := c.Param("ID")
	uint64ID, err := strconv.ParseUint(paramID, 0, 64)
	if err != nil {
		log.Panicln(err.Error())
	}
	id := uint(uint64ID)

	if err := boardHandlers.boardService.FindBoard(&boardDetails, id); err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, boardDetails)

	// var info BoardInfo

	// session := board.serverDatabase.GetSession()

	// if err := session.First(&info, req.PathValue("ID")).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// TODO struct to struct 속성 복사: struct 임베디드 고려
	// boardDetails := BoardDetails{
	// 	BoardCommon: BoardCommon{
	// 		ID:        info.ID,
	// 		CreatedAt: info.CreatedAt,
	// 		UpdatedAt: info.UpdatedAt,
	// 		DeletedAt: info.DeletedAt.Time,
	// 		Title:     info.Title,
	// 		YnUse:     info.YnUse,
	// 		Name:      info.Name,
	// 	},
	// 	Content:   info.Content,
	// 	PlainText: info.PlainText,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(boardDetails)
}

// [PATCH] /boards/{ID}
func (boardHandlers *BoardHandlers) PatchBoard(c *gin.Context) {

	// if req.Header.Get("Content-Type") != "application/json" {
	// 	http.Error(w, "content-type is not application/json", http.StatusUnsupportedMediaType)
	// 	return
	// }

	// var info BoardInfo

	// session := board.serverDatabase.GetSession()

	// if err := session.First(&info, req.PathValue("ID")).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// decoder := json.NewDecoder(req.Body)

	// if err := decoder.Decode(&info); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// session.Save(&info)
	// http.Redirect(w, req, fmt.Sprintf("/boards/%d", info.ID), http.StatusNoContent)
}

// [DELETE] /boards/{ID}
func (boardHandlers *BoardHandlers) DeleteBoard(c *gin.Context) {

	// var info BoardInfo

	// session := board.serverDatabase.GetSession()

	// if err := session.First(&info, req.PathValue("ID")).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// session.Delete(&info)
	// w.WriteHeader(http.StatusNoContent)
}
