package entry

import (
	"context"

	"icu.bughub.app/harmonyos-tool/backend"
	"icu.bughub.app/harmonyos-tool/backend/models"
)

// PM struct
type PackageManager struct {
	ctx context.Context
}

func NewPackageManager() (p *PackageManager) {
	p = &PackageManager{}
	return
}

func (p *PackageManager) ListModules(brand string) ([]models.Module, error) {
	return backend.ListModules(p.ctx, brand)
}

func (p *PackageManager) ListModuleApps(brand string, moduleId string) {

	go func() {
		apps, _ := backend.ListModuleApps(p.ctx, brand, moduleId)
		event := &models.Event{
			Ctx:  p.ctx,
			Name: models.Event_Refresh_App_List,
			Data: models.EventData{
				Type: moduleId,
				Data: apps,
			},
		}
		event.Send()
	}()
}

func (p *PackageManager) ListApps(appType int64) {
	go func() {

		var apps []models.App
		switch appType {
		case 0:
			apps, _ = backend.ListApps0(p.ctx)
		case 1:
			apps, _ = backend.ListApps1(p.ctx)
		case 2:
			apps, _ = backend.ListApps2(p.ctx)
		case 3:
			apps, _ = backend.ListApps3(p.ctx)
		case 4:
			apps, _ = backend.ListApps4(p.ctx)
		case 5:
			apps, _ = backend.ListApps5(p.ctx)
		}
		event := &models.Event{
			Ctx:  p.ctx,
			Name: models.Event_Refresh_App_List,
			Data: models.EventData{
				Type: string(rune(appType)),
				Data: apps,
			},
		}
		event.Send()
	}()
}

func (p *PackageManager) UninstallApp(packageName string, relatedIds []string) (bool, error) {
	return backend.UninstallApp(p.ctx, packageName, relatedIds)
}

func (p *PackageManager) InstallExistingApp(packageName string, relatedIds []string) (bool, error) {
	return backend.InstallExistingApp(p.ctx, packageName, relatedIds)
}

// 禁用应用
func (p *PackageManager) DisableApp(packageName string) (bool, error) {
	return backend.DisableApp(p.ctx, packageName)
}

// 启用应用
func (p *PackageManager) EnableApp(packageName string) (bool, error) {
	return backend.EnableApp(p.ctx, packageName)
}
