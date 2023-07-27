package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetTabs() *container.AppTabs {
	tabs := container.NewAppTabs()
	tabs.SetTabLocation(container.TabLocationTop)

	tabs.Append(container.NewTabItemWithIcon("Transacciones", theme.StorageIcon(), GetTabTransacciones()))
	tabs.Append(container.NewTabItemWithIcon("Cajeros", theme.ComputerIcon(), GetTabCajeros()))
	tabs.Append(container.NewTabItemWithIcon("Usuarios", theme.AccountIcon(), GetTabUsuarios()))
	tabs.Append(container.NewTabItemWithIcon("Sucursales", theme.GridIcon(), GetTabSucursales()))
	tabs.Append(container.NewTabItemWithIcon("Ajustes", theme.SettingsIcon(), widget.NewLabel("Ajustes")))

	return tabs
}
