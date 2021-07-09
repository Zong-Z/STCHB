package database

import (
	"context"
	"stchb/types"

	"github.com/go-redis/redis/v8"
)

type database interface {
	SaveUser(user types.User) error
	GetUser(userID int) (*types.User, error)
}

type redisDB struct {
	database
	Ctx    context.Context
	client *redis.Client
}

var DB *redisDB
