package types

import (
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	engine      *fiber.App
	controllers []Controller
}

func (m *Module) SetEngine(engine *fiber.App) {
	m.engine = engine
}

func (m *Module) AttachController(ctrl Controller) {
	m.controllers = append(m.controllers, ctrl)
}

func (m *Module) Init() {
	api := m.engine.Group("/api")
	for _, controller := range m.controllers {
		prefix := "/" + controller.Path()
		router := api.Group(prefix)
		controller.Routes(router)
	}
}
