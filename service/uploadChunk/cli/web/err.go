package web

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

// 所有的错误内容都在此定义
var (
	ErrInvalidID           = errors.New("id格式不正确")
	ErrNotFound            = errors.New("记录未找到")
	ErrNotAllowed          = errors.New("用户不能进行该操作")
	ErrParams              = errors.New("请求参数错误，请仔细检查参数是否异常")
	ErrFail                = errors.New("请求失败")
	ErrIllegalParams       = errors.New("请不要传递恶意参数")

	ErrTaskDone               = errors.New("任务已过期")
	ErrTaskAssigned           = errors.New("任务已指派")
	ErrTaskNotCompleted       = errors.New("任务暂未完成")
	ErrTaskActionNotPermitted = errors.New("无法对该任务进行操作")

	ErrFileNotExist = errors.New("对应的文件不存在或者已经被删除")
	ErrDuplicated	= errors.New("重复项目")
)

func RetOK(c *gin.Context) {
	RetData(c, "操作成功")
}

func RetErrNotFound(c *gin.Context) {
	RetError(c, ErrNotFound)
}

func RetErrAuthFail(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
}

func RetError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": err.Error(),
	})
}

func RetData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
