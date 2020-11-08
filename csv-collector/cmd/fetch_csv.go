// Copyright (c) 2020, Cristian Dean
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	firstYear = 2014
	tseURL    = "http://agencia.tse.jus.br/estatistica/sead/odsele/consulta_cand/consulta_cand_%d.zip"
)

var fetchCSVFilesCmd = &cobra.Command{
	Use:   "fetch-all-csv",
	Short: "Fetch all CSV files from TSE",
	Run:   fetchCSVFiles,
}

func fetchCSVFiles(cmd *cobra.Command, args []string) {
	lastYear := time.Now().Year() // TODO pegar de forma dinâmica o último ano suportado
	for year := firstYear; year <= lastYear; year++ {
		if year%2 != 0 {
			continue
		}

		zipFileURL := fmt.Sprintf(tseURL, year)
		zipFilePath := fmt.Sprintf("./consulta_cand_%d.zip", year)
		destinationPath := fmt.Sprintf("./consulta_cand_%d", year)

		// Download zip file
		log.Infof("Downloading %s", zipFileURL)
		err := downloadFile(zipFilePath, zipFileURL)
		if err != nil {
			log.Fatalf("Couldn't download %s csv files ! %v", zipFileURL, err)
		}
		log.Infof("Successfully downloaded %s", zipFilePath)

		// Extract zip file
		log.Infof("Extracting %s in %s", zipFilePath, destinationPath)
		err = extractZipFile(zipFilePath, destinationPath)
		if err != nil {
			log.Fatalf("Couldn't extract %s ! %v", zipFilePath, err)
		}

		log.Infof("Successfully extracted %s in %s", zipFilePath, destinationPath)
	}

}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractZipFile(filepath string, destination string) error {
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		return err
	}
	r, err := zip.OpenReader(filepath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		r, err := f.Open()
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		defer r.Close()
		path := fmt.Sprintf("%s/%s", destination, f.Name)
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(path, 0755); err != nil {
				return err
			}
		} else {
			if err = ioutil.WriteFile(path, buf, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
