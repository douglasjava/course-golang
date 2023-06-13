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


## Boas práticas
- Função que serão "publicas" escritas com a letra maicuscula deverão sempre ter um comentario