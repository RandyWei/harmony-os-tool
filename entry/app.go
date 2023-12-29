package entry

import (
	"context"
	"fmt"
	"os"

	"icu.bughub.app/harmonyos-tool/backend"
)

// App struct
type App struct {
	ctx      context.Context
	bundleId string
	AppDir   string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) StartUp(ctx context.Context) {
	a.ctx = ctx
	a.bundleId = "icu.bughub.app.HarmonyOSTool"

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir := configDir + string(os.PathSeparator) + a.bundleId

	a.AppDir = appDir
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) WaitForDevice() (bool, error) {
	return backend.WaitForDevice(a.AppDir)
}

func (a *App) InstallAdb() (bool, error) {
	return backend.InstallAdb(a.ctx, a.AppDir)
}

/**
 * GetAppDir
 */
func (a *App) GetAppDir() string {
	return a.AppDir
}
