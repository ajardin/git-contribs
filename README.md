# git-contribs
A lightweight tool, written in GO, to display the contributors statistics of a Git project. 

## Installation
  
### Homebrew 
```bash
brew tap ajardin/formulas
brew install git-contribs
```

### Manual
1. Download the binary corresponding to your operating system from
the [releases page](https://github.com/ajardin/git-contribs/releases).
2. Compile by yourself the binary corresponding to your operating system from the source code.

## Usage
```bash
# Default settings from a project where Git has been initialized.
git-contribs

# With a specific path to a Git project (default value is ".").
git-contribs -path="${HOME}/myproject"

# From a specific date (default value is "01 Jan 2000").
git-contribs -start="01 May 2019"

# With a specific threshold (default value is 10).
git-contribs -threshold=100

# All flags can also be used at the same time.
git-contribs -path="${HOME}/myproject" -start="01 May 2019" -threshold=100
```

## Preview
```bash
git clone git@github.com:golang/go.git
git-contribs -path="go/" -threshold=1000

+------+----------------------+---------------------+---------+------------+-----------+---------------+-----------+
| RANK |         NAME         |        EMAIL        | COMMITS | INSERTIONS | DELETIONS | LINES/COMMITS | ACTIVITY  |
+------+----------------------+---------------------+---------+------------+-----------+---------------+-----------+
|    1 | Russ Cox             | rsc@golang.org      |    6399 |    1182715 |    895122 |           324 | 3932 days |
|    2 | Robert Griesemer     | gri@golang.org      |    3044 |     391169 |    261566 |           214 | 4082 days |
|    3 | Rob Pike             | r@golang.org        |    2961 |     691834 |    522117 |           409 | 4070 days |
|    4 | Brad Fitzpatrick     | bradfitz@golang.org |    2232 |     254591 |     50782 |           136 | 3032 days |
|    5 | Ian Lance Taylor     | iant@golang.org     |    1719 |      63945 |     31969 |            55 | 3890 days |
|    6 | Josh Bleecher Snyder | josharian@gmail.com |    1211 |     144696 |    139687 |           234 | 2106 days |
|    7 | Andrew Gerrand       | adg@golang.org      |    1184 |      80560 |     75804 |           132 | 2559 days |
|    8 | Austin Clements      | austin@google.com   |    1143 |      45505 |     29075 |            65 | 1659 days |
+------+----------------------+---------------------+---------+------------+-----------+---------------+-----------+
```
