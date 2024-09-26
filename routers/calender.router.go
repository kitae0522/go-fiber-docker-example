package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-docker-example/controllers/calenderCtrl"
)

func initCalenderRouter(router fiber.Router) {
	calenderRouter := router.Group("/calender")

	calenderRouter.Post("/create", calenderCtrl.CreateCalender)
	calenderRouter.Delete("/delete", calenderCtrl.DeleteCalender)
	calenderRouter.Put("/update", calenderCtrl.UpdateCalender)
	calenderRouter.Get("/:id", calenderCtrl.GetCalender)
}
