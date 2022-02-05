# Proof-of-Work Client & Server

Pretty basic implementation.

## Local launch

1. `make init`
1. Set server addr for local launch
1. `make run-server`
1. `make run-client`

## Docker launch

1. `make init`
1. `make docker-build`
1. `make docker-run`
1. `make docker-stop` when tired of client constantly restarting.

## Protocol

Uses [Hashcash](https://en.wikipedia.org/wiki/Hashcash) because why not. Prove me wrong.

If the `X-Hashcash` header is not present in the incoming request, the server generates a challenge to solve with _nonce_ missing and sends it in `X-Hashcash` header, e.g. `X-Hashcash: 1:12:211224001325:1.2.3.4::yuRNNkxUcaU=:`.

The client then bruteforces the nonce so it satisfies the Hashcash target, concatenates it with the challenge sent by server and sends the solution using the same `X-Hashcash` header.

Should the challenge be solved correctly, the server replies with a quote from Fyodor Dostoevsky’s novels. _The Possessed_ had the greatest influence on me so far.

## Testing

`make test`

## TODO

1. Detailed statuses for failed challenges.
1. Storage to validate the hash.
1. IP forwarding via Nginx.
1. Follow the linters’ path.
