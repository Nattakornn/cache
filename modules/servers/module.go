package servers

import (
	"github.com/Nattakornn/cache/modules/monitor/monitorHandles"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r fiber.Router
	s *server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		r: r,
		s: s,
	}
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandles.MonitorHandler(m.s.cfg)

	m.r.Get("/health", handler.HealthCheck)
}
