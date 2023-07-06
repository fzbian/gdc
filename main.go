package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"julio/controllers/transacciones"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Mi Aplicación")
	window.Resize(fyne.NewSize(393, 851))

	tabs := container.NewAppTabs()
	tabs.SetTabLocation(container.TabLocationTop)

	data, err := transacciones.EnlistarTransacciones()
	if err != nil {
		println(err.Error())
	}
	//fmt.Println(len(data))

	sucursalesLabel := widget.NewLabel("Sucursal: ")
	sucursalesSelect := widget.NewSelect([]string{"SUCURSAL 1", "SUCURSAL 2", "SUCURSAL 3"}, func(s string) {})

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("TableCell")
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	// Container for select
	mainContainer := container.NewBorder(container.NewVBox(
		sucursalesLabel, sucursalesSelect), nil, nil, nil, table)

	tabs.Append(container.NewTabItemWithIcon("Transacciones", theme.StorageIcon(), mainContainer))
	tabs.Append(container.NewTabItemWithIcon("Cajeros", theme.ComputerIcon(), widget.NewLabel("Cajeros")))
	tabs.Append(container.NewTabItemWithIcon("Usuarios", theme.AccountIcon(), widget.NewLabel("Usuarios")))
	tabs.Append(container.NewTabItemWithIcon("Caja Menor", theme.GridIcon(), widget.NewLabel("Caja Menor")))
	tabs.Append(container.NewTabItemWithIcon("Ajustes", theme.SettingsIcon(), widget.NewLabel("Ajustes")))

	window.SetContent(tabs)
	window.ShowAndRun()
}

//func ConvertTo2DStringSlice(transacciones [][]models.Transacciones) [][]string {
//	// Crear la matriz bidimensional
//	result := make([][]string, len(transacciones)+1)
//
//	// Obtener los nombres de los campos como encabezados
//	header := []string{"ID", "UsuarioID", "CajeroID", "Tipo", "Descripcion", "Valor", "FechaCreacion"}
//	result[0] = header
//
//	// Recorrer cada transacción y extraer los valores de los campos en forma de strings
//	for i, t := range transacciones {
//		row := make([]string, len(header))
//		row[0] = strconv.Itoa(t.ID)
//		row[1] = strconv.Itoa(t.UsuarioID)
//		row[2] = strconv.Itoa(t.CajeroID)
//		row[3] = t.Tipo
//		row[4] = t.Descripcion
//		row[5] = strconv.Itoa(t.Valor)
//		row[6] = t.FechaCreacion.String()
//
//		result[i+1] = row
//	}
//
//	return result
//}
