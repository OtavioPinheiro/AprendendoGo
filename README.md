# Linguagem GO
Aprendendo GO Lang por meio de vídeos tutoriais.

# Tópicos
- [O que é GO Lang?](#o-que-é-go-lang)
- [Começando o estudo](#começando-o-estudo)
- [Constantes](#constantes)
  
## O que é Go Lang?
A linguagem Go é uma linguagem de programação criada pela Google e lançada em códgio livre em 2009. É uma linguagem compilada e focada em produtividade e programação concorrente, baseada em trabalhos feitos no sistema operacional chamado Inferno.

## Começando o estudo.

### Regras  e características da linguagem
1. Não é permitido declarar uma variável e não utilizá-la. Isso gerará um erro, pois fere um dos princípios da linguagem que é possuir um código limpo.
2. Go é uma lingaugem de tipagem estática.
3. Quando uma variável é declarada fora de um codeblock scope e não é atribuída a um valor neste momento, ela só poderá receber um valor dentro de um codeblock.
4. Slice array, struct, map -> Tipos de dados compostos
5. Todo valor em Go pertence também ao tipo de interface vazia, representado pela notação "interface{}"
6. Em Go strings são imutáveis. Para alterar uma string é necessário criar outra string com a alteração desejada.
7. Não existe ***While*** em GO.

### Operador curto de declaração
É representado por `:=`. É utilizado na declaração de variáveis. O diferencial é que esse operador possui tipagem automática, ou seja, não é necessário informar o tipo da variável quando utilizamos o operador curto de declaração e só podemos utilizar esse operador dentro de um *codeblocks* (escopo).

Esse operador é diferente do operador de atribuição `=`. Cuidado para não confundí-lo, pois são utilizados em contextos diferentes.

### Blank identifier
O blank identifier serve para dizer ao programa ignorar um retorno de informações recebido por uma função. É representado pelo caractere `_`.

### Variáveis
Variáveis em GO podem ser declaradas das seguintes formas:
```
  var i int -> Declaração da variável
  i = 42 -> Inicialização da variável

  var j int = 27

  k := 99
```

### Valor zero
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

### Raw string literal e Interpreted string literal
*Interpreted String Literal* são strings interpretadas.<br>**Exemplo:** `fmt.Printf("Olá!\nTudo bem?\tEspero que sim.")`

Já o *Raw String Literal* são strings cruas que não serão interpretadas.<br>**Exemplo:** ```fmt.Printf(`Olá!\nTudo bem?\tEspero que sim.`)```

**Atenção**: Para Raw Strings utiliza-se o acento grave ``.

### Tipo string (cadeia de caracteres)
- Strings são sequências de bytes.
- são imutáveis
- Uma string é uma fatia de bytes (*slice of bytes*)

### Pacote fmt
- **fmt.Print():** Exibe na tela a string passada como argumento.
- **fmt.Printf():** Exibe na tela a string passada como argumento. Permite formatações.
- **fmt.Println():** Exibe na tela a string passada como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Sprint():** Retorna uma concatenação de strings passadas como argumento.
- **fmt.Sprintf():** Retorna uma concatenação de strings passadas como argumento. Permite formatações.
- **fmt.Sprintln():** Retorna uma concatenação de strings passadas como argumento. Insere uma quebra de linha (\n) ao final da string.
- **fmt.Fprint():** Escreve a string, passada como argumento, em um arquivo.
- **fmt.Fprintf():** Escreve a string, passada como argumento, em um arquivo. Permite formatações.
- **fmt.Fprintln():** Escreve a string, passada como argumento, em um arquivo. Insere uma quebra de linha (`\n`) ao final da string.

### Iota
São números sequênciais atribuídos automáticamente pelo sistema.

**Exemplos:** ```const (a = iota, b = iota, c = iota)```

**Resultado:** ```a = 0, b = 1, c = 2```

### Switch e case
É igual para todas as linguagens que utilizam *switch case*, ou seja, o *switch* irá comparar o valor da variável informada com o valor declarado nos *cases*, caso verdadeira, executará o código dentro dos respectivos *cases*, senão seguirá para o *default*. Uma diferença é que em GO existe o termo `fallthrough` que faz com que o próximo *case* também seja executado caso o *case* com *fallthrough* seja verdadeiro.

# Referências
- Korbes, Ellen. **Aprenda Go 🇧🇷**. Aprenda Go. Disponível em: https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg
- Go by example. **Go by Example**. Disponível em: https://gobyexample.com/
- fatih. **color**. Disponível em: https://github.com/fatih/color
- GO. **The Go Programming Language Specification**. Disponível em: https://golang.org/ref/spec
- GO. **Effective Go**. Disponível em: https://golang.org/doc/effective_go