# REST API com Golang e PostgreSQL

### API de registros de livros

Esta API é construída usado o pacote [net/http](https://pkg.go.dev/net/http) e [Gorilla mux](https://github.com/gorilla/mux). È baseado em um banco de dados [PostgreSQL](https://www.postgresql.org/).

## Começando

### Clone o projeto

```bash
git clone https://github.com/Johnson49/rest-api-com-golang-e-postgreSQL
```

#### Adentre no diretório 
```shell
cd rest-api-com-golang-e-postgreSQL
```

## Criando o Banco de dados

### Localmente

Caso já possua o postgres instalado em sua máquina.

1. Vá até a pasta `sql` e execute o `snippets`.

2. Confirme se o banco de dados e a tabela foram criadas corretamente.

### Docker

Caso prefira utilizar o postgres pelo docker.

#### 1. Criamos uma imagem com o arquivo Dockerfile que está localizado na raiz do projeto.

```shell
docker image build -t db_api-golang .
```


#### 2. Após o build, criasse um container com a imagem.

```
docker container run -d -p 5432:5432 --name api_postgres db_api-golang
```


#### 3. E finalmente, com este script que copiamos para dentro do container, irá criar uma nova role chamada `docker`, o banco de dados `dbgolang` e a tabela `livros`. Além disso, também irá registrar o primeiro livro como teste.

```
docker container exec -it api_postgres bash -c "sh /tmp/init.db.sh"
```


##  Inicie o servidor 


```
go run src/main.go
```

## EndPoints

> As rotas são compostas pelo endereço base (http://localhost:8000) mais o recurso que você deseja acessa.

|Método|Rota| Funcionalidade| Acesso |
|:-------:|:-----:|:------:|:------:|
|GET | /biblioteca/ | Consulta todos os livros existentes.| Público |
|GET |  /biblioteca/livro?id= | Consulta um livro pelo seu id| Público |
|POST | /biblioteca/registro | Registrar um novo livro. | Público |
| PUT | /biblioteca/livro?id= | Atualiza os dados de um livro.| Público |
| DELETE | /biblioteca/livro?id= |  Excluir um livro. | Público |

