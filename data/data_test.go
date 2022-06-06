package data

import (
	"math"
	"testing"
)

func TestGlobalConfig(t *testing.T) {
	gc := NewGlobalConfig()
	if gc == nil {
		t.Error("Can not create GlobalConfig")
	}

	if gc.CpuCount != 12 { // FIXME: only for one(my) machine
		t.Error("Wrong init CPU count")
	}

	memory := int(math.Round(gc.RamCount.inMBytes() / 1000))
	if int32(memory) != 24 { // FIXME: only for one(my) machine
		t.Error("Wrong init RAM count. Current value: ", memory, " GB")
	}
}
