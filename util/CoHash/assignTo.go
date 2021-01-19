package CoHash

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

type AssignToStruct struct {
	muAssign  sync.RWMutex
	KeeperIDs []int
}

func (a *AssignToStruct) SetKeeperIDs(NewKeeperIDs []int) {
	a.muAssign.Lock()
	defer a.muAssign.Unlock()
	a.KeeperIDs = NewKeeperIDs
}

func (a *AssignToStruct) AppendKeeper(keeperId uint32) error {
	a.muAssign.Lock()
	defer a.muAssign.Unlock()
	for _,v := range a.KeeperIDs {
		if v == int(keeperId) {
			return errors.New("same keeperId")
		}
	}
	a.KeeperIDs = append(a.KeeperIDs, int(keeperId))
	sort.Ints(a.KeeperIDs)
	return nil
}

func (a *AssignToStruct) RemoveKeeper(keeperId uint32) error {
	a.muAssign.Lock()
	defer a.muAssign.Unlock()
	var wanted int = -1
	for i := 0; i < len(a.KeeperIDs); i++ {
		if a.KeeperIDs[i] == int(keeperId) {
			wanted = i
			break
		}
	}
	if wanted == -1 {
		return fmt.Errorf("Not Found In keeperIDs")
	}
	a.KeeperIDs = append(a.KeeperIDs[:wanted], a.KeeperIDs[wanted+1:]...)
	return nil
}

func (a *AssignToStruct) AssignTo(MurMur3uid uint32) uint32 {
	a.muAssign.RLock()
	a.muAssign.RUnlock()
	var (
		Ans int = 0
		L   int = 1
		R   int = len(a.KeeperIDs)
	)
	for L <= R {
		M := (L + R) / 2
		if a.KeeperIDs[M-1] <= int(MurMur3uid) {
			Ans = M
			L = M + 1
		} else {
			R = M - 1
		}
	}
	if Ans == len(a.KeeperIDs) {
		Ans = 0
	}
	return uint32(a.KeeperIDs[Ans])
}
