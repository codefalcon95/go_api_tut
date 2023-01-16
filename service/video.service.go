package service

import "github.com/codefalcon95/golang-gin-poc/entity"

type VideoService interface{
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}


type videoService struct {
	videos []entity.Video
}

func New() videoService {
	// return v
}