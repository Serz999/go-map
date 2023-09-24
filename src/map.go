package customMap

import (
	"fmt"
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

type Map struct {
    nextIndex int64
    mp map[int64]int64
}

func NewMap() (mp *Map) {
    return &Map{
        nextIndex: 0,
        mp: make(map[int64]int64),
    }
}

func (this *Map) Len() (l int64) {
    return int64(len(this.mp))
}

func (this *Map) Add(value int64) int64 {
    this.mp[this.nextIndex] = value
    this.nextIndex++
    return this.nextIndex - 1
}

func (this *Map) RemoveByIndex(id int64) {
    delete(this.mp, id)
}

func (this *Map) RemoveByValue(value int64) {
    for k, v := range this.mp {
        if v == value {
            delete(this.mp, k)    
            return
        }
    }
}

func (this *Map) GetByIndex(id int64) (int64, bool) {
    value, ok := this.mp[id]
    return value, ok
}

func (this *Map) GetByValue(value int64) (int64, bool) {

    for k, v := range this.mp {
        if v == value {
            return k, true
        }
    }

    return 0, false
}

func (this *Map) Clear() {
    this.mp = make(map[int64]int64)
    this.nextIndex = 0
}

func (this *Map) Print() {
    fmt.Printf("%#v", this.mp)
}
