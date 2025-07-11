// cmd/ercnftmetadataindexer/main.go
package main

import (
"flag"
"log"
"os"

"ercnftmetadataindexer/internal/ercnftmetadataindexer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ercnftmetadataindexer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
