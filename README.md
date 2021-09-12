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
- [Ponteiros](#ponteiros)
- [JSON](#json)
- [Bcrypt](#bcrypt)
- [Concorrência](#concorrência)
- [Condição de corrida](#condição-de-corrida)
- [Comandos Go](#comandos-go)
- [Canais](#canais)
- [Convergência](#convergência)
- [Divergência](#divergência)
- [Context](#context)
- [Tratamento de erros](#tratamento-de-erros)
- [Documentação](#documentação)
- [Testes em Go](#testes-em-go)
- [Benchmarks](#benchmarks)
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
[Códgio de exemplo](./exercicios/cap12/callback/cap12_aulaCallback.go)

## Closure
Closure é cercar ou capturar um escopo (*scope*) para que possamos utilizá-lo em outro contexto. Closures nos permite salvar dados entre cahmadas de funções e ao mesmo tempo isolar estes dados do resto do código.

## Recursividade
Funções recursivas são aquelas funções que chamam a si própria. Os exemplos mais comuns de recursividade são fatoriais, matrioskas, efeito droste e fractais.

[Código de exemplo](./exercicios/cap12/recursividade/cap12_aulaRecursividade.go)

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

- [**Código de exemplo 1**](./exercicios/cap12/interface_polimorfismo/01/cap12_interface_polimorfismo.go)
- [**Código de exemplo 2**](./exercicios/cap12/interface_polimorfismo/02/cap12_interface_polimorfismo_2.go)

# Ponteiros
Ponteiros, em linguagem de programação, é um objeto cujo valor aponta (ou refere-se) para outro valor armazenado em algum lugar da memória do computador que usa esse endereço de memória. Para obter o valor salvo nessa localização da memória, usa-se um processo chamado *dereferencing*. Assumindo que `x` seja uma variável qualquer e `p` seja um ponteiro que refere-se ao valor armazenado em `x`, então para exibir esse valor, por meio do ponteiro, usamos `*p`.

Outro exemplo: [cap.15 exercício 2](./exercicios/cap15/cap15_exercicio02-ponteiros.go)

Como visto no exemplo acima, existe um atalho para referenciar um valor de uma *struct*, sendo que a maneira tradicional é `(*variavel).campoDaStruct` e o atalho é `variavel.campoDaStruct`. De acordo com a [documentação](https://golang.org/ref/spec#Selectors), como exceção, se o tipo da variavel é um tipo de ponteiro nomeado e `(*variavel).campoDaStruct` é uma expressão seletora válida que aponta para um campo, mas não para um método, então é possível utilizar o atalho `variavel.campoDaStruct`.

# JSON
## Transformando em JSON
Em Go é possível transformar uma estrutura de dados (*struct*) em JSON por meio do pacote `encoding/json` usando a função `Marshal`. É importante frisar que para que os campos da *struct* sejam exportados para *JSON*, a letra inicial dos campos e do nome da *struct* precisam estar em letra maiúscula. Logo, em Go, tudo que possuir a letra inicial maiúscula poderá ser exportado e estará visível para outros pacotes (*packages*).

**Exemplo:**
[Exemplo](./exemplos/cap16/Marshal/exemplo_cap16_Marshal.go)

Ao invés de utilizar a função `Marshal` é possível usar a função `Encode` para transformar um *JSON* em *struct*. A diferença entre essas funções é a forma como se usa cada uma. Na função `Marshal` o resultado é salvo em uma variável que pode ser manipula como o usuário desejar, já na `Encode` o resultado estará vinculado a uma interface que sempre será chamada quando usarmos a variável na qual está as informações obtidas do *JSON*.
[Mais informações sobre Marshal](https://pkg.go.dev/encoding/json#Marshal)

**Exemplo:**
[Exemplo](./exemplos/cap16/Encode/exemplo_cap16_Encode.go)

## Transformando JSON para *struc* em Go
De maneira similar a transformar uma *struct* em JSON, também é possível realizar o processo inverso, ou seja, capturar um *JSON* e transportar as informações para uma *struc* em Go utilizando o pacote `enconding/json` e a função `Unmarshal` ou `Decode`.
[Mais informações sobre Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal)

**Exemplo:**
[Exemplo](./exemplos/cap16/Unmarshal/exemplo_cap16_Unmarshal.go)

Vale lembrar que no exemplo acima utilizamos *tags*. As *tags* servem para mapear os campos do *JSON* com os campos da *struc* que estará recebendo essas informações.

# Bcrypt
Bcrypt é um método de criptografia do tipo hash para senhas baseado no *BlowFish*. Foi criado por Niels Provos e David Mazières e apresentado na conferência da USENIX em 1999.

Este método apresenta uma maior segurança em relação a maioria dos outros métodos criptográficos, pois é resistente à ataques de força bruta. [Mais informações](https://pt.wikipedia.org/wiki/Bcrypt#:~:text=bcrypt%20%C3%A9%20um%20m%C3%A9todo%20de,para%20senhas%20baseado%20no%20Blowfish.&text=O%20algoritmo%20bcrypt%20foi%20implementado,fun%C3%A7%C3%A3o%20%22crypt%22%20do%20UNIX.)

[Exemplo de uso](./cap16/../exemplos/cap16/bcrypt/exemplo_cap16_bcrypt.go)

# Concorrência
Concorrência, no sentido de programação concorrente ou programação simultânea, é um paradigma de programação para a construção de programas de computador que fazem uso da execução simultânea de várias tarefas computacionais interativas, que podem ser implementadas como programas separados ou como um conjunto de *threads* criadas por um único programa. [Wikipédia - Programação concorrente](https://pt.wikipedia.org/wiki/Programa%C3%A7%C3%A3o_concorrente#:~:text=Programa%C3%A7%C3%A3o%20concorrente%20ou%20programa%C3%A7%C3%A3o%20simult%C3%A2nea,threads%20criadas%20por%20um%20%C3%BAnico)
Em palavras mais simples, concorrência trata-se da parte de software em que o mesmo possui uma ou várias funções que são executadas simultaneamente, gerando várias *threads*.

## Concorrência vs Paralelismo
**Concorrência** é sobre a execução sequencial disputada de um conjunto de tarefas independentes. Em sistema operacional o responsável pelo gerenciamento dos processos é o escalonador de processos, quanto que, em uma linguagem de programação, como *Go Lang*, o responsável por essa tarefa é chamado de *scheduler*. Escalonadores preemptivos favorecem a concorrência pausando e resumindo tarefas para que todas possam ser executadas, como é o caso dos sistemas operacionais modernos onde os processos e *threads* executam as trocas de contexto. Portanto, pode-se concluir que concorrência está relacionada com a parte do *software* e significa executar duas ou mais tarefas ao mesmo tempo.

**Paralelismo** trata-se da execução paralela de tarefas, ou seja, de forma simultânea, porém está realacionado com a parte de *hardware*, uma vez que depende da quantidade de núcleos(cores) do processador. Quanto mais núcleos, mais tarefas paralelas podem ser executadas. É uma forma de distribuir processamento em mais de um núcleo. Logo, conclui-se que paralelismo implica em concorrência, porém concorrência não implica em paralelismo, ou seja, para haver paralelismo é necessário ter concorrência, mas para ter concorrência não é necessário ter paralelismo, porque em sistemas com processadores de apenas um núcleo é possível ter processos concorrentes, basta ir pausando e resumindo uma tarefa para executar outra.

[Referência](https://www.treinaweb.com.br/blog/concorrencia-paralelismo-processos-threads-programacao-sincrona-e-assincrona)

## Goroutines & WaitGroups
Em linguagem Go, os processos que se dividem em uma ou mais tarefas concorrentes, as chamadas *Goroutines*, serão, ou sempre procurarão ser, executadas como processos concorrentes paralelos, aproveitando sempre o máximo do processador, utilizando o máximo de núcleos possíveis.

Para criar uma *Goroutine* basta acrescentar a palavra reservada `go` antes de qualquer função. É importante ressaltar que quando estamos trabalhando com *goroutines* é necessário esperar que a *goroutine* termine de executar, caso contrário o programa pode terminar, no caso a execução da `func main()`, e a execução da *goroutine* pode nem ter começado. Para isso utilizamos os *WaitGroups*. Então, temos o `sync.WaitGroup`, normalmente declarado como *package level escope*, que possui 3 funções principais:
- `Add()`: Função onde se deverá informar a quantidade total de *goroutines* na função em que está inserido. Normalmente é declarada no início da função *main*, por exemplo, antes da declaração das *goroutines*;
- `Wait()`: Função que informa para o programa esperar as *goroutines* executarem. Deve estar na última parte função *main*, por exemplo, antes do seu encerramento;
- `Done()`: Função que informa para o programa que uma determinada *goroutine* terminou a execução. Deve ser informada na última parte das *goroutines*.

[Exemplo de goroutine](./exemplos/cap18/exemplo_goroutine/goroutine.go)

**DICAS:**
1. Para saber o número de processadores basta executar `fmt.Println(runtime.NumCPU())` no início da função *main*.
2. Para saber o número de *Goroutines* basta executar `fmt.Println(runtime.NumGoroutine())`. Se for chamada no início o resultado será 1, pois a função *main* é uma *Goroutine* e se for chamada no final da função *main* irá retornar o número total de *Goroutines* chamadas durante todo o programa.

# Condição de corrida
Uma condição de corrida é uma falha num sistema ou processo em que o resultado do processo é inesperadamente dependente da sequência ou sincronia de outros eventos. Apesar de ser conhecido por "condição de corrida", uma melhor tradução seria "condição de concorrência", pois o problema está relacionado justamente ao gerenciamento da concorrência entre processos teoricamente simultâneos.

[Referência](https://pt.wikipedia.org/wiki/Condi%C3%A7%C3%A3o_de_corrida)

Em outras palavras, condição de corrida é quando o mesmo dado, ou conjunto de dados, é compartilhado por dois(ou mais) processos(*threads*) diferentes que estão sendo executados de maneira concorrente. Para resolver esse problema, Go  possui algumas aboradagens como *mutex*(exclusão mútua, garante que nenhuma outra fução irá utilizar essa variável até a mesma terminar de usá-la ou gravar algo nela), *atomic* e *channels*(canais), sendo a última a mais utilizada.

[Referência](https://golang.org/doc/effective_go#concurrency)

[Exemplo de uma condição de corrida](./exemplos/cap18/exemplo_condicaoDeCorrida/condicaoDeCorrida.go)

**OBS.:** Caso deseje visualizar se o códgio possui uma condição de corrida, adicione a *flag* `-race` na hora de executar o programa. Então fica `go run -race nomedoprograma.go`. No final será exibido na tela do terminal quantas condições de corridas há no programa informado.

## Mutex
Mutex(Exclusão mútua) garante que apenas um determinado trecho de código, ou um determinado valor, seja executado em um dado momento, evitando, assim, as condições de corrida. Os *Mutex* conseguem realizar essa tarefa "trancando" o valor de uma variável na qual irá ser acessada por várias *threads*. Utiliza-se *lock* para trancar e *unlock* para destrancar.

[Exemplo de Mutex](./exemplos/cap18/exemplo_mutex/mutex.go)

## Atomic
Da mesma maneira que os *Mutex*, *Atomic* também tem o objetivo de "trancar" um determinado valor de uma variável para evitar as condições de corrida. O pacote *Atomic* fornece funções de manipulação de memória para tipos primitivos, ideal para aplicações de baixo-nível que implementam sincronização de algorítmos. Com exceção dessas aplicações de baixo-nível, essa sincronização é melhor feita com canais(*channels*) ou com os próprios recursos do pacote de sincronização, o pacote `sync`. Em outras palavras, para evitar condições de corrida recomenda-se utilizar as funções do pacote `sync`, que é o caso do *Mutex* que já foi demonstrado, ou utilizar canais, os *channels*, *Atomic* só é recomendado para aplicações de baixo-nível.[Documentação](https://pkg.go.dev/sync/atomic)

[Exemplo de Atomic](./exemplos/cap18/exemplo_atomic/atomic.go)

# Comandos Go
Abaixo seguem alguns comandos da linguagem Go e a sua descrição.

| Comando | Descrição |
|---------|-----------|
| go version | Imprime na tela do terminal qual a versão da linguagem Go que está instalada |
| go env | Variáveis de sistema relacionadas à Linguagem Go |
| go help | Exibe uma lista de comandos utilizados na linguagem Go e suas descrições. É possível passar um nome de pacote como parâmetro para que seja exibido mais detalhes sobre o pacote. |
| go fmt | Executa `gofmt` nos pacotes principais(fontes). |
| gofmt | Esse comando realiza a formatação do código dos arquivos `.go` presentes nas pastas. |
| go run | Compila e executa um programa Go. É possível executar `go run *.go` para que todos os programas com a extensão `.go` sejam executados de uma vez. |
| go build | Compila pacotes e dependências. Ou seja, em casos de arquivos executáveis(que possuem a `func main`) gera um arquivo binário a partir do arquivo `.go` e, a partir desse arquivo, gera um executável e o salva no diretório(pasta) atual. Caso haja erros no programa `.go` o arquivo executável não será gerado e o erro será informado via terminal. Em casos de pacotes, gera um arquivo binário e, caso haja erros, informa-os e, ao final, descarta o arquivo executável.|
| go install | Compila e instala os pacotes e as dependências. De forma mais detalhada, para um arquivo executável, fará o *build*, em seguida, nomeará o arquivo com o nome do diretório atual e, por fim, salvará o arquivo binário em `$GOPATH/bin`. Para um pacote, fará o *build*, salvará o arquivo binário em `$GOPATH/pkg` e criará *archive files*, arquivos com extensão `.a` que são os arquivos pré-compilados utilizados pelos *imports*. |
| go clean | Remove arquivos. |
| go doc | Exibe a documentação para um determinado pacote. |
| go bug | Inicia um reporte de bug. |
| go fix | Executa `go tool fix` nos pacotes.|
| go generate | Gera arquivos Go a partir do processamento do arquivo fonte. |
| go get | Faz o download e a instalação dos pacotes e as dependências. |
| go list | Lista todos os pacotes instalados. |
| go test | Testa os pacotes |
| go tool | Executa uma ferramenta específica. |
| go vet | Executa `go tool vet` nos pacotes. Esse comando examina o cógido fonte dos arquivos `.go` e relata construções suspeitas, como chamadas à função Printf cujos os argumentos não se alinham com a sequência de formato esperada. O `vet` usa a heurística de que não há garantias de que todos os relatórios são problemas genuínos, mas podem conter erros que não foram capturados pelos compiladores. |
| golint | Esse comando examina os arquivos `.go` a procura de códigos mal escritos (*coding style* ruim). Caso esse comando não esteja previamente instalado, execute `go get golang.org/x/lint/golint` para instalá-lo. |

# Canais
Canais em Go é uma maneira de transmitir dados entre Goroutines. Há dois tipos de *Go Channels*, os *unbuffered*(sem buffer) e os *buffered*(com buffer). Da maneira *unbuffered* é a maneira padrão dos canais e significa que os canais só aceitarão enviar os dados de uma *goroutine* para outra se houver um receptor desses dados, ou seja, em palavras simples, uma *goroutine* se encarregará de enviar os dados pelo *channel* e outra *goroutine* se encarregará de receber esses dados. Da maneira *buffered* não é necessário ter uma *goroutine* para receber os dados enviados, porém aceita até um limite de valores para ser enviado dessa maneira.
- [FONTE 1](https://gobyexample.com/channel-buffering)
- [FONTE 2](https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855)

[Exemplo de canal unbuffered](./exemplos/cap21/unbuffered/unbufferedChannel.go)

[Exemplo de canal buffered](./exemplos/cap21/buffered/bufferedChannel.go)

## Canais direcionais
Os canais podem ser bidirecionais ou unidirecionais. Os unidirecionais servem para apenas uma determinada tarefa, ou seja, temos o *send channel*(`chan <- <tipodedado>`) que serve apenas para enviar dados e o *receive channel*(`<- chan <tipodedado>`) que serve apenas para receber dados. Essa funcionalidade impede que, por exemplo, dados sejam escritos em canais errados, em outras palavras, o *type-checking mechanisms* do compilador faz com que não seja possível, por exemplo, escrever em um canal de leitura. Já os bidirecionais podem enviar e receber dados sem problemas.

[Exemplo de unidirecional](./exemplos/cap21/canaisUnidirecionais.go)

## Range e Close()
A função close() serve para fechar a comunicação de um canal. Um canal *sender*, canal específico para enviar dados, pode ser fechado para indicar que não mais valores para serem enviados. Os canais *receivers*, específicos para receber dados, podem testar se um canal está fechado, atribuindo um segundo parâmetro à expressão, ou seja, `v, ok := <-ch`, se a variável `ok` for `false` então não há mais valores para receber, logo o canal *sender* está fechado.

A função `close()` se faz necessária quando estamos usando um *loop*(laço de repetição) em conjunto com canais. Então, quando temos `for i := range c` o *loop* recebe valores repetidamente até que o canal esteja fechado. Caso isso não ocorra, ou seja, caso o canal não seja fechado, o seguinte erro é mostrado `fatal error: all goroutines are asleep - deadlock!`. **OBS.:** Apenas os canais *sender* devem ser fechados, nunca os *receiver*. Enviar dados para um *sender* que está fechado causará um *panic*(tipo de erro em Go). Vale ressaltar que canais não são como arquivos, ou seja, geralmente canais não precisam ser fechados. Fechar canais só é necessário quando devemos dizer ao *receiver* que não há mais dados a receber, como quando usamos a função `range`, que ficará esperando o recebimento de novos dados, caso não fechemos o canal.[Referência](https://tour.golang.org/concurrency/4)

[Exemplo 01 - Fibonacci](./exemplos/cap21/Range_Close/exemplo01/close.go)

[Exemplo 02 - Contagem](./exemplos/cap21/Range_Close/exemplo02/close.go)

## Select
*Select* é como *switch*, porém para canais e não é sequencial. Ou seja, o *select* bloqueia os canais até que os dados que especificamos nele seja igual ao *case*(caso) e, aí sim, executa-os. Se houver mais de um caso que satisfaça a condição, então o *select* irá escolher um aleatoriamente.

[Exemplo 01](./exemplos/cap21/Select/exemplo01/select.go)

[Exemplo 02](./exemplos/cap21/Select/exemplo02/select.go)

[Exemplo 03](./exemplos/cap21/Select/exemplo03/select.go)

## Expressão *comma ok*
Como já mencionado antes é possível atribuir um segundo parâmetro aos canais para que possamos saber se estão abertos ou fechados e evitar que dados errados sejam lidos pelo canal, esse parâmetro é chamado de *comma ok*. O *comma ok* também é usado para *map* para que possamos saber se quando recebemos um valor 0 é porque foi passado esse valor ao *map*, ou trata-se de um valor que está faltando. Logo, para realizar essa tarefa usa-se uma variável, normalmente booleana, com nome `ok` e se ela for `true` significa que o valor foi de fato passado ao *map*, caso contrário, se for `false`, significa que o valor, no caso 0, representa um valor faltando. [Referência](https://golang.org/doc/effective_go#maps)

[Exemplo 01 exercício anterior](./exemplos/cap21/commaOK/exemplo01/commaOK_select.go)

[Exemplo 02](./exemplos/cap21/commaOK/exemplo02/commaOK_maps.go)

# Convergência
Convergência, neste caso, é quando a informação de vários canais é enviada a um número menor de canais.

[Exemplo por Todd](./exemplos/cap21/convergencia/exemplo01/convergencia.go)

[Exemplo Rob Pike](./exemplos/cap21/convergencia/exemplo02/convergencia.go)

# Divergência
Divergência, neste caso, é quando a informação de um único canal (ou poucos canais) é enviada a um número maior de canais. É o oposto de convergência.

[Exemplo 01](./exemplos/cap21/divergencia/exemplo01/divergencia.go)

[Exemplo 02](./exemplos/cap21/divergencia/exemplo02/divergencia.go)

# Context
Em servidores Go, cada solicitação(*request*) recebida é tratada em sua própria *goroutine*. Geralmente manipuladores de requisições(*Request Handlers*) iniciam *goroutines* para acessar *backends* como, por exemplo, banco de dados e servidores RPC(*Remote Procedure Call*), Chamada de Procedimento Remoto, trata-se de uma tecnologia de comunicação entre processos que permite um programa de computador chamar um procedimento em outro espaço de endereçamento(geralmente em outro computador, conectado a uma rede). [Mais sobre RPC](https://pt.wikipedia.org/wiki/Chamada_de_procedimento_remoto). O conjunto de *goroutines* trabalhando em uma requisição tipicamente precisa de acesso ao valor específico da requisição como a identidade do usuário final, tokens de autorização e a duração das requisições(*request's deadline*). Quando uma requisição é cancelada ou o tempo de duração da requisição é excedido(*time out*), todas as *goroutines* trabalhando naquela requisição deveriam sair(deixar de ser executada) rapidamente, para que o sistema possa recuperar quaisquer recursos que estavam sendo utilizados. Na Google foi desenvolvido o *package context* que torna fácil a passagem de valores de escopo de requisição(*request's scoped values*), sinais de cancelamento e *deadlines* através dos limites da API para todas as *goroutines* envolvidas no tratamento da requisição. O pacote está publicamente disponível como `context`. [Sobre Context](https://blog.golang.org/context). [Sobre o package context](https://golang.org/pkg/context)

- [Exemplo usando a função WithCancel()](./exemplos/cap21/context/exemplo01/context.go) -> [Referência](https://play.golang.org/p/Lmbyn7bO7e)
- [Exemplo 02 usando a função WithCancel()](./exemplos/cap21/context/exemplo02/context.go) -> [Referência](https://play.golang.org/p/wvGmvMzIMW)
- [Exemplo usando a função WithDeadline()](./exemplos/cap21/context/exemplo03/context.go) -> [Referência](https://play.golang.org/p/Q6mVdQqYTt)
- [Exemplo usando a função WithTimeout()](./exemplos/cap21/context/exemplo04/context.go) -> [Referência](https://play.golang.org/p/OuES9sP_yX)
- [Exemplo usando WithValue()](./exemplos/cap21/context/exemplo05/context.go) -> [Referência](https://play.golang.org/p/8JDCGk1K4P)

# Tratamento de erros
Em linguagem GO, ao contrário das outras linguagens, não há exceções(*exceptions*). A linguagem tarta essas exceções e erros de uma forma diferente, utiliza o retorno de múltiplos valores das funções para realizar esse tratamento. Então, em uma dada função, existe um valor do tipo *error type* e caso ele seja nulo(*nil*) significa que não houve nenhum erro e o programa pode prosseguir, caso contrário o erro exibido e o programa é interrompido. Vale ressaltar que é recomendado tratar os erros na mesma hora, ou seja, assim que chamar uma função já verifique se há erros e os trate. [Fonte](https://golang.org/doc/faq#exceptions)

**Exemplos:**
- Funções como `fmt.Println()` não se costuma verificar o erro, pois para fazer isso teríamos que chamar novamente a função `fmt.Println()` para exibir o erro e ao fazer isso estaríamos criando outro erro. Logo essa função é uma exceção à regra. [Exemplo fmt.Println()](./exemplos/cap23/verificandoErros/exemplo00/main.go)
- [Função fmt.Scan()](./exemplos/cap23/verificandoErros/exemplo01/main.go)
- [Escrevendo em um arquivo txt](./exemplos/cap23/verificandoErros/exemplo02/main.go)
- [io.ReadAll()](./exemplos/cap23/verificandoErros/exemplo03/main.go)

## Print & Log
Opções:
- `fmt.Println()` -> stdout
- `log.Println()` -> timestamp
- `log.SetOutput()`
- `log.Fatalln()` -> chama a função os.Exit(1)
- `log.Panicln()` -> é a função println + panic
- `panic()`

[Código de exemplo](./exemplos/cap23/tratamentoDeErros/exemplo01/main.go)

[Exemplo do SetOutput()](./exemplos/cap23/tratamentoDeErros/exemplo02/main.go)

Fonte:
- [Go Doc](https://godoc.org/builtin#panic)

## Recover
- [Blog Golang](https://blog.golang.org/defer-panic-and-recover)
- [Golang](https://golang.org/pkg/builtin/#recover)

**Exemplos:**
- [Exemplo 01](./exemplos/cap23/Recover/exemplo01/main.go)
  
## Erros com informações adicionais
Para que as funções retornem erros customizados, utilizamos:
- `return errors.New()`
- `return fmt.Errorf()` -> [Documentação](https://golang.org/pkg/builtin/#error)
  
**Exemplos:**
- [Exemplo 01 - errors.New()](./exemplos/cap23/errosCustomizados/exemplo01/main.go)
- [Exemplo 02 - var errors.New()](./exemplos/cap23/errosCustomizados/exemplo02/main.go)
  - [Outro exemplo](https://golang.org/src/pkg/bufio/bufio.go)
  - [Outro exemplo](https://golang.org/src/pkg/io/io.go)
- [Exemplo 03 - fmt.Errorf()](./exemplos/cap23/errosCustomizados/exemplo03/main.go)
- [Exemplo 04 - var fmt.Errorf()](./exemplos/cap23/errosCustomizados/exemplo04/main.go)
- [Exemplo 05](./exemplos/cap23/errosCustomizados/exemplo05/main.go)
  - [Outro exemplo](https://golang.org/pkg/net/#OpError)
  - [Outro exemplo](https://golang.org/src/pkg/net/dial.go)
  - [Outro exemplo](https://golang.org/src/pkg/net/net.go)
  - [Outro exemplo](https://golang.org/src/pkg/encoding/json/decode.go)

# Documentação
Quando falamos em documentação em Go temos as seguintes opções:
-  [golang.org](https://golang.org/doc) que possui a documentação das bibliotecas padrão do Go
-  [godoc.org](https://godoc.org) que possui a documentação das bibliotecas padrão e de outras bibliotecas.
-  Podemos usar o comando `go doc` para ler a documentação na linha de comando.
-  Outro comando é o `godoc` que também permite ler a documentação na linha de comando, mas possui uma opção que permiti abrir um servidor local para ler a documentação.

## go doc
O comando `go doc` demonstra a domcumentação de um package, const, func, type, var, método, etc. O comando aceita zero, um ou dois argumentos, sendo que se não for passado nenhum argumento(zero argumento) será exibido a documentação do pacote(*package*) da pasta atual, se a pasta não tiver nenhum pacote go será exibido um erro, "*no buildable Go source files in <caminho-da-pasta>*" e se o pacote for um comando(package main), os símbolos exportados do pacote serão eliminados da apresentação a não ser que a *flag* `-cmd` seja fornecida. Para mais informações execute o comando `go help doc`.

## godoc
O comando `godoc` é parecido com o `go doc`, porém o que os diferencia é a flag `-http` que possibilita a exibição da documentação em um servidor local. A documentação apresentada por esse comando será semelhante à do [golang.org](https://golang.org/doc). Para executar esse comando escreva no terminal `godoc -http:=<porta-do-servidor-local>`. Para mais informações acessar [godoc.org](https://pkg.go.dev/golang.org/x/tools/cmd/godoc).

## Escrevendo documentação
Para escrever a documentação de um programa que desenvolvemos, temos que seguir o seguinte padrão descrito em [golang.org](https://blog.golang.org/godoc-documenting-go-code). Basicamente, na linguagem Go, é utilizado a ferramenta `godoc` para a documentação de projetos. O Godoc simplesmente analisa o código-fonte (incluindo os comentários) e produz a documentação como HTML ou texto puro. O resultado final é uma documentação fortemente acoplada ao código que ela documenta. Por exemplo, por meio da interface da web do godoc, você pode navegar da documentação de uma função até sua implementação com um clique.

A convenção é simples: para documentar um tipo(`type`), variável, constante(`const`), função(`func`) ou mesmo um pacote(`package`), escreva um comentário regular diretamente antes de sua declaração, sem nenhuma linha em branco intermediária, ou seja, não pode haver linhas em branco entre o comentário e o código. O Godoc apresentará esse comentário como texto ao lado do item que ele documenta. Por exemplo, esta é a documentação para a função Fprint do pacote fmt:

```go
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

Observe que este comentário é uma frase completa que começa com o nome do elemento que descreve. Essa importante convenção nos permite gerar a documentação em uma variedade de formatos, de texto simples a HTML e páginas de manual do UNIX, e torna a leitura melhor quando as ferramentas a truncam para abreviar, como quando extraem a primeira linha ou frase.

Os comentários sobre as declarações do pacote devem fornecer uma visão geral da documentação do pacote. Esses comentários podem ser curtos, como uma breve descrição do pacote:

```go
// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort
```

Eles também podem ser detalhados como a visão geral do pacote gob(*gob package*). Esse pacote usa outra convenção para pacotes que precisam de grandes quantidades de documentação introdutória: o comentário do pacote é colocado em seu próprio arquivo, `doc.go`, que contém apenas esses comentários e uma cláusula de pacote.

Ao escrever comentários de pacote de qualquer tamanho, lembre-se de que sua primeira frase aparecerá na lista de pacotes do godoc.

Comentários que não são adjacentes a uma declaração de nível superior são omitidos da saída do godoc, com uma exceção notável. Comentários de nível superior que começam com a palavra "BUG (quem)" são reconhecidos como bugs conhecidos e incluídos na seção "Bugs" da documentação do pacote. A parte "quem" deve ser o nome de usuário de alguém que poderia fornecer mais informações. Por exemplo, este é um problema conhecido do pacote de bytes:

```go
// BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
```

Às vezes, um *struct*, uma função(*func*), um *type* ou até mesmo um pacote inteiro se torna redundante ou desnecessário (obsoleto), mas deve ser mantido para compatibilidade com programas existentes. Para sinalizar que um identificador não deve ser usado, adicione um parágrafo ao comentário do documento que começa com `Deprecated:` seguido por algumas informações sobre a remoção.

Para mais informações sobre como escrever a documentação acesse [golang.org](https://go.dev/blog/godoc)

## Extra
Para exibir sua documentação de forma online, acesse o [godoc.com](https://pkg.go.dev/?utm_source=godoc) e, no campo de busca por pacotes Go, cole o link completo do repositório do GitHub onde se encontra o seu software na qual deseja documentar.

# Testes em Go
Testar programas se faz necessário, uma vez que, escrever um programa que faça exatamente o que foi planejado para fazer, todas as vezes, é uma tarefa difícil, as vezes impossível, por isso é necessário testar o programa para verificar se o resultado obtido é o que se esperava receber do programa.

Em Go para realizar testes de pacotes criados previamente, precisamos criar um arquivo go dentro da pasta (ou pacote) que se deseja testar, denominado `<nomeDoPacote>_test.go`.

**IMPORTANTE:** 
1. O arquivo go a ser testado e o arquivo de teste precisam estar na mesma pasta e o arquivo de teste precisa ter o nome no formato `<nomeDoPacoteASerTestado>_test.go`.
2. Pode ser que seja necessário executar o comando `go mod init <caminhoDoGitHubDoProjeto>`. Esse comando irá criar um arquivo chamado `go.mod` que nele há o caminho do projeto e a versão da linguagem Go utilizada. Vale ressaltar que, para não ter que executar esse comando todas as vezes que for executar um teste em uma pasta diferente, execute o comando de forma que o arquivo `go.mod` fique na raíz do projeto, desta forma, só será preciso realizar esse procedimento apenas uma única vez por projeto. Se preferir executar todas as vezes para cada pasta do projeto, então o comando ficará assim: `go mod init <caminhoDoGitHubDaPastaComOTeste>`.

## Escrevendo Testes
Para escrever um teste de uma função, segue-se a seguinte regra: `func TestNomeDaFuncaoASerTestada(t *testing.T) {"Escreva seu teste aqui"}`. Observe que na função de teste, ela sempre deve começar com a palavra *Test*, seguida do nome da função que queremos testar. Essa função sempre receberá um ponteiro testing.T como parâmetro. Podemos escrever quantos testes acharmos necessários para um determinado pacote.

Para executar os testes escritos executamos o comando `go test`, lembrando que no terminal devemos estar dentro da pasta onde estão os testes e o programa principal. Ou podemos executar o comando `go test -v` para visualizar cada teste que foi executado e visualizar qual deles falhou, ou seja, esse comando nos dá uma visão mais detalhada da execução dos testes.

**Fontes:**
- [Golang Book](https://www.golang-book.com/books/intro/12)
- [Go Testing](https://pkg.go.dev/testing?utm_source=godoc)

**Exemplos:**
- [Exemplo 01](./exemplos/cap27/01-testes/exemplo01/main_test.go)

## Testes em série
É possível escrever testes em série, em Go chamamos isso de testes em tabela e o usamos para testar vários cenários de uma mesma funcionalidade.

**Exemplo:**
-[Exemplo 02](./exemplos/cap27/01-testes/exemplo02/main_test.go)

## Testes como exemplos
Nessa situação executamos os testes como exemplos, que são os mesmos exemplos que aparecem na documentação Godoc.

Exemplos de Godoc são trechos do código Go que são exibidos como documentação do pacote e que são verificados executando-os como testes. Eles também podem ser executados por um usuário que visita a página da web godoc para o pacote e clica no botão "Executar" associado. Ter documentação executável para um pacote garante que as informações não ficarão desatualizadas à medida que a API muda. [Fonte: Testable examples in Go](https://go.dev/blog/examples).

Quando formos escrever essa documentação executável devemos usar a seguinte nomenclatura: `func ExampleNomeDaFuncao() {}`. Onde o nome da função deve começar com *Example* e logo em seguida deve haver um comentário informando o resultado esperado pela execução da função, no formato `// Output: resultado`. Vale lembrar que para visualizar o exemplo na documentação executamos o código: `godoc -http :<porta>` e para executá-los, assim como fazemos para testes, rodamos o comando `go test`, ou para executar todos os testes da pasta de uma vez, usamos `go test ./ ...`.

**Lembrete:** Para instalar o *godoc* use o comando: `go get golang.org/x/tools/cmd/godoc`

**Exemplos:**

```go
package stringutil_test

import (
    "fmt"

    "golang.org/x/example/stringutil"
)

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
```

- [Exemplo 03 - Testes como exemplos](./exemplos/cap27/01-testes/exemplo03/main_test.go)

# *Benchmarks*
*Benchmarks*, em computação, é o ato de executar um programa de computador a fim de avaliar o desempenho relativo de um objeto, normalmente executando uma série de testes padrão. Em outras palavras, neste contexto, *Benchmark* trata-se de avaliar o desempenho do *software* desenvolvido, assim como suas funções e rotinas.

Assim como nos testes, as funções que serão avaliadas nos testes de *benchmarks* precisam seguir uma nomenclatura específica, neste caso é: `func BenchmarkNomeDaFuncao(b *testing.B){}`. Observa-se que estamos utilizando o pacote `testing` e passando como parâmetro um ponteiro do tipo `testing.B`, parecido com que já fazíamos com os testes comuns.

Para executar os testes de *benchmarks* usamos o comando `go test -bench .` para executar todas as funções de *benchmark* presentes na pasta ou `go test -bench <nomeDaFuncao>` para testar apenas um única função.

**Dica:**
- Execute `go help testflag` para ver a documentação sobre as flags que podem ser passadas nos comandos `go test`.

**Exemplos:**
-[Exemplo 01](./exemplos/cap27/02-02-benchmark/exemplo01/main_test.go/)

# Cobertura dos testes
Para saber quatos por cento do código foi testado, podemos executar o comando `go test -coverprofile <arquivo>`, onde `<arquivo>` é o nome do arquivo que será salvo contendo o relatório da cobertura de testes. Caso não deseje salvar o relatório de cobertura execute apenas `go test -cover`. Ainda é possível converter o arquivo com o relatório de cobertura para `.html` executando o comando `go tool cover -html=<arquivo>`. Nesse arquivo HTML, além de obetermos a porcentagem do código que está sendo coberta por testes, podemos observar que a parte do código que está em vermelho não está sendo coberta por testes, enquanto que a parte em verde está.

Vale ressaltar que você deverá acesar a pasta, pelo terminal, onde os arquivos de teste se encontram, caso queira gerar o arquivo com o relatório de cobertura, esse arquivo deverá ter a extensão `.out` e, se você estiver usando o Sistema Operacional(SO) Windows, você deverá fornecer o nome do arquivo entre aspas duplas(""). Logo temos:

| Comando | Resultado |
|---------|-----------|
| `go test -cover` | Esse comando irá avaliar quantos por cento do código foram cobertos por testes. |
| `go test -coverprofile <nomeDoArquivo>` | Esse comando irá executar o comando `go test -cover` e salvará o resultado em um arquivo cuja extensão deve ser `.out`. |
| `go tool cover -html=<nomeDoArquivo>` | Esse comando seleciona um arquivo `.out` que possui um relatório de cobertura de testes e o converte para `.html` e o exibe no navegador. O arquivo HTML não será salvo. |
| `go tool cover -help` | Exibe uma documentação de ajuda para entender melhor as flags e os comandos do `go tool cover`. |

**Observação:** Normalmente, em projetos reais, não se atinge uma cobertura de 100%, aceita-se uma porcentage de 70 a 80%. Isso pode variar dependendo do projeto.

# Referências
- Korbes, Ellen. **Aprenda Go 🇧🇷**. Aprenda Go. Disponível em: https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg
- Go by example. **Go by Example**. Disponível em: https://gobyexample.com/
- fatih. **color**. Disponível em: https://github.com/fatih/color
- GO. **The Go Programming Language Specification**. Disponível em: https://golang.org/ref/spec
- GO. **Effective Go**. Disponível em: https://golang.org/doc/effective_go
- Tutorials Point. **Tutorial GO**. Disponível em: https://](https://www.tutorialspoint.com/go/go_functions.htm)
- rakyll. **Style guideline for Go packages**. Disponível em: https://rakyll.org/style-packages/