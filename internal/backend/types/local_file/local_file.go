package local_file

import (
	"bufio"
	"dns-truth/internal/backend/records"
	"dns-truth/internal/custom_errors"
	"errors"
	"fmt"
	"os"
	"strings"
)

var localFilePath string
var format string

func Init() error {
	err := retrieveEnv()
	if err != nil {
		return err
	}
	localFilePath = os.Getenv("LOCAL_FILE_PATH")
	file, err := os.Open(localFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist, creating new file")
			file, err = os.Create(localFilePath)
			if err != nil {
				return err
			}
		} else if os.IsPermission(err) {
			fmt.Printf("Permission to %s denied", localFilePath)
			return nil
		} else {
			return nil
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
	return nil
}

func retrieveEnv() error {
	localFilePath = os.Getenv("LOCAL_FILE_PATH")
	format = os.Getenv("FORMAT")
	err := checkFormat(format)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecords(recordsList []records.ARecord) error {
	err := os.Remove(localFilePath)
	if err != nil {
		return err
	}
	file, err := os.Create(localFilePath)
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	for _, r := range recordsList {
		err = writeRecord(file, r)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeRecord(file *os.File, record records.ARecord) error {
	err := checkFormat(format)
	if err != nil {
		return err
	}

	result := strings.ReplaceAll(format, "%h", record.Hostname)
	result = strings.ReplaceAll(result, "%i", record.Ip)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(result)
	if err != nil {
		return nil
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func checkFormat(format string) error {
	var err error
	if !strings.Contains(format, "%h") {
		err = errors.New(custom_errors.ErrNoHostnameInFormat)
		return err
	}
	if !strings.Contains(format, "%i") {
		err = errors.New(custom_errors.ErrNoIpInFormat)
		return err
	}
	return nil
}
