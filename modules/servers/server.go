package servers

import (
	"encoding/json"
	"os"
	"os/signal"

	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start()
	GetServer() *server
}

type server struct {
	app *fiber.App
	cfg config.IConfig
}

func NewServer(cfg config.IConfig,
) IServer {
	return &server{
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
		cfg: cfg,
	}
}

func (s *server) GetServer() *server {
	return s
}

func (s *server) Start() {

	// Modules
	baseUrl := s.app.Group("/api/v1")
	modules := InitModule(baseUrl, s)

	modules.MonitorModule()

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		logger.Logger.Info("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	logger.Logger.Infof("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
