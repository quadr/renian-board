package board

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/securecookie"
)

type Server struct {
	sessionManager securecookie.SecureCookie
	dbPool         *redis.Pool
	postService    *PostService
}

func NewServer(config *ServerConfig) (*Server, error) {
	p := &PostService{}
	p.Register()

	hashKey := []byte(config.SecureSession.HashKey)
	blockKey := []byte(config.SecureSession.BlockKey)

	return &Server{
		sessionManager: securecookie.New(hashKey, blockKey),
		dbPool:         newPool(config),
		postService:    p,
	}
}
