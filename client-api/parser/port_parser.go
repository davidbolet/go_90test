package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/davidbolet/go_90test/client-api/model"
	"github.com/davidbolet/go_90test/client-api/repository"
)

type PortParser struct {
	repo repository.Repository
}

func NewPortParser(repo repository.Repository) *PortParser {
	return &PortParser{repo: repo}
}

func (parser *PortParser) ReadAndParseFile(filename string) error {
	file, err := os.Open("./" + filename)
	if err != nil {
		log.Println("Error opening file")
		return err
	}
	err = parser.ReadAndParse(file)
	return file.Close()
}

func (parser *PortParser) ReadAndParse(reader io.Reader) error {

	dec := json.NewDecoder(reader)
	counter := 0
	// token should be the open bracket
	_, err := dec.Token()
	if err != nil {
		log.Printf("error reading starting character: %s\n", err)
		return err
	}
	for dec.More() {
		//Should be the port key
		key, err := dec.Token()
		if err != nil {
			log.Printf("error reading port key: %s\n", err.Error())
			return err
		}
		// Now we should find a Proto struct
		var port model.Port
		err = dec.Decode(&port)
		if err != nil {
			log.Printf("error decoding to port data type: %s", err)
			return err
		}
		id := fmt.Sprint(key)
		port.Key = id
		if _, err = parser.repo.SavePort(&port); err != nil {
			log.Printf("Error saving port: %s error is %s", port.Key, err.Error())
			return err
		} else {
			counter++
		}
	}
	log.Printf("Total of %d ports processed", counter)
	return nil
}
