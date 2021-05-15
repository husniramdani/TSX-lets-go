package controllers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/husniramdani/lets-go/api/responses"
)

type Urlapi struct {
	Hostname     string `json:"host"`
	Logname      string `json:"logname"`
	Time         string `json:"time"`
	Method       string `json:"method"`
	Urls         string `json:"url"`
	Responsecode string `json:"response"`
	Bytes        string `json:"bytes"`
	Referer      string `json:"referer"`
}

func (server *Server) Tsv(w http.ResponseWriter, r *http.Request) {

	csvFile, err := os.Open("./log_19950801.tsv")

	if err != nil {
		fmt.Println(err)
	}

	row, err := bufio.NewReader(csvFile).ReadSlice('\n')
	if err != nil {
		fmt.Println(err)
	}
	_, err = csvFile.Seek(int64(len(row)), io.SeekStart)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t'

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Urlapi
	var allRecords []Urlapi

	for _, data := range csvData {
		oneRecord.Hostname = data[0]
		oneRecord.Logname = data[1]
		oneRecord.Time = data[2]
		oneRecord.Method = data[3]
		oneRecord.Urls = data[4]
		oneRecord.Responsecode = data[5]
		oneRecord.Bytes = data[6]
		oneRecord.Referer = data[7]
		allRecords = append(allRecords, oneRecord)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responses.JSON(w, http.StatusOK, allRecords)
}
