package controller

import (
	"Insomnia/app/models"
	"Insomnia/app/request"
	"Insomnia/app/response"
	"Insomnia/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RePost struct{}

var repostService *service.RePostService

func (r *RePost) CreateRePost(c *gin.Context) {
	//定义一个创建帖子请求的结构体
	req := &request.CreateRePostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	repost, err := repostService.CreateRePost(uuid, *req)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("创建re回复失败: %v", err), models.RePost{})
		return
	}

	rsp := response.GetRePostResponse{
		CreatedAt: repost.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:     repost.TUuid,
		UuId:      repost.Uuid,
		PUuid:     repost.PUuid,
		Likes:     repost.Likes,
		Body:      repost.Body,
		RUuid:     repost.RUuid,
	}
	response.FailMsgData(c, "创建回复成功", rsp)
	return
}

func (r *RePost) ReadRePost(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindRePostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	repost, err := repostService.ReadRePost(req.RUuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("获取re回复失败: %v", err), models.Post{})
		return
	}

	//获取点赞状态
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	exist, err := models.CheckLike(repost.RUuid, uuid)
	if err != nil {
		return
	}

	rsp := response.GetRePostResponse{
		CreatedAt: repost.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:     repost.TUuid,
		UuId:      repost.Uuid,
		PUuid:     repost.PUuid,
		RUuid:     repost.RUuid,
		Likes:     repost.Likes,
		Body:      repost.Body,
		Exist:     strconv.FormatBool(exist),
	}
	response.FailMsgData(c, "获取re回复成功", rsp)
	return
}

func (r *RePost) DestroyRePost(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindRePostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	err := repostService.DestroyRePost(req.RUuid)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("%v", err))
		return
	}

	response.FailMsg(c, "删除re回复成功")
	return
}

func (r *RePost) GetRePosts(c *gin.Context) {

	//定义一个获取回复请求的结构体
	req := &request.GetRePostsReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	//获取回复
	reposts, err := repostService.GetRePosts(req.PUuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("%v", err), models.RePost{})
		return
	}

	var rsp []response.GetRePostResponse
	for _, repost := range reposts {
		rp := response.GetRePostResponse{
			CreatedAt: repost.CreatedAt.Format("2006-01-02 15:04"),
			TUuid:     repost.TUuid,
			UuId:      repost.Uuid,
			PUuid:     repost.PUuid,
			Likes:     repost.Likes,
			Body:      repost.Body,
			RUuid:     repost.RUuid,
		}
		rsp = append(rsp, rp)
	}

	response.FailMsgData(c, "获取re回复成功", rsp)
	return
}

func (r *RePost) LikeRePost(c *gin.Context) {
	//定义一个创建帖子请求的结构体
	req := &request.LikesReq{}
	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)

	//切换点赞状态
	exist, err := repostService.LikeRePosts(req.Uid, uuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("点赞切换失败:%v", err), response.LikesResponse{})
		return
	}

	response.FailMsgData(c, fmt.Sprintf("点赞状态切换成功!"), response.LikesResponse{Exist: strconv.FormatBool(exist)})
	return
}
