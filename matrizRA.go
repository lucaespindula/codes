package main

import (
	"fmt"
)

func main() {
	var alunos [3][3]string

	for i := 0; i < 3; i++ {
		fmt.Printf("Digite o índice do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][0])

		fmt.Printf("Digite o RA do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][1])

		fmt.Printf("Digite o nome do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][2])
	}

	fmt.Println("Dados dos Alunos:\nÍndice  RA        Nome")
	for i := 0; i < 3; i++ {
		fmt.Printf("%s     %s     %s\n", alunos[i][0], alunos[i][1], alunos[i][2])
	}
}
