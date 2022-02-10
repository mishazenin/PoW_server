# Proof-of-Work Client & Server

![](image/Frame%2023.png)


## Local launch

1. Set server addr for local launch
1. `make run-server`
1. `make run-client`

## Docker launch

1. `make init`
1. `make docker-build`
1. `make docker-run`
1. `make docker-stop`

## Protocol

In that version [Hashcash](https://en.wikipedia.org/wiki/Hashcash) is used.

If the header in incoming request doesn't contain `X-Hashcash`, the server generate a challenge with missing _nonce_ and
put in into `X-Hashcash` header e.g. `X-Hashcash: 1:12:211224001325:1.2.3.4::yuRNNkxUcaU=:`.

So from the client side it's nessesary to bruteforce the nonce to be in fit to the Hashcash target and then sent
solution to the same `X-Hashcash` header.

If the challenge is calculated correctly, the server replies random quote from book. 



