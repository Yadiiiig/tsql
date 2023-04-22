package internal

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Database struct {
	Tables map[string]Table
}

func (s *Settings) prepareDb() error {
	dbs := make(map[string]Database)
	tables := make(map[string]Table)

	names, err := os.ReadDir(s.Location)
	if err != nil {
		return fmt.Errorf("could not read folder: %s", err.Error())
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, v := range names {
		tables = make(map[string]Table)

		dbs[v.Name()] = Database{
			Tables: tables,
		}

		tableNm, err := os.ReadDir(fmt.Sprintf("%s/%s/", s.Location, v.Name()))
		if err != nil {
			log.Printf("Could not read table: %s, %s\n", v.Name(), err.Error())
		}

		for _, tn := range tableNm {
			wg.Add(1)
			go func(db string, name string, wg *sync.WaitGroup) {
				defer wg.Done()

				table := Table{}
				table.Location = fmt.Sprintf("%s/%s/%s/", s.Location, db, name)

				mu.Lock()
				dbs[db].Tables[name] = table
				mu.Unlock()

			}(v.Name(), tn.Name(), &wg)
		}
	}

	wg.Wait()

	s.Databases = dbs

	return nil
}
