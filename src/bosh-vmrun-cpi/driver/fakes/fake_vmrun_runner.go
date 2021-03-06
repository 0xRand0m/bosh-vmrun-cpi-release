// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-vmrun-cpi/driver"
	"sync"
)

type FakeVmrunRunner struct {
	CloneStub        func(string, string, string) error
	cloneMutex       sync.RWMutex
	cloneArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	cloneReturns struct {
		result1 error
	}
	cloneReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigureStub        func() error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct {
	}
	configureReturns struct {
		result1 error
	}
	configureReturnsOnCall map[int]struct {
		result1 error
	}
	CopyFileFromHostToGuestStub        func(string, string, string, string, string) error
	copyFileFromHostToGuestMutex       sync.RWMutex
	copyFileFromHostToGuestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	copyFileFromHostToGuestReturns struct {
		result1 error
	}
	copyFileFromHostToGuestReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	HardStopStub        func(string) error
	hardStopMutex       sync.RWMutex
	hardStopArgsForCall []struct {
		arg1 string
	}
	hardStopReturns struct {
		result1 error
	}
	hardStopReturnsOnCall map[int]struct {
		result1 error
	}
	IsPlayerStub        func() bool
	isPlayerMutex       sync.RWMutex
	isPlayerArgsForCall []struct {
	}
	isPlayerReturns struct {
		result1 bool
	}
	isPlayerReturnsOnCall map[int]struct {
		result1 bool
	}
	ListStub        func() (string, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 string
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListProcessesInGuestStub        func(string, string, string) (string, error)
	listProcessesInGuestMutex       sync.RWMutex
	listProcessesInGuestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	listProcessesInGuestReturns struct {
		result1 string
		result2 error
	}
	listProcessesInGuestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunProgramInGuestStub        func(string, string, string, string, string) error
	runProgramInGuestMutex       sync.RWMutex
	runProgramInGuestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	runProgramInGuestReturns struct {
		result1 error
	}
	runProgramInGuestReturnsOnCall map[int]struct {
		result1 error
	}
	SoftStopStub        func(string) error
	softStopMutex       sync.RWMutex
	softStopArgsForCall []struct {
		arg1 string
	}
	softStopReturns struct {
		result1 error
	}
	softStopReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(string) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 string
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVmrunRunner) Clone(arg1 string, arg2 string, arg3 string) error {
	fake.cloneMutex.Lock()
	ret, specificReturn := fake.cloneReturnsOnCall[len(fake.cloneArgsForCall)]
	fake.cloneArgsForCall = append(fake.cloneArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Clone", []interface{}{arg1, arg2, arg3})
	fake.cloneMutex.Unlock()
	if fake.CloneStub != nil {
		return fake.CloneStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloneReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) CloneCallCount() int {
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	return len(fake.cloneArgsForCall)
}

func (fake *FakeVmrunRunner) CloneCalls(stub func(string, string, string) error) {
	fake.cloneMutex.Lock()
	defer fake.cloneMutex.Unlock()
	fake.CloneStub = stub
}

func (fake *FakeVmrunRunner) CloneArgsForCall(i int) (string, string, string) {
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	argsForCall := fake.cloneArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVmrunRunner) CloneReturns(result1 error) {
	fake.cloneMutex.Lock()
	defer fake.cloneMutex.Unlock()
	fake.CloneStub = nil
	fake.cloneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) CloneReturnsOnCall(i int, result1 error) {
	fake.cloneMutex.Lock()
	defer fake.cloneMutex.Unlock()
	fake.CloneStub = nil
	if fake.cloneReturnsOnCall == nil {
		fake.cloneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cloneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) Configure() error {
	fake.configureMutex.Lock()
	ret, specificReturn := fake.configureReturnsOnCall[len(fake.configureArgsForCall)]
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct {
	}{})
	fake.recordInvocation("Configure", []interface{}{})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configureReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *FakeVmrunRunner) ConfigureCalls(stub func() error) {
	fake.configureMutex.Lock()
	defer fake.configureMutex.Unlock()
	fake.ConfigureStub = stub
}

func (fake *FakeVmrunRunner) ConfigureReturns(result1 error) {
	fake.configureMutex.Lock()
	defer fake.configureMutex.Unlock()
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) ConfigureReturnsOnCall(i int, result1 error) {
	fake.configureMutex.Lock()
	defer fake.configureMutex.Unlock()
	fake.ConfigureStub = nil
	if fake.configureReturnsOnCall == nil {
		fake.configureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuest(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.copyFileFromHostToGuestMutex.Lock()
	ret, specificReturn := fake.copyFileFromHostToGuestReturnsOnCall[len(fake.copyFileFromHostToGuestArgsForCall)]
	fake.copyFileFromHostToGuestArgsForCall = append(fake.copyFileFromHostToGuestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CopyFileFromHostToGuest", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.copyFileFromHostToGuestMutex.Unlock()
	if fake.CopyFileFromHostToGuestStub != nil {
		return fake.CopyFileFromHostToGuestStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.copyFileFromHostToGuestReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuestCallCount() int {
	fake.copyFileFromHostToGuestMutex.RLock()
	defer fake.copyFileFromHostToGuestMutex.RUnlock()
	return len(fake.copyFileFromHostToGuestArgsForCall)
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuestCalls(stub func(string, string, string, string, string) error) {
	fake.copyFileFromHostToGuestMutex.Lock()
	defer fake.copyFileFromHostToGuestMutex.Unlock()
	fake.CopyFileFromHostToGuestStub = stub
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuestArgsForCall(i int) (string, string, string, string, string) {
	fake.copyFileFromHostToGuestMutex.RLock()
	defer fake.copyFileFromHostToGuestMutex.RUnlock()
	argsForCall := fake.copyFileFromHostToGuestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuestReturns(result1 error) {
	fake.copyFileFromHostToGuestMutex.Lock()
	defer fake.copyFileFromHostToGuestMutex.Unlock()
	fake.CopyFileFromHostToGuestStub = nil
	fake.copyFileFromHostToGuestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) CopyFileFromHostToGuestReturnsOnCall(i int, result1 error) {
	fake.copyFileFromHostToGuestMutex.Lock()
	defer fake.copyFileFromHostToGuestMutex.Unlock()
	fake.CopyFileFromHostToGuestStub = nil
	if fake.copyFileFromHostToGuestReturnsOnCall == nil {
		fake.copyFileFromHostToGuestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyFileFromHostToGuestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeVmrunRunner) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeVmrunRunner) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVmrunRunner) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) HardStop(arg1 string) error {
	fake.hardStopMutex.Lock()
	ret, specificReturn := fake.hardStopReturnsOnCall[len(fake.hardStopArgsForCall)]
	fake.hardStopArgsForCall = append(fake.hardStopArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HardStop", []interface{}{arg1})
	fake.hardStopMutex.Unlock()
	if fake.HardStopStub != nil {
		return fake.HardStopStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hardStopReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) HardStopCallCount() int {
	fake.hardStopMutex.RLock()
	defer fake.hardStopMutex.RUnlock()
	return len(fake.hardStopArgsForCall)
}

func (fake *FakeVmrunRunner) HardStopCalls(stub func(string) error) {
	fake.hardStopMutex.Lock()
	defer fake.hardStopMutex.Unlock()
	fake.HardStopStub = stub
}

func (fake *FakeVmrunRunner) HardStopArgsForCall(i int) string {
	fake.hardStopMutex.RLock()
	defer fake.hardStopMutex.RUnlock()
	argsForCall := fake.hardStopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVmrunRunner) HardStopReturns(result1 error) {
	fake.hardStopMutex.Lock()
	defer fake.hardStopMutex.Unlock()
	fake.HardStopStub = nil
	fake.hardStopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) HardStopReturnsOnCall(i int, result1 error) {
	fake.hardStopMutex.Lock()
	defer fake.hardStopMutex.Unlock()
	fake.HardStopStub = nil
	if fake.hardStopReturnsOnCall == nil {
		fake.hardStopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.hardStopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) IsPlayer() bool {
	fake.isPlayerMutex.Lock()
	ret, specificReturn := fake.isPlayerReturnsOnCall[len(fake.isPlayerArgsForCall)]
	fake.isPlayerArgsForCall = append(fake.isPlayerArgsForCall, struct {
	}{})
	fake.recordInvocation("IsPlayer", []interface{}{})
	fake.isPlayerMutex.Unlock()
	if fake.IsPlayerStub != nil {
		return fake.IsPlayerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isPlayerReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) IsPlayerCallCount() int {
	fake.isPlayerMutex.RLock()
	defer fake.isPlayerMutex.RUnlock()
	return len(fake.isPlayerArgsForCall)
}

func (fake *FakeVmrunRunner) IsPlayerCalls(stub func() bool) {
	fake.isPlayerMutex.Lock()
	defer fake.isPlayerMutex.Unlock()
	fake.IsPlayerStub = stub
}

func (fake *FakeVmrunRunner) IsPlayerReturns(result1 bool) {
	fake.isPlayerMutex.Lock()
	defer fake.isPlayerMutex.Unlock()
	fake.IsPlayerStub = nil
	fake.isPlayerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeVmrunRunner) IsPlayerReturnsOnCall(i int, result1 bool) {
	fake.isPlayerMutex.Lock()
	defer fake.isPlayerMutex.Unlock()
	fake.IsPlayerStub = nil
	if fake.isPlayerReturnsOnCall == nil {
		fake.isPlayerReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isPlayerReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeVmrunRunner) List() (string, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVmrunRunner) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeVmrunRunner) ListCalls(stub func() (string, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeVmrunRunner) ListReturns(result1 string, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVmrunRunner) ListReturnsOnCall(i int, result1 string, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVmrunRunner) ListProcessesInGuest(arg1 string, arg2 string, arg3 string) (string, error) {
	fake.listProcessesInGuestMutex.Lock()
	ret, specificReturn := fake.listProcessesInGuestReturnsOnCall[len(fake.listProcessesInGuestArgsForCall)]
	fake.listProcessesInGuestArgsForCall = append(fake.listProcessesInGuestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("ListProcessesInGuest", []interface{}{arg1, arg2, arg3})
	fake.listProcessesInGuestMutex.Unlock()
	if fake.ListProcessesInGuestStub != nil {
		return fake.ListProcessesInGuestStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listProcessesInGuestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVmrunRunner) ListProcessesInGuestCallCount() int {
	fake.listProcessesInGuestMutex.RLock()
	defer fake.listProcessesInGuestMutex.RUnlock()
	return len(fake.listProcessesInGuestArgsForCall)
}

func (fake *FakeVmrunRunner) ListProcessesInGuestCalls(stub func(string, string, string) (string, error)) {
	fake.listProcessesInGuestMutex.Lock()
	defer fake.listProcessesInGuestMutex.Unlock()
	fake.ListProcessesInGuestStub = stub
}

func (fake *FakeVmrunRunner) ListProcessesInGuestArgsForCall(i int) (string, string, string) {
	fake.listProcessesInGuestMutex.RLock()
	defer fake.listProcessesInGuestMutex.RUnlock()
	argsForCall := fake.listProcessesInGuestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVmrunRunner) ListProcessesInGuestReturns(result1 string, result2 error) {
	fake.listProcessesInGuestMutex.Lock()
	defer fake.listProcessesInGuestMutex.Unlock()
	fake.ListProcessesInGuestStub = nil
	fake.listProcessesInGuestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVmrunRunner) ListProcessesInGuestReturnsOnCall(i int, result1 string, result2 error) {
	fake.listProcessesInGuestMutex.Lock()
	defer fake.listProcessesInGuestMutex.Unlock()
	fake.ListProcessesInGuestStub = nil
	if fake.listProcessesInGuestReturnsOnCall == nil {
		fake.listProcessesInGuestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.listProcessesInGuestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVmrunRunner) RunProgramInGuest(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.runProgramInGuestMutex.Lock()
	ret, specificReturn := fake.runProgramInGuestReturnsOnCall[len(fake.runProgramInGuestArgsForCall)]
	fake.runProgramInGuestArgsForCall = append(fake.runProgramInGuestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("RunProgramInGuest", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.runProgramInGuestMutex.Unlock()
	if fake.RunProgramInGuestStub != nil {
		return fake.RunProgramInGuestStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runProgramInGuestReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) RunProgramInGuestCallCount() int {
	fake.runProgramInGuestMutex.RLock()
	defer fake.runProgramInGuestMutex.RUnlock()
	return len(fake.runProgramInGuestArgsForCall)
}

func (fake *FakeVmrunRunner) RunProgramInGuestCalls(stub func(string, string, string, string, string) error) {
	fake.runProgramInGuestMutex.Lock()
	defer fake.runProgramInGuestMutex.Unlock()
	fake.RunProgramInGuestStub = stub
}

func (fake *FakeVmrunRunner) RunProgramInGuestArgsForCall(i int) (string, string, string, string, string) {
	fake.runProgramInGuestMutex.RLock()
	defer fake.runProgramInGuestMutex.RUnlock()
	argsForCall := fake.runProgramInGuestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVmrunRunner) RunProgramInGuestReturns(result1 error) {
	fake.runProgramInGuestMutex.Lock()
	defer fake.runProgramInGuestMutex.Unlock()
	fake.RunProgramInGuestStub = nil
	fake.runProgramInGuestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) RunProgramInGuestReturnsOnCall(i int, result1 error) {
	fake.runProgramInGuestMutex.Lock()
	defer fake.runProgramInGuestMutex.Unlock()
	fake.RunProgramInGuestStub = nil
	if fake.runProgramInGuestReturnsOnCall == nil {
		fake.runProgramInGuestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runProgramInGuestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) SoftStop(arg1 string) error {
	fake.softStopMutex.Lock()
	ret, specificReturn := fake.softStopReturnsOnCall[len(fake.softStopArgsForCall)]
	fake.softStopArgsForCall = append(fake.softStopArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SoftStop", []interface{}{arg1})
	fake.softStopMutex.Unlock()
	if fake.SoftStopStub != nil {
		return fake.SoftStopStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.softStopReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) SoftStopCallCount() int {
	fake.softStopMutex.RLock()
	defer fake.softStopMutex.RUnlock()
	return len(fake.softStopArgsForCall)
}

func (fake *FakeVmrunRunner) SoftStopCalls(stub func(string) error) {
	fake.softStopMutex.Lock()
	defer fake.softStopMutex.Unlock()
	fake.SoftStopStub = stub
}

func (fake *FakeVmrunRunner) SoftStopArgsForCall(i int) string {
	fake.softStopMutex.RLock()
	defer fake.softStopMutex.RUnlock()
	argsForCall := fake.softStopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVmrunRunner) SoftStopReturns(result1 error) {
	fake.softStopMutex.Lock()
	defer fake.softStopMutex.Unlock()
	fake.SoftStopStub = nil
	fake.softStopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) SoftStopReturnsOnCall(i int, result1 error) {
	fake.softStopMutex.Lock()
	defer fake.softStopMutex.Unlock()
	fake.SoftStopStub = nil
	if fake.softStopReturnsOnCall == nil {
		fake.softStopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.softStopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) Start(arg1 string) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startReturns
	return fakeReturns.result1
}

func (fake *FakeVmrunRunner) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeVmrunRunner) StartCalls(stub func(string) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeVmrunRunner) StartArgsForCall(i int) string {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVmrunRunner) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVmrunRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloneMutex.RLock()
	defer fake.cloneMutex.RUnlock()
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	fake.copyFileFromHostToGuestMutex.RLock()
	defer fake.copyFileFromHostToGuestMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.hardStopMutex.RLock()
	defer fake.hardStopMutex.RUnlock()
	fake.isPlayerMutex.RLock()
	defer fake.isPlayerMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.listProcessesInGuestMutex.RLock()
	defer fake.listProcessesInGuestMutex.RUnlock()
	fake.runProgramInGuestMutex.RLock()
	defer fake.runProgramInGuestMutex.RUnlock()
	fake.softStopMutex.RLock()
	defer fake.softStopMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVmrunRunner) recordInvocation(key string, args []interface{}) {
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

var _ driver.VmrunRunner = new(FakeVmrunRunner)
