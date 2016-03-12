// This file was generated by counterfeiter
package fakes

import (
	"database/sql"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/xchapter7x/chaospeddler/service_broker"
)

type FakeGormDB struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	CreateStub        func(value interface{}) *gorm.DB
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		value interface{}
	}
	createReturns struct {
		result1 *gorm.DB
	}
	DBStub        func() *sql.DB
	dBMutex       sync.RWMutex
	dBArgsForCall []struct{}
	dBReturns     struct {
		result1 *sql.DB
	}
	FindStub        func(out interface{}, where ...interface{}) *gorm.DB
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		out   interface{}
		where []interface{}
	}
	findReturns struct {
		result1 *gorm.DB
	}
	SaveStub        func(value interface{}) *gorm.DB
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		value interface{}
	}
	saveReturns struct {
		result1 *gorm.DB
	}
	RawStub        func(sql string, values ...interface{}) *gorm.DB
	rawMutex       sync.RWMutex
	rawArgsForCall []struct {
		sql    string
		values []interface{}
	}
	rawReturns struct {
		result1 *gorm.DB
	}
	ScanStub        func(dest interface{}) *gorm.DB
	scanMutex       sync.RWMutex
	scanArgsForCall []struct {
		dest interface{}
	}
	scanReturns struct {
		result1 *gorm.DB
	}
	DeleteStub        func(value interface{}, where ...interface{}) *gorm.DB
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		value interface{}
		where []interface{}
	}
	deleteReturns struct {
		result1 *gorm.DB
	}
}

func (fake *FakeGormDB) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeGormDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeGormDB) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGormDB) Create(value interface{}) *gorm.DB {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		value interface{}
	}{value})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(value)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeGormDB) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeGormDB) CreateArgsForCall(i int) interface{} {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].value
}

func (fake *FakeGormDB) CreateReturns(result1 *gorm.DB) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeGormDB) DB() *sql.DB {
	fake.dBMutex.Lock()
	fake.dBArgsForCall = append(fake.dBArgsForCall, struct{}{})
	fake.dBMutex.Unlock()
	if fake.DBStub != nil {
		return fake.DBStub()
	} else {
		return fake.dBReturns.result1
	}
}

func (fake *FakeGormDB) DBCallCount() int {
	fake.dBMutex.RLock()
	defer fake.dBMutex.RUnlock()
	return len(fake.dBArgsForCall)
}

func (fake *FakeGormDB) DBReturns(result1 *sql.DB) {
	fake.DBStub = nil
	fake.dBReturns = struct {
		result1 *sql.DB
	}{result1}
}

func (fake *FakeGormDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		out   interface{}
		where []interface{}
	}{out, where})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(out, where...)
	} else {
		return fake.findReturns.result1
	}
}

func (fake *FakeGormDB) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeGormDB) FindArgsForCall(i int) (interface{}, []interface{}) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].out, fake.findArgsForCall[i].where
}

func (fake *FakeGormDB) FindReturns(result1 *gorm.DB) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeGormDB) Save(value interface{}) *gorm.DB {
	fake.saveMutex.Lock()
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		value interface{}
	}{value})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(value)
	} else {
		return fake.saveReturns.result1
	}
}

func (fake *FakeGormDB) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeGormDB) SaveArgsForCall(i int) interface{} {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.saveArgsForCall[i].value
}

func (fake *FakeGormDB) SaveReturns(result1 *gorm.DB) {
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeGormDB) Raw(sql string, values ...interface{}) *gorm.DB {
	fake.rawMutex.Lock()
	fake.rawArgsForCall = append(fake.rawArgsForCall, struct {
		sql    string
		values []interface{}
	}{sql, values})
	fake.rawMutex.Unlock()
	if fake.RawStub != nil {
		return fake.RawStub(sql, values...)
	} else {
		return fake.rawReturns.result1
	}
}

func (fake *FakeGormDB) RawCallCount() int {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return len(fake.rawArgsForCall)
}

func (fake *FakeGormDB) RawArgsForCall(i int) (string, []interface{}) {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return fake.rawArgsForCall[i].sql, fake.rawArgsForCall[i].values
}

func (fake *FakeGormDB) RawReturns(result1 *gorm.DB) {
	fake.RawStub = nil
	fake.rawReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeGormDB) Scan(dest interface{}) *gorm.DB {
	fake.scanMutex.Lock()
	fake.scanArgsForCall = append(fake.scanArgsForCall, struct {
		dest interface{}
	}{dest})
	fake.scanMutex.Unlock()
	if fake.ScanStub != nil {
		return fake.ScanStub(dest)
	} else {
		return fake.scanReturns.result1
	}
}

func (fake *FakeGormDB) ScanCallCount() int {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return len(fake.scanArgsForCall)
}

func (fake *FakeGormDB) ScanArgsForCall(i int) interface{} {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return fake.scanArgsForCall[i].dest
}

func (fake *FakeGormDB) ScanReturns(result1 *gorm.DB) {
	fake.ScanStub = nil
	fake.scanReturns = struct {
		result1 *gorm.DB
	}{result1}
}

func (fake *FakeGormDB) Delete(value interface{}, where ...interface{}) *gorm.DB {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		value interface{}
		where []interface{}
	}{value, where})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(value, where...)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeGormDB) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeGormDB) DeleteArgsForCall(i int) (interface{}, []interface{}) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].value, fake.deleteArgsForCall[i].where
}

func (fake *FakeGormDB) DeleteReturns(result1 *gorm.DB) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *gorm.DB
	}{result1}
}

var _ chaospeddler.GormDB = new(FakeGormDB)
