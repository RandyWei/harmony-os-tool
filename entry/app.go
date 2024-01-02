package entry

import (
	"context"
	"fmt"
	"os"

	"icu.bughub.app/harmonyos-tool/backend"
	"icu.bughub.app/harmonyos-tool/backend/models"
)

// App struct
type Application struct {
	ctx      context.Context
	bundleId string
	AppDir   string
}

// NewApp creates a new App application struct
func NewApp() *Application {
	return &Application{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Application) StartUp(ctx context.Context) {
	a.ctx = ctx
	a.bundleId = "icu.bughub.app.HarmonyOSTool"

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := configDir + string(os.PathSeparator) + a.bundleId
	//设置环境变量
	os.Setenv("PATH", os.Getenv("PATH")+":"+appDir+string(os.PathSeparator)+"platform-tools")
	a.AppDir = appDir
}

// Greet returns a greeting for the given name
func (a *Application) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *Application) WaitForDevice() ([]models.Device, error) {
	return backend.WaitForDevice(a.AppDir)
}

func (a *Application) InstallAdb() (bool, error) {
	return backend.InstallAdb(a.ctx, a.AppDir)
}

/**
 * GetAppDir
 */
func (a *Application) GetAppDir() string {
	return a.AppDir
}

func (a *Application) ListApps() ([][]models.App, error) {
	return backend.ListApps()
}

func (a *Application) ListApps1() ([]models.App, error) {
	return make([]models.App, 0), nil
}

func (a *Application) UninstallApp(packageName string) (bool, error) {
	return backend.UninstallApp(packageName)
}

func (a *Application) InstallExistingApp(packageName string) (bool, error) {
	return backend.InstallExistingApp(packageName)
}
