# Gatway de Pagamento Fincycle

![Go](https://img.shields.io/badge/go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-0db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-000.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Kafka](https://img.shields.io/badge/kafka-231F20.svg?style=for-the-badge&logo=apachekafka&logoColor=white)

_Este repositório faz parte do projeto [Sistema de pagamentos Fincycle](https://github.com/jvcalassio/fc-payment-system)_

Este serviço é responsável por receber todos os pagamentos da aplicação, e simular o funcionamento de uma operadora de cartão de crédito, aceitando ou rejeitando as transações baseado nas informações recebidas.

## Desenvolvimento

No desenvolvimento deste serviço, foram aplicadas práticas de Clean Architecture e Test Driven Development.
Cada aspecto do programa foi pensado para ser independente e testável, mantendo sempre as regras de negócio na camada mais baixa (nas entidades).

Assim, cada camada do sistema é altamente desacoplada e permite a mudança de componentes como bancos de dados, comunicação com o usuário (neste caso, o Apache Kafka), etc. bastando implementar as interfaces especificadas.

Toda entidade/adaptador/caso de uso foi feita seguindo os princípios do TDD.

## Regras de negócio

São consideradas as seguintes regras para validar uma transação:

- O valor mínimo da transação é 1
- O valor máximo da transação é 1000
- Toda transação terá apenas dois status possíveis: "approved" e "rejected"
- O cartão de crédito utilizado deve ser válido:
  - Não pode estar vencido (data e ano > data e ano atual)
  - O número deve ser válido
