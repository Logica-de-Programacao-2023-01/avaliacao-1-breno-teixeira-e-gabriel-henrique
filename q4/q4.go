package q4

import "fmt"

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	tamanho := len(prices)
	if tamanho == 0 {
		return 0, fmt.Errorf("lista vazia")
	}
	if tamanho == 1 {
		return 3, nil
	}

	crescente := true
	descresc := true

	for i := 1; i < tamanho; i++ {
		if prices[i] > prices[i-1] {
			crescente = false
		} else if prices[i] < prices[i-1] {
			descresc = false
		}

	}

	if crescente {
		return 1, nil
	} else if descresc {
		return 2, nil
	} else {
		return 3, nil
	}

	return 0, nil
}
