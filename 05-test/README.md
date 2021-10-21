# go test

## Executa os arquivos de testes da sua aplicação, dependendo dos parametros informados, pode efetuar testes de benchmark e análise da cobertura de testes.

`go test`: **Executa todos os testes do diretório atual;**

`go test -v`: **Executa todos os testes do diretório atual, com informações dos testes individuais;**

`go test -v --cover`: **Adiciona informação de cobertura de Testes no relatório final;**

`go test -v -coverprofile .test_coverage.out`: **Exporta as informações da cobertura de testes;**

`go tool cover -html=.test_coverage.out -o coverage.html`: **Gera um arquivo html com as informações de cobertura de testes;**

`go test -run ^Test_notSo -count 10`: **Roda um determinado teste por x vezes;**

`go test -bench=. -benchtime=2s`: **Roda os testes em modo Benchmarking;**

