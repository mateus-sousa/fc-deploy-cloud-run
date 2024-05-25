O endereço para validação da aplicação no Google Cloud Run é:
https://fc-deploy-cloud-run-goexpert-hzdsmkxvlq-uc.a.run.app/getweather/<cep>

Passo a passo para execução da aplicação e testes localmente:

* Para subir as dependencias do projeto execute o comando:
```
    docker-compose up -d --build
```
A aplicação estará de pé respondendo na seguinte porta: 8080.

### 1. Testes:
* Há 2 testes em nossa pilha
  * O primeiro efetua uma requisição a um CEP valido.
  * O Segundo efetua uma  requisição a um CEP inválido.
  * O Terceiro efetua uma requisição a um CEP não encontrado.    

### 2. Rodando de testes:
    * Com o ambiente e as dependencias em pé, execute o comando para entrar no container da aplicação:
```
   docker-compose exec goapp bash
```
* Dentro do container execute o comando:
```
   go test ./...
```
