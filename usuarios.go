package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

var TableUsuarios *fyne.Container

func GetButtonAddUsuarios() *widget.Button {
	return widget.NewButtonWithIcon("Agregar", theme.ContentAddIcon(), func() {
		var usuario, nombre string
		var rango, clave int

		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Usuario", Widget: widget.NewEntry()},
				{Text: "Nombre", Widget: widget.NewEntry()},
				{Text: "Clave", Widget: widget.NewEntry()},
				{Text: "Rango", Widget: widget.NewEntry()},
			},
		}
		dialogAddUsuario := dialog.NewCustomConfirm("Agregar usuario", "Aceptar", "Cancelar", form, func(b bool) {
			if b {
				usuario = form.Items[0].Widget.(*widget.Entry).Text
				nombre = form.Items[1].Widget.(*widget.Entry).Text
				clave, _ = strconv.Atoi(form.Items[2].Widget.(*widget.Entry).Text)
				rango, _ = strconv.Atoi(form.Items[3].Widget.(*widget.Entry).Text)

				if usuario == "" || nombre == "" {
					dialog.NewError(errors.New("Los campos no pueden ser vacios"), Window).Show()
				} else if clave == 0 || rango == 0 {
					dialog.NewError(errors.New("Los campos no pueden ser ceros"), Window).Show()
				} else if rango < 1 || rango > 3 {
					dialog.NewError(errors.New("El rango no es correcto"), Window).Show()
				} else if clave < 1 || clave > 9999 {
					dialog.NewError(errors.New("La clave tiene que ser de 4 digitos"), Window).Show()
				} else {
					usuario, err := CrearUsuario(usuario, nombre, clave, rango)
					if err != nil {
						dialog.NewError(err, Window).Show()
					} else {
						dialog.NewInformation("Agregado", usuario, Window).Show()
						UpdateTableUsuarios()
					}
				}
			}
		}, Window)
		dialogAddUsuario.Show()
	})
}

func GetButtonUpdateUsuarios() *widget.Button {
	return widget.NewButtonWithIcon("Actualizar", theme.ViewRefreshIcon(), func() {
		UpdateTableUsuarios()
	})
}

func GetTabUsuarios() *fyne.Container {

	TableUsuarios = container.NewMax(GetTableUsuarios())

	MainContainer := container.NewBorder(container.NewHBox(
		GetButtonUpdateUsuarios(), GetButtonAddUsuarios()), nil, nil, nil, TableUsuarios)

	return MainContainer

}

func GetTableUsuarios() *widget.Table {
	return CreateTableUsuarios(GetDataTableUsuarios())
}

func GetDataTableUsuarios() [][]string {
	usuarios, err := EnlistarUsuarios()
	if err != nil {
		println(err.Error())
	}

	return usuarios
}

func CreateTableUsuarios(data [][]string) *widget.Table {

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewMax(widget.NewLabel("template11"), widget.NewButton("template12", func() {}))
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*fyne.Container).Objects[0].(*widget.Label)
			button := obj.(*fyne.Container).Objects[1].(*widget.Button)
			label.Show()
			button.Hide()

			switch i.Col {
			case 0, 1, 2, 3, 4, 5:
				label.SetText(data[i.Row][i.Col])
			case 6:
				label.Hide()
				button.Show()
				button.SetText("Editar")
				button.OnTapped = func() {
					// TODO: si se edita el rango del mismo usuario que esta logeado, no permitir que se edite el rango
					var tipo string
					var valor any

					form := &widget.Form{
						Items: []*widget.FormItem{
							{Text: "Tipo", Widget: widget.NewSelect([]string{"Usuario", "Nombre", "Clave", "Rango"}, func(s string) {
								tipo = strings.ToLower(s)
							})},
							// el widget entry va a ser igual al la variable valor
							{Text: "Valor", Widget: widget.NewEntry()},
						},
					}
					dialogEditarUsuario := dialog.NewCustomConfirm("Editar usuario", "Aceptar", "Cancelar", form, func(b bool) {
						if b {
							id, err := strconv.Atoi(data[i.Row][0])
							if err != nil {
								println(err.Error())
							}

							if tipo == "usuario" || tipo == "nombre" {
								valor = form.Items[1].Widget.(*widget.Entry).Text
								if valor == "" {
									dialog.NewError(errors.New("El valor no puede ser vacio"), Window).Show()
								} else {
									usuario, err := EditarUsuario(id, tipo, valor)
									if err != nil {
										return
									}
									dialog.NewInformation("Actualizado", usuario, Window).Show()
									UpdateTableUsuarios()
								}
							} else if tipo == "rango" || tipo == "clave" {
								valor, err := strconv.Atoi(form.Items[1].Widget.(*widget.Entry).Text)
								if err != nil {
									dialog.NewError(errors.New("El valor no puede ser texto"), Window).Show()
								} else {
									usuario, err := EditarUsuario(id, tipo, valor)
									if err != nil {
										return
									}
									dialog.NewInformation("Actualizado", usuario, Window).Show()
									UpdateTableUsuarios()
								}
							}
						}
					}, Window)
					dialogEditarUsuario.Show()
				}
			case 7:
				label.Hide()
				button.Show()
				button.SetText("Eliminar")
				button.OnTapped = func() {
					confirm := dialog.NewConfirm("Eliminar", "¿Desea eliminar este registro?", func(b bool) {
						if b {
							result, err := EliminarUsuario(data[i.Row][1])
							if err != nil {
								dialog.NewError(err, Window).Show()
							}
							dialog.NewInformation("Eliminado", result, Window).Show()
							UpdateTableUsuarios()
						}
					}, Window)
					confirm.Show()
				}
			}

			switch i.Row {
			case 0:
				button.Hide()
				label.Show()
				label.SetText(data[i.Row][i.Col])
			}
		})

	table.SetColumnWidth(0, 30)
	table.SetColumnWidth(1, 65)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 60)
	table.SetColumnWidth(4, 60)
	table.SetColumnWidth(5, 140)

	return table
}

func UpdateTableUsuarios() {
	TableUsuarios.Objects[0] = GetTableUsuarios()
	TableUsuarios.Refresh()
}
