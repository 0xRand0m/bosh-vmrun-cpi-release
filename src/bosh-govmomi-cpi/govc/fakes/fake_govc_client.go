// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-govmomi-cpi/govc"
	"sync"
)

type FakeGovcClient struct {
	ImportOvfStub        func(string, string) (string, error)
	importOvfMutex       sync.RWMutex
	importOvfArgsForCall []struct {
		arg1 string
		arg2 string
	}
	importOvfReturns struct {
		result1 string
		result2 error
	}
	importOvfReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CloneVMStub        func(string, string) (string, error)
	cloneVMMutex       sync.RWMutex
	cloneVMArgsForCall []struct {
		arg1 string
		arg2 string
	}
	cloneVMReturns struct {
		result1 string
		result2 error
	}
	cloneVMReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UpdateVMIsoStub        func(string, string) (string, error)
	updateVMIsoMutex       sync.RWMutex
	updateVMIsoArgsForCall []struct {
		arg1 string
		arg2 string
	}
	updateVMIsoReturns struct {
		result1 string
		result2 error
	}
	updateVMIsoReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	StartVMStub        func(string) (string, error)
	startVMMutex       sync.RWMutex
	startVMArgsForCall []struct {
		arg1 string
	}
	startVMReturns struct {
		result1 string
		result2 error
	}
	startVMReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DestroyVMStub        func(string) (string, error)
	destroyVMMutex       sync.RWMutex
	destroyVMArgsForCall []struct {
		arg1 string
	}
	destroyVMReturns struct {
		result1 string
		result2 error
	}
	destroyVMReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGovcClient) ImportOvf(arg1 string, arg2 string) (string, error) {
	fake.importOvfMutex.Lock()
	ret, specificReturn := fake.importOvfReturnsOnCall[len(fake.importOvfArgsForCall)]
	fake.importOvfArgsForCall = append(fake.importOvfArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ImportOvf", []interface{}{arg1, arg2})
	fake.importOvfMutex.Unlock()
	if fake.ImportOvfStub != nil {
		return fake.ImportOvfStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.importOvfReturns.result1, fake.importOvfReturns.result2
}

func (fake *FakeGovcClient) ImportOvfCallCount() int {
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	return len(fake.importOvfArgsForCall)
}

func (fake *FakeGovcClient) ImportOvfArgsForCall(i int) (string, string) {
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	return fake.importOvfArgsForCall[i].arg1, fake.importOvfArgsForCall[i].arg2
}

func (fake *FakeGovcClient) ImportOvfReturns(result1 string, result2 error) {
	fake.ImportOvfStub = nil
	fake.importOvfReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) ImportOvfReturnsOnCall(i int, result1 string, result2 error) {
	fake.ImportOvfStub = nil
	if fake.importOvfReturnsOnCall == nil {
		fake.importOvfReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.importOvfReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) CloneVM(arg1 string, arg2 string) (string, error) {
	fake.cloneVMMutex.Lock()
	ret, specificReturn := fake.cloneVMReturnsOnCall[len(fake.cloneVMArgsForCall)]
	fake.cloneVMArgsForCall = append(fake.cloneVMArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CloneVM", []interface{}{arg1, arg2})
	fake.cloneVMMutex.Unlock()
	if fake.CloneVMStub != nil {
		return fake.CloneVMStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.cloneVMReturns.result1, fake.cloneVMReturns.result2
}

func (fake *FakeGovcClient) CloneVMCallCount() int {
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	return len(fake.cloneVMArgsForCall)
}

func (fake *FakeGovcClient) CloneVMArgsForCall(i int) (string, string) {
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	return fake.cloneVMArgsForCall[i].arg1, fake.cloneVMArgsForCall[i].arg2
}

func (fake *FakeGovcClient) CloneVMReturns(result1 string, result2 error) {
	fake.CloneVMStub = nil
	fake.cloneVMReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) CloneVMReturnsOnCall(i int, result1 string, result2 error) {
	fake.CloneVMStub = nil
	if fake.cloneVMReturnsOnCall == nil {
		fake.cloneVMReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.cloneVMReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) UpdateVMIso(arg1 string, arg2 string) (string, error) {
	fake.updateVMIsoMutex.Lock()
	ret, specificReturn := fake.updateVMIsoReturnsOnCall[len(fake.updateVMIsoArgsForCall)]
	fake.updateVMIsoArgsForCall = append(fake.updateVMIsoArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UpdateVMIso", []interface{}{arg1, arg2})
	fake.updateVMIsoMutex.Unlock()
	if fake.UpdateVMIsoStub != nil {
		return fake.UpdateVMIsoStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateVMIsoReturns.result1, fake.updateVMIsoReturns.result2
}

func (fake *FakeGovcClient) UpdateVMIsoCallCount() int {
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	return len(fake.updateVMIsoArgsForCall)
}

func (fake *FakeGovcClient) UpdateVMIsoArgsForCall(i int) (string, string) {
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	return fake.updateVMIsoArgsForCall[i].arg1, fake.updateVMIsoArgsForCall[i].arg2
}

func (fake *FakeGovcClient) UpdateVMIsoReturns(result1 string, result2 error) {
	fake.UpdateVMIsoStub = nil
	fake.updateVMIsoReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) UpdateVMIsoReturnsOnCall(i int, result1 string, result2 error) {
	fake.UpdateVMIsoStub = nil
	if fake.updateVMIsoReturnsOnCall == nil {
		fake.updateVMIsoReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.updateVMIsoReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) StartVM(arg1 string) (string, error) {
	fake.startVMMutex.Lock()
	ret, specificReturn := fake.startVMReturnsOnCall[len(fake.startVMArgsForCall)]
	fake.startVMArgsForCall = append(fake.startVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartVM", []interface{}{arg1})
	fake.startVMMutex.Unlock()
	if fake.StartVMStub != nil {
		return fake.StartVMStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startVMReturns.result1, fake.startVMReturns.result2
}

func (fake *FakeGovcClient) StartVMCallCount() int {
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	return len(fake.startVMArgsForCall)
}

func (fake *FakeGovcClient) StartVMArgsForCall(i int) string {
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	return fake.startVMArgsForCall[i].arg1
}

func (fake *FakeGovcClient) StartVMReturns(result1 string, result2 error) {
	fake.StartVMStub = nil
	fake.startVMReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) StartVMReturnsOnCall(i int, result1 string, result2 error) {
	fake.StartVMStub = nil
	if fake.startVMReturnsOnCall == nil {
		fake.startVMReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startVMReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) DestroyVM(arg1 string) (string, error) {
	fake.destroyVMMutex.Lock()
	ret, specificReturn := fake.destroyVMReturnsOnCall[len(fake.destroyVMArgsForCall)]
	fake.destroyVMArgsForCall = append(fake.destroyVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyVM", []interface{}{arg1})
	fake.destroyVMMutex.Unlock()
	if fake.DestroyVMStub != nil {
		return fake.DestroyVMStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.destroyVMReturns.result1, fake.destroyVMReturns.result2
}

func (fake *FakeGovcClient) DestroyVMCallCount() int {
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	return len(fake.destroyVMArgsForCall)
}

func (fake *FakeGovcClient) DestroyVMArgsForCall(i int) string {
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	return fake.destroyVMArgsForCall[i].arg1
}

func (fake *FakeGovcClient) DestroyVMReturns(result1 string, result2 error) {
	fake.DestroyVMStub = nil
	fake.destroyVMReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) DestroyVMReturnsOnCall(i int, result1 string, result2 error) {
	fake.DestroyVMStub = nil
	if fake.destroyVMReturnsOnCall == nil {
		fake.destroyVMReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.destroyVMReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.importOvfMutex.RLock()
	defer fake.importOvfMutex.RUnlock()
	fake.cloneVMMutex.RLock()
	defer fake.cloneVMMutex.RUnlock()
	fake.updateVMIsoMutex.RLock()
	defer fake.updateVMIsoMutex.RUnlock()
	fake.startVMMutex.RLock()
	defer fake.startVMMutex.RUnlock()
	fake.destroyVMMutex.RLock()
	defer fake.destroyVMMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGovcClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ govc.GovcClient = new(FakeGovcClient)
