package main

import (
	"errors"
	"fmt"
)

func main() {
	const C string = "Celcius"
	const R string = "Reamur"
	const F string = "Fahreinheit"
	const K string = "Kelvin"

	result, err := convertTemp(30, C, F)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hasil Konversi adalah", result)
	}
	fmt.Println(C, R, F, K)
}

func convertTemp(temp int, originUnit string, destinationUnit string) (int, error) {
	fmt.Println(temp)
	var result int
	switch originUnit {
	case "Celcius":
		
		switch destinationUnit {
		case "Celcius":
			return 0, errors.New("anda mengkonversi suhu yang sama yaitu Celcius")
		case "Reamur":
			result = temp * 4 / 5
			return result, nil
		case "Fahreinheit":
			result = temp * 9 / 5 + 32
			return result, nil
		case "Kelvin":
			result = temp + 273
			return result, nil
		default:
			return 0, errors.New("destination unit anda tidak valid")
		}

	case "Reamur":

		switch destinationUnit {
		case "Celcius":
			result = temp * 5 / 4
			return result, nil
		case "Reamur":
			return 0, errors.New("anda mengkonversi satuan suhu yang sama yaitu Reamur")
		case "Fahreinheit":
			result = temp * 9 / 4 + 32
			return result, nil
		case "Kelvin":
			result = temp * 5 / 4 + 273
			return result, nil
		default:
			return 0, errors.New("destination unit anda tidak valid")
		}

	case "Fahreinheit":

		switch destinationUnit {
		case "Celcius":
			result = (temp - 32) * 5 / 9
			return result, nil
		case "Reamur":
			result = (temp - 32) * 4 / 9
			return result, nil
		case "Fahreinheit":
			return 0, errors.New("anda mengkonversi satuan suhu yang sama yaitu Fahreinheit")
		case "Kelvin":
			result = (temp + 459) * 5 / 9
			return result, nil
		default:
			return 0, errors.New("destination unit anda tidak valid")
		}

	case "Kelvin":

		switch destinationUnit {
		case "Celcius":
			result = temp - 273
			return result, nil
		case "Reamur":
			result = (temp - 273) * 4 / 5
			return result, nil
		case "Fahreinheit":
			result = (temp * 9) / 5 - 459;
			return result, nil
		case "Kelvin":
			return 0, errors.New("anda mengkonversi satuan suhu yang sama yaitu Kelvin")
		default:
			return 0, errors.New("destination unit anda tidak valid")
		}

	default: 
		return 0, errors.New("origin unit anda tidak valid")
	} 
}