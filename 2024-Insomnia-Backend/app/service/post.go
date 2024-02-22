package service

import (
	"Insomnia/app/models"
	. "Insomnia/app/request"
	"fmt"
)

type PostService struct{}

func (p *PostService) CreatePost(uuid string, cp CreatePostReq) (models.Post, error) {
	post, err := models.CreatePost(uuid, cp)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p *PostService) ReadPost(pUuid string) (models.Post, error) {
	post, err := models.PostByPUUID(pUuid)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p *PostService) DestroyPost(pUuid string) error {
	err := models.DestroyPost(pUuid)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) GetPosts(tUuid string) ([]models.Post, error) {
	posts, err := models.Posts(tUuid)
	if err != nil {
		return []models.Post{}, fmt.Errorf("无法获取回复:%v", err)
	}
	return posts, nil
}

func (p *PostService) LikePosts(pUuid string, uuid string) (exist bool, err error) {
	// 开始事务
	tx := models.Db.Begin()
	exist, err = models.ChangeLike(pUuid, uuid)
	//检查是否出错
	if err != nil {
		// 如果出错，回滚事务
		tx.Rollback()
		return
	}
	//根据改变后的点赞类型自动增减
	err = models.UpPostLikesData(pUuid, exist)
	if err != nil {
		// 如果出错，回滚事务
		tx.Rollback()
		return
	}
	//提交事务
	tx.Commit()
	return
}
