package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetCajerosContainer() *fyne.Container {
	MainContainer := container.NewVBox()

	datos, _ := EnlistarCajeros()
	for _, fila := range datos {
		card := CreateCard(fila)
		MainContainer.Add(card)
	}

	return MainContainer
}

func CreateCard(data []string) *widget.Card {
	card := widget.NewCard("", "", nil)

	// Verificar si data tiene al menos 2 elementos
	if len(data) >= 2 {
		card.Title = fmt.Sprintf("Cajero %s", data[0])
		card.Subtitle = fmt.Sprintf("Usuario actual: %s\nSucursal: %s\nSaldo: %s", data[1], data[2], data[3])
	}

	return card
}
