// Code generated by protoc-gen-go. DO NOT EDIT.
// source: approval_queue_manager.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	approval_queue_manager.proto

It has these top-level messages:
	CreateRequest
	CreateApprovalStageRequest
	CreateResponse
	GetRequest
	GetResponse
	SearchRequest
	SearchResponse
	ActionRequest
	ActionResponse
	ApprovalQueue
	ApprovalStage
	Page
	PaginationResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionType int32

const (
	ActionType_ACTION_TYPE_UNDEFINED ActionType = 0
	ActionType_APPROVE               ActionType = 1
	ActionType_REJECT                ActionType = 2
	ActionType_SKIP                  ActionType = 3
)

var ActionType_name = map[int32]string{
	0: "ACTION_TYPE_UNDEFINED",
	1: "APPROVE",
	2: "REJECT",
	3: "SKIP",
}
var ActionType_value = map[string]int32{
	"ACTION_TYPE_UNDEFINED": 0,
	"APPROVE":               1,
	"REJECT":                2,
	"SKIP":                  3,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}
func (ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type State int32

const (
	State_STATE_UNDEFINED State = 0
	State_ASSIGNED        State = 1
	State_APPROVED        State = 2
	State_REJECTED        State = 3
	State_QUEUED          State = 4
)

var State_name = map[int32]string{
	0: "STATE_UNDEFINED",
	1: "ASSIGNED",
	2: "APPROVED",
	3: "REJECTED",
	4: "QUEUED",
}
var State_value = map[string]int32{
	"STATE_UNDEFINED": 0,
	"ASSIGNED":        1,
	"APPROVED":        2,
	"REJECTED":        3,
	"QUEUED":          4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	QueueName      string                      `protobuf:"bytes,1,opt,name=queue_name,json=queueName" json:"queue_name,omitempty"`
	ResourceId     string                      `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	ResourceType   string                      `protobuf:"bytes,3,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	ActionMeta     string                      `protobuf:"bytes,4,opt,name=action_meta,json=actionMeta" json:"action_meta,omitempty"`
	ApprovalStages *CreateApprovalStageRequest `protobuf:"bytes,5,opt,name=approval_stages,json=approvalStages" json:"approval_stages,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *CreateRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *CreateRequest) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *CreateRequest) GetActionMeta() string {
	if m != nil {
		return m.ActionMeta
	}
	return ""
}

func (m *CreateRequest) GetApprovalStages() *CreateApprovalStageRequest {
	if m != nil {
		return m.ApprovalStages
	}
	return nil
}

type CreateApprovalStageRequest struct {
	StageMeta string `protobuf:"bytes,1,opt,name=stage_meta,json=stageMeta" json:"stage_meta,omitempty"`
}

func (m *CreateApprovalStageRequest) Reset()                    { *m = CreateApprovalStageRequest{} }
func (m *CreateApprovalStageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateApprovalStageRequest) ProtoMessage()               {}
func (*CreateApprovalStageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateApprovalStageRequest) GetStageMeta() string {
	if m != nil {
		return m.StageMeta
	}
	return ""
}

type CreateResponse struct {
	ApprovalQueueId int64 `protobuf:"varint,1,opt,name=approval_queue_id,json=approvalQueueId" json:"approval_queue_id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetApprovalQueueId() int64 {
	if m != nil {
		return m.ApprovalQueueId
	}
	return 0
}

type GetRequest struct {
	ApprovalQueueId int64 `protobuf:"varint,1,opt,name=approval_queue_id,json=approvalQueueId" json:"approval_queue_id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetApprovalQueueId() int64 {
	if m != nil {
		return m.ApprovalQueueId
	}
	return 0
}

type GetResponse struct {
	Queue *ApprovalQueue `protobuf:"bytes,1,opt,name=queue" json:"queue,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetQueue() *ApprovalQueue {
	if m != nil {
		return m.Queue
	}
	return nil
}

type SearchRequest struct {
	QueryString  string `protobuf:"bytes,1,opt,name=query_string,json=queryString" json:"query_string,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	ResourceType string `protobuf:"bytes,3,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	Page         *Page  `protobuf:"bytes,4,opt,name=page" json:"page,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchRequest) GetQueryString() string {
	if m != nil {
		return m.QueryString
	}
	return ""
}

func (m *SearchRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *SearchRequest) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *SearchRequest) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

