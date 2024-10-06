# go-fiber-docker-example

## 프로젝트 구성
- `cmd/main.go`: Entry Point
- `internal/db`: Prisma 스키마 및 Prisma Client 관련 코드
- `internal/repository`: DB 접근 관련 코드. 데이터 읽기&쓰기
- `internal/service`: 비즈니스 로직
- `internal/handler`: HTTP 요청 처리
- `internal/router`: API Endpoint
- `pkg/`: 내부 사용 라이브러리

## 사용 라이브러리
- [fiber](https://github.com/gofiber/fiber) 
- [validator](https://github.com/go-playground/validator) 
- [prisma-client-go](https://github.com/steebchen/prisma-client-go) 

## CI/CD
1. Docker로 Container 관리
2. Github Actions를 이용하여, `main` branch에 `push` 될 때 Docker Image 배포
3. SSH로 개인 Vultr 인스턴스에 접속 및 Image Pull
4. Deploy!

## TODO
- [uber-go/fx](https://github.com/uber-go/fx)를 통해 `cmd/main.go`에서의 의존성 주입
- Swagger 추가
- jwt를 이용한 유저 기능 추가
- GraphQL 적용

## API Endpoints
- **POST /auth/register**: 회원가입 기능
- **GET /auth/login**: 로그인 기능 (TODO: 리턴 값 jwt 토큰으로 수정)
- **PATCH /auth/reset**: 비밀번호 초기화 기능 (TODO: jwt 토큰 헤더 추가)
- **DELETE /auth/withdraw**: 탈퇴 기능 (TODO: jwt 토큰 헤더 추가)
- **POST /schedule**: 스케쥴 등록 기능 (TODO: jwt 토큰 헤더 추가)
- **GET /schedule**: 전체 스케쥴 조회 기능 (TODO: jwt 토큰 헤더 추가)
- **GET /schedule/:id**: 특정 ID 값의 스케쥴 조회 기능 (TODO: jwt 토큰 헤더 추가)
- **PATCH /schedule/:id**: 특정 ID 값의 스케쥴 수정 기능 (TODO: jwt 토큰 헤더 추가)
- **DELETE /schedule/:id**: 특정 ID 값의 스케쥴 삭제 기능 (TODO: jwt 토큰 헤더 추가)
- **GET /user/:userTag**: 유저 프로필 조회 기능
- **POST /user**: 유저 프로필 생성 기능 (TODO: jwt 토큰 헤더 추가)
- **PATCH /user/:userTag**: 유저 프로필 수정 기능 (TODO: jwt 토큰 헤더 추가)
- **DELETE /user/:userTag**: 유저 프로필 삭제 기능 (TODO: jwt 토큰 헤더 추가)