package models

import "context"

type ApprovalQueue struct {
	Id           int64           `json:"id"`
	QueueName    string          `json:"queue_name"`
	ResourceId   string          `json:"resource_id"`
	ResourceType string          `json:"resource_type"`
	ActionMeta   string          `json:"action_meta"`
	State        State           `json:"state"`
	Stages       []ApprovalStage `json:"stages"`
}

type State string

const (
	ASSIGNED State = "ASSIGNED"
	APPROVED State = "APPROVED"
	REJECT   State = "REJECT"
)

func (a *ApprovalQueue) Action(ctx context.Context, approvalStageId int64, action string) {
	for _, stage := range a.Stages {
		if stage.Id == approvalStageId {
			stage.approve()
		}
	}

}

func (a *ApprovalQueue) GetQueueEvent() QueueEvent {
	return QueueEvent{
		State:           a.State,
		ApprovalQueueId: a.Id,
	}
}

func (a *ApprovalQueue) Evaluate(ctx context.Context) {

}