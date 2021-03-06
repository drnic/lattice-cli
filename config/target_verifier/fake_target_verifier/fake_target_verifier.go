// This file was generated by counterfeiter
package fake_target_verifier

import (
	"sync"

	"github.com/pivotal-cf-experimental/lattice-cli/config/target_verifier"
)

type FakeTargetVerifier struct {
	ValidateReceptorStub        func(name string) bool
	validateReceptorMutex       sync.RWMutex
	validateReceptorArgsForCall []struct {
		name string
	}
	validateReceptorReturns struct {
		result1 bool
	}
}

func (fake *FakeTargetVerifier) ValidateReceptor(name string) bool {
	fake.validateReceptorMutex.Lock()
	fake.validateReceptorArgsForCall = append(fake.validateReceptorArgsForCall, struct {
		name string
	}{name})
	fake.validateReceptorMutex.Unlock()
	if fake.ValidateReceptorStub != nil {
		return fake.ValidateReceptorStub(name)
	} else {
		return fake.validateReceptorReturns.result1
	}
}

func (fake *FakeTargetVerifier) ValidateReceptorCallCount() int {
	fake.validateReceptorMutex.RLock()
	defer fake.validateReceptorMutex.RUnlock()
	return len(fake.validateReceptorArgsForCall)
}

func (fake *FakeTargetVerifier) ValidateReceptorArgsForCall(i int) string {
	fake.validateReceptorMutex.RLock()
	defer fake.validateReceptorMutex.RUnlock()
	return fake.validateReceptorArgsForCall[i].name
}

func (fake *FakeTargetVerifier) ValidateReceptorReturns(result1 bool) {
	fake.ValidateReceptorStub = nil
	fake.validateReceptorReturns = struct {
		result1 bool
	}{result1}
}

var _ target_verifier.TargetVerifier = new(FakeTargetVerifier)
