package models

type ApprovalStage struct {
	Id        int64  `json:"id"`
	State     State  `json:"state"`
	Comment   string `json:"comment"`
	StageMeta string `json:"stage_meta"`
}

func (a *ApprovalStage) approve() {

}

func (a *ApprovalQueue) GetStageEvent() StageEvent {
	return StageEvent{
		State:           a.State,
		ApprovalQueueId: a.Id,
		ApprovalStageId: a.Id,
	}
}