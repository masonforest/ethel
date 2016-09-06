# Ethel
A golang framework for building and deploying ethereum contracts written in solidity.

_Ethel is still in the early prototype phase. If you're looking for something more stable check out one of these similar projects:_

* [Populus](https://github.com/pipermerriam/populus)
* [Truffle](https://github.com/ConsenSys/truffle)
* [Embark](https://github.com/iurimatias/embark-framework)

# Getting started

First install ethel.

    go get github.com/masonforest/ethel/...

Create a new project.

    ethel new my_token
    cd my_token

Run your new contract's tests.

    ethel test

# TODO

* Add a server that can connect to a contract and act as an oracle
* Add `ethel console` (will connect to a contract and boot [gore](https://github.com/motemen/gore) with a initialized session)
* Lots of other things
