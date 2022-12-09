package service

import (
	"context"

	"github.com/amsterdan/goim/internal/dao"
	"github.com/amsterdan/goim/internal/model"
	"github.com/go-redis/redis/v8"
)

type userApi struct {
	redisClient *redis.Cmdable
	dao         *dao.AllDao
}

func NewUserApi(redisC *redis.Cmdable, daoF *dao.AllDao) userApi {
	return userApi{
		redisClient: redisC,
		dao:         daoF,
	}
}

// 注册接口
func (u userApi) Register(ctx context.Context, user model.User) {

}
