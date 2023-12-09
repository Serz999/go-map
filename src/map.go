package customMap

import (
	"errors"
	"fmt"
	"reflect"
)

type Storage interface {
    Len() int64
    Add(value int64) int64
    RemoveByIndex(id int64)
    RemoveByValue(value int64)
    GetByIndex(id int64) (int64, bool)
    GetByValue(value int64) (int64, bool)
    Clear()
    Print()
}


var ErrMismatchType = errors.New("mismatched type: the type of the provided value does not match the type of items already in the storage")

type Map struct {
    nextIndex int64
    mp map[int64]any
}

func NewMap() (mp *Map) {
    return &Map{
        nextIndex: 0,
        mp: make(map[int64]any),
    }
}

func (this *Map) Len() (l int64) {
    return int64(len(this.mp))
}

func (this *Map) Add(value any) (int64, error) {
    if this.Len() > 0 {
        firstValue, _ := this.GetByIndex(0)
        if reflect.TypeOf(firstValue) != reflect.TypeOf(value) {
            return 0, ErrMismatchType 
        }
    }

    this.mp[this.nextIndex] = value
    this.nextIndex++
    return this.nextIndex - 1, nil
}

func (this *Map) RemoveByIndex(id int64) {
    delete(this.mp, id)
}

func (this *Map) RemoveByValue(value any) {
    for k, v := range this.mp {
        if v == value {
            delete(this.mp, k)    
            return
        }
    }
}

func (this *Map) RemoveAllByValue(value any) {
    for k, v := range this.mp {
        if v == value {
            delete(this.mp, k)
        }
    }
}

func (this *Map) GetByIndex(id int64) (any, bool) {
    value, ok := this.mp[id]
    return value, ok
}

func (this *Map) GetByValue(value any) (int64, bool) {
    for k, v := range this.mp {
        if v == value {
            return k, true
        }
    }

    return 0, false
}

func (this *Map) GetAllByValue(value any) (ids []int64, ok bool) { 
    for k, v := range this.mp {
        if v == value {
            ids = append(ids, k)
        }
    }

    if len(ids) > 0 {
        ok = true
    }

    return ids, ok
}

func (this *Map) getAll() (values []any, ok bool) {
    for _, v := range this.mp {
        values = append(values, v) 
    }

    if len(values) > 0 {
        ok = true
    }

    return
}

func (this *Map) Clear() {
    this.mp = make(map[int64]any)
    this.nextIndex = 0
}

func (this *Map) Print() {
    fmt.Printf("%#v", this.mp)
}
