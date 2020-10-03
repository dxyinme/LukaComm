package CoHash

import (
	"fmt"
	"sort"
	"sync"
)

var (
	muAssign sync.RWMutex
	keeperIDs []int
)


func AppendKeeper(keeperId uint32) {
	muAssign.Lock()
	defer muAssign.Unlock()
	keeperIDs = append(keeperIDs, int(keeperId))
	sort.Ints(keeperIDs)
}

func RemoveKeeper(keeperId uint32) error {
	muAssign.Lock()
	defer muAssign.Unlock()
	var wanted int = -1
	for i := 0 ; i < len(keeperIDs); i ++ {
		if keeperIDs[i] == int(keeperId) {
			wanted = i
			break
		}
	}
	if wanted == -1 {
		return fmt.Errorf("Not Found In keeperIDs")
	}
	keeperIDs = append(keeperIDs[:wanted], keeperIDs[wanted + 1:] ...)
	return nil
}

func AssignTo(MurMur3uid uint32) uint32 {
	muAssign.RLock()
	muAssign.RUnlock()
	var (
		Ans int = 0
		L int = 1
		R int = len(keeperIDs)
	)
	for ;L <= R; {
		M := (L + R) / 2
		if keeperIDs[M - 1] <= int(MurMur3uid) {
			Ans = M
			L = M + 1
		} else {
			R = M - 1
		}
	}
	if Ans == len(keeperIDs) {
		Ans = 0
	}
	return uint32(keeperIDs[Ans])
}