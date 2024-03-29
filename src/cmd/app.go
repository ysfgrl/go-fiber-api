package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ysfgrl/fibersocket"
	"github.com/ysfgrl/go-fiber-api/src/config"
	"github.com/ysfgrl/go-fiber-api/src/controller"
	"github.com/ysfgrl/go-fiber-api/src/middleware"
	"strconv"
)

var (
	app *fiber.App = nil
)

func init() {
	app = fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
		//Views:     html.NewFileSystem(http.Dir("/Users//AndroidStudioProjects/ui/build/web"), ".html"),
	})
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(middleware.PreRoute)
	controller.AuthController.Routes(app)
	controller.UserController.Routes(app)
	socketServer := fibersocket.NewSocketServer("task_server")
	app.Get("/ws/:id", socketServer.NewSocket())
	app.Static(
		"/admin", // mount address
		"/Users//AndroidStudioProjects/ui/build/web/", // path to the file folder
	)
}

func Listen() error {
	return app.Listen(config.AppConf.App.Host + ":" + strconv.Itoa(config.AppConf.App.Port))
}
