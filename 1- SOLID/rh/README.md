# RH
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/jonatasemidio/multilanguage-readme-pattern/blob/master/README.en.md)
[![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/jonatasemidio/multilanguage-readme-pattern/blob/master/README.md)

## Sobre o projeto
Esse projeto foi utilizado como exemplo para implementar os conceitos de SOLID.

- **S**ingle Responsability Principle: Toda classe deve ter apenas uma responsabilidade, como exemplo usamos a classe de Reajuste, onde todas as funções são relacionadas apenas a reajuste salarial.

- **O**pen Closed Principle: As classes devem ser abertas para extensão e fechadas para modificação, por exemplo a interface de Validação está aberta a extensão, porque podemos criar mais subclasses que implementam a função validar(), mas está fechada para modificação pois mesmo ao acrescentar novas validações as classes anteriores não serão afetadas.

- **L**iskov Substitution Principle: Uma classe 'pai' pode ser substituída por uma classe 'filha' sem introduzir efeitos colaterais e/ou quebrar a aplicação, como exemplo usamos a classe Terceirizado. Poderíamos forçar a classe Terceirizado a herdar de Funcionario mas a função de promoção só pode ser usada por Funcionario, caso seguíssemos com a herança nós criaríamos um efeito colateral indesejado (Terceirizado pode ser promovido). Também disponibilizaríamos atributos relacionados a promoção na classe Terceirizado, por isso criamos a classe Dados Pessoais para encapsular os atributos compartilhados e evitar duplicação de código.

- **I**terface Segregation Principle: Interfaces devem ter apenas funções que fazem sentido para o seu contexto. Esse projeto não implementa nenhum exemplo direto desse princípio, mas um exemplo seria se tivéssemos dois tipos de reajustes: Promoção (que é tributado) e Anuênio (que não é tributado). Somente Funcionarios podem ser promovidos, mas ambos Funcionarios e Terceirizados recebem Anuenio. Logo, não podemos usar a classe Reajuste para implementar os dois tipos de reajuste salarial, pois disponibilizaríamos a promoção para os Terceirizados. A solução seria criar uma classe chamada ReajusteTributavel, que tem a função de promoção, e herda de Reajuste, que tem a função de anuenio. Assim Funcionario utilizaria a classe ReajusteTributavel e pode usar as duas funções enquanto Terceirizado utiliza Reajuste e usa apenas a função de anuenio.

- **D**ependency Inversion Principle: Abstrações não dependem de implementações, implementações dependem de abstrações. Esse princípio está muito conectado com os princípios **S** e **O**, e podemos usar de exemplo como as classes de validação foram implementadas. Ao adicionar ou remover validações, não precisamos em nenhum momento alterar a interface Validacao, uma vez que ela é completamente abstrata.


### Run app
- Rode o comando `go run main.go` para subir o servidor. 

**Request**

`POST /reajuste`

```json
{
    "nome": "John Doe",
    "cpf": "12345678900",
    "cargo": "assistente",
    "salario": 2000,
    "data_ultimo_reajuste": "2024-01-10T10:00:00Z",
    "valor_reajuste": 100
}
```

**Response**
```
HTTP/1.1 400 Bad Request
Date: Mon, 29 Jan 2024 22:09:58 GMT
Status: 400 Bad Request
Connection: close
Content-Type: text/plain; charset=utf-8
Content-Length: 57

intervalo entre reajustes deve ser de no mínimo 6 meses
```


**Request**

`POST /reajuste`

```json
{
    "nome": "John Doe",
    "cpf": "12345678900",
    "cargo": "assistente",
    "salario": 2000,
    "data_ultimo_reajuste": "2023-01-10T10:00:00Z",
    "valor_reajuste": 100
}
```

**Response**
```
HTTP/1.1 200 Ok
Date: Mon, 29 Jan 2024 22:09:58 GMT
Status: 200 Ok
Connection: close
Content-Type: text/plain; charset=utf-8
Content-Length: 20

Salario Reajustado
```