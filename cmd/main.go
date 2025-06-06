package main

import (
	"log"

	"github.com/setiawan-chandra/simple-forum-fastcampus/internal/configs"
	"github.com/setiawan-chandra/simple-forum-fastcampus/internal/handlers/memberships"
	"github.com/setiawan-chandra/simple-forum-fastcampus/pkg/internalsql"

	"github.com/gin-gonic/gin"

	membershipsRepo "github.com/setiawan-chandra/simple-forum-fastcampus/internal/repository/memberships"
	membershipSvc "github.com/setiawan-chandra/simple-forum-fastcampus/internal/service/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config : ", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	membershipsRepo := membershipsRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipsRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
