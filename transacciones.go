package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var TableContainer *fyne.Container
var ActualSucursalTransacciones = 1

func GetSelectSucursalesTransacciones() *widget.Select {
	sucursales, err := EnlistarSucursales()
	if err != nil {
		println(err.Error())
	}

	var sucursalesNombres []string
	for _, sucursal := range sucursales {
		sucursalesNombres = append(sucursalesNombres, sucursal[1])
	}

	labelSelect := widget.NewSelect(sucursalesNombres, func(s string) {
		for _, sucursal := range sucursales {
			if sucursal[1] == s {
				sucursal, err := strconv.Atoi(sucursal[0])
				if err != nil {
					println(err.Error())
				}
				ActualSucursalTransacciones = sucursal
			}
		}
		UpdateTable()
	})

	labelSelect.SetSelected(sucursalesNombres[0])
	return labelSelect
}

func GetButtonUpdateTransacciones() *widget.Button {
	return widget.NewButtonWithIcon("Actualizar", theme.ViewRefreshIcon(), func() {
		UpdateTable()
	})
}

func GetTableTransacciones(sucursal_id int) *widget.Table {
	return CreateTable(GetDataTable(sucursal_id))
}

func GetTabTransacciones() *fyne.Container {

	TableContainer = container.NewMax(GetTableTransacciones(1))

	MainContainer := container.NewBorder(container.NewHBox(
		widget.NewLabel("Sucursal: "),
		GetSelectSucursalesTransacciones(),
		GetButtonUpdateTransacciones()), nil, nil, nil, TableContainer)

	return MainContainer
}

func GetDataTable(sucursal_id int) [][]string {
	datos, err := ElistarTransaccionesPorSucursal(sucursal_id)
	if err != nil {
		panic(err)
	}
	return datos
}

func CreateTable(data [][]string) *widget.Table {

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewMax(widget.NewLabel("template11"), widget.NewButton("template12", func() {}))
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			l := obj.(*fyne.Container).Objects[0].(*widget.Label)
			b := obj.(*fyne.Container).Objects[1].(*widget.Button)
			l.Show()
			b.Hide()
			switch i.Col {
			case 0, 1, 2, 3, 5, 6, 7, 8:
				l.SetText(data[i.Row][i.Col])
			case 4:
				if data[i.Row][4] == "Sin descripcion" {
					l.SetText("Sin descripcion")
				} else {
					l.Hide()
					b.Show()
					b.SetText("Ver")
					b.OnTapped = func() {
						dialog := dialog.NewInformation("Descripcion", data[i.Row][4], Window)
						dialog.Show()
					}
				}
			case 9:
				l.Hide()
				b.Show()
				b.OnTapped = func() {
					var tipo string
					var valor any
					form := &widget.Form{
						Items: []*widget.FormItem{
							{Text: "Tipo", Widget: widget.NewSelect([]string{"descripcion", "valor"}, func(s string) { tipo = s })},
							{Text: "Valor", Widget: widget.NewEntry()},
						},
					}
					dialog := dialog.NewCustomConfirm("Editar transaccion", "Aceptar", "Cancelar", form, func(b bool) {
						id, err := strconv.Atoi(data[i.Row][1])
						if err != nil {
							panic(err)
						}

						if tipo == "descripcion" {
							valor = form.Items[1].Widget.(*widget.Entry).Text
							EditarTransaccion(1, id, tipo, valor)
							UpdateTable()
						} else if tipo == "valor" {
							// Si el valor no es un numero
							valor, err := strconv.Atoi(form.Items[1].Widget.(*widget.Entry).Text)
							if err != nil {
								dialog := dialog.NewInformation("Error", "El valor debe ser un numero", Window)
								dialog.Show()
								return
							}
							EditarTransaccion(1, id, tipo, valor)
							UpdateTable()
						}
					}, Window)
					dialog.Show()
				}
				b.SetText(data[i.Row][i.Col])
			case 10:
				l.Hide()
				b.Show()
				b.SetText(data[i.Row][i.Col])
				b.OnTapped = func() {
					confirm := dialog.NewConfirm("Eliminar transaccion", "¿Estás seguro que deseas eliminar esta transacción?", func(ok bool) {
						if ok {
							id, err := strconv.Atoi(data[i.Row][0])
							if err != nil {
								panic(err)
							}
							EliminarTransaccion(id)
							UpdateTable()
						}
					}, Window)
					confirm.Show()
				}
			}
			switch i.Row {
			case 0:
				b.Hide()
				l.Show()
				l.SetText(data[i.Row][i.Col])
			}
		})

	table.SetColumnWidth(0, 28)
	table.SetColumnWidth(1, 55)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 105)
	table.SetColumnWidth(4, 105)
	table.SetColumnWidth(5, 75)
	table.SetColumnWidth(6, 120)
	table.SetColumnWidth(7, 150)
	table.SetColumnWidth(8, 150)
	table.SetColumnWidth(9, 105)
	table.SetColumnWidth(10, 105)

	return table
}

func UpdateTable() {
	sucursales, err := EnlistarSucursales()
	if err != nil {
		println(err.Error())
	}

	var sucursalesIds []int
	for _, sucursal := range sucursales {
		sucursal, _ := strconv.Atoi(sucursal[0])
		sucursalesIds = append(sucursalesIds, sucursal)
	}

	for _, sucursal := range sucursalesIds {
		if ActualSucursalTransacciones == sucursal {
			TableContainer.Objects[0] = GetTableTransacciones(sucursal)
			TableContainer.Refresh()
		}
	}
}
