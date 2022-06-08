package entities

type AuditLog struct {
	Origin    string      `json:"origin"`
	Level     string      `json:"level"`
	Event     string      `json:"event"`
	Type      string      `json:"type"`
	User      string      `json:"user"`
	Timestamp int64       `json:"timestamp"`
	Message   interface{} `json:"message"`
}
