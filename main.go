package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"icu.bughub.app/harmonyos-tool/backend/utils"
	"icu.bughub.app/harmonyos-tool/entry"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := entry.NewApp()

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := configDir + string(os.PathSeparator) + app.BundleId

	utils.MkDir(appDir)

	//设置环境变量
	os.Setenv("PATH", os.Getenv("PATH")+":"+appDir+string(os.PathSeparator)+"platform-tools")
	app.AppDir = appDir

	myLog := logger.NewFileLogger(fmt.Sprintf("%s/%s", appDir, "HarmonyOS工具箱.log"))

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "HarmonyOS工具箱",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.StartUp,
		Bind: []interface{}{
			app,
		},
		Logger:             myLog,
		LogLevel:           logger.INFO,
		LogLevelProduction: logger.INFO,
	})

	if err != nil {
		runtime.LogInfo(app.Ctx, err.Error())
		println("Error:", err.Error())
	}
}
