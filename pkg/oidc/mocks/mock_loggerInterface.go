// Code generated by mockery v2.46.3. DO NOT EDIT.

package oidcmocks

import (
	mock "github.com/stretchr/testify/mock"

	zap "go.uber.org/zap"
)

// MockLoggerInterface is an autogenerated mock type for the LoggerInterface type
type MockLoggerInterface struct {
	mock.Mock
}

type MockLoggerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoggerInterface) EXPECT() *MockLoggerInterface_Expecter {
	return &MockLoggerInterface_Expecter{mock: &_m.Mock}
}

// Debugw provides a mock function with given fields: message, keysAndValues
func (_m *MockLoggerInterface) Debugw(message string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLoggerInterface_Debugw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugw'
type MockLoggerInterface_Debugw_Call struct {
	*mock.Call
}

// Debugw is a helper method to define mock.On call
//   - message string
//   - keysAndValues ...interface{}
func (_e *MockLoggerInterface_Expecter) Debugw(message interface{}, keysAndValues ...interface{}) *MockLoggerInterface_Debugw_Call {
	return &MockLoggerInterface_Debugw_Call{Call: _e.mock.On("Debugw",
		append([]interface{}{message}, keysAndValues...)...)}
}

func (_c *MockLoggerInterface_Debugw_Call) Run(run func(message string, keysAndValues ...interface{})) *MockLoggerInterface_Debugw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLoggerInterface_Debugw_Call) Return() *MockLoggerInterface_Debugw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLoggerInterface_Debugw_Call) RunAndReturn(run func(string, ...interface{})) *MockLoggerInterface_Debugw_Call {
	_c.Call.Return(run)
	return _c
}

// Errorw provides a mock function with given fields: message, keysAndValues
func (_m *MockLoggerInterface) Errorw(message string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLoggerInterface_Errorw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorw'
type MockLoggerInterface_Errorw_Call struct {
	*mock.Call
}

// Errorw is a helper method to define mock.On call
//   - message string
//   - keysAndValues ...interface{}
func (_e *MockLoggerInterface_Expecter) Errorw(message interface{}, keysAndValues ...interface{}) *MockLoggerInterface_Errorw_Call {
	return &MockLoggerInterface_Errorw_Call{Call: _e.mock.On("Errorw",
		append([]interface{}{message}, keysAndValues...)...)}
}

func (_c *MockLoggerInterface_Errorw_Call) Run(run func(message string, keysAndValues ...interface{})) *MockLoggerInterface_Errorw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLoggerInterface_Errorw_Call) Return() *MockLoggerInterface_Errorw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLoggerInterface_Errorw_Call) RunAndReturn(run func(string, ...interface{})) *MockLoggerInterface_Errorw_Call {
	_c.Call.Return(run)
	return _c
}

// Infow provides a mock function with given fields: message, keysAndValues
func (_m *MockLoggerInterface) Infow(message string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLoggerInterface_Infow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infow'
type MockLoggerInterface_Infow_Call struct {
	*mock.Call
}

// Infow is a helper method to define mock.On call
//   - message string
//   - keysAndValues ...interface{}
func (_e *MockLoggerInterface_Expecter) Infow(message interface{}, keysAndValues ...interface{}) *MockLoggerInterface_Infow_Call {
	return &MockLoggerInterface_Infow_Call{Call: _e.mock.On("Infow",
		append([]interface{}{message}, keysAndValues...)...)}
}

func (_c *MockLoggerInterface_Infow_Call) Run(run func(message string, keysAndValues ...interface{})) *MockLoggerInterface_Infow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLoggerInterface_Infow_Call) Return() *MockLoggerInterface_Infow_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLoggerInterface_Infow_Call) RunAndReturn(run func(string, ...interface{})) *MockLoggerInterface_Infow_Call {
	_c.Call.Return(run)
	return _c
}

// Warnw provides a mock function with given fields: message, keysAndValues
func (_m *MockLoggerInterface) Warnw(message string, keysAndValues ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, keysAndValues...)
	_m.Called(_ca...)
}

// MockLoggerInterface_Warnw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnw'
type MockLoggerInterface_Warnw_Call struct {
	*mock.Call
}

// Warnw is a helper method to define mock.On call
//   - message string
//   - keysAndValues ...interface{}
func (_e *MockLoggerInterface_Expecter) Warnw(message interface{}, keysAndValues ...interface{}) *MockLoggerInterface_Warnw_Call {
	return &MockLoggerInterface_Warnw_Call{Call: _e.mock.On("Warnw",
		append([]interface{}{message}, keysAndValues...)...)}
}

func (_c *MockLoggerInterface_Warnw_Call) Run(run func(message string, keysAndValues ...interface{})) *MockLoggerInterface_Warnw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLoggerInterface_Warnw_Call) Return() *MockLoggerInterface_Warnw_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLoggerInterface_Warnw_Call) RunAndReturn(run func(string, ...interface{})) *MockLoggerInterface_Warnw_Call {
	_c.Call.Return(run)
	return _c
}

// With provides a mock function with given fields: args
func (_m *MockLoggerInterface) With(args ...interface{}) *zap.SugaredLogger {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for With")
	}

	var r0 *zap.SugaredLogger
	if rf, ok := ret.Get(0).(func(...interface{}) *zap.SugaredLogger); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zap.SugaredLogger)
		}
	}

	return r0
}

// MockLoggerInterface_With_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'With'
type MockLoggerInterface_With_Call struct {
	*mock.Call
}

// With is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockLoggerInterface_Expecter) With(args ...interface{}) *MockLoggerInterface_With_Call {
	return &MockLoggerInterface_With_Call{Call: _e.mock.On("With",
		append([]interface{}{}, args...)...)}
}

func (_c *MockLoggerInterface_With_Call) Run(run func(args ...interface{})) *MockLoggerInterface_With_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLoggerInterface_With_Call) Return(_a0 *zap.SugaredLogger) *MockLoggerInterface_With_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLoggerInterface_With_Call) RunAndReturn(run func(...interface{}) *zap.SugaredLogger) *MockLoggerInterface_With_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoggerInterface creates a new instance of MockLoggerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoggerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoggerInterface {
	mock := &MockLoggerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
