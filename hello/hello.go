package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	/*Há algumas maneiras de você declarar variáveis em GO e uma delas é utilizando a tipagem juntamente ao VAR;
	Ex: var nome string = "Meu nome é Gabriel" - Essa é a maneira mais explícita de declaração, sendo necessário declarar
	que é uma variável e seu tipo.

	Outra maneira é a declaração que o próprio GO trata de referenciar a tipagem, EX: var nome = "Meu nome é Gabriel",
	neste modo de declaração não é necessário passar o tipo pois o próprio GO ficará responsável por definir a tipagem de acordo com o valor que você preencheu.

	Outra maneira que também existe é a declaração de variável curta, não sendo necessário explicitar que é uma variável e nem a tipagem, EX: nome := "Meu nome é Gabriel".
	*/
	showIntroduce()

	showMenu()

	//O & na frente serve para buscar o ponteiro. É possível ver diretamente o ponteiro através do &command, ex: fmt.Println("O endereço é:", &command)
	//Neste caso do ScanF é passado o ponteiro da variável para passar o valor diretamente neste ponteiro.
	//fmt.Scanf("%d", &command)

	//Na função Scan você não precisa esperar o tipo do valor que você espera, no caso o %d, que estaria esperando um do tipo int, pois já
	//entende que a variável é do tipo INT

	command := readCommand()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Showing logs...")
	// } else if command == 3 {
	// 	fmt.Println("Exiting programm...")
	// } else {
	// 	fmt.Println("Idk the command...")
	// }

	//Em GOLANG não precisa declarar o break. O compilador não irá reclamar, mas não é necessário.
	switch command {
	case 1:
		initializeMonitoring()
	case 2:
		fmt.Println("Showing logs...")
	case 3:
		exitProgramm()
	default:
		fmt.Println("Idk the command...")
		os.Exit(-1)
	}
}

func showIntroduce() {
	name, _ := showNamesAndAge() // A utilização do "_" serve para declarar uma variável que não tenho interesse em usar em funções de múltiplos valores.
	version := 1.0

	fmt.Println("Hello Word, sr", name)
	fmt.Println("The version program", version)
}

// No GOLANG você declara o retorno de uma função depois dos parâmetros.
func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}

func showMenu() {

	fmt.Println("------------- / ------------- /")
	fmt.Println("1 - Initialize monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit programm")

}

func exitProgramm() {
	fmt.Println("Exiting programm...")
	os.Exit(0)
}

func initializeMonitoring() {
	fmt.Println("Monitoring...")

	site := "https://www.alura.com.br"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("O site ", site, "está online com sucesso!")
	} else {
		fmt.Println("O site ", site, "está com problema!")
	}
}

// Para função de múltiplos retornos, você deve declarar os 2 tipos, exemplo:
// E para receber este retorno, deve-se respectivamente esperar os seus tipos, ex: name, age := showNamesAndAge()
func showNamesAndAge() (string, int) {
	return "Gabriel", 24
}
