package models

type ApprovalQueue struct {
	Id           int64           `json:"id"`
	QueueName    string          `json:"queue_name"`
	ResourceId   string          `json:"resource_id"`
	ResourceType string          `json:"resource_type"`
	ActionMeta   string          `json:"action_meta"`
	State        State           `json:"state"`
	Stages       []ApprovalStage `json:"stages"`
}

type ApprovalStage struct {
	Id        int64  `json:"id"`
	State     State  `json:"state"`
	Comment   string `json:"comment"`
	StageMeta string `json:"stage_meta"`
}

type State string

const (
	ASSIGNED State = "ASSIGNED"
	APPROVED State = "APPROVED"
	REJECT   State = "REJECT"
)
