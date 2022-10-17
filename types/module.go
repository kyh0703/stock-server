package types

import (
	"fmt"

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
	for ctrl := range m.controllers {
		fmt.Println(ctrl.Path())
	}
}
