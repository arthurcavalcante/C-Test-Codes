package main

import (
	"fmt"
)

func main() {

	var Coleta_1, Coleta_2, Coleta_3, Coleta_4 int
	inicio_matriz(Coleta_1, Coleta_2, Coleta_3, Coleta_4)

}

func inicio_matriz(Coleta_1, Coleta_2, Coleta_3, Coleta_4 int) {

	var Opcao int

	fmt.Printf("|=========================|\n")
	fmt.Printf("|===| Programa Matriz |===|\n")
	fmt.Printf("|=========================|\n")
	fmt.Printf("|==| Selecione a Opção |==|\n")
	fmt.Printf("|=========================|\n")
	fmt.Printf("|==|   1º Ler matriz   |==|\n")
	fmt.Printf("|==| 2º Imprimir matriz|==|\n")
	fmt.Printf("|==|3º matriz simetrica|==|\n")
	fmt.Printf("|==|4º mult matrizes   |==|\n")
	fmt.Printf("|=========================|\n")
	fmt.Println("| Digite a opção:         |")
	fmt.Printf("|=========================|\n")
	fmt.Scan(&Opcao)

	if Opcao == 1 {
		ler_matriz(Coleta_1, Coleta_2, Coleta_3, Coleta_4)
	} else if Opcao == 2 {
		imprimir_matriz(matriz, matriz2, Coleta_1, Coleta_2, Coleta_3, Coleta_4)
	}
}

func ler_matriz(Coleta_1, Coleta_2, Coleta_3, Coleta_4 int) {

	var Coleta_Matriz, Coleta_Matriz2, Qtde_matriz int

	for {
		fmt.Println("Quantas matrizes você deseja criar? Máximo 2!")
		fmt.Scan(&Qtde_matriz)

		if Qtde_matriz > 2 || Qtde_matriz < 0 {
			fmt.Printf("Quantidade de matrizes inválida, digite novamente!\n")
		} else {
			break
		}
	}

	if Qtde_matriz == 1 {

		fmt.Println("Digite o número de linhas da matriz:")
		fmt.Scan(&Coleta_1)
		fmt.Println("Digite o número de colunas da matriz:")
		fmt.Scan(&Coleta_2)

		matriz := make([][]int, Coleta_1)
		for y := range matriz {
		matriz[y] = make([]int, Coleta_2)
		}

		for i := 0; i < Coleta_1; i++ {
			for j := 0; j < Coleta_2; j++ {
				fmt.Printf("Insira o valor da linha %d e da coluna %d: ", i+1, j+1)
				fmt.Scan(&Coleta_Matriz)
				matriz[i][j] = Coleta_Matriz
			}
		}
	} else if Qtde_matriz == 2 {

		fmt.Println("Digite o número de linhas da 1ª matriz:")
		fmt.Scan(&Coleta_1)
		fmt.Println("Digite o número de colunas da 1ª matriz:")
		fmt.Scan(&Coleta_2)
		fmt.Println("Digite o número de linhas da 2ª matriz:")
		fmt.Scan(&Coleta_3)
		fmt.Println("Digite o número de colunas da 2ª matriz:")
		fmt.Scan(&Coleta_4)
		
		matriz := make([][]int, Coleta_1)
		for y := range matriz {
		matriz[y] = make([]int, Coleta_2)
		}

		matriz2 := make([][]int, Coleta_3)
		for f := range matriz2 {
		matriz2[f] = make([]int, Coleta_4)
		}

		for p := 0; p < Coleta_1; p++ {
			for q := 0; q < Coleta_2; q++ {
				fmt.Printf("Insira o valor da linha %d e da coluna %d da 1ª matriz: ", p+1, q+1)
				fmt.Scan(&Coleta_Matriz)
				matriz[p][q] = Coleta_Matriz
			}
		}

		for k := 0; k < Coleta_3; k++ {
			for l := 0; l < Coleta_4; l++ {
				fmt.Printf("Insira o valor da linha %d e da coluna %d da 2ª matriz: ", k+1, l+1)
				fmt.Scan(&Coleta_Matriz2)
				matriz2[k][l] = Coleta_Matriz2
			}
		}
	}
}

func imprimir_matriz(matriz[][]int, matriz2[][]int, Coleta_1, Coleta_2, Coleta_3, Coleta_4 int) {

	for i := 0; i < Coleta_1; i++ {
		for j := 0; j < Coleta_2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
	}

	for k := 0; k < Coleta_3; k++ {
		for l := 0; l < Coleta_4; l++ {
			fmt.Printf("%d", matriz2[k][l])
		} 
	}
}
