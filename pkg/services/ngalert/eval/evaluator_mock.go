// Code generated by mockery v2.10.0. DO NOT EDIT.

package eval

import (
	backend "github.com/grafana/grafana-plugin-sdk-go/backend"
	expr "github.com/grafana/grafana/pkg/expr"

	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"

	time "time"
)

// FakeEvaluator is an autogenerated mock type for the Evaluator type
type FakeEvaluator struct {
	mock.Mock
}

type FakeEvaluator_Expecter struct {
	mock *mock.Mock
}

func (_m *FakeEvaluator) EXPECT() *FakeEvaluator_Expecter {
	return &FakeEvaluator_Expecter{mock: &_m.Mock}
}

// ConditionEval provides a mock function with given fields: condition, now, expressionService
func (_m *FakeEvaluator) ConditionEval(condition *models.Condition, now time.Time, expressionService *expr.Service) (Results, error) {
	ret := _m.Called(condition, now, expressionService)

	var r0 Results
	if rf, ok := ret.Get(0).(func(*models.Condition, time.Time, *expr.Service) Results); ok {
		r0 = rf(condition, now, expressionService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Results)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Condition, time.Time, *expr.Service) error); ok {
		r1 = rf(condition, now, expressionService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeEvaluator_ConditionEval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConditionEval'
type FakeEvaluator_ConditionEval_Call struct {
	*mock.Call
}

// ConditionEval is a helper method to define mock.On call
//  - condition *models.Condition
//  - now time.Time
//  - expressionService *expr.Service
func (_e *FakeEvaluator_Expecter) ConditionEval(condition interface{}, now interface{}, expressionService interface{}) *FakeEvaluator_ConditionEval_Call {
	return &FakeEvaluator_ConditionEval_Call{Call: _e.mock.On("ConditionEval", condition, now, expressionService)}
}

func (_c *FakeEvaluator_ConditionEval_Call) Run(run func(condition *models.Condition, now time.Time, expressionService *expr.Service)) *FakeEvaluator_ConditionEval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Condition), args[1].(time.Time), args[2].(*expr.Service))
	})
	return _c
}

func (_c *FakeEvaluator_ConditionEval_Call) Return(_a0 Results, _a1 error) *FakeEvaluator_ConditionEval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// QueriesAndExpressionsEval provides a mock function with given fields: orgID, data, now, expressionService
func (_m *FakeEvaluator) QueriesAndExpressionsEval(orgID int64, data []models.AlertQuery, now time.Time, expressionService *expr.Service) (*backend.QueryDataResponse, error) {
	ret := _m.Called(orgID, data, now, expressionService)

	var r0 *backend.QueryDataResponse
	if rf, ok := ret.Get(0).(func(int64, []models.AlertQuery, time.Time, *expr.Service) *backend.QueryDataResponse); ok {
		r0 = rf(orgID, data, now, expressionService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backend.QueryDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, []models.AlertQuery, time.Time, *expr.Service) error); ok {
		r1 = rf(orgID, data, now, expressionService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeEvaluator_QueriesAndExpressionsEval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueriesAndExpressionsEval'
type FakeEvaluator_QueriesAndExpressionsEval_Call struct {
	*mock.Call
}

// QueriesAndExpressionsEval is a helper method to define mock.On call
//  - orgID int64
//  - data []models.AlertQuery
//  - now time.Time
//  - expressionService *expr.Service
func (_e *FakeEvaluator_Expecter) QueriesAndExpressionsEval(orgID interface{}, data interface{}, now interface{}, expressionService interface{}) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	return &FakeEvaluator_QueriesAndExpressionsEval_Call{Call: _e.mock.On("QueriesAndExpressionsEval", orgID, data, now, expressionService)}
}

func (_c *FakeEvaluator_QueriesAndExpressionsEval_Call) Run(run func(orgID int64, data []models.AlertQuery, now time.Time, expressionService *expr.Service)) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].([]models.AlertQuery), args[2].(time.Time), args[3].(*expr.Service))
	})
	return _c
}

func (_c *FakeEvaluator_QueriesAndExpressionsEval_Call) Return(_a0 *backend.QueryDataResponse, _a1 error) *FakeEvaluator_QueriesAndExpressionsEval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
