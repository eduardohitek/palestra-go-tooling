## go build

## Compila o código e suas dependência e gera um executável baseado no parametros informados;

`go build .`: **Cria um executável baseado nas informações do seu go.mod;**

`go build -o=foo .`: **O nome do executável gerado será foo;**

`GOOS=windows GOARCH=amd64 go build -o=windows_amd64/ .`: **Efetua Cross-Compilation e gera um executável para Plataforma Windows;**

`go tool dist list`: **Lista as opções para Cross-Compilation;**
