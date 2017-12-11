package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"utils"
)

type storage struct {
	path string
}

func (s *storage) load(ptr interface{}) {
	logger.Printf("[storage] try loading storage file '%s'\n", s.path)
	file, err := os.Open(s.path)
	defer file.Close()
	if os.IsNotExist(err) {
		logger.Printf("[storage] storage file '%s' does not exist, so use an empty one\n", s.path)
		return
	}
	if err != nil {
		utils.Panic(err)
	}

	if err := json.NewDecoder(file).Decode(ptr); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode storage file '%s'. You might want to remove it.\n", s.path)
		utils.Panicf("[storage] failed to decode storage file '%s': %v", s.path, err)
	}
	logger.Printf("[storage] storage file '%s' decoded successfully\n", s.path)
}

func (s *storage) dump(ptr interface{}) {
	logger.Printf("[storage] try dumping to storage file '%s'...\n", s.path)
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		utils.Panic(err)
	}

	if err := json.NewEncoder(file).Encode(ptr); err != nil {
		utils.Panic(err)
	}
	logger.Printf("[storage] storage file '%s' written successfully\n", s.path)
}
