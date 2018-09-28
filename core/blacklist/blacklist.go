package blacklist

import (
	"sync"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

const send_to_lock = "0x7777777777777777777777777777777777777777"
const send_to_unlock = "0x8888888888888888888888888888888888888888"
const delay = 2

var w map[common.Address]bool

func init() {
	w = make(map[common.Address]bool)
	w[common.HexToAddress(send_to_lock)] = true
	w[common.HexToAddress(send_to_unlock)] = true
}

type Blacklist struct {
	CurHeight int64
	M         map[common.Address]int64
	sync.RWMutex
}

func (b *Blacklist) Add(addr common.Address) (bool, error) {
	b.Lock()
	defer b.Unlock()
	fmt.Println("is blocked:", &b)
	b.M[addr] = -1
	return true, nil
}

func (b *Blacklist) Remove(addr common.Address) (bool, error) {
	b.Lock()
	defer b.Unlock()
	if _, ok := b.M[addr]; ok {
		if b.M[addr] == -1 {
			b.M[addr] = b.CurHeight
		}
	}
	return true, nil
}

func (b *Blacklist) compact() {
	for k, v := range b.M {
		if v != -1 && v+delay < b.CurHeight {
			b.del(k)
		}
	}
}

func (b *Blacklist) del(k common.Address) {
	b.Lock()
	defer b.Unlock()
	delete(b.M, k)
}

func (b *Blacklist) IsBlocked(from common.Address, to *common.Address) bool {
	b.RLock()
	defer b.RUnlock()
	fmt.Println("is blocked:", &b)
	h, ok := b.M[from]
	if !ok {
		return false;
	}
	fmt.Println(h+delay <= b.CurHeight)
	return (h == -1 || h+delay >= b.CurHeight) && (to == nil || !w[*to]);
}

func IsLockTx(to common.Address) bool {
	return to == common.HexToAddress(send_to_lock)
}

func IsUnlockTx(to common.Address) bool {
	return to == common.HexToAddress(send_to_unlock)
}
