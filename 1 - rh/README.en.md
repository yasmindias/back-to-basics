# RH
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/jonatasemidio/multilanguage-readme-pattern/blob/master/README.en.md)
[![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/jonatasemidio/multilanguage-readme-pattern/blob/master/README.md)

## About the project
This project was used as an examples to implement the SOLID concepts.

- **S**ingle Responsability Principle: Every class must have only one responsability, as exemplified by the Reajuste class, where the only functions are related to the responsability of reajuste.
- **O**pen Closed Principle: Classes must be open for extension and closed for modification, as exemplified by the Validation interface and its 'subclasses', we can add more validations without the need to modify any other file.
- **L**iskov Substitution Principle: A 'parent' class can be substituted by any 'child' class without introducing side effects and/or breaking the app, as exemplified by the Terceirizado class. We could force Terceirizado to extend the Funcionario class but only Funcionario should be able to be promoted, while Terceirizado can't. So it was created the Dados Pessoais class, so it can encapsulate the shared attributes and avoid code duplication but without implementing methods that shoulnd't be shared.
- **I**terface Segregation Principle: Interfaces should only have functions that make sense for their context. This project doesn't implement a direct example of this principle, but one example would be if there were two types of reajuste: Promoção (which is taxed) and Anuenio (which is not taxed). Only Funcionarios can be promoted but both Funcionarios and Terceirizados can receive Anuenio. So we can't use the same Reajuste class for both.  We would need to create a second classed called ReajusteTributavel that extends Reajuste, so Funcionario would be able to use ReajusteTributavel and be able to be promoted and receive anuenio while Terceirizado would use only Reajuste and receive only the anuenio.
- **D**ependency Inversion Principle: Abstractions must not depend on implementations, implementations should depend on abstractions. This principle is very connected with **S** and **O**, we can see it in how the validation classes were implemented.


### Run app
- Run `go run main.go` to run the server. 

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