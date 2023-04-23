package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"tsql/internal"
)

func main() {
	var setupTable, buildData bool

	flag.BoolVar(&setupTable, "setupTable", false, "Set to true to run the table setup function")
	flag.BoolVar(&buildData, "buildSample", false, "Set to true to run the table setup function")

	flag.Parse()

	settings, err := internal.Init("")
	if err != nil {
		log.Fatal(err)
	}

	if setupTable {
		err = setupTb(&settings, "foo", "orgs")
		if err != nil {
			log.Fatal(err)
		}

		return
	} else if buildData {
		err = buildSampleData(&settings, "foo", "orgs")
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	fmt.Println("--------------------")

	for name, db := range settings.Databases {
		fmt.Println("Found database:", name)
		for tName, table := range db.Tables {
			fmt.Printf("Table: %s\n", tName)
			fmt.Println("Location:", table.Location)
			for _, field := range table.Fields {
				fmt.Printf("> field_name: %s, field_type: %s\n", field.Name, field.Type)
			}

		}
		fmt.Println()
	}
}

func setupTb(s *internal.Settings, db, tbName string) error {
	table, err := s.BuildTable(db, tbName, []internal.Field{
		{"name", "string"},
		{"url", "string"},
		{"employees", "int"},
		{"enabled", "bool"},
		{"score", "float32"},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Created table: %s in database: %s\n", tbName, db)
	fmt.Println("Stored layout:", table.Layout)
	fmt.Println("Full structure:")

	for _, v := range table.Fields {
		fmt.Printf("> field_name: %s, field_type: %s\n", v.Name, v.Type)
	}

	return nil
}

func buildSampleData(s *internal.Settings, db, tbName string) error {
	rand.Seed(time.Now().UnixNano())

	min := 0.0
	max := 1.0

	for f := 0; f < 2; f++ {
		file, err := os.Create(fmt.Sprintf("%s%s/%s/%s", s.Location, db, tbName, fmt.Sprintf("%d.tsql", f)))
		if err != nil {
			return err
		}

		var b bytes.Buffer

		for i := 0; i < 2500; i++ {
			value := min + rand.Float64()*(max-min)
			b.WriteString(fmt.Sprintf("%d:foo%d/foo%d.com/%d/%t/%f/", i, i, i, i, rand.Intn(2) == 0, value))

		}

		err = ioutil.WriteFile(file.Name(), b.Bytes(), 0644)
		if err != nil {
			return err
		}

	}

	return nil
}
