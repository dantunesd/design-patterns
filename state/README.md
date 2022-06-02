# State

## Uma conta (despeza) comum, pode conter alguns estados, como Em aberto (Para ser paga), Paga (Quando de fato ja realizamos o pagamento) e Vencida (quando o tempo limite para ser pago foi atingido).

## Dependendo em qual estado ela estiver, algumas ações podem ser feitas, e outras não. Por Exemplo, após uma conta vencida, não podemos pagar ela (não sem fazer algo a mais), porém, numa conta aberta podemos.

## Aberto -> Pago ou Vencido
## Pago -> Nenhum
## Vencido -> Nenhum

## Implementando o State, conseguimos controlar para qual proximo estado algum objeto poderá ir a partir do seu estado atual. Basta criarmos objetos que representam cada estado, e delegar a eles os comportamentos e as validações para os proximos estados.