package controller

import (
	"BlockApp/db"
	"BlockApp/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type IBlockController interface {
//	Get()
//}

type BlockController struct {
}

// Block 接口
// @Summary 分页返回 BlockModel 数据
// @Description 这个接口返回分页的 BlockModel 列表
// @Tags 示例
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} []model.BlockModel
// @Router /blockModel [get]
func (c *BlockController) Get(ctx *gin.Context) {
	var blocks []model.BlockModel
	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	// 校验分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询数据库
	result := db.PgsqlDB.Limit(pageSize).Offset(offset).Find(&blocks)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blocks)

}
