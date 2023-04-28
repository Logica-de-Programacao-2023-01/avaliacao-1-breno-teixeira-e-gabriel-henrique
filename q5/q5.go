package q5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	var vogais = "AEIOUaeiou"

	var consoantes = "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma sequência de caracteres: ")
	scanner.Scan()
	sequencia := scanner.Text()
	for _, vogal := range vogais {
		sequencia = strings.ReplaceAll(sequencia, string(vogal), "")

	}
	for _, consoantes := range consoantes {
		sequencia = strings.ReplaceAll(sequencia, string(consoantes), "."+string(consoantes))

	}
	sequencia = strings.ToLower(sequencia)

	return ""
}
