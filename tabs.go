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
	tabs.Append(container.NewTabItemWithIcon("Cajeros", theme.ComputerIcon(), widget.NewLabel("Cajeros")))
	tabs.Append(container.NewTabItemWithIcon("Usuarios", theme.AccountIcon(), widget.NewLabel("Usuarios")))
	tabs.Append(container.NewTabItemWithIcon("Caja Menor", theme.GridIcon(), widget.NewLabel("Caja Menor")))
	tabs.Append(container.NewTabItemWithIcon("Ajustes", theme.SettingsIcon(), widget.NewLabel("Ajustes")))

	return tabs
}
