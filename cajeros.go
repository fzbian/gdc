package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var CardsContainer *fyne.Container
var ActualSucursalCajeros = 1

func GetSelectSucursal() *widget.Select {
	labelSelect := widget.NewSelect([]string{"SUCURSAL 1", "SUCURSAL 2", "SUCURSAL 3"}, func(s string) {
		switch s {
		case "SUCURSAL 1":
			ActualSucursalCajeros = 1
			UpdateCards()
		case "SUCURSAL 2":
			ActualSucursalCajeros = 2
			UpdateCards()
		case "SUCURSAL 3":
			ActualSucursalCajeros = 3
			UpdateCards()
		}
	})
	labelSelect.SetSelected("SUCURSAL 1")
	return labelSelect
}

func GetButtonUpdateCards() *widget.Button {
	return widget.NewButtonWithIcon("Actualizar", theme.ViewRefreshIcon(), func() {
		UpdateCards()
	})
}

// Acomodar el contenido en la vista
func GetCajerosContainer() *fyne.Container {
	Top := container.NewVBox(GetSelectSucursal(), GetButtonUpdateCards())
	return container.NewBorder(Top, nil, nil, nil, GetCajerosCards(ActualSucursalCajeros))
}

// Funcion para obtener la data de la base de datos
func GetDataCajeros(sucursal_id int) [][]string {
	datos, err := EnlistarCajerosPorSucursalId(sucursal_id)
	if err != nil {
		panic(err)
	}
	return datos
}

func GetCajerosCards(sucursal_id int) *fyne.Container {
	CardsContainer = container.NewHBox()

	dataCajeros := GetDataCajeros(sucursal_id)

	// Verificar si hay datos disponibles
	if len(dataCajeros) == 0 {
		// No hay datos, puedes mostrar un mensaje o realizar alguna acción apropiada
		return CardsContainer
	}

	// Iterar sobre los datos y crear las tarjetas correspondientes y ponerlas en el contenedor
	for _, data := range dataCajeros {
		card := CreateCard(data)
		CardsContainer.Add(card)
	}

	return CardsContainer
}

// FATAL ERROR HERE
func UpdateCards() {
	if ActualSucursalCajeros == 1 {
		TableContainer.Objects[0] = GetTableTransacciones(1)
		TableContainer.Refresh()
	} else if ActualSucursalCajeros == 2 {
		TableContainer.Objects[0] = GetTableTransacciones(2)
		TableContainer.Refresh()
	} else if ActualSucursalCajeros == 3 {
		TableContainer.Objects[0] = GetTableTransacciones(3)
		TableContainer.Refresh()
	}
}

func CreateCard(data []string) *widget.Card {
	card := widget.NewCard(fmt.Sprintf("Cajero #%s", data[0]), "", nil)
	card.SetContent(container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Usuario actual: %s\nSucursal: %s\nSaldo: %s", data[1], data[2], data[3])),
		widget.NewButton("Editar", func() {
			println("Editar")
		},
		),
		widget.NewButton("Eliminar", func() {
			println("Eliminar")
		},
		),
	))

	return card
}
