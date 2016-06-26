# sen

![](icon.png)

Separated API automation testing made easy

# Usage

```

$ sen --help

NAME:
   sen - A small cli written in Go to help automation test

USAGE:
   sen [global options] command [command options] [arguments...]

VERSION:
   1.0

AUTHOR(S):
    <dev@dwarvesf.com>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --Input file CSV "input"	Input file CSV with given format
   --help, -h			show help
   --version, -v		print the version

$ sen testcases.csv
Running test case: Get new project
Test passed
Running test case: Login
Test passed
Running test case: Create Top up
Test passed
---------------------------------------
Number test failed: 0
Number test passed: 3

```
For format of csv file, please check testcases.csv


