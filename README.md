# go-fiber-docker-example

## 프로젝트 구성
- *cmd/main.go*: Entry Point
- *internal/db*: Prisma 스키마 및 Prisma Client 관련 코드
- *internal/repository*: DB 접근 관련 코드. 데이터 읽기&쓰기
- *internal/service*: 비즈니스 로직
- *internal/handler*: HTTP 요청 처리
- *pkg/*: 내부 사용 라이브러리

## 사용 라이브러리
- [fiber](https://github.com/gofiber/fiber) 
- [validator](https://github.com/go-playground/validator) 
- [prisma-client-go](https://github.com/steebchen/prisma-client-go) 

## CI/CD
1. Docker로 Container 관리
2. Github Actions를 이용하여, `main` branch에 `push` 될 때 Docker Image 배포
3. SSH로 개인 Vultr 인스턴스에 접속 및 Image Pull
4. Deploy!

