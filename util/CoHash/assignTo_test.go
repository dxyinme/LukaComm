package CoHash

import (
	"sync"
	"testing"
)

func TestAssignTo(t *testing.T) {
	var (
		expected uint32
		result   uint32
	)
	ats := AssignToStruct{
		muAssign:  sync.RWMutex{},
		KeeperIDs: []int{},
	}
	ats.AppendKeeper(12)
	ats.AppendKeeper(229)
	ats.AppendKeeper(15)

	expected = 229
	result = ats.AssignTo(48)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}
	expected = 12
	result = ats.AssignTo(314)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

	ats.AppendKeeper(379)

	expected = 379
	result = ats.AssignTo(314)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

	_ = ats.RemoveKeeper(15)
	expected = 229
	result = ats.AssignTo(14)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

}
