package models

var (
	AWS_ASSISTANT_RESOURCES = []string{"s3", "iam-user"} // 동작을 실행항 AWS Resource 입니다.
	AWS_ASSISTANT_COMMANDS  = []string{"list", "search"} // AWS Resource에 명령 입니다.
)

type Commands interface {
	List() ([]string, error)
	Search(prefix string) ([]string, error)
}

type MCPRequestParmas struct {
	Resource string `json:"resource"`
	Command  string `json:"command"`
	Item     string `json:"item"`
}
