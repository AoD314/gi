package storage

import (
	"testing"
)

const (
	device = "/dev/sda"
)

func TestErasePartsOnDevice(t *testing.T) {
	output, err := ErasePartsOnDevice(device)

	if err != nil {
		t.Error("Can not get erase command !!!")
	}

	if output != "parted /dev/sda" {
		t.Errorf("Wrong erase command: [%s]", output)
	}

}
