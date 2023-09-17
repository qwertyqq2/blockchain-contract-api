# Contract API

## Запуск

### Docker

    docker-compose up - d


### Build

    go build main.go

Установить Ganache по ссылке https://trufflesuite.com/ganache/

    ganache-cli
    
    ./main -host=localhost

### Описание
Сервис позволяющий взаимодейстовать с смарт-контрактом. Сам контракт, разещеный в contract/Coin.sol, способен выполнять три фукнции:

1) mint(count) - выпустить токены в количестве count
2) send(receiver, count) - перевести токены в количестве amount на адрес receiver
3) getBalance(address) - вернуть баланс пользователя с адресом address

### Спека

1) POST /contract : деплоить новый контракт
2) POST /contract/mint {amount: string}:  выпустить токены
3) POST /contract/send {to: string, amount: string}: перевести токены
4) GET /contract/:addr : вернуть баланс пользователя по адресу