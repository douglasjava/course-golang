## Anotações

- Todas as variaveis e métodos são visiveis nos arquivos do mesmo pacote.
- O arquivo executavel obrigatóriamento tem que está no pacote main
- métodos publicos escreve com a primeira letra minuscula onde a 
visibilidade será somente dentro do pacote que a função estiver, com a letra maiculas 
podemos exportar em outrso pacotes
- Fortemente tipada
- O erro é um tipo 
- funções podem ter mais de um retorno

## Comandos
> go run arquivo -> Executa o código
> 
> go mod init nome_modulo -> cria um modulo
> 
> go build -> compila o modulo
> 
> go install -> compila e salva um build na raiz da instalação
> 
> go mod tidy -> remove todas as dependencias não utilizadas
> 
>

## Funções
go -> go funcao() -> executa esse método, mas não esperar o retorno


## Boas práticas
- Função que serão "publicas" escritas com a letra maicuscula deverão sempre ter um comentario


## Criando uma api simples

- Initialize a new Go module by running the following command in your project directory:
> go mod init <module-name>
- Create a new file called main.go in your project directory.
>   package main

    import (
    "fmt"
    "log"
    "net/http"
    )
    
    func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
    }
    
    func main() {
    http.HandleFunc("/", helloWorldHandler)
    
        log.Println("Server listening on port 8080...")
        log.Fatal(http.ListenAndServe(":8080", nil))
    }
- In your terminal, navigate to the project directory and build the Go program by running the following command:
> go build
- successfully building the program, run it using the following command:
> ./<module-name>