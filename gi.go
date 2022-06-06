package main

import (
	"flag"
	"fmt"
	"main/config"
	"main/network"
	"main/storage"
)

func main() {

	var installConfigFile = flag.String("c", "configure.yaml", "Path to install config file (default: configure.yaml)")
	flag.Parse()

	cfg, err := config.InstallConfigLoad(*installConfigFile)

	if err != nil {
		fmt.Errorf("Error: ", err)
		return
	}

	fmt.Println(cfg)

	storage.EraseDiskAndCreateParts()

	//fmt.Println(cfg.Storage.Device)
	// fmt.Println(cfg.Install.TimeZone)
	//fmt.Println(cfg)

	//gcfg := InitGlobalConfig()
	//ConfigFile(".bashrc", gcfg)

	// isOk := getConnectionStatus()
	// if isOk != true {
	// 	return
	//}

	//////////////////////////////////////////////////////////////

	url := "https://mirror.yandex.ru/gentoo-distfiles/releases/amd64/autobuilds/current-stage3-amd64-openrc/"
	names, err := network.GetLatestStage3(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("names:")
	for _, name := range names {
		fmt.Println(name)
	}

	newURL := url + names[0]

	outputFileName := "latest-stage3.tar.xz"
	if !FileExists(outputFileName) {
		err = network.DownloadFile(outputFileName, newURL)
		if err != nil {
			panic(err)
		}
		fmt.Println("\ndownloaded succesfully")
	} else {
		fmt.Println("\nalready downloaded")
	}

	//////////////////////////////////////////////////////////////

	// // checking file on disk
	// if !fileExists(name) {
	// 	newURL := "https://gentoo.osuosl.org/releases/amd64/autobuilds/current-stage3-amd64/" + name
	// 	err := downloadFile(name, newURL)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("\ndownloaded: " + name)
	// } else {
	// 	fmt.Println("already downloaded: " + name)
	// }

	// testName := "test.tar.xz"
	// unpack(testName)

	// fmt.Println("unpack: " + name)
}
