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
	HasVMStub        func(string) (bool, error)
	hasVMMutex       sync.RWMutex
	hasVMArgsForCall []struct {
		arg1 string
	}
	hasVMReturns struct {
		result1 bool
		result2 error
	}
	hasVMReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetVMResourcesStub        func(string, int, int) error
	setVMResourcesMutex       sync.RWMutex
	setVMResourcesArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	setVMResourcesReturns struct {
		result1 error
	}
	setVMResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CreateEphemeralDiskStub        func(string, int) error
	createEphemeralDiskMutex       sync.RWMutex
	createEphemeralDiskArgsForCall []struct {
		arg1 string
		arg2 int
	}
	createEphemeralDiskReturns struct {
		result1 error
	}
	createEphemeralDiskReturnsOnCall map[int]struct {
		result1 error
	}
	CreateDiskStub        func(string, int) error
	createDiskMutex       sync.RWMutex
	createDiskArgsForCall []struct {
		arg1 string
		arg2 int
	}
	createDiskReturns struct {
		result1 error
	}
	createDiskReturnsOnCall map[int]struct {
		result1 error
	}
	AttachDiskStub        func(string, string) error
	attachDiskMutex       sync.RWMutex
	attachDiskArgsForCall []struct {
		arg1 string
		arg2 string
	}
	attachDiskReturns struct {
		result1 error
	}
	attachDiskReturnsOnCall map[int]struct {
		result1 error
	}
	DetachDiskStub        func(string, string) error
	detachDiskMutex       sync.RWMutex
	detachDiskArgsForCall []struct {
		arg1 string
		arg2 string
	}
	detachDiskReturns struct {
		result1 error
	}
	detachDiskReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyDiskStub        func(string) error
	destroyDiskMutex       sync.RWMutex
	destroyDiskArgsForCall []struct {
		arg1 string
	}
	destroyDiskReturns struct {
		result1 error
	}
	destroyDiskReturnsOnCall map[int]struct {
		result1 error
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

func (fake *FakeGovcClient) HasVM(arg1 string) (bool, error) {
	fake.hasVMMutex.Lock()
	ret, specificReturn := fake.hasVMReturnsOnCall[len(fake.hasVMArgsForCall)]
	fake.hasVMArgsForCall = append(fake.hasVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HasVM", []interface{}{arg1})
	fake.hasVMMutex.Unlock()
	if fake.HasVMStub != nil {
		return fake.HasVMStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasVMReturns.result1, fake.hasVMReturns.result2
}

func (fake *FakeGovcClient) HasVMCallCount() int {
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	return len(fake.hasVMArgsForCall)
}

func (fake *FakeGovcClient) HasVMArgsForCall(i int) string {
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	return fake.hasVMArgsForCall[i].arg1
}

func (fake *FakeGovcClient) HasVMReturns(result1 bool, result2 error) {
	fake.HasVMStub = nil
	fake.hasVMReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) HasVMReturnsOnCall(i int, result1 bool, result2 error) {
	fake.HasVMStub = nil
	if fake.hasVMReturnsOnCall == nil {
		fake.hasVMReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasVMReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGovcClient) SetVMResources(arg1 string, arg2 int, arg3 int) error {
	fake.setVMResourcesMutex.Lock()
	ret, specificReturn := fake.setVMResourcesReturnsOnCall[len(fake.setVMResourcesArgsForCall)]
	fake.setVMResourcesArgsForCall = append(fake.setVMResourcesArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetVMResources", []interface{}{arg1, arg2, arg3})
	fake.setVMResourcesMutex.Unlock()
	if fake.SetVMResourcesStub != nil {
		return fake.SetVMResourcesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setVMResourcesReturns.result1
}

func (fake *FakeGovcClient) SetVMResourcesCallCount() int {
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	return len(fake.setVMResourcesArgsForCall)
}

func (fake *FakeGovcClient) SetVMResourcesArgsForCall(i int) (string, int, int) {
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	return fake.setVMResourcesArgsForCall[i].arg1, fake.setVMResourcesArgsForCall[i].arg2, fake.setVMResourcesArgsForCall[i].arg3
}

func (fake *FakeGovcClient) SetVMResourcesReturns(result1 error) {
	fake.SetVMResourcesStub = nil
	fake.setVMResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) SetVMResourcesReturnsOnCall(i int, result1 error) {
	fake.SetVMResourcesStub = nil
	if fake.setVMResourcesReturnsOnCall == nil {
		fake.setVMResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setVMResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) CreateEphemeralDisk(arg1 string, arg2 int) error {
	fake.createEphemeralDiskMutex.Lock()
	ret, specificReturn := fake.createEphemeralDiskReturnsOnCall[len(fake.createEphemeralDiskArgsForCall)]
	fake.createEphemeralDiskArgsForCall = append(fake.createEphemeralDiskArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateEphemeralDisk", []interface{}{arg1, arg2})
	fake.createEphemeralDiskMutex.Unlock()
	if fake.CreateEphemeralDiskStub != nil {
		return fake.CreateEphemeralDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createEphemeralDiskReturns.result1
}

func (fake *FakeGovcClient) CreateEphemeralDiskCallCount() int {
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	return len(fake.createEphemeralDiskArgsForCall)
}

func (fake *FakeGovcClient) CreateEphemeralDiskArgsForCall(i int) (string, int) {
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	return fake.createEphemeralDiskArgsForCall[i].arg1, fake.createEphemeralDiskArgsForCall[i].arg2
}

func (fake *FakeGovcClient) CreateEphemeralDiskReturns(result1 error) {
	fake.CreateEphemeralDiskStub = nil
	fake.createEphemeralDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) CreateEphemeralDiskReturnsOnCall(i int, result1 error) {
	fake.CreateEphemeralDiskStub = nil
	if fake.createEphemeralDiskReturnsOnCall == nil {
		fake.createEphemeralDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createEphemeralDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) CreateDisk(arg1 string, arg2 int) error {
	fake.createDiskMutex.Lock()
	ret, specificReturn := fake.createDiskReturnsOnCall[len(fake.createDiskArgsForCall)]
	fake.createDiskArgsForCall = append(fake.createDiskArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateDisk", []interface{}{arg1, arg2})
	fake.createDiskMutex.Unlock()
	if fake.CreateDiskStub != nil {
		return fake.CreateDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createDiskReturns.result1
}

func (fake *FakeGovcClient) CreateDiskCallCount() int {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	return len(fake.createDiskArgsForCall)
}

func (fake *FakeGovcClient) CreateDiskArgsForCall(i int) (string, int) {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	return fake.createDiskArgsForCall[i].arg1, fake.createDiskArgsForCall[i].arg2
}

func (fake *FakeGovcClient) CreateDiskReturns(result1 error) {
	fake.CreateDiskStub = nil
	fake.createDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) CreateDiskReturnsOnCall(i int, result1 error) {
	fake.CreateDiskStub = nil
	if fake.createDiskReturnsOnCall == nil {
		fake.createDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) AttachDisk(arg1 string, arg2 string) error {
	fake.attachDiskMutex.Lock()
	ret, specificReturn := fake.attachDiskReturnsOnCall[len(fake.attachDiskArgsForCall)]
	fake.attachDiskArgsForCall = append(fake.attachDiskArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AttachDisk", []interface{}{arg1, arg2})
	fake.attachDiskMutex.Unlock()
	if fake.AttachDiskStub != nil {
		return fake.AttachDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.attachDiskReturns.result1
}

func (fake *FakeGovcClient) AttachDiskCallCount() int {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	return len(fake.attachDiskArgsForCall)
}

func (fake *FakeGovcClient) AttachDiskArgsForCall(i int) (string, string) {
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	return fake.attachDiskArgsForCall[i].arg1, fake.attachDiskArgsForCall[i].arg2
}

func (fake *FakeGovcClient) AttachDiskReturns(result1 error) {
	fake.AttachDiskStub = nil
	fake.attachDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) AttachDiskReturnsOnCall(i int, result1 error) {
	fake.AttachDiskStub = nil
	if fake.attachDiskReturnsOnCall == nil {
		fake.attachDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.attachDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) DetachDisk(arg1 string, arg2 string) error {
	fake.detachDiskMutex.Lock()
	ret, specificReturn := fake.detachDiskReturnsOnCall[len(fake.detachDiskArgsForCall)]
	fake.detachDiskArgsForCall = append(fake.detachDiskArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DetachDisk", []interface{}{arg1, arg2})
	fake.detachDiskMutex.Unlock()
	if fake.DetachDiskStub != nil {
		return fake.DetachDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.detachDiskReturns.result1
}

func (fake *FakeGovcClient) DetachDiskCallCount() int {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return len(fake.detachDiskArgsForCall)
}

func (fake *FakeGovcClient) DetachDiskArgsForCall(i int) (string, string) {
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	return fake.detachDiskArgsForCall[i].arg1, fake.detachDiskArgsForCall[i].arg2
}

func (fake *FakeGovcClient) DetachDiskReturns(result1 error) {
	fake.DetachDiskStub = nil
	fake.detachDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) DetachDiskReturnsOnCall(i int, result1 error) {
	fake.DetachDiskStub = nil
	if fake.detachDiskReturnsOnCall == nil {
		fake.detachDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.detachDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) DestroyDisk(arg1 string) error {
	fake.destroyDiskMutex.Lock()
	ret, specificReturn := fake.destroyDiskReturnsOnCall[len(fake.destroyDiskArgsForCall)]
	fake.destroyDiskArgsForCall = append(fake.destroyDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyDisk", []interface{}{arg1})
	fake.destroyDiskMutex.Unlock()
	if fake.DestroyDiskStub != nil {
		return fake.DestroyDiskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyDiskReturns.result1
}

func (fake *FakeGovcClient) DestroyDiskCallCount() int {
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
	return len(fake.destroyDiskArgsForCall)
}

func (fake *FakeGovcClient) DestroyDiskArgsForCall(i int) string {
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
	return fake.destroyDiskArgsForCall[i].arg1
}

func (fake *FakeGovcClient) DestroyDiskReturns(result1 error) {
	fake.DestroyDiskStub = nil
	fake.destroyDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGovcClient) DestroyDiskReturnsOnCall(i int, result1 error) {
	fake.DestroyDiskStub = nil
	if fake.destroyDiskReturnsOnCall == nil {
		fake.destroyDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
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
	fake.hasVMMutex.RLock()
	defer fake.hasVMMutex.RUnlock()
	fake.setVMResourcesMutex.RLock()
	defer fake.setVMResourcesMutex.RUnlock()
	fake.createEphemeralDiskMutex.RLock()
	defer fake.createEphemeralDiskMutex.RUnlock()
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	fake.attachDiskMutex.RLock()
	defer fake.attachDiskMutex.RUnlock()
	fake.detachDiskMutex.RLock()
	defer fake.detachDiskMutex.RUnlock()
	fake.destroyDiskMutex.RLock()
	defer fake.destroyDiskMutex.RUnlock()
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
