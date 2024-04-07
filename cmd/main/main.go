package main

import (
	"context"
	"letsgo-mini-is/internal/adapter/config"
	"letsgo-mini-is/internal/adapter/handler/http"
	"letsgo-mini-is/internal/adapter/storage/mysql"
	"letsgo-mini-is/internal/adapter/storage/mysql/repository"
	"letsgo-mini-is/internal/core/service"
	"log"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("environment variables를 불러오던 중 에러가 발생했습니다 : ", err)
	}

	db, err := mysql.New(context.Background(), config.DB)
	if err != nil {
		log.Fatal("Database를 초기화 및 생성 중 에러가 발행했습니다 : ", err)
	}
	defer db.Close()

	pictureRepository := repository.NewPictureRepository(db)
	pictureService := service.NewPictureService(pictureRepository)
	pictureHandler := http.NewPictureHandler(pictureService)

	router, err := http.NewRouter(pictureHandler)
	if err != nil {
		log.Fatal("router를 초기화 하던 중 에러가 발생했습니다 : ", err)
	}

	if err := router.Run(config.HTTP.Port); err != nil {
		log.Fatal("서버가 실행 중 에러가 발생했습니다 : ", err)
	}
}
