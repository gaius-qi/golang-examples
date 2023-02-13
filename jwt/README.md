# JWT token

## Generate RSA key pair

```shell
$ openssl genrsa -out private.pem 4096
$ openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```
