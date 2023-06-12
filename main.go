package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-fiber-api/config"
	"go-fiber-api/controller"
	"go-fiber-api/helpers"
	"go-fiber-api/repository"
	"go-fiber-api/routes"
	"go-fiber-api/services"
	"log"
)

type MyApp struct {
	app        *fiber.App
	config     *config.Config
	repository *repository.MongoRepository
	redis      *helpers.RedisHelper
	minio      *helpers.MinioHelper
	ctx        *context.Context
	routes     []routes.BaseRoute
}

func (myApp *MyApp) Start() {

	myApp.app.Use(cors.New())
	myApp.app.Use(logger.New())
	myApp.app.Use(routes.PreRoute)
	for i, _ := range myApp.routes {
		myApp.routes[i].RegisterRoutes(myApp.app)
	}
	log.Fatal(myApp.app.Listen(":8080"))
}

func NewApp() MyApp {

	ctx := context.TODO()

	_config := config.Config{}
	err := _config.Init(".")
	if err != nil {
		fmt.Println("db err")
	}

	mongoRepository := repository.MongoRepository{
		Config: &_config,
	}
	if err := mongoRepository.Connect(); err != nil {
		fmt.Println("db err")
	}

	opt, err := redis.ParseURL(_config.RedisUrl)
	if err != nil {
		panic(err)
	}
	redisClient := helpers.RedisHelper{
		Client: redis.NewClient(opt),
	}

	if err := redisClient.Ping(ctx); err != nil {
		panic(err)
	}

	miniHelper := helpers.MinioHelper{
		Config: &_config,
		Ctx:    ctx,
	}

	if err := miniHelper.Connect(); err != nil {
		panic("mini0 err")
	}

	if err := miniHelper.NewBucket("test"); err != nil {
		panic("mini0 err")
	}

	myApp := MyApp{
		app:        fiber.New(),
		ctx:        &ctx,
		config:     &_config,
		repository: &mongoRepository,
		redis:      &redisClient,
		minio:      &miniHelper,
		routes:     make([]routes.BaseRoute, 0),
	}

	userService := services.UserService{
		Collection: myApp.repository.GetCollection("users"),
		Cache:      myApp.redis,
		Minio:      myApp.minio,
		Ctx:        ctx,
	}
	userRoute := routes.NewUserRoute(controller.NewUserController(userService))
	authRoute := routes.NewAuthRoute(controller.NewAuthController(userService))

	myApp.routes = append(myApp.routes, userRoute)
	myApp.routes = append(myApp.routes, authRoute)

	return myApp
}

func main() {
	myApp := NewApp()
	myApp.Start()
}
