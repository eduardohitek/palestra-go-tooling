# gofmt

## Serve para formatar os arquivos seguindo as guidelines oficiais do Go.
## Além disso efetuar alguns refactoring nos arquivos

`gofmt -w *.go`: **Formata dos os arquivos .go do diretório;**

`gofmt -w -r 'bytes.Compare(bytea, byteb) > 0 -> bytes.Equal(bytea, byteb)' *.go`: **Substitui a chamada .Compare por .Equal;**

`gofmt -w -r 'bytea -> byteA' *.go`: **Substitui o nome da variável;**
