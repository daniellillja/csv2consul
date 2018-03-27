package main

import (
	"encoding/csv"
	"io"
	"log"
)

// KvpEnvironmentReader is the interface that reads model objects from file.
type KvpEnvironmentReader interface {
	KvpEnvironmentRead() ([]EnvironmentKvp, error)
}

// NewKvpEnvironmentReader returns a reader that reads from r.
func NewKvpEnvironmentReader(r io.Reader) KvpEnvironmentReader {
	return modelReader{
		strReader: r,
	}
}

type modelReader struct {
	strReader io.Reader
}

// Load loads models from CSV records.
func (m modelReader) KvpEnvironmentRead() ([]EnvironmentKvp, error) {
	log.Println("Loading records from CSV")
	csvReader := csv.NewReader(m.strReader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	header := records[0]

	// Load defined environments.
	log.Println("Loading defined environments")
	envMap := make(map[int]string)
	for i := 1; i < len(header); i++ {
		log.Printf("Found environment: %s", header[i])
		envMap[i-1] = header[i]
	}

	// Populate model.
	var models []EnvironmentKvp
	log.Println("Populating model")
	for i := 1; i < len(records); i++ {
		record := records[i]
		for e := 0; e < len(envMap); e++ {
			key := record[0]
			env := envMap[e]
			value := record[e+1]
			log.Printf("Found record: key=%s, environment=%s, value=%s", key, env, value)
			model := EnvironmentKvp{
				Key:         record[0],
				Environment: envMap[e],
				Value:       record[e+1],
			}
			models = append(models, model)
		}
	}

	return models, nil
}
