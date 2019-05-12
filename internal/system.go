package contributions

import (
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/olekukonko/tablewriter"
)

// check causes the current program to exit if an error occurred.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getMaxParallelism retrieves the maximum number of concurrent workers.
func getMaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	if maxProcs < numCPU {
		return maxProcs
	} else {
		return numCPU
	}
}

// getRegexResults retrieves regular expression results with namespaces.
func getRegexResults(str string, reg *regexp.Regexp) map[string]string {
	result := make(map[string]string)
	matches := reg.FindStringSubmatch(str)

	if len(matches) > 0 {
		for i, name := range reg.SubexpNames() {
			if i > 0 && name != "" {
				result[name] = matches[i]
			}
		}
	}

	return result
}

// render prints a table with all
func render() {
	data := make([][]string, 0)

	for i, c := range contributors {
		data = append(data, []string{
			fmt.Sprint(i + 1),
			c.name,
			c.email,
			fmt.Sprint(c.commits),
			fmt.Sprint(c.Insertions()),
			fmt.Sprint(c.Deletions()),
			fmt.Sprint((c.Insertions() + c.Deletions()) / c.commits),
			c.Activity(),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Email", "Commits", "Insertions", "Deletions", "Lines/Commits", "Activity"})
	table.AppendBulk(data)
	table.Render()
}
