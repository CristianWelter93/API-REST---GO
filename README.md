# api-rest-go

Estudos sobre api Restful utilizando GO. Foi criada uma api que é basicamente um CRUD de pessoa, os dados são consultados e salvos em um banco MySql.

## Pacotes instalados

Necessário a instalação das seguintes dependências.

### Mux

Para então poder executar o comando que instala de fato o pacote

~~~bash
go get github.com/gorilla/mux
~~~

### gorm

Um ORM para realizar a conexão com o MySql. É necessário a previa existência do banco de dados e do DataBase a ser utilizado para poder ser realizada a conexão.

~~~bash
go get github.com/jinzhu/gorm
~~~

### mysql driver

~~~bash
go get github.com/go-sql-driver/mysql
~~~

## Instruções para execução

Entrar na pasta do projeto e executar o comando

~~~bash
go run main.go
~~~

Isto irá executar o código e permitir o acesso a api pelo endereço localhost:8080

Utilizando o Postman ou outro produto similar é possível fazer as seguintes operações:

### GetAll - Buscar todas as pessoas cadastradas

### GetById - Busca uma pessoa em específico através do seu Id

### Create - Adicionar uma nova pessoa

### UpdateById - Atualizar um ou mais dados de uma pessoa já cadastrada

### UpdateNameById - Atualiza o nome de uma pessoa já cadastrada

### DeleteById - Excluir dados de uma pessoa já cadastrada
