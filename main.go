package main

import (
	"docker-api-service/docker"
	"docker-api-service/routers"
	"docker-api-service/setting"
	"fmt"
	"log"
	"net/http"
)

func init() {
	docker.Setup()
}
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:              fmt.Sprint(":", setting.HttpPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 << 20, //最大1M
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
