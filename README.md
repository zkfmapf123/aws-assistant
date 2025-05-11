# AWS Assistant (AWS 조수)

AWS 리소스를 쉽게 관리할 수 있는 MCP(Multi-Cloud Protocol) 기반의 AWS 관리 도구입니다.

## 주요 기능

- **S3 버킷 관리**
  - [x] 버킷 목록 조회
  - [x] 버킷 검색
  - [ ] 버킷 생성
  - [ ] 버킷 삭제

- **IAM 사용자 관리**
  - [x] 사용자 목록 조회
  - [x] 사용자 검색
  - [ ] 사용자 생성
  - [ ] 사용자 삭제

## 설치 방법

1. Go 1.22 이상 설치
2. 프로젝트 클론
```bash
git clone https://github.com/zkfmapf123/aws-assistant.git
```

3. 의존성 설치
```bash
go mod download
```

4. 빌드
```bash
make _build
```

## 설정 방법

- AWS 자격 증명 설정
   - AWS_ACCESS_KEY_ID: AWS 액세스 키
   - AWS_SECRET_ACCESS_KEY: AWS 시크릿 키
   - AWS_REGION: AWS 리전 (예: ap-northeast-2)

```json
// code ~/Library/Application\ Support/Claude/claude_desktop_config.json
{
  "mcpServers": {
    "aws-assistant": {
      "command": "/path/to/assistant",
      "env": {
        "AWS_ACCESS_KEY_ID": "YOUR_ACCESS_KEY_ID",
        "AWS_SECRET_ACCESS_KEY": "YOUR_SECRET_ACCESS_KEY",
        "AWS_REGION": "ap-northeast-2"
      }
    }
  }
}
```

## 사용 방법

### S3 버킷 관리

1. 버킷 목록 조회
```bash
aws-assistant s3 list
```

2. 버킷 검색
```bash
aws-assistant s3 search <검색어>
```

### IAM 사용자 관리

1. 사용자 목록 조회
```bash
aws-assistant iam-user list
```

2. 사용자 검색
```bash
aws-assistant iam-user search <검색어>
```

## 프로젝트 구조

```
.
├── cmd/
│   └── aws_assistant/    # 메인 애플리케이션
├── internal/
│   ├── handlers/        # 요청 핸들러
│   ├── helpers/         # 유틸리티 함수
│   └── vendor/          # AWS 서비스 클라이언트
├── models/              # 데이터 모델
├── config/             # 설정 파일
└── Makefile           # 빌드 스크립트
```

## 개발 환경 설정

1. VSCode 설정
   - Go 확장 설치
   - 설정 파일: `.vscode/settings.json`

2. 개발 도구
   - golangci-lint: 코드 품질 검사
   - goimports: 코드 포맷팅

## 테스트

```bash
make test
```
