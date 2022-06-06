package data

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

type StorageSection struct {
	Device  string     `yaml:"device"`
	Scheme  string     `yaml:"scheme"`
	Install []DiskPart `yaml:"install"`
	Volumes []Volume   `yaml:"volumes"`
}

type DiskPart struct {
	MountToPoint string `yaml:"mount"`
	Size         string `yaml:"size"`
	Filesystem   string `yaml:"filesystem"`
}

type InstallSection struct {
	RootMount  string   `yaml:"root_mount"`
	InitSystem string   `yaml:"init_system"`
	TimeZone   string   `yaml:"time_zone"`
	Locale     []string `yaml:"locale"`
}

type Volume struct {
	MountToPoint string `yaml:"mount"`
	Device       string `yaml:"device"`
	Filesystem   string `yaml:"filesystem"`
	Options      string `yaml:"opts"`
}

type UserSection struct {
	Login        string `yaml:"login"`
	Password     string `yaml:"password"`
	RootPassword string `yaml:"root_password"`
}

type InstallConfig struct {
	Storage  StorageSection `yaml:"storage"`
	Install  InstallSection `yaml:"install"`
	User     UserSection    `yaml:"user"`
	Packages []string       `yaml:"packages"`
}

type Size struct {
	value uint64
}

func (s Size) inBytes() uint64 {
	return s.value
}

func (s Size) inKBytes() float64 {
	return float64(s.value) / 1024
}

func (s Size) inMBytes() float64 {
	return float64(s.value) / (1024 * 1024)
}

func (s Size) inGBytes() float64 {
	return float64(s.value) / (1024 * 1024 * 1024)
}

type GlobalConfig struct {
	CpuCount int
	RamCount Size
}

func NewGlobalConfig() *GlobalConfig {

	out, err := exec.Command("cat", "/proc/meminfo").Output()
	if err != nil {
		log.Fatal(err)
	}

	output := string(out)

	re, _ := regexp.Compile(`MemTotal:\ +(\d+) kB`)
	res := re.FindAllStringSubmatch(output, -1)

	s := res[0][1]
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("Can not convert [" + s + "] to int")
	}
	var memory_total uint64 = uint64(i) * 1000

	return &GlobalConfig{CpuCount: runtime.NumCPU(), RamCount: Size{memory_total}}
}
