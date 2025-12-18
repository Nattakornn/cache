package servers

import (
	"github.com/Nattakornn/cache/modules/middlewares/middlewaresHandlers"
	"github.com/Nattakornn/cache/modules/middlewares/middlewaresRepositories"
	"github.com/Nattakornn/cache/modules/middlewares/middlewaresUseCases"
	"github.com/Nattakornn/cache/modules/monitor/monitorHandles"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r fiber.Router
	s *server
	m middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r: r,
		s: s,
		m: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	useCase := middlewaresUseCases.MiddlewaresUseCase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, useCase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandles.MonitorHandler(m.s.cfg)

	m.r.Get("/health", handler.HealthCheck)
}
