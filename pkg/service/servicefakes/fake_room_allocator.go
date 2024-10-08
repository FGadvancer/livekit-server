// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
)

type FakeRoomAllocator struct {
	CreateRoomStub        func(context.Context, *livekit.CreateRoomRequest, bool) (*livekit.Room, *livekit.RoomInternal, bool, error)
	createRoomMutex       sync.RWMutex
	createRoomArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.CreateRoomRequest
		arg3 bool
	}
	createRoomReturns struct {
		result1 *livekit.Room
		result2 *livekit.RoomInternal
		result3 bool
		result4 error
	}
	createRoomReturnsOnCall map[int]struct {
		result1 *livekit.Room
		result2 *livekit.RoomInternal
		result3 bool
		result4 error
	}
	CreateRoomEnabledStub        func() bool
	createRoomEnabledMutex       sync.RWMutex
	createRoomEnabledArgsForCall []struct {
	}
	createRoomEnabledReturns struct {
		result1 bool
	}
	createRoomEnabledReturnsOnCall map[int]struct {
		result1 bool
	}
	SelectRoomNodeStub        func(context.Context, livekit.RoomName, livekit.NodeID) error
	selectRoomNodeMutex       sync.RWMutex
	selectRoomNodeArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.NodeID
	}
	selectRoomNodeReturns struct {
		result1 error
	}
	selectRoomNodeReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateCreateRoomStub        func(context.Context, livekit.RoomName) error
	validateCreateRoomMutex       sync.RWMutex
	validateCreateRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	validateCreateRoomReturns struct {
		result1 error
	}
	validateCreateRoomReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoomAllocator) CreateRoom(arg1 context.Context, arg2 *livekit.CreateRoomRequest, arg3 bool) (*livekit.Room, *livekit.RoomInternal, bool, error) {
	fake.createRoomMutex.Lock()
	ret, specificReturn := fake.createRoomReturnsOnCall[len(fake.createRoomArgsForCall)]
	fake.createRoomArgsForCall = append(fake.createRoomArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.CreateRoomRequest
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.CreateRoomStub
	fakeReturns := fake.createRoomReturns
	fake.recordInvocation("CreateRoom", []interface{}{arg1, arg2, arg3})
	fake.createRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeRoomAllocator) CreateRoomCallCount() int {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	return len(fake.createRoomArgsForCall)
}

func (fake *FakeRoomAllocator) CreateRoomCalls(stub func(context.Context, *livekit.CreateRoomRequest, bool) (*livekit.Room, *livekit.RoomInternal, bool, error)) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = stub
}

