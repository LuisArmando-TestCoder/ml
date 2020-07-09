package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func GetCSVRecords(datasetPath string) map[string][]string {
	var labels []string
	records := make(map[string][]string)
	csvfile, err := os.Open(datasetPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	for i:= 0; ; i++ {
		isFirstRecord := i == 0
		record, err := r.Read()
		isLastLine := err == io.EOF
		if isLastLine {
			break
		}
		if isFirstRecord {
			for _, label := range record {
				labels = append(labels, label)
			}
			continue
		}
		for i, label := range labels {
			records[label] = append(records[label], record[i])
		}
	}
	return records
}