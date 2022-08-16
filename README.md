# Hello server


## Generate private/public key

reference: https://github.com/denji/golang-tls

### Generate private key (.key)

```bash
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

### Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

```bash
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## Run HTTP/HTTPS server 

reference: https://stackoverflow.com/questions/26090301/run-both-http-and-https-in-same-program

```go
//  Start HTTP
go func() {
    err_http := http.ListenAndServe(fmt.Sprintf(":%d", port), http_r)
    if err_http != nil {
        log.Fatal("Web server (HTTP): ", err_http)
    }
 }()

//  Start HTTPS
err_https := http.ListenAndServeTLS(fmt.Sprintf(":%d", ssl_port),     "D:/Go/src/www/ssl/public.crt", "D:/Go/src/www/ssl/private.key", https_r)
if err_https != nil {
    log.Fatal("Web server (HTTPS): ", err_https)
}
```

## insecure access on MacOS with Chrome

- link: https://www.boolsee.pe.kr/macos-chrome-err-cert-revoked/

```
1. NET::ERR_CERT_REVOKED 화면의 빈 공간의 아무 곳에서 마우스의 ‘선택’ 단추를 클릭.
2. 키보드로 ‘thisisunsafe’ 문자열 입력. (화면에 보이지 않으니 그냥 치세요.)
3. 접속하고자 하는 화면이 보이면 성공. 보이지 않으면 화면 Refresh 하시고 다시 시도해 보십시오.
```