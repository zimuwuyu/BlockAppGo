package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//type IBlockController interface {
//	Get()
//}

type BlockController struct {
}

// Block 示例接口
// @Summary 返回 Block
// @Description 这个接口返回 "Block!"
// @Tags 示例
// @Produce json
// @Success 200 {string} string "Block!"
// @Router /blockModel [get]
func (c *BlockController) Get(ctx *gin.Context) {
	// 你也可以使用一个结构体
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// 注意 msg.Name 在 JSON 中变成了 "user"
	// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
	ctx.JSON(http.StatusOK, msg)
}
