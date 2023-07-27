package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var CardsContainer = container.NewHBox()
var ActualSucursalCajeros = 1

func GetSelectSucursalesCajeros() *widget.Select {
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
				ActualSucursalCajeros = sucursal
			}
		}
		UpdateCardsCajeros()
	})

	// Set the default value
	labelSelect.SetSelected(sucursalesNombres[ActualSucursalCajeros-1])
	// Return the select widget
	return labelSelect
}

func GetButtonUpdateCardsCajeros() *widget.Button {
	return widget.NewButtonWithIcon("Actualizar", theme.ViewRefreshIcon(), func() {
		UpdateCardsCajeros()
	})
}

func GetTopContainerCajeros() *fyne.Container {
	DataModify := container.NewHBox(
		widget.NewLabel("Modificar"),
		widget.NewSeparator(),
		widget.NewButton("Agregar", func() {
			var sucursal = 1

			sucursales, err := EnlistarSucursales()
			if err != nil {
				println(err.Error())
			}
			var sucursalIDsMasNombre []string
			for _, sucursal := range sucursales {
				sucursalIDsMasNombre = append(sucursalIDsMasNombre, fmt.Sprintf("%s - %s", sucursal[0], sucursal[1]))
			}

			form := &widget.Form{
				Items: []*widget.FormItem{
					{Text: "Sucursal", Widget: widget.NewSelect(sucursalIDsMasNombre, func(s string) {
						for _, forsucursal := range sucursales {
							if fmt.Sprintf("%s - %s", forsucursal[0], forsucursal[1]) == s {
								sucursalId, err := strconv.Atoi(forsucursal[0])
								if err != nil {
									println(err.Error())
								}
								sucursal = sucursalId
							}
						}
					},
					)},
				},
			}
			dialog := dialog.NewCustomConfirm("Agregar Cajero", "Agregar", "Cancelar", form, func(b bool) {
				if b {
					result, err := CrearCajero(sucursal)
					if err != nil {
						println(err.Error())
						dialogError := dialog.NewError(err, Window)
						dialogError.Show()
					}
					dialogInfo := dialog.NewInformation("Informacion", result, Window)
					dialogInfo.Show()
					UpdateCardsCajeros()
				}
			}, Window)
			dialog.Show()
		}),
	)
	ContainerModify := container.NewHBox(
		widget.NewLabel("Sucursal: "),
		GetSelectSucursalesCajeros(),
		GetButtonUpdateCardsCajeros(),
	)
	return container.NewVBox(
		DataModify,
		ContainerModify,
	)
}

func GetCardsCajeros(sucursal_id int) *fyne.Container {
	return CreateCardsCajeros(GetDataCards(sucursal_id))
}

func GetTabCajeros() *fyne.Container {

	CardsContainer = container.NewHBox(GetCardsCajeros(ActualSucursalCajeros))

	MainContainer := container.NewBorder(
		GetTopContainerCajeros(), nil, nil, nil, CardsContainer)

	return MainContainer
}

func GetDataCards(sucursal_id int) [][]string {
	datos, err := EnlistarCajerosPorSucursalId(sucursal_id)
	if err != nil {
		panic(err)
	}
	return datos
}

