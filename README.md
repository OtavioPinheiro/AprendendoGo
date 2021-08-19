# Linguagem GO
Aprendendo GO Lang por meio de vídeos tutoriais.

# Tópicos
- [O que é GO Lang?](#o-que-é-go-lang)
- [Regras e características da linguagem](#regras-e-caractersticas-da-linguagem)
- [Operador curto de declaração](#operador-curto-de-declarao)
- [Blank Identifier](#blank-identifier)
- [Variáveis](#variveis)
- [String literal](#raw-string-literal-e-interpreted-string-literal)
- [Tipo String](#tipo-string-cadeia-de-caracteres)
- [Pacote fmt](#pacote-fmt)
- [Iota](#iota)
- [Switch case](#switch-e-case)
- [Função Range](#função-range)
- [Dados compostos](#dados-compostos)
- [Manipulando Slices](#manipulando-slices)
- [Função make()](#função-maket-len-cap)
- [Slice multidimensional](#slice-multidimensional)
- [Funções](#funções)
- [Defer](#defer)
- [Métodos](#métodos)
- [Interface e polimorfismo](#interface-e-polimorfismo)
- [Referências](#referências)
  
# O que é Go Lang?
A linguagem Go é uma linguagem de programação criada pela Google e lançada em códgio livre em 2009. É uma linguagem compilada e focada em produtividade e programação concorrente, baseada em trabalhos feitos no sistema operacional chamado Inferno.

# Regras e características da linguagem
1. Não é permitido declarar uma variável e não utilizá-la. Isso gerará um erro, pois fere um dos princípios da linguagem que é possuir um código limpo.
2. Go é uma lingaugem de tipagem estática.
3. Quando uma variável é declarada fora de um codeblock scope e não é atribuída a um valor neste momento, ela só poderá receber um valor dentro de um codeblock.
4. Slice array, struct, map -> Tipos de dados compostos
5. Todo valor em Go pertence também ao tipo de interface vazia, representado pela notação "interface{}"
6. Em Go strings são imutáveis. Para alterar uma string é necessário criar outra string com a alteração desejada.
7. Não existe ***While*** em GO.

# Operador curto de declaração
É representado por `:=`. É utilizado na declaração de variáveis. O diferencial é que esse operador possui tipagem automática, ou seja, não é necessário informar o tipo da variável quando utilizamos o operador curto de declaração e só podemos utilizar esse operador dentro de um *codeblocks* (escopo).

Esse operador é diferente do operador de atribuição `=`. Cuidado para não confundí-lo, pois são utilizados em contextos diferentes.

# Blank identifier
O blank identifier serve para dizer ao programa ignorar um retorno de informações recebido por uma função. É representado pelo caractere `_`.

# Variáveis
Variáveis em GO podem ser declaradas das seguintes formas:
```go
  var i int -> Declaração da variável
  i = 42 -> Inicialização da variável

  var j int = 27

  k := 99
```

## Valor zero
Sempre que criarmos uma variável e não inicializarmos ela, por padrão ela virá com *"valor zero"*
Valores zeros nos diferentes tipos de variáveis:
- ints: 0;
- floats: 0.0;
- booleans: false;
- strings: ""
- pointers, functions, interfaces, slices, channels, maps: nil

**Observações**
- Use `:=` sempre que possível;
- Use var para package level scope.

# Raw string literal e Interpreted string literal
*Interpreted String Literal* são strings interpretadas.<br>**Exemplo:** `fmt.Printf("Olá!\nTudo bem?\tEspero que sim.")`

Já o *Raw String Literal* são strings cruas que não serão interpretadas.<br>**Exemplo:** ```fmt.Printf(`Olá!\nTudo bem?\tEspero que sim.`)```

**Atenção**: Para Raw Strings utiliza-se o acento grave ``.

# Tipo string (cadeia de caracteres)
- Strings são sequências de bytes.
- são imutáveis
- Uma string é uma fatia de bytes (*slice of bytes*)

# Pacote fmt
- **fmt.Print():** Exibe na tela a string passada como argumento.
- **fmt.Printf():** Exibe na tela a string passada como argumento. Permite formatações.
- **fmt.Println():** Exibe na tela a string passada como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Sprint():** Retorna uma concatenação de strings passadas como argumento.
- **fmt.Sprintf():** Retorna uma concatenação de strings passadas como argumento. Permite formatações.
- **fmt.Sprintln():** Retorna uma concatenação de strings passadas como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Fprint():** Escreve a string, passada como argumento, em um arquivo.
- **fmt.Fprintf():** Escreve a string, passada como argumento, em um arquivo. Permite formatações.
- **fmt.Fprintln():** Escreve a string, passada como argumento, em um arquivo. Insere uma quebra de linha (`\n`) ao final da string.

# Iota
São números sequênciais atribuídos automáticamente pelo sistema.

**Exemplos:**
```go 
const (a = iota, b = iota, c = iota)
```

**Resultado:** `a = 0, b = 1, c = 2`

# Switch e case
É igual para todas as linguagens que utilizam *switch case*, ou seja, o *switch* irá comparar o valor da variável informada com o valor declarado nos *cases*, caso verdadeira, executará o código dentro dos respectivos *cases*, senão seguirá para o *default*. Uma diferença é que em GO existe o termo `fallthrough` que faz com que o próximo *case* também seja executado caso o *case* com *fallthrough* seja verdadeiro.

**Detalhes:**
- O switch statement (variável que vem logo após a palavra reservada *switch*) pode ou não ser informado nessa etapa. Porém senão for, terá que ser informado dentro dos *cases*.

  **Exemplo 1:**
  ```go
  switch isso {
    case "aquilo":
      fmt.Println("isso é igual aquilo")
    case "isto aqui":
      fmt.Println("isso é igual a isto aqui")
  }
  ```

  **Exemplo 2:**
  ```go
  switch {
    case isso == "aquilo":
      fmt.Println("isso é igual aquilo")
    case isso == "isto aqui":
      fmt.Println("isso é igual a isto aqui")
  }
  ```

- Não é necessário ter o ***break*** dentro dos ***cases*** para interrompê-los, como em outras linguagens.

# Função range
Percorre todo o array ou slice até o final. Comumente usado no loop for.

**Exemplo 1:**
```go
slice := []int{20, 21, 22, 23}
total := 0
for _, valor := range slice {
  total += valor
}
fmt.Println("O valor total é: ", total)
```
<pre><code><span style = "color: blue">-- Output: O valor total é: 86
</span></code></pre>

**Exemplo 2:**
```go
slice := []string{"morango", "uva", "pêra", "maçã", "kiwi"}

for índice, valor := range slice {
  fmt.Println("No índice", índice, "temos o valor:", valor)
}
```
<pre><code><span style = "color: blue">-- Output: No índice 0 temos o valor: morango
No índice 1 temos o valor: uva
No índice 2 temos o valor: pêra
No índice 3 temos o valor: maçã
No índice 4 temos o valor: kiwi
</span></code></pre>

# Dados compostos
Dados compostos são qualquer tipo de dados que podem ser construídos em um programa utilizando dados primitivos da programação ou outro tipo de dados compostos.

- ***Arrays:*** Arrays são vetores de número finito. Podem ser vetores de *strings*, *integers*, *floats*, etc. Arrays são dados compostos.
- ***Slices:*** Slices são ponteiros de arrays (conjunto de dados) que podem ser compostos por *strings*, *integers*, *floats*, etc. Logo *slices* são um tipo de dados compostos. Slices, quando são declarados, são como um array de tamanho "infinito".

**Exemplo de declaração de um array e de um slice:**
  ```go
  array := [5]int{1, 2, 3, 4, 5}
  slice := []int{1, 2, 3, 4, 5}
  ```

# Manipulando Slices
- **Atribuição:** Podemos reatribuir um dado valor do slice simplesmente informando o índice entre colchetes e em seguida passando o novo valor.<br>**Exemplo:** `slice[3] = "banana"`
- **Adicionando novos elementos:** Não é possível adicionar novos elementos ao slice sem usar a função append, pois o slice é criado com um **número fixo** de elementos.

  **Exemplo:**
  ```go
  slice := []string{"morango", "uva", "pêra", "maçã", "kiwi"}
  
  //Maneira errada de se adicionar um novo elemento ao slice.</span>
  slice[5] = "maracujá"
  
  //Maneira certa de se adicionar um novo elemento ao slice.</span>
  slice = append(slice, "maracujá")
  ```

  **Importante:** A função *append* adiciona novos elementos ao final do *slice* e recebe como argumentos um *slice* e pode receber um número infinito de elementos desde que sejam do mesmo tipo dos elementos presentes no *slice* ao qual deseja-se adicioná-los.

- **Excluindo elementos:** Para excluir elementos de um *slice* devemos usar a função append junto com a operação *slice*, nesse caso significa corte, ou seja, pegar uma parte do dado composto chamado *slice* (quase igual a um *array*) e unir com outra parte, porém deixando o elemento a ser excluído de fora.
**Exemplo:**
  ```go
  frutasFavoritas := []string{"morango", "uva", "melão", "kiwi", "maracujá"}
  frutasFavoritasAtualizada := append(frutasFavoritas[:2], frutasFavoritas[3:]...)
  ```

## Slice slices (fatia de fatias)
Operação *slice* (cortar, fatiar) criar um sub-array de um array principal, ou seja, cria um outro conjunto de dados a partir de um já existente.

**Exemplo:**
```go
frutasFavoritas := []string{"morango", "limão", "laranja", "uva", "kiwi", "melão", "banana", "maçã", "pêra", "goiaba"}
frutas_acidas := frutasFavoritas[:5]
frutas_n_acidas := frutasFavoritas[5:]
```

**Importante:** Na hora de realizar o corte no conjunto de dados sempre lembrar que o **último índice não será incluído**, ou seja, `[:5]` neste exemplo o corte será aplicado desde o início do conjunto de dados até o índice 5, porém o valor deste índice não será incluído no corte. E `[5:]` neste outro exemplo significa que o corte começará no índice 5, então o valor dele estará incluído no corte, e irá até o final do conjunto de dados, uma outra de maneira de se realizar a mesma tarefa seria usar a função *len()*, então teríamos `newSlice := slice[2:len(slice)]`.

**Importantíssimo:** **CUIDADO** ao usar *slices*, pois quando se realiza um corte de um slice pré-existente para criar um novo, o *slice* antigo ficará com os valores originais alterados devido ao corte realizado.

**Exemplo:**
```go
  package main
  import (
      "fmt"
  )
  func main() {
      primeiroslice := []int{1, 2, 3, 4, 5}
      fmt.Println(primeiroslice)
      segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)
      fmt.Println(segundoslice)
      fmt.Println(primeiroslice)
  }
```
  **Resultado:**
  <pre><code>[1 2 <span style = "color:blue"><b>3</b></span> 4 5]
[1 2 5]
[1 2 <span style = "color: red"><b>5</b></span> 4 5]</code></pre>

**Dica:** Se for precisar criar um *slice* a partir de outro *slice* e não quer que os dados do *slice* original se percam, utilize um laço *for* para realizar essa tarefa e copie elemento a elemento do *slice* antigo para o novo. Ou então utilize a mesma variável para realizar o corte e criar um slice subjacente.

# Função ***make([]T, len, cap)***.
Cria um slice do tamanho e da capacidade informados como parâmetros. O tamanho (*length*) trata-se de quantos elementos o slice possui inicialmente e a capacidade (*capacity*) trata-se do tamanho máximo ou limite do slice. Caso seja atribuído mais elementos que a capacidade informada, um novo slice será criado com o dobro da capacity e os dados serão copiados para lá.

```go
package main

import (
  "fmt"
)

func main() {
  slice := make([]int, 5, 10)
  
  slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

  slice = append(slice, 6)
  slice = append(slice, 7)
  slice = append(slice, 8)
  slice = append(slice, 9)
  slice = append(slice, 10)

  fmt.Println(slice, len(slice), cap(slice))

  slice = append(slice, 11)

  fmt.Println(slice, len(slice), cap(slice))
}
```

**Importante:** Sempre que um slice mudar de tamanho (*length*), um novo array será criado e os dados serão copiados para ele.

# Slice Multidimensional
*Slices Multidimensionais* são como matrizes em Linguagem C.
- **Exemplo:**
  ```go
  package main

  import (
    "fmt"
  )

  func main() {
    ss := [][]int {
      []int {1, 2, 3},
      []int {4, 5, 6},
      []int {7, 8, 9},
    }
    fmt.Println(ss)
  }
  ```

# Maps
São uma estrutura de dados baseados em chave-valor, assim como são os dicionários em python ou os arquivos json em JavaScript. Essas estruturas não possuem ordem, ou seja, são **desordenadas**.
- Declaração: `map[chave]valor{ chave: valor }`
- Acesso: `m[chave]`
- Observações: Uma chave sem valor retornará, por padrão, o valor zero. Porém, as vezes, esse comportamento não é o desejável, então existe a solução *comma ok idiom.*
- **Exemplo:**
  ```go
  package main
  
  import (
      "fmt"
  )

  func main() {
      amigos := map[string]int{
          "alfredo": 5551234,
          "joana": 9996674,
      }

      fmt.Println(amigos)
      fmt.Println(amigos["joana"])

      amigos["gopher"] = 44444444

      fmt.Println(amigos)
      fmt.Println(amigos["gopher"], "\n\n")

      //comma ok idiom
      if será, ok := amigos["gasparzinho"]; !ok {
          fmt.Println("não tem!")
      } else {
          fmt.Println(será, ok)
      }
  }
  ```
- Adicionando elementos: `map[chave] = valor`

## Maps - range e delete.
- A função ***range*** em *Maps* irá percorrer todas as *keys* e *values* presentes no *Maps*. Lembrando que sempre o primeiro valor será a chave (*key*) e o segundo valor será o valor (*value*).

- A função ***delete***, como o nome sugere, irá deletar um item presente no *Map*, lembrando que caso não haja nenhum item com a chave especificada, a função **não irá retornar um erro ou uma exception!** Para esta função devem ser passados o *map* que se dejeta remover o item e a chave desse item.

**Exemplo:**
```go
package main

import (
    "fmt"
)

func main() {
    qualquercoisa := map[int]string {
        10: "dez",
        100: "cem",
        300: "trezentos",
        1000: "mil",
    }
    fmt.Println(qualquercoisa)
    delete(qualquercoisa, 300)
    fmt.Println(qualquercoisa)
}
```

# Structs
*Structs* são estruturas de dados, cuja os dados podem ser difrentes.
- Declaração:
  ```go
  package main

  import (
    "fmt"
  )

  type cliente struct {
    nome string
    sobrenome string
    fumante bool
  }

  func main() {
    cliente1 := cliente{
      nome: "João",
      sobrenome: "da Silva",
      fumante: false,
    }

    cliente2 := cliente{"Joana", "Pereira", true}
    
    fmt.Println(cliente1)
    fmt.Println(cliente2)
  }
  ```

## Struct embutido
*Structs* embutidos nada mais são do que um *Struct* dentro de outro.
- **Exemplo:**
  ```go
  package main

  import (
    "fmt"
  )

  type pessoa struct {
    nome string
    idade int
  }

  type profissional struct {
    pessoa
    titulo string
    salario int
  }

  func main() {
    pessoa1 := pessoa{
      nome: "Alfredo",
      idade: 30,
    }

    pessoa2 := profissional{
      pessoa: pessoa{
        nome: "Maria",
        idade: 22,
      }
      titulo: "Pizzaiola",
      salario: 10000,
    }

    pessoa3 := profissional{ pessoa: pessoa{ "José", 40 }, "Político", 100000 }
    
    fmt.Println(pessoa1)
    fmt.Println(pessoa2)
  }
  ```

## Acessar Valores
Para acessar valores específicos da *struct*, usamos a notação de ponto, como em linguagens orientadas a objeto. Caso seja *structs* embutidos é só passar o nome da *struct* embutida (interna) e depois o nome do campo. E se o nome do campo da *struct* embutida seja diferente de todos os nomes do campo da *struct* principal, então o campo da *struct* interna é promovida para campo da *struct* principal.

**Exemplos:**
1. `fmt.Println(cliente1.sobrenome)`
2. `fmt.Println(pessoa2.titulo)`
3. `fmt.Println(pessoa2.pessoa.nome)`
4. `fmt.Println(pessoa2.nome)`

## Struct anônimo
Não se declara o tipo do *struct* anônimo, apenas declara-se o valor. Esse tipo de *struct* não é reutilizável.
- **Exemplo:**
  ```go
  package main

  import (
    "fmt"
  )

  func main() {
    x := struct { 
      nome string
      idade int
    }{
      nome: "Mário",
      idade: 50
    }
  }
  ```

# Funções
Uma função é um grupo de instruções que juntas executam uma tarefa. Cada programa em Go tem pelo menos uma função, que é a ***main()***. Normalmente cada função deve executar uma tarefa específica e é usada para dividir o código e reaproveitar algumas partes dele, evitando repetições. Uma declaração de uma função informa ao compilador sobre o nome da função, tipo de retorno e os parâmetros que a função deve receber. Uma definição de uma função está associada aos valores que os argumentos (parâmetros) irão receber quando a função é chamada.

A linguagem Go por padrão possui várias funções pré-prontas para serem utilizadas. As funções, em Go, também são conhecidas como método, sub-rotina ou procedimento.

**IMPORTANTE:** Em Go, os argumentos, que são passados para a função, são do tipo ***pass by value***, ou seja, o argumento não irá receber a variável em si, mas sim o valor que a variável possui, diferente do que acontece com as funções utilizadas na linguagem ***Python***, por exemplo. Há apenas um caso em Go que será ***pass by reference***.

**Exemplos:**
- Funções que não recebem argumentos:
  ```go
    package main
    
    import (
      "fmt"
    )

    func main() {
      basica()
    }

    func basica() {
      fmt.Println("Oi, bom dia!")
    }
  ```

- Funções que recebem argumentos:
  ```go
    package main
    
    import (
      "fmt"
    )

    func main() {
      argumento("tarde")
    }

    func argumento(s string) {
      if s == "manhã" {
        fmt.Println("Oi, bom dia!")
      } else if s == "tarde" {
        fmt.Println("Oi, boa tarde!")
      } else {
        fmt.Println("Oi, bom noite!")
      }
    }
  ```

- Funções que retornam valores:
  ```go
    package main
    
    import (
      "fmt"
    )

    func main() {
      valor := soma(10, 10)
      fmt.Println(valor)
    }

    func soma(x, y int) int {
      return x + y
    }
  ```

- Funções com múltiplos retornos e parâmetros variádicos:
  ```go
    package main
    
    import (
      "fmt"
    )

    func main() {
      total, quantos := soma(10, 10)
      fmt.Println(total, quantos)
    }

    func soma(x ...int) (int, int) {
      soma := 0
      for _, v := range x {
        soma += v
      }
      return soma, len(x)
    }
  ```

**IMPORTANTE:** Quando for utilizar um parâmetro variádico em um argumento de função, ele deve ser o último a ser passado. Essa função variádica pode receber nenhum valor como argumento.

## Enumerando um slice como argumento para uma função
Não é possível passar um slice para uma função, mesmo que esta seja variádica. Pois a função espera um único argumento por vez e não um slice. Logo para ressolver esse problema é necessário usar o enumerador de slice (operador ...), assim a função irá receber um único argumento por vez.

**Exemplo:**
```go
  package main

  import (
    "fmt"
  )

  func main() {
    si := []int{1, 2, 3, 4, 5}
    total := soma(si...)
    fmt.Println(total)
  }

  func soma(x ...int) int {
    total := 0
    for _, v := range x {
      total += v
    }
    return total
  }
```

## Função anônima
Função anônima é aquela que não possui nome e podemos chamá-la imediatamente após a sua declaração. **Veja o exemplo:**

```go
package main

import (
  "fmt"
)

func main() {
  x := 387

  func(x int) {
    fmt.Println(x, "vezes 10 é:")
    fmt.Println(x * 10)
  }(x)
}
```

## Função como expressão
É quando atribuímos uma função a uma variável.

**Exemplo:**

```go
package main

import (
  "fmt"
)

func main() {
  x := 100

  y = func (x int) int {
    return x * 3
  }
  fmt.Println(x, "vezes 3 é:", y(x))
}
```

## Retornando uma função
Uma função pode retornar uma outra função. Essa prática é muito comum e muito utilizada na linguagem Go.

**Exemplo:**

```go
package main

import(
  "fmt"
)

func main() {
  x := retornaumafuncao()
  z := x(3)
  y := retornaumafuncao()(3) //segunda maneira de se chamar uma função e passar um parâmetro a ela.
  fmt.Println(z)
  fmt.Println(y)
}

func retornaumafuncao() func(int) int {
  return func(i int) int {
    return i * 10
  }
}
```

## Função Callback
Função Callback é quando passamos uma função como argumento de outra função.
[Códgio de exemplo](./exercicios/cap12/cap12_aulaCallback.go)

## Closure
Closure é cercar ou capturar um escopo (*scope*) para que possamos utilizá-lo em outro contexto. Closures nos permite salvar dados entre cahmadas de funções e ao mesmo tempo isolar estes dados do resto do código.

## Recursividade
Funções recursivas são aquelas funções que chamam a si própria. Os exemplos mais comuns de recursividade são fatoriais, matrioskas, efeito droste e fractais.

[Código de exemplo](./exercicios/cap12/cap12_aulaRecursividade.go)

# Defer
*Defer*, do inglês, significa adiar. Na linguagem de programação Go, *Defer* é um *statement* (instrução) que é colocada antes de uma outra instrução, fazendo com que esta seja executada por último. Se houverem mais de um *Defer*, então o primeiro *Defer* que o compilador encontrar será o último a ser executado, como se fosse um FILO (First In, Last Out).

*Defer* é frequentemente usado em programas que manipulam arquivos. Então, quando abrimos arquivos em Go, usamos o *Defer*, junto com a função para fechar o arquivo, logo em seguida, deste modo, não esquecemos de fechar o arquivo e evitamos problemas de consumo de memória.

**IMPORTANTE:** Quando há um *return* e um *defer* juntos em uma parte do código, o *defer* será executado primeiro e depois o *return*.

**Exemplos:**
```go
  package main

  import (
    "fmt"
  )

  func main() {
    defer fmt.Println("Os últimos...")
    fmt.Println("Serão os primeiros")
  }
```

```go
  package main

  import (
    "fmt"
  )

  func main() {
    defer fmt.Println("der Anfang")
    defer fmt.Println("das Ende ist")
    defer fmt.Println("das Ende und")
    defer fmt.Println("Der Anfang ist")
  }
```

# Métodos
Em Linguagem Go, um método é uma função anexada a um tipo, ou seja, é uma função que traz uma funcionalidade para um tipo específico.

**Exemplo:**
```go
package main

import (
  "fmt"
)

type pessoa struct {
  nome string
  idade int
}

func (p pessoa) oibomdia() {
  fmt.Println(p.nome, "diz bom dia!")
}

func main() {
  jose := pessoa{"José", 30}
  jose.oibomdia()
}
```

# Interface e polimorfismo
**Interfaces** são um conjunto de métodos. Se um tipo (`type`, entenda *type* como um objeto) tiver os mesmos métodos que uma interface, a linguagem Go, automaticamente, atribuirá o tipo da interface. Fazendo um paralelo com outras linguagens, como *Java*, é como se utilizássemos o `implements` para implementar uma interface, porém em Go não é necessário realizar essa tarefa, pois ela será realizada automaticamente.

**Polimorfismo** em Go significa utilizar uma mesma função para mais de um tipo (`type`) se os mesmos fizerem parte da mesma interface.

- [**Código de exemplo 1**](./exercicios/cap12/cap12_interface_polimorfismo.go)
- [**Código de exemplo 2**](./exercicios/cap12/cap12_interface_polimorfismo_2.go)

# Ponteiros
Ponteiros, em linguagem de programação, é um objeto cujo valor aponta (ou refere-se) para outro valor armazenado em algum lugar da memória do computador que usa esse endereço de memória. Para obter o valor salvo nessa localização da memória, usa-se um processo chamado *dereferencing*. Assumindo que `x` seja uma variável qualquer e `p` seja um ponteiro que refere-se ao valor armazenado em `x`, então para exibir esse valor, por meio do ponteiro, usamos `*p`.

Outro exemplo: [cap.15 exercício 2](./exercicios/cap15/cap15_exercicio02-ponteiros.go)

Como visto no exemplo acima, existe um atalho para referenciar um valor de uma *struct*, sendo que a maneira tradicional é `(*variavel).campoDaStruct` e o atalho é `variavel.campoDaStruct`. De acordo com a [documentação](https://golang.org/ref/spec#Selectors), como exceção, se o tipo da variavel é um tipo de ponteiro nomeado e `(*variavel).campoDaStruct` é uma expressão seletora válida que aponta para um campo, mas para um método, então é possível utilizar o atalho `variavel.campoDaStruct`.

# JSON
Em Go é possível transformar uma estrutura de dados (*struc*) em JSON por meio do pacote `encoding/json`. É importante frisar que para que os campos da *struc* sejam exportados para *JSON*, a letra inicial dos campos e do nome da *struc* precisam estar em letra maiúscula. Logo, em Go, tudo que possuir a letra inicial maiúscula poderá ser exportado e estará visível para outros pacotes (*packages*).

# Referências
- Korbes, Ellen. **Aprenda Go 🇧🇷**. Aprenda Go. Disponível em: https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg
- Go by example. **Go by Example**. Disponível em: https://gobyexample.com/
- fatih. **color**. Disponível em: https://github.com/fatih/color
- GO. **The Go Programming Language Specification**. Disponível em: https://golang.org/ref/spec
- GO. **Effective Go**. Disponível em: https://golang.org/doc/effective_go
- Tutorials Point. **Tutorial GO**. Disponível em: https://](https://www.tutorialspoint.com/go/go_functions.htm)