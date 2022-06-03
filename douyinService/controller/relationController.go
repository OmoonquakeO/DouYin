package controller

import (
	"github.com/DouYin/common/constant"
	"github.com/DouYin/common/entity/request"
	"github.com/DouYin/common/entity/response"
	"github.com/DouYin/common/entity/vo"
	"github.com/DouYin/service/service"
	"github.com/DouYin/service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var relationService service.RelationService

const (
	FOCUS   = 1
	NoFOCUS = 2
)

// RelationAction
// @Description: 登录用户对其他用户进行关注或取消关注
// @param: c
func RelationAction(c *gin.Context) {
	var relationReq request.RelationReq
	_ = c.ShouldBindQuery(&relationReq)
	user := utils.GetUserContext(c)
	if relationReq.ActionType == NoFOCUS {
		if /*CommentVos,*/ err := relationService.RelationAction(relationReq, user.Id); !err {
			response.FailWithMessage("操作失败", c)
		} else {
			c.JSON(http.StatusOK, response.Response{StatusCode: constant.Success, StatusMsg: "操作成功"})
		}
	} else if relationReq.ActionType == FOCUS {
		if /*CommentVos,*/ err := relationService.AddAction(relationReq, user.Id); !err {
			response.FailWithMessage("操作失败", c)
		} else {
			c.JSON(http.StatusOK, response.Response{StatusCode: constant.FAIL, StatusMsg: "操作成功"})
		}
	}
}

// FollowList
// @Description: 登录用户关注的所有用户列表
// @param: c
func FollowList(c *gin.Context) {
	userId := utils.String2Uint64(c.Query("user_id"))
	userList := relationService.FollowList(userId)
	c.JSON(http.StatusOK, vo.UserListVo{
		Response: response.Response{StatusCode: response.SUCCESS, StatusMsg: "操作成功"},
		UserList: userList,
	})
}

// FollowerList
// @Description: 所有关注登录用户的粉丝列表
// @param: c
func FollowerList(c *gin.Context) {
	userId := utils.String2Uint64(c.Query("user_id"))
	userList := relationService.FollowerList(userId)
	c.JSON(http.StatusOK, vo.UserListVo{
		Response: response.Response{StatusCode: response.SUCCESS, StatusMsg: "操作成功"},
		UserList: userList,
	})
}
