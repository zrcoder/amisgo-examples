package main

import (
	"embed"

	"github.com/zrcoder/amisgo"
)

//go:embed assets/*
var assetsFS embed.FS

var appConfig *amisgo.Config

func init() {
	appConfig = amisgo.GetDefaultConfig()
	appConfig.Theme = amisgo.ThemeDark
	appConfig.Lang = amisgo.LangEn
	appConfig.StaticDir = "assets"
	appConfig.StaticFS = assetsFS
	appConfig.Icon = "/assets/favicon.ico"
}
