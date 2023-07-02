package utils

import "errors"

func SumarSaldos(entrada_billetes_100, entrada_billetes_50, entrada_billetes_20, entrada_billetes_10, entrada_billetes_5, entrada_billetes_2, entrada_billetes_1, entrada_monedas_1000, entrada_monedas_500, entrada_monedas_200, entrada_monedas_100, entrada_monedas_50 int) (int, error) {
	// Verificar que todas las variables sean mayores o iguales a cero
	if entrada_billetes_100 < 0 || entrada_billetes_50 < 0 || entrada_billetes_20 < 0 || entrada_billetes_10 < 0 || entrada_billetes_5 < 0 || entrada_billetes_2 < 0 || entrada_billetes_1 < 0 || entrada_monedas_1000 < 0 || entrada_monedas_500 < 0 || entrada_monedas_200 < 0 || entrada_monedas_100 < 0 || entrada_monedas_50 < 0 {
		return 0, errors.New("Las entradas no pueden ser negativas")
	}

	// Definir el valor de cada denominación
	valores := map[string]int{
		"billetes_100": 100000,
		"billetes_50":  50000,
		"billetes_20":  20000,
		"billetes_10":  10000,
		"billetes_5":   5000,
		"billetes_2":   2000,
		"billetes_1":   1000,
		"monedas_1000": 1000,
		"monedas_500":  500,
		"monedas_200":  200,
		"monedas_100":  100,
		"monedas_50":   50,
	}

	// Calcular el total de cada denominación y sumarlos
	suma := 0
	for denominacion, valor := range valores {
		switch denominacion {
		case "billetes_100":
			suma += entrada_billetes_100 * valor
		case "billetes_50":
			suma += entrada_billetes_50 * valor
		case "billetes_20":
			suma += entrada_billetes_20 * valor
		case "billetes_10":
			suma += entrada_billetes_10 * valor
		case "billetes_5":
			suma += entrada_billetes_5 * valor
		case "billetes_2":
			suma += entrada_billetes_2 * valor
		case "billetes_1":
			suma += entrada_billetes_1 * valor
		case "monedas_1000":
			suma += entrada_monedas_1000 * valor
		case "monedas_500":
			suma += entrada_monedas_500 * valor
		case "monedas_200":
			suma += entrada_monedas_200 * valor
		case "monedas_100":
			suma += entrada_monedas_100 * valor
		case "monedas_50":
			suma += entrada_monedas_50 * valor
		}
	}

	return suma, nil
}
