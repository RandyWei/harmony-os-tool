package main

import (
	"embed"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"icu.bughub.app/harmonyos-tool/backend/utils"
	"icu.bughub.app/harmonyos-tool/entry"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	pm := entry.NewPackageManager()
	// Create an instance of the app structure
	app := entry.NewApp(pm)

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := configDir + string(os.PathSeparator) + app.BundleId

	utils.MkDir(appDir)

	//设置环境变量
	delimiter := ":"
	if runtime.GOOS == "windows" {
		delimiter = ";"
	}
	os.Setenv("PATH", os.Getenv("PATH")+delimiter+appDir+string(os.PathSeparator)+"platform-tools")
	app.AppDir = appDir
	logDir := appDir + string(os.PathSeparator) + "log"
	utils.MkDir(logDir)
	//获取系统当前时间作为日志文件名
	logName := time.Now().Format("20060102")
	myLog := logger.NewFileLogger(fmt.Sprintf("%s/%s%s", logDir, logName, ".log"))

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
			pm,
		},
		Logger:             myLog,
		LogLevel:           logger.INFO,
		LogLevelProduction: logger.INFO,
	})

	if err != nil {
		wailsRuntime.LogInfo(app.Ctx, err.Error())
		println("Error:", err.Error())
	}
}
