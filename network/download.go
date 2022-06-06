package network

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const gentoo_repo_url = "https://https://mirror.yandex.ru/gentoo-distfiles/releases/amd64/autobuilds/current-stage3-amd64-openrc/"

// https://bouncer.gentoo.org/fetch/root/all/releases/amd64/autobuilds/20210905T170549Z/stage3-amd64-openrc-20210905T170549Z.tar.xz

func GetLatestStage3(url string) ([]string, error) {
	// FIXME: change name from openrc to general
	re := regexp.MustCompile(`a href="stage3-amd64-openrc-(\d+)T(\d+)Z\.tar\.xz"`)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	text, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	pageContent := string(text)

	names := re.FindAll([]byte(pageContent), -1)

	if len(names) > 0 {
		var output []string

		for _, name := range names {
			output = append(output, string(name[8:len(name)-1]))
			fmt.Printf("DEBUG: %s\n", output[0])
		}

		return output, nil
	}
	Err := fmt.Errorf("not found stage name in %s", pageContent)
	return nil, Err
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 42))
	mb := float64(wc.Total) / (1024.0 * 1024.0)
	fmt.Printf("\rDownloading ... %6.2f MB complete", mb)
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	return err
}

func Unpack(filepath string) error {
	// tar xpvf stage3-*.tar.bz2 --xattrs-include='*.*' --numeric-owner
	cmd := exec.Command("7z", "x", filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("unpack error: " + err.Error())
	}
	return err
}
