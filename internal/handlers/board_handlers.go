package handlers

import (
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// [GET] /boards
func getBoards(c *gin.Context) {

	// logger := c.MustGet("logger").(*log.Logger)
	// db := c.MustGet("db").(*gorm.DB)

	// session := board.serverDatabase.GetSession()
	// session = session.Scopes(database.Paginate(req))

	// for column, value := range req.URL.Query() {
	// TODO column-value 유효성 확인?
	// 	session = session.Scopes(database.WhereEqual(column, value[0])) // TODO value list 처리
	// }
	// var boardInfos []BoardInfo
	// if err := session.Find(&boardInfos).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// listBoardSummary := []BoardSummary{}
	// for _, info := range boardInfos {
	// TODO struct to struct 속성 복사: struct 임베디드 고려
	// 	boardSummary := BoardSummary{
	// 		BoardCommon: BoardCommon{
	// 			ID:        info.ID,
	// 			CreatedAt: info.CreatedAt,
	// 			UpdatedAt: info.UpdatedAt,
	// 			DeletedAt: info.DeletedAt.Time,
	// 			Title:     info.Title,
	// 			YnUse:     info.YnUse,
	// 			Name:      info.Name,
	// 		},
	// 		ContentSummary: utils.Substr(info.PlainText, 255),
	// 	}
	// 	listBoardSummary = append(listBoardSummary, boardSummary)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(listBoardSummary)
}

// [POST] /boards
func postBoards(c *gin.Context) {

	// if req.Header.Get("Content-Type") != "application/json" {
	// 	http.Error(w, "content-type is not application/json", http.StatusUnsupportedMediaType)
	// 	return
	// }

	// var info BoardInfo
	// session := board.serverDatabase.GetSession()
	// decoder := json.NewDecoder(req.Body)

	// if err := decoder.Decode(&info); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// if err := session.Create(&info).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	// 	return
	// }

	// http.Redirect(w, req, fmt.Sprintf("/boards/%d", info.ID), http.StatusCreated)
}

// [*] /boards/{ID}
func boardsById(c *gin.Context) {
	// switch req.Method {
	// case http.MethodGet:
	// 	board.getBoardsById(w, req)
	// case http.MethodPatch:
	// 	board.patchBoardsById(w, req)
	// case http.MethodDelete:
	// 	board.deleteBoardsById(w, req)
	// default:
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// }
}

// [GET] /boards/{ID}
func getBoardsById(c *gin.Context) {

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
func patchBoardsById(c *gin.Context) {

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
func deleteBoardsById(c *gin.Context) {

	// var info BoardInfo

	// session := board.serverDatabase.GetSession()

	// if err := session.First(&info, req.PathValue("ID")).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// session.Delete(&info)
	// w.WriteHeader(http.StatusNoContent)
}
