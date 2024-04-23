# Exemplo de uso do módulo auxiliar

Este documento fornece um guia rápido sobre como usar o módulo `helpers` dentro de uma aplicação Go, demonstrando as funcionalidades com a estrutura `Octopus` do pacote `test`.

## Configurar

Certifique-se de ter o módulo `helpers` instalado e importado corretamente em seu projeto. 
```shell
go get github.com/flambra/helpers
```
O exemplo a seguir usa a estrutura `Octopus` do subpacote `test`.

## Programa principal

Abaixo está um exemplo completo de um programa Go que usa a estrutura `Octopus` para manipular e exibir dados do Octopus:

```go
package main

import (
	"fmt"
	"github.com/flambra/helpers/test"
)

func main() {
	// Create an instance of Octopus
	oct := test.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}

	// Display the initial state of the octopus
	fmt.Println(oct.String())

	// Reset the octopus's name and color
	oct.Reset()

	// Display the state of the octopus after reset
	fmt.Println(oct.String())

    // Prints a message to show that the helpers module is working
	test.Print()
}
```
