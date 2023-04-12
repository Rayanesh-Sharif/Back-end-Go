package session

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisStorage struct {
	client *redis.Client
}

const accessTokenPrefix = "access:"
const refreshTokenPrefix = "refresh:"

func NewStorageFromRedis(options *redis.Options) (Storage, error) {
	r := redis.NewClient(options)
	err := r.Ping(context.Background()).Err()
	if err != nil {
		return redisStorage{}, err
	}
	return redisStorage{r}, nil
}

func (r redisStorage) Store(userID uint32, ttl time.Duration) (accessToken, refreshToken string, err error) {
	// Generate tokens
	accessToken = newToken()
	refreshToken = newToken()
	// Insert them in cache
	err = r.client.Set(context.Background(), accessTokenPrefix+accessToken, userID, ttl).Err()
	if err != nil {
		return "", "", errors.Wrap(err, "cannot set access token")
	}
	err = r.client.Set(context.Background(), refreshTokenPrefix+refreshToken, accessToken, ttl).Err()
	if err != nil {
		return "", "", errors.Wrap(err, "cannot set refresh token")
	}
	return
}

func (r redisStorage) Get(accessToken string) (userID uint32, err error) {
	id, err := r.client.Get(context.Background(), accessTokenPrefix+accessToken).Uint64()
	if errors.Is(err, redis.Nil) {
		return 0, ErrNotFound
	}
	return uint32(id), err
}

func (r redisStorage) Delete(accessToken string) error {
	// Refresh token will be checked in Refresh method
	return r.client.Del(context.Background(), accessTokenPrefix+accessToken).Err()
}

func (r redisStorage) Refresh(refreshToken string, ttl time.Duration) (newAccessToken string, err error) {
	// Get the access token associated with refresh token
	accessTokenData := r.client.Get(context.Background(), refreshTokenPrefix+refreshToken)
	if accessTokenData.Err() != nil {
		return "", errors.Wrap(accessTokenData.Err(), "cannot get access token of refresh token")
	}
	r.client.Expire(context.Background(), refreshTokenPrefix+refreshToken, ttl)
	accessToken := accessTokenData.Val()
	// Renew TTL
	changedTTL := r.client.Expire(context.Background(), accessTokenPrefix+accessToken, ttl)
	if changed, _ := changedTTL.Result(); !changed {
		return "", ErrNotFound
	}
	// Change it
	newAccessToken = newToken()
	r.client.Rename(context.Background(), accessTokenPrefix+accessToken, accessTokenPrefix+newAccessToken)
	return
}

func (r redisStorage) Close() {
	_ = r.client.Close()
}
