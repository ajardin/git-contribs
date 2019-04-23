package contributions

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	contributors []*Contributor
)

// Contributor contains all details about a contributor.
type Contributor struct {
	name       string
	email      string
	commits    int
	insertions int
	deletions  int
	start      string
	end        string
}

// AddInsertions increases the number of insertions.
func (c *Contributor) AddInsertions(insertions int) {
	c.insertions += insertions
}

// Insertions retrieves the number of insertions.
func (c Contributor) Insertions() int {
	return c.insertions
}

// AddDeletions increases the number of deletions.
func (c *Contributor) AddDeletions(deletions int) {
	c.deletions += deletions
}

// Deletions retrieves the number of deletions.
func (c Contributor) Deletions() int {
	return c.deletions
}

// SetStart defines the date of the first commit.
func (c *Contributor) SetStart(start string) {
	c.start = start
}

// SetEnd defines the date of the last commit.
func (c *Contributor) SetEnd(end string) {
	c.end = end
}

// Activity retrieves the activity in days.
func (c Contributor) Activity() string {
	start := getDateFromTimestamp(c.start)
	end := getDateFromTimestamp(c.end)
	delta := end.Sub(start)

	return fmt.Sprintf("%d days", int64((delta.Hours()/24)+0.5))
}

// getDateFromTimestamp retrieves a date from a commit timestamp.
func getDateFromTimestamp(t string) time.Time {
	date := getCommitDate(t)
	i, err := strconv.ParseInt(strings.TrimSpace(date), 10, 64)
	check(err)

	return time.Unix(i, 0)
}
