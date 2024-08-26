## Go-crawl :brazil:

Apenas um crawler que acessa um determinado site, pega todos os links da página, visita cada um deles - inclusive os links que existem dentro da página visitada no momento - e armazena no banco de dados (MongoDB)

#### Como executar?

Obs: certifique-se de ter o Go instalado na sua máquina.

Antes de executar, suba uma instância do MongoDB em um container Docker:

`docker run -d --name mongodb -p 27017:27017 mongo`

Depois, instale as dependências do projeto:

`go mod tidy`

E, em seguida:

`go run main.go`

Para verificar as inserções no banco, você pode acessar o container:

`docker exec -it mongodb /bin/bash`

Depois, vá até a collection e verifique a quantidade de inserções:

- `mongosh`
- `show dbs;`
- `show collections;`
- `use crawler` (nome da collection que está sendo utilizada)
- `db.links.countDocuments({})`

---------------------------

## Go-crawl :uk:

Just a crawler that accesses a given url site, takes all links on the page, visits each one of them - include links that exist within the page currently visited - and stores them in the database (MongoDB)

#### How to run?

Note: verify if Go was installed on your machine.

Before run, run a MongoDB instance in a docker container:

`docker run -d --name mongodb -p 27017:27017 mongo`

Install dependencies;

`go mod tidy`

And then:

`go run main.go`

To verify insertions on database, you can access MongoDB container:

`docker exec -it mongodb /bin/bash`

After go to the collection and count insertions:

- `mongosh`
- `show dbs;`
- `show collections;`
- `use crawler` (name of collection)
- `db.links.countDocuments({})`