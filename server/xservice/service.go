package xservice

import (
	"boostPuzzle/server/models"
	"boostPuzzle/server/rpc"
)

type service struct {
	rpc rpc.RPC
}

type Service interface {
	GetAlbumByUser(username, offset string) (*models.MediaAlbum, error)
}

func NewService(rpc rpc.RPC) Service {
	return service{
		rpc: rpc,
	}
}

func (s service) GetAlbumByUser(username, offset string) (*models.MediaAlbum, error) {
	return s.rpc.GetAlbum(username, offset)
}
