package goargs

import "fmt"

type UsageList struct {
	FileName       string
	Path           string
	SpacingLength  int
	StartSpacing   string
	BetweenSpacing string
	List           []*Usage
}

type Usage struct {
	flag string
	desc string
}

func defaultTemplate(usageList UsageList) error {

	fmt.Printf("Usage: %s %s\n\n", usageList.FileName, usageList.Path)
	var row = fmt.Sprintf("%s%%%ds%s%%s\n", usageList.StartSpacing, usageList.SpacingLength, usageList.BetweenSpacing)
	for _, usage := range usageList.List {
		fmt.Printf(row, usage.flag, usage.desc)
	}

	return nil
}