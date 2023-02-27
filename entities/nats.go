package entities

// NatsPermissions represents a list of allowed subjects the active user token may subscribe / publish to after connection to nats
type NatsPermissions struct {
	Email       string `json:"email"`
	Permissions struct {
		Publish struct {
			Deny  []string `json:"deny"`
			Allow []string `json:"allow"`
		} `json:"publish"`
		Subscribe struct {
			Deny  []string `json:"deny"`
			Allow []string `json:"allow"`
		} `json:"subscribe"`
		Responses struct {
			Max int `json:"max"`
			TTL int `json:"ttl"`
		} `json:"responses"`
	} `json:"permissions"`
}
