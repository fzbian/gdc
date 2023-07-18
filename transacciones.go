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

// Funciones para crear los componentes de la vista
func GetLabelsTransacciones() *widget.Label {
	return widget.NewLabel("Sucursal: ")
}

// Funcion para crear el select de sucursales
func GetSelectTransacciones() *widget.Select {
	labelSelect := widget.NewSelect([]string{"SUCURSAL 1", "SUCURSAL 2", "SUCURSAL 3"}, func(s string) {
		switch s {
		case "SUCURSAL 1":
			ActualSucursalTransacciones = 1
			UpdateTable()
		case "SUCURSAL 2":
			ActualSucursalTransacciones = 2
			UpdateTable()
		case "SUCURSAL 3":
			ActualSucursalTransacciones = 3
			UpdateTable()
		}
	})
	labelSelect.SetSelected("SUCURSAL 1")
	return labelSelect
}

// Funcion para crear el boton de actualizar
func GetButtonUpdateTransacciones() *widget.Button {
	return widget.NewButtonWithIcon("Actualizar", theme.ViewRefreshIcon(), func() {
		UpdateTable()
	})
}

// Funcion para crear la tabla con la data
func GetTableTransacciones(sucursal_id int) *widget.Table {
	return CreateTable(GetDataTable(sucursal_id))
}

// Funcion para crear el contenedor principal
func GetTabTransacciones() *fyne.Container {

	TableContainer = container.NewMax(GetTableTransacciones(1))

	MainContainer := container.NewBorder(container.NewHBox(
		GetLabelsTransacciones(),
		GetSelectTransacciones(),
		GetButtonUpdateTransacciones()), nil, nil, nil, TableContainer)

	return MainContainer
}

// Funcion para obtener la data de la base de datos
func GetDataTable(sucursal_id int) [][]string {
	datos, err := ElistarTransaccionesPorSucursal(sucursal_id)
	if err != nil {
		panic(err)
	}
	return datos
}

// Funcion para crear la tabla con la data
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
	if ActualSucursalTransacciones == 1 {
		TableContainer.Objects[0] = GetTableTransacciones(1)
		TableContainer.Refresh()
	} else if ActualSucursalTransacciones == 2 {
		TableContainer.Objects[0] = GetTableTransacciones(2)
		TableContainer.Refresh()
	} else if ActualSucursalTransacciones == 3 {
		TableContainer.Objects[0] = GetTableTransacciones(3)
		TableContainer.Refresh()
	}
}
