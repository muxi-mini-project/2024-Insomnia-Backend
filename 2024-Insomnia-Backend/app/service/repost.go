package service

import (
	"Insomnia/app/models"
	. "Insomnia/app/request"
	"fmt"
)

type RePostService struct{}

func (r *RePostService) CreateRePost(uuid string, crp CreateRePostReq) (models.RePost, error) {
	repost, err := models.CreateRePost(uuid, crp)
	if err != nil {
		return models.RePost{}, err
	}
	return repost, nil
}

func (r *RePostService) ReadRePost(rUuid string) (models.RePost, error) {
	repost, err := models.RePostByRUUID(rUuid)
	if err != nil {
		return models.RePost{}, err
	}
	return repost, nil
}

func (r *RePostService) DestroyRePost(rUuid string) error {
	err := models.DestroyRePost(rUuid)
	if err != nil {
		return err
	}
	return nil
}

func (r *RePostService) GetRePosts(pUuid string) ([]models.RePost, error) {
	reposts, err := models.RePosts(pUuid)
	if err != nil {
		return []models.RePost{}, fmt.Errorf("无法获取re回复:%v", err)
	}
	return reposts, nil
}

func (r *RePostService) LikeRePosts(rUuid string, uuid string) (exist bool, err error) {
	// 开始事务
	tx := models.Db.Begin()
	exist, err = models.ChangeLike(rUuid, uuid)
	//检查是否出错
	if err != nil {
		// 如果出错，回滚事务
		tx.Rollback()
		return
	}
	//根据改变后的点赞类型自动增减
	err = models.UpRePostLikesData(rUuid, exist)
	if err != nil {
		// 如果出错，回滚事务
		tx.Rollback()
		return
	}
	//提交事务
	tx.Commit()
	return
}
