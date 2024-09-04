package config

import (
	"encoding/json"
	"errors"
	"time"

	"braces.dev/errtrace"
	cache "github.com/patrickmn/go-cache"
)

type Cache struct {
	Code    int         `json:"Code"`
	Message string      `json:"Messsage"`
	Data    interface{} `json:"Data"`
}

type AppCache struct {
	Client *cache.Cache
}

var MyCache CacheItf

func InitCache() {
	MyCache = &AppCache{
		Client: cache.New(5*time.Minute, 10*time.Minute),
	}
}

type CacheItf interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
	IsExist(key string) bool
	GetAndConvertToStruct(key string) (Cache, error)
}

func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errtrace.Wrap(err)
	}

	r.Client.Set(key, b, expiration)
	return nil
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.Client.Get(key)
	if !exist {
		return nil, errtrace.Wrap(errors.New("data not found"))
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errtrace.Wrap(errors.New("format is not arr of bytes"))
	}

	return resByte, nil
}

func (r *AppCache) IsExist(key string) bool {
	_, exist := r.Client.Get(key)
	return exist
}

func (r *AppCache) GetAndConvertToStruct(key string) (Cache, error) {
	var cacheData Cache
	byteData, err := r.Get(key)
	if err != nil {
		return cacheData, err
	}

	if err := json.Unmarshal(byteData, &cacheData); err != nil {
		return cacheData, err
	}

	return cacheData, nil
}
