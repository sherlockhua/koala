package idgen

import (
	"fmt"
	"sync"

	"github.com/sony/sonyflake"
)

var (
	sf   *sonyflake.Sonyflake
	once sync.Once
)

func Init() (err error) {

	st := sonyflake.Settings{
		//StartTime: time.Now(),
	}
	once.Do(func() {
		sf, err = sonyflake.New(st)
		if err != nil {
			panic(fmt.Sprintf("init id generator failed, err:%v", err))
		}
	})
	return
}

func NewID() (id uint64, err error) {
	return sf.NextID()
}
