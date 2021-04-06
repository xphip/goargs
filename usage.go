package goargs

import "fmt"

// TODO: doc
type UsageList struct {
	FileName       string
	Path           string
	SpacingLength  int
	CurrentUsage   string
	List           []*Usage
}

// TODO: doc
type Usage struct {
	flag string
	desc string
}

func defaultTemplate(usageList UsageList) error {

	fmt.Printf("\nUsage: %s %s\n\n", usageList.FileName, usageList.Path)

	if usageList.CurrentUsage != "" {
		fmt.Printf("%s\n", usageList.CurrentUsage)

		if len(usageList.List) > 0 {
			fmt.Printf("\n")
		}
	}

	for _, usage := range usageList.List {
		// Align: "%-d" to left, "%d" to right
		fmt.Printf(fmt.Sprintf(" %%%ds  %%s\n", usageList.SpacingLength),
			usage.flag,
			usage.desc)
	}

	return nil
}