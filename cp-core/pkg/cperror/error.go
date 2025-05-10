package cperror

import "errors"

var ErrKeyNotExistsRedis = errors.New("key not exists in redis")
