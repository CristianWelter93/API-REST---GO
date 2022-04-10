# api-rest-go

Estudos sobre api Restful utilizando GO

## Execute

### Mux

É necessário instalar o pacote do mux, ele implement um roteador e um despachante de solicitação.
No meu caso foi necessário primeiro rodar

~~~bash
go env -w GO111MODULE=off
~~~

Para então poder executar o comando que instala de fato o pacote

~~~bash
go get github.com/gorilla/mux
~~~
