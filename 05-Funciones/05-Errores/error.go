package main

import (
	"errors"
	"fmt"
)

func divison(dividiendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir entre 0")
	} else {
		return dividiendo / divisor, nil
	}

}

func main() {
	resultado, err := divison(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("resultado ", resultado)
	}

}
