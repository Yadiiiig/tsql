package internal

import (
	"fmt"
	"strconv"
)

type Lines struct {
	Data map[int]map[string]Metadata
	Mtd  Metadata
}

type Metadata struct {
	Value string

	StartIndex int
	EndIndex   int

	FileId int
}

func (t *Table) Unmarshal(input *string, fieldIndices map[int]string) error {
	var entry, prevToken string
	var currentId int

	_ = prevToken
	// length := len(*input)
	// amountFields := len(t.Fields)

	foundFields := 0

	lines := Lines{}
	currentMdt := Metadata{}

	currentLine := make(map[string]Metadata)
	lines.Data = make(map[int]map[string]Metadata)

	for k, char := range *input {
		if char == ':' {
			start := k - len(entry)

			if len(currentLine) != 0 || prevToken == "/" {
				currentMdt.EndIndex = start - 1

				lines.Mtd = currentMdt
				lines.Data[currentId] = currentLine

				currentMdt = Metadata{}
				currentLine = make(map[string]Metadata)
				foundFields = 0
			}

			currentMdt.StartIndex = start

			id, err := strconv.Atoi(entry)
			if err != nil {
				return err
			}

			currentId = id
			foundFields++

			entry = ""
			prevToken = ":"

			continue
		} else if char == '/' {
			fieldMtd := Metadata{
				Value:      entry,
				StartIndex: k - len(entry),
				EndIndex:   k - 1,
			}

			currentLine[fieldIndices[foundFields]] = fieldMtd
			foundFields++

			entry = ""
			prevToken = "/"
		} else {
			entry += string(char)
		}
	}

	for k, v := range lines.Data {
		res := ""
		res += fmt.Sprintf("tid: %d | ", k)

		for name, mtd := range v {
			res += fmt.Sprintf("%s: %s | ", name, mtd.Value)
		}

		fmt.Println(res)
	}

	fmt.Println(len(lines.Data))
	return nil
}
