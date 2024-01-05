package entry

import (
	"context"
	"fmt"
	"os"

	"icu.bughub.app/harmonyos-tool/backend"
	"icu.bughub.app/harmonyos-tool/backend/models"
	"icu.bughub.app/harmonyos-tool/backend/utils"
)

// App struct
type Application struct {
	ctx      context.Context
	Ctx      context.Context
	BundleId string
	AppDir   string
}

// NewApp creates a new App application struct
func NewApp() *Application {
	return &Application{
		BundleId: "icu.bughub.app.HarmonyOSTool",
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Application) StartUp(ctx context.Context) {
	a.ctx = ctx
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *Application) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *Application) WaitForDevice() ([]models.Device, error) {
	return backend.WaitForDevice(a.ctx, a.AppDir)
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

func (a *Application) ListModules(brand string) ([]models.Module, error) {
	return backend.ListModules(a.ctx, brand)
}

func (a *Application) ListModuleApps(moduleId string) {

	go func() {
		apps, _ := backend.ListModuleApps(a.ctx, moduleId)
		event := &models.Event{
			Ctx:  a.ctx,
			Name: models.Event_Refresh_App_List,
			Data: models.EventData{
				Type: moduleId,
				Data: apps,
			},
		}
		event.Send()
	}()
}

func (a *Application) ListApps(appType int64) {
	go func() {

		var apps []models.App
		switch appType {
		case 0:
			apps, _ = backend.ListApps0(a.ctx)
		case 1:
			apps, _ = backend.ListApps1(a.ctx)
		case 2:
			apps, _ = backend.ListApps2(a.ctx)
		case 3:
			apps, _ = backend.ListApps3(a.ctx)
		case 4:
			apps, _ = backend.ListApps4(a.ctx)
		case 5:
			apps, _ = backend.ListApps5(a.ctx)
		}
		event := &models.Event{
			Ctx:  a.ctx,
			Name: models.Event_Refresh_App_List,
			Data: models.EventData{
				Type: string(rune(appType)),
				Data: apps,
			},
		}
		event.Send()
	}()
}

func (a *Application) UninstallApp(packageName string, relatedIds []string) (bool, error) {
	return backend.UninstallApp(a.ctx, packageName, relatedIds)
}

func (a *Application) InstallExistingApp(packageName string, relatedIds []string) (bool, error) {
	return backend.InstallExistingApp(a.ctx, packageName, relatedIds)
}

// 禁用应用
func (a *Application) DisableApp(packageName string) (bool, error) {
	return backend.DisableApp(a.ctx, packageName)
}

// 启用应用
func (a *Application) EnableApp(packageName string) (bool, error) {
	return backend.EnableApp(a.ctx, packageName)
}

// 查看内存情况
func (a *Application) Free() (string, error) {
	return backend.Free(a.ctx)
}

// 查看进程情况
func (a *Application) Top() (models.TopInfo, error) {
	return backend.Top(a.ctx)
}

// 获取adb安装目录，如果不存在则返回空
func (a *Application) GetAdbDir() *models.DirModel {
	filePath := a.AppDir + string(os.PathSeparator) + "platform-tools"
	return utils.GetDir(filePath)
}

// 获取adb安装目录，如果不存在则返回空
func (a *Application) GetLogDir() *models.DirModel {
	filePath := a.AppDir + string(os.PathSeparator) + "log"
	return utils.GetDir(filePath)
}

// 删除log目录
func (a *Application) RemoveLog() {
	filePath := a.AppDir + string(os.PathSeparator) + "log"
	dir, _ := os.ReadDir(filePath)
	for _, file := range dir {
		os.RemoveAll(filePath + string(os.PathSeparator) + file.Name())
	}
}

// 删除adb安装目录
func (a *Application) RemoveAdb() {
	filePath := a.AppDir + string(os.PathSeparator) + "platform-tools"
	dir, _ := os.ReadDir(filePath)
	for _, file := range dir {
		os.RemoveAll(filePath + string(os.PathSeparator) + file.Name())
	}
}

// 这个方法只是为了在前端生成models.EventData类
func (a *Application) EventTest() *models.EventData {
	return nil
}

// 这个方法只是为了在前端生成models.App类
func (a *Application) AppTest() *models.App {
	return nil
}

// 这个方法只是为了在前端生成models.Process类
func (a *Application) ProcessTest() *models.Process {
	return nil
}