func CreateCardsCajeros(datos [][]string) *fyne.Container {
	cardsContainer := container.NewMax()

	if len(datos) == 0 {
		return cardsContainer
	}

	for _, data := range datos {
		currentCajeroData := data // Create a local variable to capture the current cajero data for this iteration

		card := widget.NewCard(fmt.Sprintf("Cajero #%s", currentCajeroData[0]), "", nil)
		card.SetContent(container.NewVBox(
			widget.NewLabel(fmt.Sprintf("Usuario actual: %s\nSucursal: %s\nSaldo: %s", currentCajeroData[1], currentCajeroData[2], currentCajeroData[3])),
			widget.NewButton("Editar", func() {
				// Use the local variable currentCajeroData instead of data
				var selectedType, sucursal, usuario string
				var saldo int

				formEdit := &widget.Form{}

				// Convertir currentCajeroData[0] a int
				cajeroID, err := strconv.Atoi(currentCajeroData[0])
				if err != nil {
					println(err.Error())
				}

				dialogEdit := dialog.NewCustomConfirm("Editar transaccion", "Aceptar", "Cancelar", formEdit, func(b bool) {
					if b {
						switch selectedType {
						case "Sucursal":
							sucursalID, err := strconv.Atoi(sucursal)
							if err != nil {
								println(err.Error())
							}
							result, err := EditarCajero(cajeroID, "sucursal_id", sucursalID)
							if err != nil {
								println(err.Error())
								dialogError := dialog.NewError(err, Window)
								dialogError.Show()
							} else {
								dialogInfo := dialog.NewInformation("Informacion", result, Window)
								dialogInfo.Show()
								UpdateCardsCajeros()
							}
						case "Usuario":
							usuarioID, err := strconv.Atoi(usuario)
							if err != nil {
								println(err.Error())
							}
							result, err := EditarCajero(cajeroID, "usuario_id", usuarioID)
							if err != nil {
								println(err.Error())
								dialogError := dialog.NewError(err, Window)
								dialogError.Show()
							} else {
								dialogInfo := dialog.NewInformation("Informacion", result, Window)
								dialogInfo.Show()
								UpdateCardsCajeros()
							}
						case "Saldo":
							result, err := EditarCajero(cajeroID, "saldo", saldo)
							println(fmt.Sprintf("CajeroID: %d\nSaldo: %d", cajeroID, saldo))
							if err != nil {
								println(err.Error())
								dialogError := dialog.NewError(err, Window)
								dialogError.Show()
							} else {
								dialogInfo := dialog.NewInformation("Informacion", result, Window)
								dialogInfo.Show()
								UpdateCardsCajeros()
							}
						}
					} else {
						println("Cancelado")
					}

				}, Window)

				// Enlisrar sucursales para el select
				sucursales, err := EnlistarSucursales()
				if err != nil {
					println(err.Error())
				}
				var sucursalIDsMasNombre []string
				for _, sucursal := range sucursales {
					sucursalIDsMasNombre = append(sucursalIDsMasNombre, fmt.Sprintf("%s - %s", sucursal[0], sucursal[1]))
				}

				// Enlistar usuarios para el select
				usuarios, err := EnlistarUsuarios()
				if err != nil {
					println(err.Error())
				}
				var usuariosNombres []string
				for _, usuario := range usuarios {
					usuariosNombres = append(usuariosNombres, fmt.Sprintf("%s - %s", usuario[0], usuario[2]))
				}

				formTipo := &widget.Form{
					Items: []*widget.FormItem{
						{Text: "Tipo", Widget: widget.NewSelect([]string{"Sucursal", "Usuario", "Saldo"}, func(s string) {
							selectedType = s
							switch selectedType {
							case "Sucursal":
								formEdit.Items = []*widget.FormItem{}
								formEdit.Append("Sucursal", widget.NewSelect(sucursalIDsMasNombre, func(s string) {
									sucursal = string(s[0])
								}))
							case "Usuario":
								formEdit.Items = []*widget.FormItem{}
								formEdit.Append("Usuario", widget.NewSelect(usuariosNombres, func(s string) {
									usuario = string(s[0])
								}))
							case "Saldo":
								formEdit.Items = []*widget.FormItem{}
								formEdit.Append("Saldo", widget.NewEntry())
								// El saldo del entry va a ser igual a la variable saldo
								formEdit.Items[0].Widget.(*widget.Entry).OnChanged = func(s string) {
									saldo, err = strconv.Atoi(s)
									if err != nil {
										println(err.Error())
									}
								}
							}
						})},
					},
				}

				dialogTipo := dialog.NewCustomConfirm("¿Que deseas modificar?", "Siguiente", "Cancelar", formTipo, func(b bool) {
					if b {
						dialogEdit.Show()
					} else {
						dialogEdit.Hide()
					}
				}, Window)

				dialogTipo.Show()
			}),
			widget.NewButton("Eliminar", func() {
				confirm := dialog.NewConfirm("Eliminar", "¿Estás seguro que deseas eliminar este cajero?", func(b bool) {
					if b {
						cajeroID, _ := strconv.Atoi(currentCajeroData[0])
						result, err := EliminarCajero(cajeroID)
						if err != nil {
							println(err.Error())
							dialogError := dialog.NewError(err, Window)
							dialogError.Show()
						}
						dialogInfo := dialog.NewInformation("Informacion", result, Window)
						dialogInfo.Show()
						UpdateCardsCajeros()
					}
				}, Window)
				confirm.Show()
			}),
		))
		cardsContainer.Add(card)
	}

	return cardsContainer
}

func UpdateCardsCajeros() {
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
			// Eliminar las tarjetas anteriores del contenedor
			CardsContainer.RemoveAll()

			// Obtener las nuevas tarjetas y agregarlas al contenedor
			newCardsContainer := CreateCardsCajeros(GetDataCards(ActualSucursalCajeros))
			for _, obj := range newCardsContainer.Objects {
				CardsContainer.Add(obj)
			}

			CardsContainer.Refresh()
			break // Rompemos el bucle después de actualizar para evitar problemas
		}
	}
}

func ShowObjectsContainer(cont *fyne.Container) {
	fmt.Println("Objetos del contenedor:")
	objects := cont.Objects
	for _, obj := range objects {
		fmt.Println("-------------------------")
		switch w := obj.(type) {
		case *widget.Card:
			fmt.Printf("Tipo: *widget.Card\n")
			fmt.Printf("Título: %s\n", w.Title)
			// Puedes agregar más detalles sobre el card si es necesario
		case *fyne.Container:
			fmt.Printf("Tipo: *container.VBox\n")
			// Puedes agregar más detalles sobre el VBox si es necesario
		default:
			fmt.Printf("Tipo: %T\n", obj)
			// Puedes agregar más detalles sobre otros tipos de objetos aquí si es necesario
		}
		fmt.Println("-------------------------")
	}
}
