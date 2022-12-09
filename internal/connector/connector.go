package connector

import (
	"context"
	"fmt"

	"github.com/amsterdan/goim/conf"
	"github.com/amsterdan/goim/internal/dao"
	"github.com/amsterdan/goim/pkg/logs"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Connector struct {
	AllDao      dao.AllDao
	RedisClient *redis.Client
}

func New(cfg conf.Config) (*Connector, error) {
	ctx := context.Background()
	db, err := newDB(cfg.DB)
	if err != nil {
		panic(err)
	}

	c := &Connector{
		AllDao: dao.NewAllDao(db),
	}
	logs.Ctx(ctx).Info("数据库已初始化")

	// 初始化redis
	if cfg.Redis.Host != "" {
		redisC, err := newRedis(cfg)
		if err != nil {
			return c, err
		}

		c.RedisClient = redisC
		logs.Ctx(ctx).Info("redis已初始化")
	}

	return c, nil
}

// 关闭链接
func (c *Connector) Close() error {
	if err := c.RedisClient.Close(); err != nil {
		return err
	}

	return nil
}

// db
func newDB(cfg conf.DBConf) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		cfg.User, cfg.Pwd, cfg.Addr, cfg.Port, cfg.Database, "Asia%2fShanghai",
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return
}

// redis
func newRedis(cfg conf.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Pwd,
		DB:       cfg.Redis.Database,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return client, err
	}
	return client, nil
}
