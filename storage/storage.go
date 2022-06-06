package storage

import (
	"fmt"
	"log"
	"main/data"
	"os/exec"
	"strings"
)

func EraseDiskAndCreateParts() {
	commands := []string{
		"parted -s /dev/sda mklabel msdos",
		"parted -s /dev/sda mkpart primary ext2 1MiB 256MiB",
		"parted -s /dev/sda set 1 boot on",
		"parted -s /dev/sda mkpart primary linux-swap 256MiB 512MiB",
		"parted -s /dev/sda mkpart primary ext4 512MiB 100%",
	}

	for _, cmd := range commands {
		cmds := strings.Split(cmd, " ")
		out, err := exec.Command(cmds[0], cmds[1:]...).Output()
		if err != nil {
			log.Fatal(err)
		}
		output := string(out)
		fmt.Println("> [", cmd, "] -> [", output, "]")
	}
}

func ErasePartsOnDevice(device string) (string, error) {
	command := fmt.Sprintf("parted %s", device)
	return command, nil
}

func CreatePartsOnDevice(device string, parts []data.DiskPart) []string {
	// var cmd []string

	// base_cmd := fmt.Sprintf("parted %s", device)

	// cmd = append(cmd, fmt.Sprintf("%s %s"))

	cmd := []string{
		"parted -s /dev/sda mklabel msdos",
		"parted -s /dev/sda mkpart primary ext2 1MiB 256MiB",
		"parted -s /dev/sda set 1 boot on",
		"parted -s /dev/sda mkpart primary linux-swap 256MiB 512MiB",
		"parted -s /dev/sda mkpart primary ext4 512MiB 100%",
	}
	return cmd
}

//	"parted /dev/sda mklabel {gpt,msdos}" // https://www.gnu.org/software/parted/manual/html_node/mklabel.html

//	"parted /dev/sda mkpart primary ext2 1MiB 256MiB" // https://www.gnu.org/software/parted/manual/html_node/mkpart.html
//	<part-type> is one of ‘primary’, ‘extended’ or ‘logical’, and may be specified only with ‘msdos’

//	"parted /dev/sda set 1 boot on",
//	"parted /dev/sda mkpart primary ext2 256MiB 512MiB",
//	"parted /dev/sda mkpart primary ext4 512MiB 100%",
