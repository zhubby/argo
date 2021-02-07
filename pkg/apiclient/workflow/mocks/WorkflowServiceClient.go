// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo/v3/pkg/apis/workflow/v1alpha1"

	workflow "github.com/argoproj/argo/v3/pkg/apiclient/workflow"
)

// WorkflowServiceClient is an autogenerated mock type for the WorkflowServiceClient type
type WorkflowServiceClient struct {
	mock.Mock
}

// CreateWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) CreateWorkflow(ctx context.Context, in *workflow.WorkflowCreateRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowCreateRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowCreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) DeleteWorkflow(ctx context.Context, in *workflow.WorkflowDeleteRequest, opts ...grpc.CallOption) (*workflow.WorkflowDeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *workflow.WorkflowDeleteResponse
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowDeleteRequest, ...grpc.CallOption) *workflow.WorkflowDeleteResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workflow.WorkflowDeleteResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowDeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) GetWorkflow(ctx context.Context, in *workflow.WorkflowGetRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowGetRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LintWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) LintWorkflow(ctx context.Context, in *workflow.WorkflowLintRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowLintRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowLintRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflows provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) ListWorkflows(ctx context.Context, in *workflow.WorkflowListRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowList
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowListRequest, ...grpc.CallOption) *v1alpha1.WorkflowList); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodLogs provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) PodLogs(ctx context.Context, in *workflow.WorkflowLogRequest, opts ...grpc.CallOption) (workflow.WorkflowService_PodLogsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 workflow.WorkflowService_PodLogsClient
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowLogRequest, ...grpc.CallOption) workflow.WorkflowService_PodLogsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workflow.WorkflowService_PodLogsClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowLogRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResubmitWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) ResubmitWorkflow(ctx context.Context, in *workflow.WorkflowResubmitRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowResubmitRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowResubmitRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResumeWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) ResumeWorkflow(ctx context.Context, in *workflow.WorkflowResumeRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowResumeRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowResumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetryWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) RetryWorkflow(ctx context.Context, in *workflow.WorkflowRetryRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowRetryRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowRetryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) SetWorkflow(ctx context.Context, in *workflow.WorkflowSetRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowSetRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowSetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) StopWorkflow(ctx context.Context, in *workflow.WorkflowStopRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowStopRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowStopRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) SubmitWorkflow(ctx context.Context, in *workflow.WorkflowSubmitRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowSubmitRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowSubmitRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuspendWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) SuspendWorkflow(ctx context.Context, in *workflow.WorkflowSuspendRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowSuspendRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowSuspendRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminateWorkflow provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) TerminateWorkflow(ctx context.Context, in *workflow.WorkflowTerminateRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowTerminateRequest, ...grpc.CallOption) *v1alpha1.Workflow); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowTerminateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchEvents provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) WatchEvents(ctx context.Context, in *workflow.WatchEventsRequest, opts ...grpc.CallOption) (workflow.WorkflowService_WatchEventsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 workflow.WorkflowService_WatchEventsClient
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WatchEventsRequest, ...grpc.CallOption) workflow.WorkflowService_WatchEventsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workflow.WorkflowService_WatchEventsClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WatchEventsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchWorkflows provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) WatchWorkflows(ctx context.Context, in *workflow.WatchWorkflowsRequest, opts ...grpc.CallOption) (workflow.WorkflowService_WatchWorkflowsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 workflow.WorkflowService_WatchWorkflowsClient
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WatchWorkflowsRequest, ...grpc.CallOption) workflow.WorkflowService_WatchWorkflowsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workflow.WorkflowService_WatchWorkflowsClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WatchWorkflowsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkflowLogs provides a mock function with given fields: ctx, in, opts
func (_m *WorkflowServiceClient) WorkflowLogs(ctx context.Context, in *workflow.WorkflowLogRequest, opts ...grpc.CallOption) (workflow.WorkflowService_WorkflowLogsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 workflow.WorkflowService_WorkflowLogsClient
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.WorkflowLogRequest, ...grpc.CallOption) workflow.WorkflowService_WorkflowLogsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workflow.WorkflowService_WorkflowLogsClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *workflow.WorkflowLogRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
