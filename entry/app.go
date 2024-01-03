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

func (a *Application) ListApps1() ([]models.App, error) {
	return backend.ListApps1()
}

func (a *Application) ListApps2() ([]models.App, error) {
	return backend.ListApps2()
}
func (a *Application) ListApps3() ([]models.App, error) {
	return backend.ListApps3()
}
func (a *Application) ListApps4() ([]models.App, error) {
	return backend.ListApps4()
}
func (a *Application) ListApps5() ([]models.App, error) {
	return backend.ListApps5()
}
func (a *Application) ListApps(appType int64) {
	go func() {

		var apps []models.App
		switch appType {
		case 1:
			apps, _ = backend.ListApps1()
		case 2:
			apps, _ = backend.ListApps2()
		case 3:
			apps, _ = backend.ListApps3()
		case 4:
			apps, _ = backend.ListApps4()
		case 5:
			apps, _ = backend.ListApps5()
		}
		event := &models.Event{
			Ctx:  a.ctx,
			Name: models.Event_Refresh_App_List,
			Data: models.EventData{
				Type: appType,
				Data: apps,
			},
		}
		fmt.Printf("ListApps:%d\n", appType)
		event.Send()
	}()
}

func (a *Application) UninstallApp(packageName string, relatedIds []string) (bool, error) {
	return backend.UninstallApp(packageName, relatedIds)
}

func (a *Application) InstallExistingApp(packageName string, relatedIds []string) (bool, error) {
	return backend.InstallExistingApp(packageName, relatedIds)
}

// 这个方法只是为了在前端生成models.EventData类
func (a *Application) EventTest() *models.EventData {
	return nil
}
