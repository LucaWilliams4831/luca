package stardust

import (
	"math/big"
<<<<<<< HEAD
	"github.com/sipt/panda141035/log"
=======

	"github.com/sipt/shuttle/log"
>>>>>>> ba7fba0 (built with no error)
)

func FromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

func Recover(fs ...func()) {
	if err := recover(); err != nil {
		log.Logger.Errorf("[PANIC] %v", err)
		for _, f := range fs {
			f()
		}
	}
}
