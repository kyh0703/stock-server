package types

import (
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	controllers []Controller
}

func (m *Module) AppendControllers(controllers ...Controller) {
	m.controllers = append(m.controllers, controllers...)
}

func (m *Module) Init(engine *fiber.App) {
	api := engine.Group("/api")
	for _, controller := range m.controllers {
		prefix := "/" + controller.Path()
		router := api.Group(prefix)
		controller.Index(router)
	}
}
