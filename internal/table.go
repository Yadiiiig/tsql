package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Table struct {
	Location string
	Layout   string

	Fields []Field
}

type Field struct {
	Name string
	Type string
}

/*
# defining the structure of a table:
tstr://tid:field_name:field_type/..:..

# indices (no clue yet)
tind://


# notes:
# structure will be static, data is being stored in the correct order
# tid: required default field, acts as identifier

*/
func (s *Settings) BuildTable(db, name string, fields []Field) (Table, error) {
	t := Table{
		Fields: fields,
	}

	if err := os.Mkdir(fmt.Sprintf("%s%s/%s/", s.Location, db, name), os.ModePerm); err != nil {
		return t, err
	}

	file, err := os.Create(fmt.Sprintf("%s%s/%s/.%s", s.Location, db, name, "metadata"))
	if err != nil {
		return t, err
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%s://%s:", "tstr", "tid"))

	for _, v := range fields {
		b.WriteString(fmt.Sprintf("%s:%s/", v.Name, v.Type))
	}

	t.Layout = b.String()

	err = ioutil.WriteFile(file.Name(), b.Bytes(), 0644)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (s *Settings) ReadTable(name string, table *Table) error {
	file, err := os.Open(fmt.Sprintf("%s/.%s", table.Location, "metadata"))
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch line[0:4] {
		case "tstr":
			table.Layout = line[7:]
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fields := []Field{}
	fields = append(fields, Field{"tid", "int"})

	for _, v := range strings.Split(table.Layout[4:], "/") {
		tmp := strings.Split(v, ":")
		if len(tmp) == 2 {
			fields = append(fields, Field{tmp[0], tmp[1]})
		}

	}

	table.Fields = fields

	dataFiles := []string{}
	err = filepath.Walk(table.Location, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == ".metadata" || info.IsDir() {
			return nil
		}

		dataFiles = append(dataFiles, path)

		return nil
	})
	if err != nil {
		return err
	}

	err = ReadTableData(name, table, dataFiles)
	if err != nil {
		return err
	}

	return nil
}

func ReadTableData(name string, table *Table, paths []string) error {
	fieldIndices := make(map[int]string)
	for k, v := range table.Fields {
		fieldIndices[k] = v.Name
	}

	for i := 0; i < len(paths); i++ {
		file, err := os.Open(paths[i])
		if err != nil {
			return err
		}

		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("here")

			return err
		}

		err = table.Unmarshal(&line, fieldIndices)
		if err != nil {
			return err
		}
	}

	return nil
}
