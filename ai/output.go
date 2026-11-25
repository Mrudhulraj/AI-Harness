package ai

type EngineExecOutput struct {
	Command     string `json:"cmd"`
	Explanation string `json:"exp"`
	Executable  bool   `json:"exec"`
}

type EngineChatStreamOutput struct {
	content    string
	last       bool
	interrupt  bool
	executable bool
}
