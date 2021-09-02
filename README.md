# jwt
jwt cli to decode and encode jwt tokens

## usage

### decode
```
Usage:
  jwt decode [flags] <token>

Flags:
      --alg string      algorithm for signature validation, if it is empty, no validation is done
  -h, --help            help for decode
      --indent          indent output
      --secret string   secret to validate token

Global Flags:
      --silent   suppress logs (warn, debug, ...)
```

Example
```shell
jwt decode --secret hello --alg HS256 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.ElsKKULlzGtesThefMuj2_a6KIY9L5i2zDrBLHV-e0M
valid: true, header: {"alg":"HS256","typ":"JWT"}, claims: {"iat":1516239022,"name":"John Doe","sub":"1234567890"}
```
