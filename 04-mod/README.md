# go mod

## Comando para criar uma configuração básica de uma aplicação e para o gerenciamento das suas dependências;

`go mod init nome_do_app`: **Criar o arquivo go.mod para a aplicação atual;**

`go get -u github.com/google/uuid`: **Adicionar a última versão do pacote selecionado como dependência da aplicação;**

`go get -u github.com/google/uuid@latest`: **Adicionar a última versão do pacote selecionado como dependência da aplicação;**

`go get -u github.com/google/uuid@v1.0.0`: **Adiciona a versão específica do pacote selecionado;**

`go list -m -u all`: **Lista se há possíveis atualizações na lista de pacotes;**

`go mod tidy`: **Lista a lista de dependência, removendo pacotes não utilizados;**
