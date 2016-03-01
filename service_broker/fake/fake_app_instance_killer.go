// This file was generated by counterfeiter
package fake

import (
	"sync"

	"github.com/xchapter7x/chaospeddler/service_broker"
)

type FakeAppInstanceKiller struct {
	KillPercentStub        func(chaospeddler.ServiceBinding, int) (map[string]int, error)
	killPercentMutex       sync.RWMutex
	killPercentArgsForCall []struct {
		arg1 chaospeddler.ServiceBinding
		arg2 int
	}
	killPercentReturns struct {
		result1 map[string]int
		result2 error
	}
}

func (fake *FakeAppInstanceKiller) KillPercent(arg1 chaospeddler.ServiceBinding, arg2 int) (map[string]int, error) {
	fake.killPercentMutex.Lock()
	fake.killPercentArgsForCall = append(fake.killPercentArgsForCall, struct {
		arg1 chaospeddler.ServiceBinding
		arg2 int
	}{arg1, arg2})
	fake.killPercentMutex.Unlock()
	if fake.KillPercentStub != nil {
		return fake.KillPercentStub(arg1, arg2)
	} else {
		return fake.killPercentReturns.result1, fake.killPercentReturns.result2
	}
}

func (fake *FakeAppInstanceKiller) KillPercentCallCount() int {
	fake.killPercentMutex.RLock()
	defer fake.killPercentMutex.RUnlock()
	return len(fake.killPercentArgsForCall)
}

func (fake *FakeAppInstanceKiller) KillPercentArgsForCall(i int) (chaospeddler.ServiceBinding, int) {
	fake.killPercentMutex.RLock()
	defer fake.killPercentMutex.RUnlock()
	return fake.killPercentArgsForCall[i].arg1, fake.killPercentArgsForCall[i].arg2
}

func (fake *FakeAppInstanceKiller) KillPercentReturns(result1 map[string]int, result2 error) {
	fake.KillPercentStub = nil
	fake.killPercentReturns = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

var _ chaospeddler.AppInstanceKiller = new(FakeAppInstanceKiller)