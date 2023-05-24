package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	result := rand.Intn(100) + 1
	inp := 0
	fmt.Println("adivinhe um numero inteiro entre 1 e 100:")
	s_n := "x"

	n_Tentativas := make(map[int]int)
	tentativa := 1

	for {
		fmt.Scanln(&inp)
		n_Tentativas[tentativa]++
		if inp == result {
			fmt.Println("parabens você ganhou!!")
			fmt.Printf("voce usou %d tentativas\n", n_Tentativas[tentativa])
			fmt.Println("deseja jogar novamente? (s/n)")
			fmt.Scanln(&s_n)
			if s_n == "n" {
				break
			} else {
				tentativa++
				result = rand.Intn(100) + 1
				fmt.Println("adivinhe um numero inteiro entre 1 e 100:")
			}

		} else if inp > result {
			fmt.Println("O resultado é menor, tente denovo:")
		} else if inp < result {
			fmt.Println("O resultado é maior, tente denovo:")
		}
	}
	n_Tentativas_ord := make([]int, 0, len(n_Tentativas))

	for n_index := range n_Tentativas {
		n_Tentativas_ord = append(n_Tentativas_ord, n_index)

	}

	sort.Ints(n_Tentativas_ord)

	for _, num_tent := range n_Tentativas_ord {
		fmt.Printf("Você gastou %d tentativas na tentativa %d \n", n_Tentativas[num_tent], num_tent)

	}

}
