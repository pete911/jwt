# jwt
jwt cli to decode and encode jwt tokens

## download

- [binary](https://github.com/pete911/jwt/releases)

## build/install

### brew

- add tap `brew tap pete911/tap`
- install `brew install jwt`

### go

- install [go](https://golang.org/doc/install)
- build `make build`
- install `make install`

## release

Releases are published when the new tag is created e.g.
`git tag -m "add super cool feature" v1.0.0 && git push --follow-tags`

## usage

### decode
```
Usage:
  jwt decode [flags] [token]

Flags:
      --alg string   algorithm for signature validation, if it is empty, no validation is done
  -h, --help         help for decode
      --indent       indent output
      --key string   key to validate token

Global Flags:
      --silent   suppress logs (warn, debug, ...)
```

Example
```shell
jwt decode --key hello --alg HS256 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.ElsKKULlzGtesThefMuj2_a6KIY9L5i2zDrBLHV-e0M
valid: true, header: {"alg":"HS256","typ":"JWT"}, claims: {"iat":1516239022,"name":"John Doe","sub":"1234567890"}
```

### encode
```
Usage:
  jwt encode [flags] [claims]

Flags:
      --alg string   algorithm for signature validation, if it is empty, no validation is done
  -h, --help         help for encode
      --key string   key to validate token

Global Flags:
      --silent   suppress logs (warn, debug, ...)
```

Example
```shell
jwt encode --key hello --alg HS256 '{"hello":"ok"}'
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoZWxsbyI6Im9rIn0.ycJ6FlOpr9qbWVQsVqQr9Pls1F_QfoOE7fk9pUvYNm8
```

### encode/decode stdin
Input can be read from stdin instead of passed as an argument as well.

```shell
jwt encode --key hello --alg HS256 '{"hello":"ok"}' | jwt decode --key hello --alg HS256
valid: true, header: {"alg":"HS256","typ":"JWT"}, claims: {"hello":"ok"}
```
