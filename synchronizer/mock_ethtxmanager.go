// Code generated by mockery. DO NOT EDIT.

package synchronizer

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"
)

// ethTxManagerMock is an autogenerated mock type for the ethTxManager type
type ethTxManagerMock struct {
	mock.Mock
}

type ethTxManagerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ethTxManagerMock) EXPECT() *ethTxManagerMock_Expecter {
	return &ethTxManagerMock_Expecter{mock: &_m.Mock}
}

// Reorg provides a mock function with given fields: ctx, fromBlockNumber, dbTx
func (_m *ethTxManagerMock) Reorg(ctx context.Context, fromBlockNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, fromBlockNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, fromBlockNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ethTxManagerMock_Reorg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reorg'
type ethTxManagerMock_Reorg_Call struct {
	*mock.Call
}

// Reorg is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - dbTx pgx.Tx
func (_e *ethTxManagerMock_Expecter) Reorg(ctx interface{}, fromBlockNumber interface{}, dbTx interface{}) *ethTxManagerMock_Reorg_Call {
	return &ethTxManagerMock_Reorg_Call{Call: _e.mock.On("Reorg", ctx, fromBlockNumber, dbTx)}
}

func (_c *ethTxManagerMock_Reorg_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, dbTx pgx.Tx)) *ethTxManagerMock_Reorg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *ethTxManagerMock_Reorg_Call) Return(_a0 error) *ethTxManagerMock_Reorg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ethTxManagerMock_Reorg_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) error) *ethTxManagerMock_Reorg_Call {
	_c.Call.Return(run)
	return _c
}

// newEthTxManagerMock creates a new instance of ethTxManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newEthTxManagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ethTxManagerMock {
	mock := &ethTxManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
