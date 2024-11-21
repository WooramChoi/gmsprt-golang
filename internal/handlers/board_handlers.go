package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	// TODO generate pageable query
	query := make(map[string]interface{})
	for key, value := range c.Request.URL.Query() {
		switch key {
		case "page":
			pageable.Page, _ = strconv.Atoi(value[0])
		case "page_size":
			pageable.PageSize, _ = strconv.Atoi(value[0])
		case "sort":
			for _, sortVal := range value {
				sort := gorm_scopes.Order{}
				if strings.Contains(sortVal, ",") {
					splitSortVal := strings.Split(sortVal, ",")
					sort.Column = splitSortVal[0]
					sort.IsDESC = (strings.ToLower(splitSortVal[1]) == "desc")
				} else {
					sort.Column = sortVal
					sort.IsDESC = false
				}
				pageable.Sort = append(pageable.Sort, sort)
			}
		default:
			// TODO board table 에 존재하는 컬럼만 걸러내야함
			// TODO + Join Column 에 대한 고민도 필요함
			query[key] = value[0]
		}
	}

	if err := boardHandlers.boardService.FindBoards(&listBoardSummary, &pageable, query); err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"pageable": pageable,
		"content":  listBoardSummary,
	})
}

// [POST] /boards
// curl -X POST -d '{"title": "Test01", "content": "<p>Content 01</p>", "plain_text": "Content 01", "name": "name01", "pwd": "1123"}' http://127.0.0.1:9000/boards
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
}

// [PATCH] /boards/{ID}
// curl -X PATCH -d '{"title": "Test01-1", "content": "<p>Content 01-1</p>", "plain_text": "Content 01-1"}' http://127.0.0.1:9000/boards
func (boardHandlers *BoardHandlers) PatchBoard(c *gin.Context) {

	var boardDetails services.BoardDetails

	paramID := c.Param("ID")
	uint64ID, err := strconv.ParseUint(paramID, 0, 64)
	if err != nil {
		log.Panicln(err.Error())
	}
	id := uint(uint64ID)

	

	var boardModify services.BoardModify
	if err := c.BindJSON(&boardModify); err != nil {
		log.Println(err.Error())
	}

	if err := boardHandlers.boardService.ModifyBoard(&boardDetails, id, &boardModify); err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, boardDetails)

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
