package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

// CtxUserIDKey 是用于在上下文中存储用户ID的键
const CtxUserIDKey = "userID"

// ErrorUsrNotLogin 是用户未登录时返回的错误
var ErrorUsrNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录的用户ID
// 参数:
//   - c: gin的上下文
// 返回值:
//   - userID: 用户ID
//   - err: 可能的错误，如用户未登录
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUsrNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUsrNotLogin
		return
	}
	return
}

// getPageInfo 从请求中获取分页信息
// 参数:
//   - c: gin的上下文
// 返回值:
//   - page: 页码，默认为1
//   - size: 每页大小，默认为10
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 0
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
