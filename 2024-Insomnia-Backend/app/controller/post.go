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

type Post struct{}

var postService *service.PostService

func (p *Post) CreatePost(c *gin.Context) {
	//定义一个创建帖子请求的结构体
	req := &request.CreatePostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	post, err := postService.CreatePost(uuid, *req)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("创建回复失败: %v", err), models.Post{})
		return
	}

	rsp := response.GetPostResponse{
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:     post.TUuid,
		UuId:      post.Uuid,
		Likes:     post.Likes,
		Body:      post.Body,
	}
	response.FailMsgData(c, "创建回复成功", rsp)
	return
}

func (p *Post) ReadPost(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindPostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	post, err := postService.ReadPost(req.PUuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("获取回复失败: %v", err), models.Post{})
		return
	}

	//获取点赞状态
	exist, err := models.CheckLike(post.PUuid, post.Uuid)
	if err != nil {
		return
	}

	rsp := response.GetPostResponse{
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:     post.TUuid,
		UuId:      post.Uuid,
		Likes:     post.Likes,
		Body:      post.Body,
		Exist:     strconv.FormatBool(exist),
	}
	response.FailMsgData(c, "获取回复成功", rsp)
	return
}

func (p *Post) DestroyPost(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindPostReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	err := postService.DestroyPost(req.PUuid)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("%v", err))
		return
	}

	response.FailMsg(c, "删除回复成功")
	return
}

func (p *Post) GetPosts(c *gin.Context) {

	//定义一个获取回复请求的结构体
	req := &request.GetPostsReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	//获取回复
	posts, err := postService.GetPosts(req.TUuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("%v", err), models.Post{})
		return
	}

	var rsp []response.GetPostResponse
	for _, post := range posts {
		p := response.GetPostResponse{
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04"),
			TUuid:     post.TUuid,
			UuId:      post.Uuid,
			Likes:     post.Likes,
			Body:      post.Body,
		}
		rsp = append(rsp, p)
	}

	response.FailMsgData(c, "获取回复成功", rsp)
	return
}

func (p *Post) LikePost(c *gin.Context) {
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
	exist, err := postService.LikePosts(req.Uid, uuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("点赞切换失败:%v", err), response.LikesResponse{})
		return
	}

	response.FailMsgData(c, fmt.Sprintf("点赞状态切换成功!"), response.LikesResponse{Exist: strconv.FormatBool(exist)})
	return
}
