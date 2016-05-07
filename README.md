# wmic-xml-command
Very simple repo to run WMIC (Windows Management Instrumentation Command-line) command and parse the returned XML

## Quick Start

Install with `go get -u go-zero-boilerplate/wmic-xml-command`

```
package main

import (
    "fmt"
    "log"
    "strings"

    wmic_command "github.com/go-zero-boilerplate/wmic-xml-command"
)

func main() {
    prodID := 123 //Choose any currently running process

    columnNames := []string{
        "Caption",
        "ProcessId",
    }

    wmicArgs := []string{
        "process",
        "where",
        fmt.Sprintf("(ProcessId=%d)", prodID),
        "get",
        strings.Join(columnNames, ","),
    }

    responseXml, err := wmic_command.Run(wmicArgs)
    if err != nil {
        log.Fatal(err)
    }

    for _, res := range responseXml.Results {
        for _, prop := range res.Properties {
            fmt.Println(fmt.Sprintf("PROP: %+v", prop))
        }
    }
}
```