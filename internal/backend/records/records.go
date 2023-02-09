package records

import "fmt"

type ARecord struct {
	Hostname string
	Ip       string
}

var records []ARecord

func Init() {
	records = make([]ARecord, 0)
}

func GetRecords() []ARecord {
	return records
}

func GetFirstRecordByHostname(hostname string) ARecord {
	for _, record := range records {
		if record.Hostname == hostname {
			return record
		}
	}
	return ARecord{}
}

func GetRecordsByHostname(hostname string) []ARecord {
	var recordsByHostname []ARecord
	for _, record := range records {
		if record.Hostname == hostname {
			recordsByHostname = append(recordsByHostname, record)
		}
	}
	return recordsByHostname
}

func GetRecordsByIp(ip string) []ARecord {
	var recordsByIp []ARecord
	for _, record := range records {
		if record.Ip == ip {
			recordsByIp = append(recordsByIp, record)
		}
	}
	return recordsByIp
}

func GetFirstRecordByIp(ip string) ARecord {
	for _, record := range records {
		if record.Ip == ip {
			return record
		}
	}
	return ARecord{}
}

func AddRecord(record ARecord) {
	records = append(records, record)
}

func PrintRecords() {
	for _, record := range records {
		fmt.Printf("%s %s\n", record.Hostname, record.Ip)
	}
}
