package backend

import (
	"dns-truth/internal/backend/local_file"
	"dns-truth/internal/backend/records"
	"os"
)

func Init() {
	backendType := os.Getenv("BACKEND_TYPE")
	if backendType == "local_file" {
		local_file.Init()
	} else {
		panic("Invalid backend type")
	}
	records.PrintRecords()
}
