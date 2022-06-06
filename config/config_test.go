package config

import (
	"testing"
)

const (
	config_filename = "../configure.yaml"
)

func TestReadDefaultConfigFile(t *testing.T) {

	cfg, err := InstallConfigLoad(config_filename)
	if err != nil {
		t.Error("Can not open file: [" + config_filename + "]")
	}

	user := cfg.User.Login

	if user != "user" {
		t.Error("Wrong user name in configure.yaml (Value=[" + user + "], Expected=[user])")
	}

	device := cfg.Storage.Device

	if device != "/dev/sda" {
		t.Error("Wrong device name in configure.yaml (Value=[" + device + "], Expected=[/dev/sda])")
	}

	install_devs := cfg.Storage.Install
	for _, dev := range install_devs {
		if dev.MountToPoint == "/" && dev.Filesystem != "ext4" {
			t.Error("Wrong filesystem from / (Value=[" + dev.Filesystem + "], Expected=[ext4])")
		}
	}
}
