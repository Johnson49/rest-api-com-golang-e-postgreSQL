# API REST 

> Status do Projeto: :heavy_check_mark: Concluído.


## Informações gerais
API de livros que usa um arquivo json como banco de dados.


## Tecnologias 
> O projeto é criado com:

* Go
* gorilla/mux


## Setup: 
> Para rodar este projeto, clone localmente e depois instale as dependências.

```go
$ git clone https://github.com/Johnson49/api-rest-simples
$ cd api-rest-simples
$ go get -u github.com/gorilla/mux
$ go run main.go
```  

## EndPoints
> As rotas são compostas pelo endereço base (localhost:port) mais o recurso que você deseja acessa.

|Request|URL| Detalhes|
|:-------:|:-----:|:------:|
|GET | /biblioteca/ | Busca todos os livros|
|GET |  /biblioteca/1 | Busca o livro de ID 1|
|POST | /biblioteca/add | Adiciona um livro |
| PUT | /biblioteca/1 | atualiza o livro de ID 1|
| DELETE | /biblioteca/1 | remove o livro de ID 1|


## Licença

The [MIT License]() (MIT)

Copyright :copyright: 2022 - **API REST**
