package models

type QueueEvent struct {
	State           State `json:"state"`
	ApprovalQueueId int64 `json:"approval_queue_id"`
}

type StageEvent struct {
	State           State `json:"state"`
	ApprovalQueueId int64 `json:"approval_queue_id"`
	ApprovalStageId int64 `json:"approval_stage_id"`
}
