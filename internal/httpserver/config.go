package httpserver

import (
	"companies-test-task/internal/service"
	"fmt"
)

type Config struct {
	Service service.Config

	Port int
}

func validateCfg(cfg Config) error {
	if cfg.Port == 0 {
		return fmt.Errorf("http listen port not specified")
	}

	return nil
}
