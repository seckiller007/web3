package handler

import (
	"demo4/internal/logic"
	"demo4/internal/model"
	"demo4/pkg/db"
	error2 "demo4/pkg/error"
	"demo4/pkg/response"
	"errors"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreatePost(c *gin.Context) {
	params := model.Post{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	userId := c.MustGet("userID").(uint)
	params.UserID = userId
	err := logic.PostLogic.CreatePost(&params)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "创建文章失败")
		return
	}
	response.Success(c, nil, "创建文章成功")
}

// 分页查询文章
func PostPage(c *gin.Context) {
	params := db.QueryParams{}                         //page?page=1&pageSize=5
	if err := c.ShouldBindQuery(&params); err != nil { //ShouldBindQuery 只能解析 URL 中的查询参数  page?page=1&pageSize=5
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	currentUserId := c.MustGet("userID").(uint)
	result, err := logic.PostLogic.PostPage(&params, currentUserId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "分页查询文章失败")
		return
	}
	response.Success(c, result, "查询文章成功")
}

// 查询文章详情
func PostById(c *gin.Context) {
	postId, exist := c.GetQuery("postId")
	if !exist {
		response.Error(c, error2.ErrInvalidParams)
	}
	currentUserId := c.MustGet("userID").(uint)
	post, err := logic.PostLogic.PostById(postId, currentUserId)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "查询文章失败")
		return
	}
	response.Success(c, post, "查询文章成功")
}

// 修改文章
func EditPost(c *gin.Context) {
	params := model.Post{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	currentUserId := c.MustGet("userID").(uint)
	err := logic.PostLogic.EditPost(&params, &currentUserId)
	if err != nil {
		if errors.Is(err, error2.ErrUnauthorized) {
			response.Error(c, error2.ErrUnauthorized)
		} else {
			response.Fail(c, error2.ErrSystem.Code, "更新文章失败")
		}
		return
	}
	response.Success(c, nil, "更新文章成功")
}

// 删除文章
func DelPost(c *gin.Context) {
	postId, exist := c.GetQuery("postId")
	if !exist {
		response.Error(c, error2.ErrInvalidParams)
	}
	currentUserId := c.MustGet("userID").(uint)
	err := logic.PostLogic.DelPost(&postId, &currentUserId)
	if err != nil {
		if errors.Is(err, error2.ErrUnauthorized) {
			response.Error(c, error2.ErrUnauthorized)
		} else {
			response.Fail(c, error2.ErrSystem.Code, "删除文章失败")
		}
		return
	}
	response.Success(c, nil, "删除文章成功")
}
