package handler

import (
	"demo4/internal/logic"
	"demo4/internal/model"
	error2 "demo4/pkg/error"
	"demo4/pkg/response"

	"github.com/gin-gonic/gin"
)

func UserPage(c *gin.Context) {
	params := model.UserPageReq{}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	result, err := logic.UserLogic.Page(&params)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "分页查询用户失败")
		return
	}
	response.Success(c, result, "查询成功")
}

// 注册
func Register(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	if err := logic.UserLogic.Register(&user); err != nil {
		response.Fail(c, error2.ErrSystem.Code, "注册失败")
		return
	}
	response.Success(c, nil, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	req := model.UserLoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, error2.ErrInvalidParams)
		return
	}
	logic.UserLogic.Login(&req)
	resp, err := logic.UserLogic.Login(&req)
	if err != nil {
		response.Fail(c, error2.ErrSystem.Code, "登录失败")
		return
	}
	response.Success(c, resp, "登录成功")

}
