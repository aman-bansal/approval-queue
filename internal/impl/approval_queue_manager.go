package impl

import (
	"context"
	"github.com/aman-bansal/approval-queue/config"
	"github.com/aman-bansal/approval-queue/pkg/grpc"
)

func NewApprovalQueueManager(conf *config.ApprovalQueueConfig) grpc.ApprovalQueueManagerServer {
	return &ApprovalQueueManager{}
}

type ApprovalQueueManager struct {
	
}

func (a ApprovalQueueManager) Create(context.Context, *grpc.CreateRequest) (*grpc.CreateResponse, error) {
	panic("implement me")
}

func (a ApprovalQueueManager) Get(context.Context, *grpc.GetRequest) (*grpc.GetResponse, error) {
	panic("implement me")
}

func (a ApprovalQueueManager) Search(context.Context, *grpc.SearchRequest) (*grpc.SearchResponse, error) {
	panic("implement me")
}

func (a ApprovalQueueManager) Action(context.Context, *grpc.ActionRequest) (*grpc.ActionResponse, error) {
	panic("implement me")
}

