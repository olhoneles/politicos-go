// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/olhoneles/politicos-go/csv-collector/collector"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/charmap"
)

var importCSVFilesCmd = &cobra.Command{
	Use:   "import-csv",
	Short: "Imports the data from all CSV files",
	Run: func(cmd *cobra.Command, args []string) {
		years := map[int]politicos.Candidature{
			2014: &politicos.Cand2014{},
			2016: &politicos.Cand2016{},
			2018: &politicos.Cand2018{},
			2020: &politicos.Cand2020{},
		}

		for year, data := range years {
			base_dir := fmt.Sprintf("./consulta_cand_%d", year)

			files, err := collector.WalkMatch(base_dir, "*.csv")
			if err != nil {
				log.Fatalf("Couldn't read files! %v", err)
			}

			for _, file := range files {
				log.Printf("Processing... %s", file)
				err := processCSVFile(file, data)
				if err != nil {
					log.Fatalf("Couldn't process CSV file! %v", err)
				}
			}
		}
	},
}

// FIXME
func deepCopy(v interface{}) (interface{}, error) {
	log.Debug("[Collector] deepCopy")

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return nil, err
	}
	return vptr.Elem().Interface(), err
}

func processCSVFile(filename string, cand politicos.Candidature) error {
	log.Debug("[Collector] processCSVFile")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	csvfile, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer csvfile.Close()

	r := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(csvfile))
	r.Comma = ';'

	// FIXME: skip CSV header
	_, err = r.Read()
	if err != nil {
		log.Println(err)
	}

	cands := []interface{}{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			continue
		}

		cand.New(record)

		// FIXME
		newCand, err := deepCopy(cand)
		if err != nil {
			continue
		}

		cands = append(cands, newCand)

		if len(cands) == 1000 {
			_, err := dbInstance.InsertMany(&politicos.Candidatures{}, cands)
			if err != nil {
				log.Fatalf("Couldn't insert data: %v", err)
			}
			cands = []interface{}{}
		}
	}

	if len(cands) > 0 {
		_, err = dbInstance.InsertMany(&politicos.Candidatures{}, cands)
		if err != nil {
			log.Fatalf("Couldn't insert data: %v", err)
		}
	}

	return nil
}
