package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

func (s *Settings) ReadTable(table *Table) error {
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

	/*
			// code for tid retrieval

			end := 7
		    for end < len(table.Layout) && table.Layout[end] != ':' {
		        end++
		    }
		    tid := str[7:end]
		    result := table.Layout[end+1:]
	*/
	fields = append(fields, Field{"tid", "int"})

	for _, v := range strings.Split(table.Layout[4:], "/") {
		tmp := strings.Split(v, ":")
		if len(tmp) == 2 {
			fields = append(fields, Field{tmp[0], tmp[1]})
		}

	}

	table.Fields = fields

	return nil
}
