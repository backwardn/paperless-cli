package paperless

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

var documentselection int

func (p Paperless) writeFile(document Document) {
	downloadURL := p.DownloadString(document.DownloadURL)

	client := http.Client{Timeout: time.Second * 5}
	log.Debugf("downloading from: %v", downloadURL)
	req := ReturnAuthenticatedRequest(p.Username, p.Password)
	req.Method = "GET"
	urlPtr, _ := url.Parse(downloadURL)
	req.URL = urlPtr

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error downloading file: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(document.FileName)
	if err != nil {
		log.Panicln(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Panicln(err)
	}
}

// DownloadFiles pulls either a paperless document or allows selection from a list
func (p Paperless) DownloadFiles(filename string) {
	documents, err := Paperless.GetDocument(p, filename, true)
	if err != nil {
		log.Errorf("Error finding file")
		return
	}
	if len(documents) == 0 {
		log.Panicln("No matches for", filename)
		return
	}
	if len(documents) > 1 {
		fmt.Println("Multiple files found. Please select from:")
		for index, element := range documents {
			fmt.Println(strconv.Itoa(index) + ": " + element.FileName)
		}
		var selection string
		fmt.Scanln(&selection)
		documentselection, err = strconv.Atoi(selection)
		if err != nil {
			log.Error("Please pass a number back for selection")
			return
		}
	}

	document := documents[documentselection]
	p.writeFile(document)
}