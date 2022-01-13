# Generate a new localhost certificate with Go

Select appropriate command depending on install location:

```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

# or

go run /opt/homebrew/opt/go/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```