type SearchResponse struct {
	Queue []*ApprovalQueue `protobuf:"bytes,1,rep,name=queue" json:"queue,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SearchResponse) GetQueue() []*ApprovalQueue {
	if m != nil {
		return m.Queue
	}
	return nil
}

type ActionRequest struct {
	ApprovalQueueId int64      `protobuf:"varint,1,opt,name=approval_queue_id,json=approvalQueueId" json:"approval_queue_id,omitempty"`
	ApprovalStageId int64      `protobuf:"varint,2,opt,name=approval_stage_id,json=approvalStageId" json:"approval_stage_id,omitempty"`
	Comment         string     `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	ActionType      ActionType `protobuf:"varint,4,opt,name=action_type,json=actionType,enum=ActionType" json:"action_type,omitempty"`
}

func (m *ActionRequest) Reset()                    { *m = ActionRequest{} }
func (m *ActionRequest) String() string            { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()               {}
func (*ActionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ActionRequest) GetApprovalQueueId() int64 {
	if m != nil {
		return m.ApprovalQueueId
	}
	return 0
}

func (m *ActionRequest) GetApprovalStageId() int64 {
	if m != nil {
		return m.ApprovalStageId
	}
	return 0
}

func (m *ActionRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ActionRequest) GetActionType() ActionType {
	if m != nil {
		return m.ActionType
	}
	return ActionType_ACTION_TYPE_UNDEFINED
}

type ActionResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *ActionResponse) Reset()                    { *m = ActionResponse{} }
func (m *ActionResponse) String() string            { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()               {}
func (*ActionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ActionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ApprovalQueue struct {
	Id             int64          `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	QueueName      string         `protobuf:"bytes,2,opt,name=queue_name,json=queueName" json:"queue_name,omitempty"`
	ResourceId     string         `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	ResourceType   string         `protobuf:"bytes,4,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	ActionMeta     string         `protobuf:"bytes,5,opt,name=action_meta,json=actionMeta" json:"action_meta,omitempty"`
	State          State          `protobuf:"varint,6,opt,name=state,enum=State" json:"state,omitempty"`
	ApprovalStages *ApprovalStage `protobuf:"bytes,7,opt,name=approval_stages,json=approvalStages" json:"approval_stages,omitempty"`
}

func (m *ApprovalQueue) Reset()                    { *m = ApprovalQueue{} }
func (m *ApprovalQueue) String() string            { return proto.CompactTextString(m) }
func (*ApprovalQueue) ProtoMessage()               {}
func (*ApprovalQueue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ApprovalQueue) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApprovalQueue) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *ApprovalQueue) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ApprovalQueue) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ApprovalQueue) GetActionMeta() string {
	if m != nil {
		return m.ActionMeta
	}
	return ""
}

func (m *ApprovalQueue) GetState() State {
	if m != nil {
		return m.State
	}
	return State_STATE_UNDEFINED
}

func (m *ApprovalQueue) GetApprovalStages() *ApprovalStage {
	if m != nil {
		return m.ApprovalStages
	}
	return nil
}

type ApprovalStage struct {
	Id        int64  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	State     State  `protobuf:"varint,3,opt,name=state,enum=State" json:"state,omitempty"`
	Comment   string `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	StageMeta string `protobuf:"bytes,5,opt,name=stageMeta" json:"stageMeta,omitempty"`
}

func (m *ApprovalStage) Reset()                    { *m = ApprovalStage{} }
func (m *ApprovalStage) String() string            { return proto.CompactTextString(m) }
func (*ApprovalStage) ProtoMessage()               {}
func (*ApprovalStage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ApprovalStage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApprovalStage) GetState() State {
	if m != nil {
		return m.State
	}
	return State_STATE_UNDEFINED
}

func (m *ApprovalStage) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ApprovalStage) GetStageMeta() string {
	if m != nil {
		return m.StageMeta
	}
	return ""
}

type Page struct {
	From int64 `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Page) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Page) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type PaginationResponse struct {
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	From       int64 `protobuf:"varint,2,opt,name=from" json:"from,omitempty"`
	Size       int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *PaginationResponse) Reset()                    { *m = PaginationResponse{} }
func (m *PaginationResponse) String() string            { return proto.CompactTextString(m) }
func (*PaginationResponse) ProtoMessage()               {}
func (*PaginationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PaginationResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *PaginationResponse) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *PaginationResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateApprovalStageRequest)(nil), "CreateApprovalStageRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*ActionRequest)(nil), "ActionRequest")
	proto.RegisterType((*ActionResponse)(nil), "ActionResponse")
	proto.RegisterType((*ApprovalQueue)(nil), "ApprovalQueue")
	proto.RegisterType((*ApprovalStage)(nil), "ApprovalStage")
	proto.RegisterType((*Page)(nil), "Page")
	proto.RegisterType((*PaginationResponse)(nil), "PaginationResponse")
	proto.RegisterEnum("ActionType", ActionType_name, ActionType_value)
	proto.RegisterEnum("State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for ApprovalQueueManager service

type ApprovalQueueManagerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc1.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc1.CallOption) (*GetResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc1.CallOption) (*SearchResponse, error)
	Action(ctx context.Context, in *ActionRequest, opts ...grpc1.CallOption) (*ActionResponse, error)
}

type approvalQueueManagerClient struct {
	cc *grpc1.ClientConn
}

func NewApprovalQueueManagerClient(cc *grpc1.ClientConn) ApprovalQueueManagerClient {
	return &approvalQueueManagerClient{cc}
}

func (c *approvalQueueManagerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc1.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc1.Invoke(ctx, "/ApprovalQueueManager/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalQueueManagerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc1.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc1.Invoke(ctx, "/ApprovalQueueManager/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalQueueManagerClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc1.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc1.Invoke(ctx, "/ApprovalQueueManager/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalQueueManagerClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc1.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := grpc1.Invoke(ctx, "/ApprovalQueueManager/Action", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApprovalQueueManager service

type ApprovalQueueManagerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Action(context.Context, *ActionRequest) (*ActionResponse, error)
}

func RegisterApprovalQueueManagerServer(s *grpc1.Server, srv ApprovalQueueManagerServer) {
	s.RegisterService(&_ApprovalQueueManager_serviceDesc, srv)
}

func _ApprovalQueueManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalQueueManagerServer).Create(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalQueueManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalQueueManagerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalQueueManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalQueueManagerServer).Get(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalQueueManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalQueueManagerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalQueueManager_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalQueueManagerServer).Search(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalQueueManager/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalQueueManagerServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalQueueManager_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalQueueManagerServer).Action(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalQueueManager/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalQueueManagerServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApprovalQueueManager_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "ApprovalQueueManager",
	HandlerType: (*ApprovalQueueManagerServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ApprovalQueueManager_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApprovalQueueManager_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ApprovalQueueManager_Search_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _ApprovalQueueManager_Action_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "approval_queue_manager.proto",
}

func init() { proto.RegisterFile("approval_queue_manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x63, 0x27, 0xe9, 0x1e, 0x27, 0x4e, 0x18, 0x40, 0xf2, 0x86, 0xa2, 0x5d, 0x0c, 0x17,
	0xab, 0xc0, 0x4e, 0x50, 0x56, 0x02, 0x24, 0xb8, 0x09, 0x89, 0xa9, 0xbc, 0x68, 0xd3, 0xd4, 0x4e,
	0x91, 0x40, 0x42, 0xd6, 0xd4, 0x19, 0xdc, 0x88, 0xfa, 0xa7, 0xf6, 0x18, 0xa9, 0xbc, 0x06, 0x0f,
	0xc2, 0x33, 0xf0, 0x1c, 0xbc, 0x0a, 0x17, 0x68, 0x66, 0xec, 0x24, 0x4e, 0x5a, 0x55, 0xa0, 0xbd,
	0xf3, 0x7c, 0x73, 0xe6, 0x3b, 0xe7, 0x3b, 0x7f, 0x86, 0x53, 0x92, 0xa6, 0x59, 0xf2, 0x1b, 0xb9,
	0xf1, 0x6f, 0x0b, 0x5a, 0x50, 0x3f, 0x22, 0x31, 0x09, 0x69, 0x86, 0xd3, 0x2c, 0x61, 0x89, 0xf5,
	0xb7, 0x02, 0xbd, 0x59, 0x46, 0x09, 0xa3, 0x2e, 0xbd, 0x2d, 0x68, 0xce, 0xd0, 0x87, 0x00, 0xd2,
	0x30, 0x26, 0x11, 0x35, 0x95, 0xe7, 0xca, 0x8b, 0x27, 0xee, 0x13, 0x81, 0x2c, 0x48, 0x44, 0xd1,
	0x33, 0xd0, 0x33, 0x9a, 0x27, 0x45, 0x16, 0x50, 0x7f, 0xb3, 0x36, 0x9b, 0xe2, 0x1e, 0x2a, 0xc8,
	0x59, 0xa3, 0x8f, 0xa1, 0xb7, 0x35, 0x60, 0x77, 0x29, 0x35, 0x55, 0x61, 0xd2, 0xad, 0xc0, 0xd5,
	0x5d, 0x2a, 0x58, 0x48, 0xc0, 0x36, 0x49, 0xec, 0x47, 0x94, 0x11, 0x53, 0x93, 0x2c, 0x12, 0x7a,
	0x43, 0x19, 0x41, 0x73, 0xe8, 0x6f, 0xe3, 0xce, 0x19, 0x09, 0x69, 0x6e, 0xb6, 0x9e, 0x2b, 0x2f,
	0xf4, 0xc9, 0x07, 0x58, 0x86, 0x3b, 0x2d, 0x6f, 0x3d, 0x7e, 0x59, 0xc6, 0xee, 0x1a, 0x64, 0x1f,
	0xcd, 0xad, 0xaf, 0x61, 0xf8, 0xb0, 0x35, 0x57, 0x2a, 0xa8, 0x65, 0x0c, 0xa5, 0x52, 0x81, 0xf0,
	0x10, 0xac, 0x6f, 0xc0, 0xa8, 0x32, 0x93, 0xa7, 0x49, 0x9c, 0x53, 0x34, 0x82, 0x77, 0x0e, 0x92,
	0xb9, 0x59, 0x8b, 0x77, 0xaa, 0xbb, 0x8d, 0xf6, 0x82, 0xe3, 0xce, 0xda, 0xfa, 0x0a, 0xe0, 0x8c,
	0xb2, 0xca, 0xd5, 0x7f, 0x79, 0xf9, 0x0a, 0x74, 0xf1, 0xb2, 0x74, 0xfa, 0x09, 0xb4, 0xc4, 0x0b,
	0x61, 0xae, 0x4f, 0x0c, 0x3c, 0xdd, 0xb7, 0x77, 0xe5, 0xa5, 0xf5, 0x87, 0x02, 0x3d, 0x8f, 0x92,
	0x2c, 0xb8, 0xae, 0x5c, 0x7e, 0x04, 0xdd, 0xdb, 0x82, 0x66, 0x77, 0x7e, 0xce, 0xb2, 0x4d, 0x1c,
	0x96, 0xfa, 0x74, 0x81, 0x79, 0x02, 0x7a, 0x4b, 0xb5, 0x7c, 0x0a, 0x5a, 0x4a, 0x42, 0x2a, 0x8a,
	0xa8, 0x4f, 0x5a, 0x78, 0xc9, 0x53, 0x2c, 0x20, 0xeb, 0x0b, 0x30, 0xaa, 0xa0, 0x8e, 0xd5, 0xa8,
	0x0f, 0xab, 0xf9, 0x53, 0x81, 0xde, 0x54, 0x34, 0xc3, 0xff, 0x48, 0x60, 0xcd, 0x56, 0x16, 0xb8,
	0x14, 0xb7, 0x67, 0x2b, 0x1a, 0xc1, 0x59, 0x23, 0x13, 0x3a, 0x41, 0x12, 0x45, 0x34, 0x66, 0xa5,
	0xb6, 0xea, 0x88, 0x3e, 0xdb, 0xb6, 0xa8, 0x50, 0xce, 0xd5, 0x19, 0x13, 0x1d, 0xcb, 0xb0, 0xb8,
	0xf0, 0xaa, 0x5f, 0xf9, 0xb7, 0x35, 0x02, 0xa3, 0x0a, 0xb8, 0x54, 0x6a, 0x42, 0x27, 0x2f, 0x82,
	0x80, 0xe6, 0xb9, 0x88, 0xf3, 0xc4, 0xad, 0x8e, 0xd6, 0x3f, 0x5c, 0xdd, 0x7e, 0xcc, 0xc8, 0x80,
	0xa6, 0x53, 0xc9, 0x69, 0x3a, 0xeb, 0x83, 0x19, 0x6c, 0x3e, 0x32, 0x83, 0xea, 0xe3, 0x75, 0xd3,
	0x1e, 0x9f, 0xc1, 0xd6, 0xd1, 0x0c, 0x9e, 0x42, 0x2b, 0x67, 0x84, 0x51, 0xb3, 0x2d, 0xb4, 0xb7,
	0xb1, 0xc7, 0x4f, 0xae, 0x04, 0xd1, 0x97, 0xc7, 0x13, 0xda, 0x39, 0xe8, 0x50, 0x39, 0x6d, 0x87,
	0x43, 0x59, 0xec, 0xd4, 0x0b, 0xe4, 0x48, 0xfd, 0xd6, 0xaf, 0x7a, 0x9f, 0xdf, 0xbd, 0x8a, 0x69,
	0xf5, 0x8a, 0x9d, 0xc2, 0x6e, 0x7a, 0x4b, 0x39, 0x7b, 0xe3, 0x8c, 0x41, 0xe3, 0x9d, 0x89, 0x10,
	0x68, 0xbf, 0x64, 0x49, 0x54, 0xfa, 0x13, 0xdf, 0x1c, 0xcb, 0x37, 0xbf, 0xd3, 0xb2, 0x49, 0xc4,
	0xb7, 0xf5, 0x33, 0xa0, 0x25, 0x09, 0x37, 0x31, 0xa9, 0x55, 0xf5, 0x19, 0xe8, 0x2c, 0x61, 0xe4,
	0xc6, 0x0f, 0x92, 0x22, 0x66, 0x25, 0x09, 0x08, 0x68, 0xc6, 0x91, 0x2d, 0x7d, 0xf3, 0x1e, 0x7a,
	0x75, 0x47, 0x3f, 0x7a, 0x0d, 0xb0, 0x6b, 0x25, 0xf4, 0x14, 0xde, 0x9f, 0xce, 0x56, 0xce, 0xf9,
	0xc2, 0x5f, 0xfd, 0xb8, 0xb4, 0xfd, 0xcb, 0xc5, 0xdc, 0xfe, 0xce, 0x59, 0xd8, 0xf3, 0x41, 0x03,
	0xe9, 0xd0, 0x99, 0x2e, 0x97, 0xee, 0xf9, 0x0f, 0xf6, 0x40, 0x41, 0x00, 0x6d, 0xd7, 0x7e, 0x6d,
	0xcf, 0x56, 0x83, 0x26, 0x3a, 0x01, 0xcd, 0xfb, 0xde, 0x59, 0x0e, 0xd4, 0x91, 0x0b, 0x2d, 0x91,
	0x22, 0xf4, 0x2e, 0xf4, 0xbd, 0xd5, 0x74, 0x55, 0x27, 0xe8, 0xc2, 0xc9, 0xd4, 0xf3, 0x9c, 0x33,
	0x7e, 0x52, 0xc4, 0x49, 0xd2, 0xcd, 0x07, 0x4d, 0x7e, 0x92, 0x7c, 0xf6, 0x7c, 0xa0, 0x72, 0xf6,
	0x8b, 0x4b, 0xfb, 0xd2, 0x9e, 0x0f, 0xb4, 0xc9, 0x5f, 0x0a, 0xbc, 0x57, 0x6b, 0xd2, 0x37, 0xf2,
	0xbf, 0x81, 0x3e, 0x85, 0xb6, 0x5c, 0x8b, 0xc8, 0xc0, 0xb5, 0x3f, 0xc7, 0xb0, 0x8f, 0xeb, 0xfb,
	0xd2, 0x6a, 0x20, 0x0b, 0xd4, 0x33, 0xca, 0x90, 0x8e, 0x77, 0xbb, 0x70, 0xd8, 0xc5, 0x7b, 0xeb,
	0xcd, 0x6a, 0x70, 0x42, 0xb9, 0x24, 0x90, 0x81, 0x6b, 0x2b, 0x6c, 0xd8, 0xc7, 0xf5, 0xed, 0x21,
	0x8d, 0x65, 0xda, 0x90, 0x81, 0x6b, 0x1b, 0x62, 0xd8, 0xc7, 0xf5, 0x01, 0xb4, 0x1a, 0xdf, 0x7e,
	0xfe, 0x13, 0x0e, 0x37, 0xec, 0xba, 0xb8, 0xc2, 0x41, 0x12, 0x8d, 0x49, 0x44, 0xe2, 0x97, 0x57,
	0x24, 0xce, 0xc9, 0xcd, 0xb8, 0x6a, 0xc9, 0x97, 0x62, 0xa8, 0xc6, 0xe9, 0xaf, 0xe1, 0x38, 0xcc,
	0xd2, 0xe0, 0xaa, 0x2d, 0xfe, 0x8a, 0xaf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x74, 0x68,
	0x43, 0x35, 0x07, 0x00, 0x00,
}
