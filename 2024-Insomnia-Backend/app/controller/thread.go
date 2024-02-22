package controller

import (
	"Insomnia/app/models"
	"Insomnia/app/request"
	"Insomnia/app/response"
	"Insomnia/app/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Thread struct{}

var threadService *service.ThreadService

func (t *Thread) CreateThread(c *gin.Context) {
	//定义一个创建帖子请求的结构体
	req := &request.CreateThreadReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("无法解析的表单: %v", err))
		return
	}
	fmt.Println(req.Images)
	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	thread, err := threadService.CreateThread(uuid, *req)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("创建帖子失败: %v", err), models.Thread{})
		return
	}
	// 反序列化 JSON 字符串并输出
	var retrievedSlice []string
	if err := json.Unmarshal([]byte(thread.Images), &retrievedSlice); err != nil {
		log.Fatal("反序列化数据失败:", err)
	}

	rsp := response.GetThreadResponse{
		CreatedAt:  thread.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:      thread.TUuid,
		Topic:      thread.Topic,
		Title:      thread.Title,
		UuId:       thread.Uuid,
		Likes:      thread.Likes,
		Body:       thread.Body,
		PostNumber: thread.PostNumber,
		Images:     retrievedSlice,
	}
	response.FailMsgData(c, "创建帖子成功", rsp)
	return
}

func (t *Thread) ReadThread(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindThreadReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	thread, err := threadService.ReadThread(req.TUuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("获取帖子失败: %v", err), models.Thread{})
		return
	}

	// 反序列化 JSON 字符串并输出
	var retrievedSlice []string
	if err := json.Unmarshal([]byte(thread.Images), &retrievedSlice); err != nil {
		log.Fatal("反序列化数据失败:", err)
	}

	//获取点赞状态
	exist, err := models.CheckLike(thread.TUuid, thread.Uuid)
	if err != nil {
		return
	}

	rsp := response.GetThreadResponse{
		CreatedAt:  thread.CreatedAt.Format("2006-01-02 15:04"),
		TUuid:      thread.TUuid,
		Topic:      thread.Topic,
		Title:      thread.Title,
		UuId:       thread.Uuid,
		Likes:      thread.Likes,
		Body:       thread.Body,
		PostNumber: thread.PostNumber,
		Images:     retrievedSlice,
		Exist:      strconv.FormatBool(exist),
	}
	response.FailMsgData(c, "获取帖子成功", rsp)
	return
}

func (t *Thread) DestroyThread(c *gin.Context) {
	//定义一个寻找到帖子请求的结构体
	req := &request.FindThreadReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	err := threadService.DestroyThread(req.TUuid)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("%v", err))
		return
	}

	response.FailMsg(c, "删除帖子成功")
	return
}

func (t *Thread) GetThreads(c *gin.Context) {

	//定义一个获取帖子请求的结构体
	req := &request.GetThreadsReq{}

	//使用ShouldBind去解析获得的结构体,蛙趣真清晰啊
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("params invalid error: %v", err))
		return
	}

	//获取帖子
	threads, err := threadService.GetThreads(req.Topic)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("%v", err), models.Thread{})
		return
	}

	var rsp []response.GetThreadResponse
	for _, thread := range threads {
		threads := response.GetThreadResponse{
			CreatedAt:  thread.CreatedAt.Format("2006-01-02 15:04"),
			TUuid:      thread.TUuid,
			Topic:      thread.Topic,
			Title:      thread.Title,
			UuId:       thread.Uuid,
			Likes:      thread.Likes,
			Body:       thread.Body,
			PostNumber: thread.PostNumber,
		}
		rsp = append(rsp, threads)
	}

	response.FailMsgData(c, "获取帖子成功", rsp)
	return
}

func (t *Thread) GetMyThreads(c *gin.Context) {

	Uuid, _ := c.Get("Uuid")
	uuid := Uuid.(string)
	//获取帖子
	threads, err := threadService.GetMyThreads(uuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("%v", err), models.Thread{})
		return
	}

	var rsp []response.GetThreadResponse
	for _, thread := range threads {
		threads := response.GetThreadResponse{
			CreatedAt:  thread.CreatedAt.Format("2006-01-02 15:04"),
			TUuid:      thread.TUuid,
			Topic:      thread.Topic,
			Title:      thread.Title,
			UuId:       thread.Uuid,
			Likes:      thread.Likes,
			Body:       thread.Body,
			PostNumber: thread.PostNumber,
		}
		rsp = append(rsp, threads)
	}

	response.FailMsgData(c, "获取帖子成功", rsp)
	return
}

func (t *Thread) LikeThread(c *gin.Context) {
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
	exist, err := threadService.LikeThreads(req.Uid, uuid)
	if err != nil {
		response.FailMsgData(c, fmt.Sprintf("点赞切换失败:%v", err), response.LikesResponse{})
		return
	}

	response.FailMsgData(c, fmt.Sprintf("点赞状态切换成功!"), response.LikesResponse{Exist: strconv.FormatBool(exist)})
	return
}
