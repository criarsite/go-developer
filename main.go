package main

import (
	"fmt"
	"modulo/projetos"
)

var escopoGlobal = "Variavel de escopo global"

func main() {

	//fundamentos()

	var temperaturaDeEbulicacao = projetos.Conversor()
	fmt.Println(temperaturaDeEbulicacao)
}

func fundamentos() {
	fmt.Println("Yauh Seja Louvado")

	//tipos numericos com tamanhos definidos
	var a uint8 = 255                 //uint8: numeros inteiros de 0 a 255
	var b int16 = 32767               //int16: numeros inteiros de -32768 a 32767
	var c int32 = 2147483647          //int32: numeros inteiros de -2147483648 a 2147483647
	var d int64 = 9223372036854775807 //int64: numeros inteiros de -9223372036854775808 a 9223372036854775807

	var e uint8 = 255                 //uint8: numeros inteiros de 0 a 255
	var f int16 = 32767               //int16: numeros inteiros de -32768 a 32767
	var g int32 = 2147483647          //int32: numeros inteiros de -2147483648 a 2147483647
	var h int64 = 9223372036854775807 //int64: numeros inteiros de -9223372036854775808 a 9223372036854775807

	fmt.Println("O valor de a `:", a)
	fmt.Println("O valor de b `:", b)
	fmt.Println("O valor de c `:", c)
	fmt.Println("O valor de d `:", d)
	fmt.Println("O valor de e `:", e)
	fmt.Println("O valor de f `:", f)
	fmt.Println("O valor de g `:", g)
	fmt.Println("O valor de h `:", h)

	var ab float32 = 3.14
	var bc float64 = 3.14159265358979323846

	fmt.Println("O valor de ab `:", ab)
	fmt.Println("O valor de bc `:", bc)

	//Booleanos  operadores logicos que vao resultar verdadeiro ou falso
	/*
		&& - AND  E
		|| - OR   OU
		! - NOT   Diferente
	*/
	var x bool = true
	var y bool = false
	fmt.Println("O valor de x `:", x)
	fmt.Println("O valor de y `:", y)

	//tipo string: Cadeia de caracteres
	var z string = "Hello, World!"
	fmt.Println("O valor de z `:", z)

	//retorna o tamanho da string
	fmt.Println("O tamanho de z `:", len(z))
	fmt.Println(z[2]) //retorna o valor do caracter na posição 2 da string

	println("top demais")

	//Inferencia de tipos é a capacidade do go identificar o tipo da variavel
	var nun = 10         //o go identifica que a variavel é do tipo int
	var nome = "Yausha"  //o go identifica que a variavel é do tipo string
	var flutuante = 10.4 // o go identifica que a variavel é do tipo float64

	fmt.Println("O valor de a `:", nun)
	fmt.Println("O valor de nome `:", nome)
	fmt.Println("O valor de flutuante `:", flutuante)

	//Palavras reservadas são palavras que não podem ser usadas como identificadores, pois são usadas pelo compilador para identificar a sintaxe da linguagem. Exemplo: func, var, package, import, return, if, else, for, switch, case, default, break, continue, goto, defer, go, select, struct, interface, map, chan, range.
	/*
				        Em Go (Golang), as palavras reservadas podem ser organizadas por grupos de funcionalidades, como controle de fluxo, declaração de funções, tipos de dados, variáveis e concorrência.

		Abaixo está a organização correta, com algumas correções nas descrições.

		## 1. Controle de fluxo e decisões

		| Palavra       | Função                                                                                   |
		| ------------- | ---------------------------------------------------------------------------------------- |
		| `if`          | Avalia uma condição e executa um bloco se ela for verdadeira.                            |
		| `else`        | Define o bloco executado quando a condição do `if` é falsa.                              |
		| `switch`      | Seleciona um bloco de execução com base em uma expressão ou condição.                    |
		| `case`        | Define uma possibilidade de execução dentro de um `switch`.                              |
		| `default`     | Define o caso padrão de um `switch` ou `select`, quando aplicável.                       |
		| `fallthrough` | Permite continuar a execução no próximo `case` de um `switch`, sem avaliar sua condição. |

		## 2. Laços de repetição e controle de iteração

		| Palavra    | Função                                                                                  |
		| ---------- | --------------------------------------------------------------------------------------- |
		| `for`      | Executa um bloco de código repetidamente. É a estrutura de repetição do Go.             |
		| `range`    | Itera sobre elementos de arrays, slices, strings, maps e canais, conforme o contexto.   |
		| `break`    | Interrompe o `for`, o `switch` ou o `select` mais interno, ou o indicado por um rótulo. |
		| `continue` | Pula para a próxima iteração do `for`, ou para o ponto indicado por um rótulo.          |
		| `goto`     | Transfere a execução para um rótulo dentro da mesma função.                             |

		## 3. Declaração de pacotes, importações e funções

		| Palavra   | Função                                                                       |
		| --------- | ---------------------------------------------------------------------------- |
		| `package` | Declara o pacote ao qual o arquivo pertence.                                 |
		| `import`  | Importa pacotes para uso no arquivo.                                         |
		| `func`    | Declara uma função ou método.                                                |
		| `return`  | Encerra a execução de uma função e, quando necessário, retorna valores.      |
		| `defer`   | Agenda a execução de uma chamada para o encerramento da função que a contém. |

		## 4. Declaração de variáveis e constantes

		| Palavra | Função                                                                      |
		| ------- | --------------------------------------------------------------------------- |
		| `var`   | Declara uma ou mais variáveis.                                              |
		| `const` | Declara constantes, cujos valores não podem ser alterados após a definição. |

		## 5. Tipos de dados e estruturas

		| Palavra     | Função                                                                     |
		| ----------- | -------------------------------------------------------------------------- |
		| `type`      | Declara um novo tipo ou associa um nome a um tipo existente.               |
		| `struct`    | Declara uma estrutura que agrupa campos de diferentes tipos.               |
		| `interface` | Define um conjunto de métodos que um tipo pode implementar implicitamente. |
		| `map`       | Representa um tipo de coleção de pares chave-valor.                        |

		## 6. Concorrência e comunicação entre goroutines

		| Palavra  | Função                                                                 |
		| -------- | ---------------------------------------------------------------------- |
		| `go`     | Inicia a execução de uma chamada de função em uma nova goroutine.      |
		| `chan`   | Declara um tipo de canal utilizado para comunicação entre goroutines.  |
		| `select` | Aguarda operações em canais e executa um dos casos que estiver pronto. |


		## 7 Funções predefinidas

		| Função   | Finalidade                                         |
		| -------- | -------------------------------------------------- |
		| `make`   | Inicializa slices, maps e channels                 |
		| `new`    | Aloca um valor e retorna um ponteiro               |
		| `len`    | Retorna o tamanho ou quantidade de elementos       |
		| `cap`    | Retorna a capacidade de arrays, slices ou channels |
		| `append` | Adiciona elementos a um slice                      |
		| `delete` | Remove uma chave de um map                         |


		## Resumo dos grupos

		Controle de fluxo

		`if`, `else`, `switch`, `case`, `default`, `fallthrough`

		Repetição e desvios

		`for`, `range`, `break`, `continue`, `goto`

		Pacotes e funções

		`package`, `import`, `func`, `return`, `defer`

		Variáveis e constantes

		`var`, `const`

		Tipos e estruturas

		`type`, `struct`, `interface`, `map`

		Concorrência

		`go`, `chan`, `select`

		Importante: essa lista contém 25 palavras reservadas de Go, mais as funcoes predefinidas, que são 6 funções, totalizando 31 palavras reservadas e funções predefinidas.


	*/

	/*
			# Expressões e comandos em Go (Golang)

		Em Go, os **comandos são instruções que dizem ao programa o que fazer**, enquanto as expressões são combinações de valores, variáveis e operadores que produzem um resultado.

		Por exemplo:
		- `var idade int` → declara uma variável.
		- `idade := 20` → cria uma variável e atribui um valor.
		- `idade > 18` → expressão que verifica se a idade é maior que 18.
		- `if idade > 18` → executa um bloco de código se a condição for verdadeira.
		- `for` → repete um bloco de código.

		Vamos organizar os principais comandos por grupos, com explicações simples e exemplos.

		## 1. Comandos de declaração

		Servem para **criar variáveis, constantes e outros elementos** que o programa utiliza.

		### `var` — declara uma variável

		O comando `var` declara uma variável. Podemos indicar o tipo de dado e, opcionalmente, um valor inicial.

		```go
		package main

		import "fmt"

		func main() {
		    var nome string = "Carlos"
		    var idade int = 25

		    fmt.Println(nome)
		    fmt.Println(idade)
		}
		```

		**Resultado:**
		```text
		Carlos
		25
		```

		Explicação:
		- `var` → declara uma variável.
		- `nome` → nome da variável.
		- `string` → tipo de texto.
		- `idade` → nome de outra variável.
		- `int` → tipo de número inteiro.

		Também podemos declarar uma variável sem atribuir um valor. Go atribui automaticamente o valor zero correspondente ao tipo.

		```go
		var idade int
		fmt.Println(idade)
		```

		Resultado:
		```text
		0
		```

		### `const` — declara uma constante

		Uma constante guarda um valor que não pode ser alterado durante a execução do programa.

		```go
		const PI = 3.14

		fmt.Println(PI)
		```

		Resultado:
		```text
		3.14
		```

		Por exemplo, podemos utilizar constantes para representar valores fixos, como o número de dias de uma semana.

		```go
		const DIAS_DA_SEMANA = 7
		```

		**Diferença:**
		- `var` → declara uma variável cujo valor pode ser alterado.
		- `const` → declara uma constante cujo valor não pode ser alterado.

		---

		## 2. Comandos de atribuição

		Servem para **atribuir valores a variáveis ou modificar valores existentes**.

		### `:=` — declaração curta com atribuição

		O operador `:=` declara uma nova variável e atribui um valor a ela ao mesmo tempo.

		```go
		package main

		import "fmt"

		func main() {
		    nome := "Ana"
		    idade := 30

		    fmt.Println(nome)
		    fmt.Println(idade)
		}
		```

		Resultado:
		```text
		Ana
		30
		```

		Go identifica automaticamente o tipo do valor atribuído:
		- `"Ana"` → `string`.
		- `30` → `int`.

		**Atenção:** `:=` só pode ser utilizado dentro de funções, como `main()`. Não pode ser utilizado para declarar variáveis diretamente no nível do pacote.

		### `=` — atribuição de um valor

		O operador `=` atribui ou substitui o valor de uma variável já declarada.

		```go
		idade := 20

		idade = 25

		fmt.Println(idade)
		```

		Resultado:
		```text
		25
		```

		Primeiro, a variável recebe o valor `20`. Depois, o valor é substituído por `25`.

		Não confunda:

		```go
		idade := 20 // Declara a variável e atribui um valor.
		idade = 25  // Altera o valor da variável existente.
		```

		### Outros operadores de atribuição

		| Operador | Significado | Exemplo | Resultado |
		|---|---|---|---|
		| `=` | Atribui um valor | `x = 10` | `10` |
		| `+=` | Soma e atribui | `x += 5` | Soma 5 a `x` |
		| `-=` | Subtrai e atribui | `x -= 2` | Subtrai 2 de `x` |
		| `*=` | Multiplica e atribui | `x *= 3` | Multiplica `x` por 3 |
		| `/=` | Divide e atribui | `x /= 2` | Divide `x` por 2 |
		| `%=` | Calcula o resto e atribui | `x %= 3` | Guarda o resto da divisão por 3 |

		Exemplo:

		```go
		x := 10

		x += 5

		fmt.Println(x)
		```

		Resultado:
		```text
		15
		```

		O comando `x += 5` equivale a `x = x + 5`.

		---

		## 3. Comandos condicionais

		Servem para **tomar decisões com base em condições**.

		### `if` — executa código se uma condição for verdadeira

		O `if` verifica uma condição. Se ela for verdadeira, executa o bloco de código correspondente.

		```go
		idade := 20

		if idade >= 18 {
		    fmt.Println("É maior de idade.")
		}
		```

		Resultado:
		```text
		É maior de idade.
		```

		A condição `idade >= 18` verifica se a idade é maior ou igual a 18.

		### `else` — executa código quando a condição é falsa

		```go
		idade := 16

		if idade >= 18 {
		    fmt.Println("É maior de idade.")
		} else {
		    fmt.Println("É menor de idade.")
		}
		```

		Resultado:
		```text
		É menor de idade.
		```

		### `else if` — verifica outra condição

		Utilizamos `else if` quando precisamos testar várias condições diferentes.

		```go
		nota := 15

		if nota >= 18 {
		    fmt.Println("Excelente!")
		} else if nota >= 10 {
		    fmt.Println("Aprovado!")
		} else {
		    fmt.Println("Reprovado!")
		}
		```

		Resultado:
		```text
		Aprovado!
		```

		Explicação:
		1. Se a nota for igual ou superior a 18, apresenta `Excelente!`.
		2. Caso contrário, se for igual ou superior a 10, apresenta `Aprovado!`.
		3. Se nenhuma das condições anteriores for verdadeira, apresenta `Reprovado!`.

		### `switch` — escolhe entre vários casos

		O `switch` permite executar diferentes blocos de código conforme o valor de uma expressão.

		É útil quando precisamos comparar um valor com várias possibilidades.

		```go
		dia := 2

		switch dia {
		case 1:
		    fmt.Println("Segunda-feira")
		case 2:
		    fmt.Println("Terça-feira")
		case 3:
		    fmt.Println("Quarta-feira")
		default:
		    fmt.Println("Outro dia")
		}
		```

		Resultado:
		```text
		Terça-feira
		```

		Explicação:
		- `switch` → inicia a seleção.
		- `case` → define uma possibilidade.
		- `default` → executa quando nenhum dos casos anteriores corresponde ao valor.

		**Importante:** em Go, não é necessário utilizar `break` em cada `case` para impedir que o programa execute automaticamente o caso seguinte. O comportamento normal do `switch` já é terminar após executar o caso correspondente.

		### Diferença entre `if` e `switch`

		- `if` → ideal para verificar condições, como `idade >= 18`.
		- `switch` → ideal para escolher entre várias possibilidades, como dias da semana, opções de menu ou estados de um pedido.

		---

		## 4. Comandos de repetição

		Servem para **repetir a execução de um bloco de código**.

		### `for` — repete um bloco de código

		Em Go, o `for` é o principal comando de repetição. Pode ser utilizado de diferentes maneiras.

		#### A) `for` com contador

		```go
		for i := 1; i <= 5; i++ {
		    fmt.Println(i)
		}
		```

		Resultado:
		```text
		1
		2
		3
		4
		5
		```

		O comando possui três partes:

		```go
		for inicialização; condição; atualização {
		    // Código repetido
		}
		```

		- `i := 1` → cria o contador e inicia-o em 1.
		- `i <= 5` → verifica se o contador é menor ou igual a 5.
		- `i++` → aumenta o contador em 1 a cada repetição.

		O ciclo termina quando a condição deixa de ser verdadeira.

		#### B) `for` como um `while`

		Go não possui um comando `while` separado. Podemos utilizar `for` para repetir enquanto uma condição for verdadeira.

		```go
		contador := 1

		for contador <= 3 {
		    fmt.Println(contador)
		    contador++
		}
		```

		Resultado:
		```text
		1
		2
		3
		```

		#### C) `for` infinito

		Quando não existe uma condição de paragem, o ciclo continua indefinidamente, até que alguma instrução o interrompa ou o programa termine.

		```go
		for {
		    fmt.Println("A executar...")
		    break
		}
		```

		Resultado:
		```text
		A executar...
		```

		Neste exemplo, `break` interrompe imediatamente o ciclo.

		### `break` — interrompe um ciclo

		O `break` termina imediatamente o ciclo em execução ou um `switch`.

		```go
		for i := 1; i <= 10; i++ {
		    if i == 4 {
		        break
		    }

		    fmt.Println(i)
		}
		```

		Resultado:
		```text
		1
		2
		3
		```

		Quando `i` chega a `4`, o `break` interrompe o ciclo.

		### `continue` — passa para a próxima repetição

		O `continue` ignora o restante código da repetição atual e passa para a seguinte.

		```go
		for i := 1; i <= 5; i++ {
		    if i == 3 {
		        continue
		    }

		    fmt.Println(i)
		}
		```

		Resultado:
		```text
		1
		2
		4
		5
		```

		O número `3` não é apresentado porque o `continue` ignora essa repetição.

		**Diferença:**
		- `break` → termina o ciclo.
		- `continue` → salta a repetição atual, mas mantém o ciclo em funcionamento.

		---

		## 5. Expressões e operadores

		As expressões combinam valores, variáveis e operadores para produzir resultados. São frequentemente utilizadas nas condições dos comandos `if`, `switch` e `for`.

		### Operadores aritméticos

		Permitem realizar cálculos matemáticos.

		| Operador | Significado | Exemplo | Resultado |
		|---|---|---|---|
		| `+` | Adição | `10 + 5` | `15` |
		| `-` | Subtração | `10 - 5` | `5` |
		| `*` | Multiplicação | `10 * 5` | `50` |
		| `/` | Divisão | `10 / 2` | `5` |
		| `%` | Resto da divisão | `10 % 3` | `1` |
		| `++` | Incremento de 1 | `i++` | Aumenta `i` em 1 |
		| `--` | Decremento de 1 | `i--` | Diminui `i` em 1 |

		Atenção: quando dois valores são inteiros, a divisão também pode produzir um resultado inteiro, descartando a parte decimal.

		```go
		fmt.Println(10 / 3)
		```

		Resultado:
		```text
		3
		```

		### Operadores relacionais

		Comparam valores e produzem um resultado booleano: `true` ou `false`.

		| Operador | Significado | Exemplo | Resultado |
		|---|---|---|---|
		| `==` | Igual a | `5 == 5` | `true` |
		| `!=` | Diferente de | `5 != 3` | `true` |
		| `>` | Maior que | `10 > 5` | `true` |
		| `<` | Menor que | `3 < 5` | `true` |
		| `>=` | Maior ou igual a | `5 >= 5` | `true` |
		| `<=` | Menor ou igual a | `3 <= 2` | `false` |

		Exemplo:

		```go
		idade := 20

		fmt.Println(idade >= 18)
		```

		Resultado:
		```text
		true
		```

		### Operadores lógicos

		Permitem combinar ou inverter condições.

		| Operador | Significado | Exemplo | Resultado |
		|---|---|---|---|
		| `&&` | E: ambas as condições devem ser verdadeiras | `true && false` | `false` |
		| `\|\|` | OU: pelo menos uma condição deve ser verdadeira | `true \|\| false` | `true` |
		| `!` | NÃO: inverte o valor lógico | `!true` | `false` |

		Exemplo:

		```go
		idade := 25
		temBilhete := true

		if idade >= 18 && temBilhete {
		    fmt.Println("Pode entrar.")
		}
		```

		Resultado:
		```text
		Pode entrar.
		```

		Para entrar, a pessoa precisa ter pelo menos 18 anos **e** possuir bilhete.

		---

		## 6. `range` — percorre elementos de uma coleção

		O `range` permite percorrer elementos de estruturas como slices, arrays, mapas e strings.

		É frequentemente utilizado em conjunto com o `for`.

		```go
		nomes := []string{"Ana", "Carlos", "Maria"}

		for indice, nome := range nomes {
		    fmt.Println(indice, nome)
		}
		```

		Resultado:
		```text
		0 Ana
		1 Carlos
		2 Maria
		```

		Explicação:
		- `range nomes` → percorre os elementos da coleção.
		- `indice` → recebe a posição do elemento.
		- `nome` → recebe o valor do elemento.

		Se apenas precisarmos dos valores, podemos ignorar o índice:

		```go
		for _, nome := range nomes {
		    fmt.Println(nome)
		}
		```

		O símbolo `_` descarta um valor que não pretendemos utilizar.

		---

		## 7. `func` — declara uma função

		Uma função agrupa instruções para executar uma tarefa específica. Ajuda a organizar o código e a reutilizar operações.

		```go
		package main

		import "fmt"

		func saudar() {
		    fmt.Println("Olá!")
		}

		func main() {
		    saudar()
		}
		```

		Resultado:
		```text
		Olá!
		```

		Explicação:
		- `func` → declara uma função.
		- `saudar` → nome da função.
		- `()` → indica os parâmetros; neste caso, não existem.
		- `{}` → contém as instruções executadas pela função.
		- `saudar()` → chama a função.

		### Função com parâmetros e retorno

		```go
		func somar(a int, b int) int {
		    return a + b
		}

		func main() {
		    resultado := somar(10, 5)
		    fmt.Println(resultado)
		}
		```

		Resultado:
		```text
		15
		```

		Neste exemplo:
		- `a` e `b` são parâmetros inteiros.
		- O último `int` indica que a função devolve um número inteiro.
		- `return` devolve o resultado da soma.
		- `somar(10, 5)` chama a função com os valores `10` e `5`.

		### `return` — devolve um resultado

		O `return` termina a execução da função atual e, quando aplicável, devolve um ou mais valores.

		```go
		func dobro(numero int) int {
		    return numero * 2
		}
		```

		Se chamarmos `dobro(4)`, a função devolve `8`.

		---

		## 8. `defer` — adia a execução de uma instrução

		O `defer` agenda a execução de uma instrução para quando a função que a contém estiver prestes a terminar.

		```go
		func exemplo() {
		    defer fmt.Println("Finalizado")

		    fmt.Println("A executar")
		}

		func main() {
		    exemplo()
		}
		```

		Resultado:
		```text
		A executar
		Finalizado
		```

		É frequentemente utilizado para fechar ficheiros ou libertar recursos depois de serem utilizados.

		---

		## 9. `go` — inicia uma goroutine

		O comando `go` permite iniciar a execução de uma função numa goroutine, um mecanismo leve de concorrência da linguagem.

		```go
		package main

		import (
		    "fmt"
		    "time"
		)

		func tarefa() {
		    fmt.Println("Tarefa executada")
		}

		func main() {
		    go tarefa()

		    time.Sleep(100 * time.Millisecond)
		}
		```

		Resultado possível:
		```text
		Tarefa executada
		```

		Neste exemplo, `go tarefa()` inicia a função numa goroutine. A pequena espera permite que a tarefa tenha tempo para executar antes de o programa terminar.

		**Importante:** `time.Sleep` é utilizado aqui apenas para simplificar o exemplo. Em programas reais, mecanismos como canais (`channels`) ou `sync.WaitGroup` são geralmente mais adequados para coordenar a execução concorrente.

		---

		## 10. `select` — seleciona operações em canais

		O `select` permite esperar por operações em canais (`channels`), escolhendo uma das operações que estiver pronta.

		É especialmente útil em programas que utilizam concorrência.

		```go
		package main

		import "fmt"

		func main() {
		    canal := make(chan string)

		    go func() {
		        canal <- "Mensagem recebida"
		    }()

		    mensagem := <-canal
		    fmt.Println(mensagem)
		}
		```

		Resultado:
		```text
		Mensagem recebida
		```

		Este exemplo demonstra a comunicação através de um canal. Para utilizar `select`, poderíamos escrever:

		```go
		select {
		case mensagem := <-canal:
		    fmt.Println(mensagem)
		default:
		    fmt.Println("Nenhuma mensagem disponível")
		}
		```

		O `case` permite definir a operação a executar. O `default` é opcional e executa imediatamente se nenhuma das outras operações estiver pronta.

		---

		# Resumo: principais comandos e expressões em Go

		| Grupo | Comando ou operador | Para que serve |
		|---|---|---|
		| Declaração | `var` | Declara uma variável. |
		| Declaração | `const` | Declara uma constante. |
		| Atribuição | `:=` | Declara uma variável e atribui um valor dentro de uma função. |
		| Atribuição | `=` | Atribui um valor a uma variável existente. |
		| Atribuição | `+=`, `-=`, `*=`, `/=`, `%=` | Realizam uma operação e atribuem o resultado. |
		| Condicionais | `if` | Executa código se uma condição for verdadeira. |
		| Condicionais | `else` | Executa código quando a condição anterior é falsa. |
		| Condicionais | `else if` | Verifica uma condição alternativa. |
		| Condicionais | `switch` | Seleciona entre vários casos. |
		| Condicionais | `case` | Define um caso num `switch` ou `select`. |
		| Condicionais | `default` | Define o caso padrão num `switch` ou `select`. |
		| Repetição | `for` | Repete um bloco de código. |
		| Repetição | `range` | Percorre elementos de uma coleção. |
		| Controlo de fluxo | `break` | Interrompe um ciclo ou `switch`. |
		| Controlo de fluxo | `continue` | Passa para a próxima repetição do ciclo. |
		| Controlo de fluxo | `fallthrough` | Permite continuar para o próximo `case` de um `switch`, com restrições. |
		| Funções | `func` | Declara uma função. |
		| Funções | `return` | Termina a função e devolve valores, quando aplicável. |
		| Funções | `defer` | Adia uma instrução até ao fim da função. |
		| Concorrência | `go` | Inicia uma goroutine. |
		| Concorrência | `select` | Aguarda operações em canais. |

		### Atenção a uma diferença importante

		Nem tudo o que aparece nesta tabela é uma palavra reservada de Go.

		- **Palavras reservadas:** `var`, `const`, `if`, `else`, `switch`, `case`, `default`, `for`, `range`, `break`, `continue`, `fallthrough`, `func`, `return`, `defer`, `go` e `select`.
		- **Operadores:** `:=`, `=`, `+`, `-`, `*`, `/`, `%`, `==`, `!=`, `&&`, `||` e outros.
		- **Identificadores predefinidos:** `true`, `false`, `nil`, `make`, `new`, `len`, `cap`, `append` e `delete`. Não são palavras reservadas, mas têm significados predefinidos na linguagem.

		**Para estudar na ordem certa:** começa por `var`, `:=`, `=`, `if`, `else`, `switch`, `for`, `range`, `break` e `continue`. Depois, avança para funções (`func`, `return`), operadores e, por fim, concorrência (`go`, `select` e canais).
	*/

	/*
			Variaveis e Constantes
			Variáveis são espaços de memória que armazenam valores que podem mudar durante a execução do programa. Em Go, as variáveis são declaradas usando a palavra-chave `var`, seguida pelo nome da variável e seu tipo. Por exemplo:
		    Constantes são valores fixos que não podem ser alterados após a sua definição. Em Go, as constantes são declaradas usando a palavra-chave `const`, seguida pelo nome da constante e seu valor. Por exemplo:

	*/

	var idade2 int
	idade2 = 39
	fmt.Println(idade2)

	const cpf string = "123.456.789-00"
	fmt.Println(cpf)

	/*
		Blocos e Escopo
		Blocos sao conjuntos de instrucoes delimitados por chaves {}. O escopo define a visibilidade e o tempo de vida das variaveis e constantes dentro do programa. Variaveis declaradas dentro de um bloco sao acessiveis apenas dentro desse bloco, enquanto variaveis declaradas fora de blocos sao acessiveis em todo o programa.
		Escopo de variáveis e constantes em Go é determinado pelo local onde são declaradas. Variáveis declaradas dentro de uma função têm escopo local, enquanto variáveis declaradas fora de funções têm escopo global. Por exemplo:
	*/

	var escopolocal = escopoGlobal
	fmt.Println(escopolocal)

	/*
		Arrays Fatias e Mapas
		Arrays sao coleçoes de elementos do mesmo tipo com tamanho fixo.
		Fatias sao coleções dinâmicas que podem crescer ou diminuir de tamanho.
		Mapas sao coleções de pares chave-valor, onde cada chave é única e mapeia para um valor. Por exemplo:
		 os arrays sao declarados usando a sintaxe [n]tipo, onde n é o tamanho fixo do array.

	*/
	// exemplo de array
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	array[4] = 10 // altera o valor do elemento na posicao 4
	fmt.Println("Array: ", array)
	fmt.Println("Elemento na posicao 3: ", array[3])

	// exemplo de fatia
	var fatia = []int{1, 2, 3, 4, 5}
	fmt.Println("Fatia: ", fatia)

	// exemplo de mapa
	var mapa = map[string]int{"um": 2, "dois": 2, "tres": 3}
	fmt.Println("Mapa: ", mapa)

	//Abstracoes
	//o conceito de abstracao consisite em esconder detalhes de algo, no acaso os detalhes desnecessrios

	//Paramentros
	//paramentros sao valores que sao passados para uma funcao ou metodo para que ele possa realizar uma tarefa especifica. Por exemplo, em uma funcao de soma, os numeros a serem somados sao passados como parametros.

	//Objetos e Classes (Go nao tem classes, mas tem structs)
	//objetos sao instancias de classes, que sao moldes que definem as caracteristicas e comportamentos de um objeto. Por exemplo, uma classe "Carro" pode ter caracteristicas como cor e modelo, e comportamentos como acelerar e frear. Um objeto "meuCarro" seria uma instancia da classe "Carro" com valores especificos para suas caracteristicas.

	//Structs em Go são tipos de dados compostos que permitem agrupar campos de diferentes tipos. Elas são usadas para criar objetos com propriedades e métodos associados. Por exemplo, podemos definir uma struct "Carro" com campos como "Marca", "Modelo" e "Ano". Em seguida, podemos criar instâncias dessa struct para representar carros específicos.

	//Exemplo de uma struct em Go:
	type Carro struct {
		Marca  string
		Modelo string
		Ano    int
	}

	//Exemplo de um objeto criado a partir da estruct "Carro"
	carro1 := Carro{Marca: "Toyota", Modelo: "Corolla", Ano: 2020}
	carro2 := Carro{Marca: "Honda", Modelo: "Civic", Ano: 2021}
	fmt.Println("Carro 1:", carro1)
	fmt.Println("Carro 2:", carro2)

	//Sistema de tipos
	// é um conjunto de regras que atribuem uma prorpriedade de tipo para as varias construcoes, tais como variaveis,, expressoes, funcoes, que um programa de computador e composto

	//Interfaces e Exceções

	var temperatura = Temperatura()
	fmt.Println(temperatura)

	/*Formatação

		%d → inteiro (decimal)
		%f e %g → números flutuantes
		%s → texto (string)
		%t → verdadeiro ou falso (true/false)
		%T → tipo da variável
		%v → valor genérico
		%q → valor entre aspas
	    %.2f → Números flutuantes com 2 casas decimais
		%p  → Endereço de memória de um ponteiro
	*/

	//Conversao de tipos

	var inteiro = 10
	var texto = "texto"

	var convertido float64 = float64(inteiro)
	fmt.Printf("valor de inteirot = %d : texto = %s ", inteiro, texto)
	fmt.Println(convertido)
	fmt.Printf("tipo da variavel %T", convertido)

	//Descobrindo o tipo de uma variavel %T

	//Pacote FMT format
	/*

			Função	Finalidade	Mostra no ecrã?
		fmt.Print()	Apresenta valores sem quebra de linha automática.
		fmt.Println()	Apresenta valores com quebra de linha.
		fmt.Printf()	Apresenta valores formatados.

		fmt.Scan()	Lê valores introduzidos pelo utilizador.
		fmt.Scanln()	Lê valores até ao fim da linha.
		fmt.Scanf()	Lê dados utilizando um formato especificado.

		fmt.Sprint()	Converte e combina valores em texto.
		fmt.Sprintln()	Combina valores em texto, acrescentando uma quebra de linha.
		fmt.Sprintf()	Cria uma string formatada.

		fmt.Fprint()	Escreve num destino específico.
		fmt.Fprintf()	Escreve texto formatado num destino específico.
		fmt.Fprintln()	Escreve num destino específico com quebra de linha.
	*/

}

// palavra reservada + nome + tipo recebe valor
const ebulicaoF float64 = 212.0

func Temperatura() float64 {
	fmt.Println("Temperatura em Celsius:")
	var temperaturaF float64 = ebulicaoF

	var tempC float64 = (temperaturaF - 32) * 5 / 9 //Converte F para C
	return tempC
	//fmt.Println("a temperatura de ebulicao da agua em F =", temperaturaF)
	//fmt.Println("a temperatura de ebulicao da agua em C =", tempC)

}
