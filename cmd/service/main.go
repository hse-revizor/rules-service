package main

import (
	"github.com/hse-revizor/rules-service/internal/di"
	"github.com/hse-revizor/rules-service/internal/utils/config"
	"github.com/hse-revizor/rules-service/internal/utils/flags"
	"github.com/slipneff/gogger"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(cfg)
	gogger.ConfigureZeroLogger()

	container.GetRuleService()
}
