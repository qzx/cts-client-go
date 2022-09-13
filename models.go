package cts

type TaskList struct {
	RequestID string     `json:"request_id"`
	Tasks     []TaskItem `json:"tasks"`
}

// Task -
type TaskItem struct {
	Condition   Condition   `json:"condition"`
	Description string      `json:"description"`
	Enabled     bool        `json:"enabled"`
	Module      string      `json:"module"`
	ModuleInput ModuleInput `json:"module_input"`
	Name        string      `json:"name"`
	Providers   []string    `json:"providers"`
	Variables   Variables   `json:"variables"`
	Version     string      `json:"version"`
}

type Task struct {
	Task TaskItem `json:"task"`
}

type TaskResponse struct {
	RequestID string   `json:"request_id"`
	Task      TaskItem `json:"task"`
}

type Condition struct {
	ConsulKv ConsulKv `json:"consul_kv,omitempty"`
	Services Services `json:"services,omitempty"`
}

type ConsulKv struct {
	Datacenter       string `json:"datacenter"`
	Namespace        string `json:"namespace"`
	Path             string `json:"path"`
	Recurse          bool   `json:"recurse"`
	UseAsModuleInput bool   `json:"use_as_module_input"`
}

type Services struct {
	CtsUserDefinedMeta CtsUserDefinedMeta `json:"cts_user_defined_meta"`
	Datacenter         string             `json:"datacenter"`
	Filter             string             `json:"filter"`
	Names              []string           `json:"names"`
	Namespace          string             `json:"namespace"`
	UseAsModuleInput   bool               `json:"use_as_module_input"`
}

type CtsUserDefinedMeta struct{}

type ModuleInput struct{}

type Variables struct{}

type UpdateResponse struct {
	Inspect Inspect `json:"inspect"`
}

type Inspect struct {
	ChangesPresent bool   `json:"changes_present"`
	Plan           string `json:"plan"`
}

type DeleteResponse struct {
	RequestID string `json:"request_id"`
}
