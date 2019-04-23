package contributions

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"sync"
)

var (
	Path      string
	Start     string
	Threshold int
)

// analyze starts the analysis process.
func Analyze() {
	raw := getRawStatistics()
	extractContributors(raw)

	var wg sync.WaitGroup

	if len(contributors) > 0 {
		for _, author := range contributors {
			wg.Add(1)

			go func(author *Contributor) {
				defer wg.Done()

				loadContributorStats(author)
				loadContributorActivity(author)
			}(author)
		}
	}

	wg.Wait()

	deduplicate()
	render()
}

// getRawStatistics retrieves raw statistics from Git command line.
func getRawStatistics() []byte {
	bin := "git"
	args := []string{
		"shortlog",
		"--since=" + Start,
		"--email",
		"--no-merges",
		"--numbered",
		"--summary",
		"HEAD",
	}

	cmd := exec.Command(bin, args...)
	cmd.Dir = Path

	out, err := cmd.Output()
	check(err)

	return out
}

// extractContributors extracts contributors details from from raw statistics.
func extractContributors(raw []byte) {
	emails := make(map[string]bool)

	scanner := bufio.NewScanner(bytes.NewReader(raw))
	for scanner.Scan() {
		reg := regexp.MustCompile(`\s*(?P<commits>\d+)\s*(?P<name>.+)\s<(?P<email>.+)>`)
		results := getRegexResults(scanner.Text(), reg)

		if len(results) > 0 && emails[results["email"]] == false {
			commits, _ := strconv.Atoi(results["commits"])
			if commits < Threshold {
				break
			}

			contributors = append(contributors, &Contributor{
				name:    results["name"],
				email:   results["email"],
				commits: commits,
			})

			// Avoid to add the same contributor multiple times if the username has changed
			emails[results["email"]] = true
		}
	}
}

// loadContributorStats retrieves contributions details of the specified contributor.
func loadContributorStats(author *Contributor) {
	bin := "git"
	args := []string{
		"log",
		"--shortstat",
		"--author=" + author.email,
		"--since=" + Start,
		"--no-merges",
		"HEAD",
	}

	cmd := exec.Command(bin, args...)
	cmd.Dir = Path

	out, err := cmd.Output()
	check(err)

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		reg := regexp.MustCompile(`\s?\d+ file(s?) changed(, (?P<insertions>\d+) insertion(s)?\(\+\))?(, (?P<deletions>\d+) deletion(s)?\(-\))?`)
		results := getRegexResults(scanner.Text(), reg)

		if len(results) > 0 {
			insertions, _ := strconv.Atoi(results["insertions"])
			author.AddInsertions(insertions)

			deletions, _ := strconv.Atoi(results["deletions"])
			author.AddDeletions(deletions)
		}
	}
}

// loadContributorActivity retrieves activity range of the specified contributor.
func loadContributorActivity(author *Contributor) {
	bin := "git"
	args := []string{
		"log",
		"--pretty=format:%H",
		"--author=" + author.email,
		"--since=" + Start,
		"--no-merges",
		"--reverse",
		"HEAD",
	}

	cmd := exec.Command(bin, args...)
	cmd.Dir = Path

	out, err := cmd.Output()
	check(err)

	line := 0

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		if line == 0 {
			author.start = scanner.Text()
		} else {
			author.end = scanner.Text()
		}

		line++
	}
}

// getCommitDate retrieves the date of the specified commit.
func getCommitDate(commit string) string {
	bin := "git"
	args := []string{
		"show",
		"-s",
		"--format=%at",
		commit,
	}

	cmd := exec.Command(bin, args...)
	cmd.Dir = Path

	out, err := cmd.Output()
	check(err)

	return fmt.Sprintf("%s", out)
}

// deduplicate merges all duplicated contributors.
func deduplicate() {
	for i1, c1 := range contributors {
		for i2, c2 := range contributors {
			if i2 <= i1 {
				continue
			}

			if c2.name == c1.name {
				// Update main entry
				c1.commits += c2.commits
				c1.AddInsertions(c2.Insertions())
				c1.AddDeletions(c2.Deletions())

				// Remove duplicated entry
				contributors = append(contributors[:i2], contributors[i2+1:]...)
			}
		}
	}
}
