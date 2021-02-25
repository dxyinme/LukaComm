package CoHash

import (
	"log"
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


func TestAssignTo_2(t *testing.T) {
	ats := AssignToStruct{
		muAssign:  sync.RWMutex{},
		KeeperIDs: []int{},
	}
	test1 := &UID{Uid: "test1"}
	test2 := &UID{Uid: "test2"}

	log.Println("test1:", test1.GetHash())
	log.Println("test2:", test2.GetHash())

	err := ats.AppendKeeper(574203662)
	if err != nil {
		t.Fatal(err)
	}
	err = ats.AppendKeeper(3155393867)
	if err != nil {
		t.Fatal(err)
	}
}
