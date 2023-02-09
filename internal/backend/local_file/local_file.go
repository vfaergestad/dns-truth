package local_file

import (
	"bufio"
	"dns-truth/internal/backend/records"
	"fmt"
	"os"
	"strings"
)

func Init() {

	localFilePath := os.Getenv("LOCAL_FILE_PATH")
	file, err := os.Open(localFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist, creating new file")
			file, err = os.Create(localFilePath)
		} else if os.IsPermission(err) {
			fmt.Printf("Permission to %s denied", localFilePath)
			panic(err)
		} else {
			panic(err)
		}
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		recordLine := scanner.Text()
		data := strings.Split(recordLine, " ")
		hostname := strings.TrimSpace(data[0])
		ip := strings.TrimSpace(data[1])
		record := records.ARecord{
			Hostname: hostname,
			Ip:       ip,
		}
		records.AddRecord(record)
	}

}
