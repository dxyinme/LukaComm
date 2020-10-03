package CoHash

import "testing"


func TestAssignTo(t *testing.T) {
	var (
		expected 	uint32
		result		uint32
	)
	AppendKeeper(12)
	AppendKeeper(15)
	AppendKeeper(229)

	expected = 229
	result = AssignTo(48)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}
	expected = 12
	result = AssignTo(314)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

	AppendKeeper(379)

	expected = 379
	result = AssignTo(314)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

	_ = RemoveKeeper(15)
	expected = 229
	result = AssignTo(14)
	if result != expected {
		t.Errorf("expected %v but %v", expected, result)
	}

}
