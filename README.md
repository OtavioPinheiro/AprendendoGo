# Linguagem GO
Aprendendo GO Lang por meio de vídeos tutoriais.

# Tópicos
- [O que é GO Lang?](#o-que-é-go-lang)
- [Começando o estudo](#começando-o-estudo)
- [Constantes](#constantes)
  
## O que é Go Lang?
A linguagem Go é uma linguagem de programação criada pela Google e lançada em códgio livre em 2009. É uma linguagem compilada e focada em produtividade e programação concorrente, baseada em trabalhos feitos no sistema operacional chamado Inferno.

## Regras  e características da linguagem
1. Não é permitido declarar uma variável e não utilizá-la. Isso gerará um erro, pois fere um dos princípios da linguagem que é possuir um código limpo.
2. Go é uma lingaugem de tipagem estática.
3. Quando uma variável é declarada fora de um codeblock scope e não é atribuída a um valor neste momento, ela só poderá receber um valor dentro de um codeblock.
4. Slice array, struct, map -> Tipos de dados compostos
5. Todo valor em Go pertence também ao tipo de interface vazia, representado pela notação "interface{}"
6. Em Go strings são imutáveis. Para alterar uma string é necessário criar outra string com a alteração desejada.
7. Não existe ***While*** em GO.

## Operador curto de declaração
É representado por `:=`. É utilizado na declaração de variáveis. O diferencial é que esse operador possui tipagem automática, ou seja, não é necessário informar o tipo da variável quando utilizamos o operador curto de declaração e só podemos utilizar esse operador dentro de um *codeblocks* (escopo).

Esse operador é diferente do operador de atribuição `=`. Cuidado para não confundí-lo, pois são utilizados em contextos diferentes.

## Blank identifier
O blank identifier serve para dizer ao programa ignorar um retorno de informações recebido por uma função. É representado pelo caractere `_`.

## Variáveis
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

## Raw string literal e Interpreted string literal
*Interpreted String Literal* são strings interpretadas.<br>**Exemplo:** `fmt.Printf("Olá!\nTudo bem?\tEspero que sim.")`

Já o *Raw String Literal* são strings cruas que não serão interpretadas.<br>**Exemplo:** ```fmt.Printf(`Olá!\nTudo bem?\tEspero que sim.`)```

**Atenção**: Para Raw Strings utiliza-se o acento grave ``.

## Tipo string (cadeia de caracteres)
- Strings são sequências de bytes.
- são imutáveis
- Uma string é uma fatia de bytes (*slice of bytes*)

## Pacote fmt
- **fmt.Print():** Exibe na tela a string passada como argumento.
- **fmt.Printf():** Exibe na tela a string passada como argumento. Permite formatações.
- **fmt.Println():** Exibe na tela a string passada como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Sprint():** Retorna uma concatenação de strings passadas como argumento.
- **fmt.Sprintf():** Retorna uma concatenação de strings passadas como argumento. Permite formatações.
- **fmt.Sprintln():** Retorna uma concatenação de strings passadas como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Fprint():** Escreve a string, passada como argumento, em um arquivo.
- **fmt.Fprintf():** Escreve a string, passada como argumento, em um arquivo. Permite formatações.
- **fmt.Fprintln():** Escreve a string, passada como argumento, em um arquivo. Insere uma quebra de linha (`\n`) ao final da string.

## Iota
São números sequênciais atribuídos automáticamente pelo sistema.

**Exemplos:**
```go 
const (a = iota, b = iota, c = iota)
```

**Resultado:** `a = 0, b = 1, c = 2`

## Switch e case
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

## Função range
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

## Dados compostos
Dados compostos são qualquer tipo de dados que podem ser construídos em um programa utilizando dados primitivos da programação ou outro tipo de dados compostos.

- ***Arrays:*** Arrays são vetores de número finito. Podem ser vetores de *strings*, *integers*, *floats*, etc. Arrays são dados compostos.
- ***Slices:*** Slices são ponteiros de arrays (conjunto de dados) que podem ser compostos por *strings*, *integers*, *floats*, etc. Logo *slices* são um tipo de dados compostos. Slices, quando são declarados, são como um array de tamanho "infinito".

**Exemplo de declaração de um array e de um slice:**
  ```go
  array := [5]int{1, 2, 3, 4, 5}
  slice := []int{1, 2, 3, 4, 5}
  ```

## Manipulando Slices
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

### Slice slices (fatia de fatias)
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
  <pre><code>[1 2 <span style = "color:blue">3</span> 4 5]
[1 2 5]
[1 2 <span style = "color: red">5</span> 4 5]</code></pre>

**Dica:** Se for precisar criar um *slice* a partir de outro *slice* e não quer que os dados do *slice* original se percam, utilize um laço *for* para realizar essa tarefa e copie elemento a elemento do *slice* antigo para o novo. Ou então utilize a mesma variável para realizar o corte e criar um slice subjacente.

### Função ***make([]T, len, cap)***.
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

## Slice Multidimensional
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

# Referências
- Korbes, Ellen. **Aprenda Go 🇧🇷**. Aprenda Go. Disponível em: https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg
- Go by example. **Go by Example**. Disponível em: https://gobyexample.com/
- fatih. **color**. Disponível em: https://github.com/fatih/color
- GO. **The Go Programming Language Specification**. Disponível em: https://golang.org/ref/spec
- GO. **Effective Go**. Disponível em: https://golang.org/doc/effective_go