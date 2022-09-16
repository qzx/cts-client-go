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
	Services        *Services        `json:"services,omitempty"`
	CatalogServices *CatalogServices `json:"catalog_services,omitempty"`
	ConsulKv        *ConsulKv        `json:"consul_kv,omitempty"`
	Schedule        *Schedule        `json:"schedule,omitempty"`
}

type Services struct {
	CtsUserDefinedMeta *CtsUserDefinedMeta `json:"cts_user_defined_meta,omitempty"`
	Datacenter         string              `json:"datacenter,omitempty"`
	Namespace          string              `json:"namespace,omitempty"`
	Filter             string              `json:"filter,omitempty"`
	Regexp             string              `json:"regexp,omitempty"`
	Names              []string            `json:"names,omitempty"`
	UseAsModuleInput   bool                `json:"use_as_module_input,omitempty"`
}

type CatalogServices struct {
	CtsUserDefinedMeta *CtsUserDefinedMeta `json:"cts_user_defined_meta,omitempty"`
	Datacenter         string              `json:"datacenter,omitempty"`
	Namespace          string              `json:"namespace,omitempty"`
	Regexp             string              `json:"regexp,omitempty"`
	NodeMeta           *NodeMeta           `json:"node_meta,omitempty"`
	UseAsModuleInput   *bool               `json:"use_as_module_input,omitempty"`
}

type ConsulKv struct {
	Datacenter       string `json:"datacenter,omitemptyr"`
	Namespace        string `json:"namespace,omitempty"`
	Path             string `json:"path,omitempty"`
	Recurse          bool   `json:"recurse,omitempty"`
	UseAsModuleInput *bool  `json:"use_as_module_input,omitempty"`
}

type Schedule struct {
	Cron string `json:"cron,omitempty"`
}

type CtsUserDefinedMeta map[string]string

type NodeMeta map[string]string

type ModuleInput struct {
	Services *Services `json:"services,omitempty"`
	ConsulKv *ConsulKV `json:"consul_kv,omitempty"`
}

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
