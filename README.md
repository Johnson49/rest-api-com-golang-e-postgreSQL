# REST API com Golang e PostgreSQL


Esta API é construída usado o pacote [net/http](https://pkg.go.dev/net/http) e [Gorilla mux](https://github.com/gorilla/mux). È baseado em um banco de dados [PostgreSQL](https://www.postgresql.org/).

## Começando

### Clone o projeto e instale as dependências

`git clone`

####  Instale as dependências

```go
cd 
go mod download

```

## Criando o Banco de dados

> ⚠️ Antes de prosseguir, confirme se possui o postgresSQL instalado na máquina.

1. Vá até a pasta `sql` e execute a `query` já pronta.

2. Confirme se o banco de dados e a tabela foram criadas corretamente.

##  Inicie o servidor 

`go run src/main.go`


## EndPoints

> As rotas são compostas pelo endereço base (http://localhost:8000) mais o recurso que você deseja acessa.

|Método|Rota| Funcionalidade| Acesso |
|:-------:|:-----:|:------:|:------:|
|GET | /biblioteca/ | Consulta todos os livros existentes.| Público |
|GET |  /biblioteca/livro?id= | Consulta um livro pelo seu id| Público |
|POST | /biblioteca/registro | Registrar um novo livro. | Público |
| PUT | /biblioteca/livro?id= | Atualiza os dados de um livro.| Público |
| DELETE | /biblioteca/livro?id= |  Excluir um livro. | Público |

