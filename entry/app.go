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
	pm       *PackageManager
	BundleId string
	AppDir   string
}

// NewApp creates a new App application struct
func NewApp(pm *PackageManager) *Application {
	return &Application{
		pm:       pm,
		BundleId: "icu.bughub.app.HarmonyOSTool",
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Application) StartUp(ctx context.Context) {
	a.ctx = ctx
	a.Ctx = ctx
	a.pm.ctx = ctx
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
func (a *Application) EventExposed() *models.EventData {
	return nil
}

// 这个方法只是为了在前端生成models.App类
func (a *Application) AppExposed() *models.App {
	return nil
}

// 这个方法只是为了在前端生成models.Process类
func (a *Application) ProcessExposed() *models.Process {
	return nil
}
