# AnnoWiz File Server 

## 개요

REST API를 통해 데이터를 받아 저장하는 파일 서버

## 빌드 방법

1. Go 패키지를 설치한다. https://go.dev/dl/ 
2. `$ make build` 명령으로 빌드하면 `bin` 디렉토리에 윈도우, 리눅스, 맥용 바이너리가 생성된다. 


## 기능 

REST API를 통해 HTTP 또는 HTTPS 서버로 데이터가 전송되면 하드코딩된 `uploaded` 디렉토리 아래에 데이터를 압축해제하여 저장한다. 

## Makefile target

### gencert

- `$ make gencert` 명령으로 HTTPS server를 위한 `server.key`, `server.crt` 를 생성한다. 
- 인증서의 유효기간은 365일이다. 

### build

- `$ make build` 명령으로 윈도우, 리눅스, 맥용 바이너리를 생성한다. 

### upload

- IDC(Internet Data Center)의 테스트 서버에 리눅스용 서버 바이너리를 올리도록 하드코딩 되어있다. 
- 테스트 서버의 비밀번호가 필요하며 정현석, 설상훈에게 문의 바란다. 

### remote

- `$ make remote` 명령으로 IDC 테스트 서버에 원격 접속한다. 

### run 

- `$ make run` 명령으로 로컬에서 port 8888 및 8443(각각 HTTP, HTTPS)로 서버를 실행한다. 
- `nodemon` 설치가 필요하다. 참고 링크: https://github.com/Rezvitsky/nodemon-golang-example
- `nodemon` 으로 실행하여 소스코드가 변경되면 새로이 빌드하여 재실행한다. 

### delete

- `$ make delete` 명령으로 `uploaded` 디렉토리의 모든 파일을 지운다. 


## 참고: insecure access on MacOS with Chrome

MacOS에서 크롬 브라우저로 서버에 접근하려 하면 인증서 오류가 발생한다. 대처법을 아래에 공유한다. 

- 링크: https://www.boolsee.pe.kr/macos-chrome-err-cert-revoked/

```
1. NET::ERR_CERT_REVOKED 화면의 빈 공간의 아무 곳에서 마우스의 ‘선택’ 단추를 클릭.
2. 키보드로 ‘thisisunsafe’ 문자열 입력. (화면에 보이지 않으니 그냥 치세요.)
3. 접속하고자 하는 화면이 보이면 성공. 보이지 않으면 화면 Refresh 하시고 다시 시도해 보십시오.
```