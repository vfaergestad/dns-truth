package backend

import (
	"dns-truth/internal/backend/records"
	"dns-truth/internal/backend/types/local_file"
	"fmt"
	"os"
)

func Init() {
	backendType := os.Getenv("BACKEND_TYPE")
	if backendType == "local_file" {
		err := local_file.Init()
		if err != nil {
			panic(err)
		}
	} else {
		panic("Invalid backend type")
	}
	records.PrintRecords()
	fmt.Println("Testing change of records")
	recordList := records.GetRecords()
	err := local_file.UpdateRecords(recordList)
	if err != nil {
		panic(err)
	}
}
