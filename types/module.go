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
	for _, controller := range m.controllers {
		prefix := "/" + controller.Path()
		router := m.engine.Group(prefix)
		controller.Routes(router)
	}
}
