package service

import "github.com/codefalcon95/golang-gin-poc/entity"

type VideoService interface{
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}


type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (s *videoService) Save(v entity.Video) entity.Video{
	s.videos = append(s.videos, v);
	return v
}

func (s *videoService) FindAll() []entity.Video{
	return s.videos;
}