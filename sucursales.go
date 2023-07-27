package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var SucursalesContainer = container.NewHBox()

func GetButtonUpdateCardsSucursales() *widget.Button {
	return widget.NewButton("Actualizar", func() {
		UpdateCardsSucursales()
	})
}

func GetButtonCreateSucrusal() *widget.Button {
	return widget.NewButtonWithIcon("Agregar", theme.ContentAddIcon(), func() {
		var name string
		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Nombre", Widget: widget.NewEntry()},
			},
		}
		form.Items[0].Widget.(*widget.Entry).OnChanged = func(s string) {
			name = s
		}
		dialog.NewCustomConfirm("Agregar Sucursal", "Agregar", "Cancelar", form, func(b bool) {
			if b {
				sucursal, err := CrearSucursal(name)
				if err != nil {
					dialog.ShowError(err, Window)
				} else {
					dialog.ShowInformation("Informacion", sucursal, Window)
					UpdateCardsSucursales()
				}
			}
		}, Window).Show()
	})
}

func GetCardsContainerSucursales() *fyne.Container {
	return CreateCardsSucursales(GetDataSucursales())
}

func GetTabSucursales() *fyne.Container {
	SucursalesContainer = container.NewHBox(GetCardsContainerSucursales())

	return container.NewBorder(
		container.NewVBox(
			GetButtonUpdateCardsSucursales(),
			GetButtonCreateSucrusal()),
		nil, nil, nil, SucursalesContainer)
}

func GetDataSucursales() [][]string {
	datos, err := EnlistarSucursales()
	if err != nil {
		println(err.Error())
	}
	return datos
}

func CreateCardsSucursales(datos [][]string) *fyne.Container {
	sucursalesContainer := container.NewHBox()

	if len(datos) == 0 {
		return sucursalesContainer
	}

	for _, data := range datos {
		card := widget.NewCard(fmt.Sprintf("%s (ID: %s)", data[1], data[0]), "", nil)
		card.SetContent(
			container.NewVBox(
				widget.NewButtonWithIcon("Modificar", theme.ContentAddIcon(), func() {
					var name string
					form := &widget.Form{
						Items: []*widget.FormItem{
							{Text: "Nombre", Widget: widget.NewEntry()},
						},
					}
					form.Items[0].Widget.(*widget.Entry).OnChanged = func(s string) {
						name = s
					}
					dialog.NewCustomConfirm("Modificar Sucursal", "Modificar", "Cancelar", form, func(b bool) {
						if b {
							id, err := strconv.Atoi(data[0])
							if err != nil {
								dialog.ShowError(err, Window)
							}
							sucursal, err := EditarSucursal(id, "nombre", name)
							if err != nil {
								dialog.ShowError(err, Window)
							} else {
								dialog.ShowInformation("Informacion", sucursal, Window)
								UpdateCardsSucursales()
							}
						}
					}, Window).Show()
				}),
				widget.NewButtonWithIcon("Eliminar", theme.ContentRemoveIcon(), func() {
					dialog.NewCustomConfirm("Eliminar Sucursal", "Eliminar", "Cancelar", widget.NewLabel(fmt.Sprintf("¿Esta seguro que desea eliminar la sucursal %s?", data[1])), func(b bool) {
						if b {
							sucursal, err := EliminarSucursal(data[1])
							if err != nil {
								dialog.ShowError(err, Window)
							} else {
								dialog.ShowInformation("Informacion", sucursal, Window)
								UpdateCardsSucursales()
							}
						}
					}, Window).Show()
				}),
			))
		sucursalesContainer.Add(card)
	}

	return sucursalesContainer
}

func UpdateCardsSucursales() {
	SucursalesContainer.RemoveAll()
	newCardsContainer := CreateCardsSucursales(GetDataSucursales())
	for _, obj := range newCardsContainer.Objects {
		SucursalesContainer.Add(obj)
	}
	SucursalesContainer.Refresh()
}
