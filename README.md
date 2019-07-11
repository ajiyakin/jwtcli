# jwtcli
CLI app to help you extract JWT token

## Install

Before you execute the installation command, make sure that you have Golang installed in your machine with
minimum version is `v1.12.4`.

`go get -u -v github.com/ajiyakin/jwtcli`

After execute the command above, the binary file will be placed under `$GOPATH/bin/` folder called `jwtcli`.
Make sure that folder is included within your `$PATH` environment variable in order to be able to execute it
from anywhere.

## Usage

```
App to decode information within a JWT token
Usage:
  jwtcli JWT_TOKEN
example:
  jwtcli eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU
```