package main

import (
	"flag"
	"fmt"
	"log"
	"tsql/internal"
)

func main() {
	var setupFlag bool

	flag.BoolVar(&setupFlag, "setup", false, "Set to true to run setup function")
	flag.Parse()

	settings, err := internal.Init("")
	if err != nil {
		log.Fatal(err)
	}

	if setupFlag {
		err = setup(&settings, "foo", "orgs")
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	for k, v := range settings.Databases {
		fmt.Println("Found database:", k)
		for k := range v.Tables {
			fmt.Printf("Table: %s\n", k)
		}
		fmt.Println()
	}
}

func setup(s *internal.Settings, db, tbName string) error {
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