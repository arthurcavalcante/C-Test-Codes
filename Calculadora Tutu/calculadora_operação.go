package main

import (
	"fmt"
	"strings"
	)

var MatOperacoes := map[string]Operacoes {
	"sum": Sum{},
	"+": Sum{},

	"sub": Sub{},
	"-": Sub{},
	
	"mul": Mul{},
	"*": Mul{},

	"div": Div{},
	"/": Div{},


	"pow": Pow{},
	"^": Pow{},

	"rot": Rot{},
	"&": Rot{},
	
} 

func coleta_dados(MatOperacoes) (Float32, err) {

	UserImput := UserImputOperation("Operacao: ")
	
	if (UserImput == nil) {
 		fmt.println("Vazio, coloque alguma coisa nos numeros da operação")
		return
	}

}	

func UserImputAsString() (string) {
	var UserImput string
	fmt.Scan(&UserImput)
	return UserImput
}
	
func SeparateString(UserImput) []string {

	var sep []string

	for key, _ := range MatOperacoes {

		for strings.Contains(UserImput, key) {

			sep = strings.Split(op,key)
			sep = append(sep,key)
			return sep

		}
	}
	return nil
}

func Calculo_Converter(sep) (float32, string) {

	Float = strconv.ParseFloat(sep[0], 32)
	Operando_01 = float32(Float)

	Float1 = strconv.ParseFloat(sep[2], 32)
	Operando_02 = float32(Float)

	Operador = sep[1]

}




