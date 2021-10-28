# go proff

## Faz o profiling da sua aplicação, ou seja, analisa a utilização dos recursos computacionais, a fim de identificar gargalos da execução.

## Existem 3 formas:
1. Adicionar o pacote `"net/http/pprof"` em uma aplicação Web;
2. Inicializar a colete de informações manualmente em um determinado ponto do código, através das funções de profiling. Exemplo: `pprof.StartCPUProfile`
3. Analisar o export de Teste Unitário; Exemplo: `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`

Formas de profiling em Go: [link](https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz)


`brew install graphviz`


`go tool pprof -http=:8081 http://127.0.0.1:8080/debug/pprof/profile`
`go tool pprof -http=:8081 http://127.0.0.1:8080/debug/pprof/heap`
