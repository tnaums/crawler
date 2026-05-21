package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0, len(pages))

	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pdSlice := make([]PageData, 0, len(keys))

	for _, k := range keys {
		//		fmt.Println(k, pages[k])
		pdSlice = append(pdSlice, pages[k])
	}

	pdMarshalled, err := json.MarshalIndent(pdSlice, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling PageData into json")
	}
	os.WriteFile(filename, pdMarshalled, 0644)

	return nil
}
