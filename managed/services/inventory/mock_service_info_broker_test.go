// Code generated by mockery. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	reform "gopkg.in/reform.v1"

	models "github.com/percona/pmm/managed/models"
)

// mockServiceInfoBroker is an autogenerated mock type for the serviceInfoBroker type
type mockServiceInfoBroker struct {
	mock.Mock
}

// GetInfoFromService provides a mock function with given fields: ctx, q, service, agent
func (_m *mockServiceInfoBroker) GetInfoFromService(ctx context.Context, q *reform.Querier, service *models.Service, agent *models.Agent) error {
	ret := _m.Called(ctx, q, service, agent)

	if len(ret) == 0 {
		panic("no return value specified for GetInfoFromService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reform.Querier, *models.Service, *models.Agent) error); ok {
		r0 = rf(ctx, q, service, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockServiceInfoBroker creates a new instance of mockServiceInfoBroker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockServiceInfoBroker(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockServiceInfoBroker {
	mock := &mockServiceInfoBroker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