func (fake *FakeRoomAllocator) CreateRoomArgsForCall(i int) (context.Context, *livekit.CreateRoomRequest, bool) {
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	argsForCall := fake.createRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomAllocator) CreateRoomReturns(result1 *livekit.Room, result2 *livekit.RoomInternal, result3 bool, result4 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	fake.createRoomReturns = struct {
		result1 *livekit.Room
		result2 *livekit.RoomInternal
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRoomAllocator) CreateRoomReturnsOnCall(i int, result1 *livekit.Room, result2 *livekit.RoomInternal, result3 bool, result4 error) {
	fake.createRoomMutex.Lock()
	defer fake.createRoomMutex.Unlock()
	fake.CreateRoomStub = nil
	if fake.createRoomReturnsOnCall == nil {
		fake.createRoomReturnsOnCall = make(map[int]struct {
			result1 *livekit.Room
			result2 *livekit.RoomInternal
			result3 bool
			result4 error
		})
	}
	fake.createRoomReturnsOnCall[i] = struct {
		result1 *livekit.Room
		result2 *livekit.RoomInternal
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRoomAllocator) CreateRoomEnabled() bool {
	fake.createRoomEnabledMutex.Lock()
	ret, specificReturn := fake.createRoomEnabledReturnsOnCall[len(fake.createRoomEnabledArgsForCall)]
	fake.createRoomEnabledArgsForCall = append(fake.createRoomEnabledArgsForCall, struct {
	}{})
	stub := fake.CreateRoomEnabledStub
	fakeReturns := fake.createRoomEnabledReturns
	fake.recordInvocation("CreateRoomEnabled", []interface{}{})
	fake.createRoomEnabledMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomAllocator) CreateRoomEnabledCallCount() int {
	fake.createRoomEnabledMutex.RLock()
	defer fake.createRoomEnabledMutex.RUnlock()
	return len(fake.createRoomEnabledArgsForCall)
}

func (fake *FakeRoomAllocator) CreateRoomEnabledCalls(stub func() bool) {
	fake.createRoomEnabledMutex.Lock()
	defer fake.createRoomEnabledMutex.Unlock()
	fake.CreateRoomEnabledStub = stub
}

func (fake *FakeRoomAllocator) CreateRoomEnabledReturns(result1 bool) {
	fake.createRoomEnabledMutex.Lock()
	defer fake.createRoomEnabledMutex.Unlock()
	fake.CreateRoomEnabledStub = nil
	fake.createRoomEnabledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRoomAllocator) CreateRoomEnabledReturnsOnCall(i int, result1 bool) {
	fake.createRoomEnabledMutex.Lock()
	defer fake.createRoomEnabledMutex.Unlock()
	fake.CreateRoomEnabledStub = nil
	if fake.createRoomEnabledReturnsOnCall == nil {
		fake.createRoomEnabledReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.createRoomEnabledReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRoomAllocator) SelectRoomNode(arg1 context.Context, arg2 livekit.RoomName, arg3 livekit.NodeID) error {
	fake.selectRoomNodeMutex.Lock()
	ret, specificReturn := fake.selectRoomNodeReturnsOnCall[len(fake.selectRoomNodeArgsForCall)]
	fake.selectRoomNodeArgsForCall = append(fake.selectRoomNodeArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
		arg3 livekit.NodeID
	}{arg1, arg2, arg3})
	stub := fake.SelectRoomNodeStub
	fakeReturns := fake.selectRoomNodeReturns
	fake.recordInvocation("SelectRoomNode", []interface{}{arg1, arg2, arg3})
	fake.selectRoomNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomAllocator) SelectRoomNodeCallCount() int {
	fake.selectRoomNodeMutex.RLock()
	defer fake.selectRoomNodeMutex.RUnlock()
	return len(fake.selectRoomNodeArgsForCall)
}

func (fake *FakeRoomAllocator) SelectRoomNodeCalls(stub func(context.Context, livekit.RoomName, livekit.NodeID) error) {
	fake.selectRoomNodeMutex.Lock()
	defer fake.selectRoomNodeMutex.Unlock()
	fake.SelectRoomNodeStub = stub
}

func (fake *FakeRoomAllocator) SelectRoomNodeArgsForCall(i int) (context.Context, livekit.RoomName, livekit.NodeID) {
	fake.selectRoomNodeMutex.RLock()
	defer fake.selectRoomNodeMutex.RUnlock()
	argsForCall := fake.selectRoomNodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoomAllocator) SelectRoomNodeReturns(result1 error) {
	fake.selectRoomNodeMutex.Lock()
	defer fake.selectRoomNodeMutex.Unlock()
	fake.SelectRoomNodeStub = nil
	fake.selectRoomNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomAllocator) SelectRoomNodeReturnsOnCall(i int, result1 error) {
	fake.selectRoomNodeMutex.Lock()
	defer fake.selectRoomNodeMutex.Unlock()
	fake.SelectRoomNodeStub = nil
	if fake.selectRoomNodeReturnsOnCall == nil {
		fake.selectRoomNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.selectRoomNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomAllocator) ValidateCreateRoom(arg1 context.Context, arg2 livekit.RoomName) error {
	fake.validateCreateRoomMutex.Lock()
	ret, specificReturn := fake.validateCreateRoomReturnsOnCall[len(fake.validateCreateRoomArgsForCall)]
	fake.validateCreateRoomArgsForCall = append(fake.validateCreateRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.ValidateCreateRoomStub
	fakeReturns := fake.validateCreateRoomReturns
	fake.recordInvocation("ValidateCreateRoom", []interface{}{arg1, arg2})
	fake.validateCreateRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoomAllocator) ValidateCreateRoomCallCount() int {
	fake.validateCreateRoomMutex.RLock()
	defer fake.validateCreateRoomMutex.RUnlock()
	return len(fake.validateCreateRoomArgsForCall)
}

func (fake *FakeRoomAllocator) ValidateCreateRoomCalls(stub func(context.Context, livekit.RoomName) error) {
	fake.validateCreateRoomMutex.Lock()
	defer fake.validateCreateRoomMutex.Unlock()
	fake.ValidateCreateRoomStub = stub
}

func (fake *FakeRoomAllocator) ValidateCreateRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.validateCreateRoomMutex.RLock()
	defer fake.validateCreateRoomMutex.RUnlock()
	argsForCall := fake.validateCreateRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRoomAllocator) ValidateCreateRoomReturns(result1 error) {
	fake.validateCreateRoomMutex.Lock()
	defer fake.validateCreateRoomMutex.Unlock()
	fake.ValidateCreateRoomStub = nil
	fake.validateCreateRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomAllocator) ValidateCreateRoomReturnsOnCall(i int, result1 error) {
	fake.validateCreateRoomMutex.Lock()
	defer fake.validateCreateRoomMutex.Unlock()
	fake.ValidateCreateRoomStub = nil
	if fake.validateCreateRoomReturnsOnCall == nil {
		fake.validateCreateRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateCreateRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoomAllocator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRoomMutex.RLock()
	defer fake.createRoomMutex.RUnlock()
	fake.createRoomEnabledMutex.RLock()
	defer fake.createRoomEnabledMutex.RUnlock()
	fake.selectRoomNodeMutex.RLock()
	defer fake.selectRoomNodeMutex.RUnlock()
	fake.validateCreateRoomMutex.RLock()
	defer fake.validateCreateRoomMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoomAllocator) recordInvocation(key string, args []interface{}) {
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

var _ service.RoomAllocator = new(FakeRoomAllocator)
