TripRoute
============

### Descrição
    TripRoute é um serviço para calcular a melhor rota entre um ponto de partida e um destino levando em consideração o custo das conexções.

### Linguagens
- golang: 1.14.3
  
### Frameworks
- fiber/fiber: 1.13.3


Instalation
-----------

### Building

- Server.

```sh
$ make build-server
```

- CLI
```sh
$ make build-cmd
```

- Build Docker.

```sh
$ make build-docker
```

### Testando

```sh
$ make test
$ make cover
```


### Executando

- Launch Server.
```sh
$ make run-server
```
- Executar CLI.
``` sh
$ make run-cmd
```
- Launch Server via Docker.
```sh
$ make docker-up
```

Exemplo
-------

Após iniciar o servidor é possível ter acesso ao serviço através da porta 3000

```sh
curl --location --request GET 'localhost:3000/findroute/?start=GRU&end=ORL'
```

API
---
### GET /findroute/
Localiza a rota com o menor custo independente da quantidade de conexções

query params:
  - ?start=string   // Ponto de origem da viagem
  - ?end=string     // Destino final

### POST /createroute/
Insere uma nova rota na base de dados

form-data:
   - start=string    // Ponto de origem da viagem
   - end=string      // Destino final
   - cost=int       // Custo desta conexão

Na pasta api/ esta o arquivo de colection do postman para execução de testes.

_________

Config
------
Os parâmetros de configuração do serviço estão no arquivo .env na raiz do projeto
### ./.env
    - port: default 3000
    - config-file: default input-file.txt


Arquitetura
-------------

Aplicação em 3 camadas:

  - Regra de negocio da empresa: responsável pelos calculos de rotas e validação de dados.
  - Regra de negocoi da aplicação: responsável por orquestrar os detalhes de implementação e as regras de negocio da empresa.
  - Camada de acesso a dados: Controla o acesso ao arquivo que possui os dados de rotas e custos.

### Estrutura de pastas do projeto:

- api: Arquivos para documentação de contratos de dados
- bin: Binários, scripts e outros executáveis
- cmd: Arquivos fontes de entry points do projeto
- pkg: Pacotes do projeto em golang
    - pkg/controller: Regra de negocio da aplicação
    - pkg/graph: Regra de negocio da empresa
    - pkg/repository: camada de acesso a dados
    - pkg/test: Utilitários para testes

A exposição da API REST é feita através da biblioteca Fiber, que utiliza a fasthttp que é a biblioteca http para golang com os melhores resultados em benchmarks.