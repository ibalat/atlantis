// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: ProjectCommandRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	"reflect"
	"time"
)

type MockProjectCommandRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandRunner(options ...pegomock.Option) *MockProjectCommandRunner {
	mock := &MockProjectCommandRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectCommandRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectCommandRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectCommandRunner) Apply(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Apply", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) ApprovePolicies(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("ApprovePolicies", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) Import(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Import", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) Plan(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Plan", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) PolicyCheck(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("PolicyCheck", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) StateRm(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("StateRm", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) Version(ctx command.ProjectContext) command.ProjectResult {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandRunner().")
	}
	_params := []pegomock.Param{ctx}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Version", _params, []reflect.Type{reflect.TypeOf((*command.ProjectResult)(nil)).Elem()})
	var _ret0 command.ProjectResult
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(command.ProjectResult)
		}
	}
	return _ret0
}

func (mock *MockProjectCommandRunner) VerifyWasCalledOnce() *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectCommandRunner {
	return &VerifierMockProjectCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectCommandRunner struct {
	mock                   *MockProjectCommandRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectCommandRunner) Apply(ctx command.ProjectContext) *MockProjectCommandRunner_Apply_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Apply", _params, verifier.timeout)
	return &MockProjectCommandRunner_Apply_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Apply_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Apply_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Apply_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) ApprovePolicies(ctx command.ProjectContext) *MockProjectCommandRunner_ApprovePolicies_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ApprovePolicies", _params, verifier.timeout)
	return &MockProjectCommandRunner_ApprovePolicies_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_ApprovePolicies_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_ApprovePolicies_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_ApprovePolicies_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) Import(ctx command.ProjectContext) *MockProjectCommandRunner_Import_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Import", _params, verifier.timeout)
	return &MockProjectCommandRunner_Import_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Import_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Import_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Import_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) Plan(ctx command.ProjectContext) *MockProjectCommandRunner_Plan_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Plan", _params, verifier.timeout)
	return &MockProjectCommandRunner_Plan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Plan_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Plan_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Plan_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) PolicyCheck(ctx command.ProjectContext) *MockProjectCommandRunner_PolicyCheck_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PolicyCheck", _params, verifier.timeout)
	return &MockProjectCommandRunner_PolicyCheck_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_PolicyCheck_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_PolicyCheck_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_PolicyCheck_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) StateRm(ctx command.ProjectContext) *MockProjectCommandRunner_StateRm_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "StateRm", _params, verifier.timeout)
	return &MockProjectCommandRunner_StateRm_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_StateRm_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_StateRm_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_StateRm_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}

func (verifier *VerifierMockProjectCommandRunner) Version(ctx command.ProjectContext) *MockProjectCommandRunner_Version_OngoingVerification {
	_params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", _params, verifier.timeout)
	return &MockProjectCommandRunner_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectCommandRunner_Version_OngoingVerification struct {
	mock              *MockProjectCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectCommandRunner_Version_OngoingVerification) GetCapturedArguments() command.ProjectContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockProjectCommandRunner_Version_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
	}
	return
}
