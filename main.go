package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/urfave/cli"
	"github.com/ysfgrl/fibersocket"
	"go-fiber-api/src/config"
	"go-fiber-api/src/controller"
	"go-fiber-api/src/helpers"
	"go-fiber-api/src/repository/mongo_repository"
	"go-fiber-api/src/routes"
	"go-fiber-api/src/services"
	"go-fiber-api/src/worker"
	"go-fiber-api/src/worker/tasks"
	"os"
)

type myApp struct {
	app    *fiber.App
	ctx    *context.Context
	routes []routes.Route
}

func (myApp *myApp) Start() error {
	myApp.app.Use(cors.New())
	myApp.app.Use(logger.New())
	myApp.app.Use(routes.PreRoute)
	for i, _ := range myApp.routes {
		myApp.routes[i].RegisterRoutes(myApp.app)
	}
	socketServer := fibersocket.NewSocketServer("task_server")
	myApp.app.Get("/ws/:id", socketServer.NewSocket())

	myApp.app.Static(
		"/admin", // mount address
		"/Users//AndroidStudioProjects/ui/build/web/", // path to the file folder
	)
	return myApp.app.Listen(config.App.Host + ":" + config.App.Port)
}

func newFiberApp(confType string) *myApp {
	ctx := context.TODO()
	err := config.InitConf(confType)
	if err != nil {
		panic(err.ToJson())
	}
	helpers.InitHelpers()

	newApp := myApp{
		app: fiber.New(fiber.Config{
			BodyLimit: 100 * 1024 * 1024,
			//Views:     html.NewFileSystem(http.Dir("/Users//AndroidStudioProjects/ui/build/web"), ".html"),
		}),
		ctx:    &ctx,
		routes: make([]routes.Route, 0),
	}

	userRepo := mongo_repository.NewUserRepo(helpers.Mongo.GetCollection("users"))
	userService := services.NewUserService(userRepo)

	userRoute := routes.NewUserRoute(controller.NewUserController(userService))
	authRoute := routes.NewAuthRoute(controller.NewAuthController(userService))

	resourceRepo := mongo_repository.NewResourceRepo(helpers.Mongo.GetCollection("resources"))
	resourceService := services.NewResourceService(resourceRepo)
	resourceRoute := routes.NewResourceRoute(controller.NewResourceController(resourceService))

	newApp.routes = append(newApp.routes, userRoute)
	newApp.routes = append(newApp.routes, authRoute)
	newApp.routes = append(newApp.routes, resourceRoute)
	newApp.routes = append(newApp.routes, routes.NewTaskRoute(controller.NewTaskController(services.NewTaskService())))

	newWorker := worker.NewWorker()
	newWorker.AddTask(tasks.NewDownloadTask(resourceService))
	newWorker.AddTask(tasks.NewTestTask())
	newWorker.RegisterAll()
	return &newApp
}

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "ysfgrl"
	cliApp.Usage = "go fiber api"
	cliApp.Version = "0.0.0"

	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "run workers",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						newFiberApp("dev")
						if err := worker.Launch(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "send",
			Usage: "test workers ",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						newFiberApp("dev")
						if err := worker.Test(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "app",
			Usage: "run fiber app",
			Subcommands: []cli.Command{
				{
					Name:  "dev",
					Usage: "one of [dev, prod, test]",
					Action: func(c *cli.Context) error {
						app := newFiberApp("dev")
						//schema := models.TaskState{
						//	TaskName: "test",
						//	TaskUUID: uuid.New().String(),
						//	Results:  nil,
						//	Args:     nil,
						//}
						if err := app.Start(); err != nil {
							return cli.NewExitError(err, 1)
						}
						return nil
					},
				},
			},
		},
	}
	_ = cliApp.Run(os.Args)
}
