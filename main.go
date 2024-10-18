package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sqweek/dialog"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var (
	currentVersion  string
	newestVersion   string
	architecture    string
	operatingSystem string
	downloadLink    string
)

func init() {
	currentVersion = runtime.Version()
	architecture = runtime.GOARCH
	operatingSystem = runtime.GOOS
}

func main() {
	resp, err := http.Get("https://go.dev/dl/")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	span := doc.Find("h3.toggleButton span").First()
	newestVersion = span.Text()

	downloadLink = fmt.Sprintf("https://golang.org/dl/%s.%s-%s.%s", newestVersion, operatingSystem, architecture, getExtension())

	if currentVersion == newestVersion {
		fmt.Println("You are up to date")
	} else {
		if messageBox := fmt.Sprintf("New version available\nCurrent version: %s\nNewest version: %s\nOpening the link after you press OK.", currentVersion, newestVersion); showMessageBox(messageBox) {
			openBrowser(downloadLink)
		}
	}
}

func getExtension() string {
	if operatingSystem == "windows" {
		return "msi"
	} else if operatingSystem == "darwin" {
		return "pkg"
	}
	return "tar.gz"
}

func showMessageBox(message string) bool {
	dialog.Message(message).Title("golangUpdater").Info()
	return true
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	err := cmd.Start()
	if err != nil {
		return
	}
}
