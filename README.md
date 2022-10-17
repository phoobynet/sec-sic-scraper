# SEC SIC Scraper

## Installation

```bash
go get github.com/phoobynet/sec-sic-scraper
```

## Usage
```go

import (
	"fmt"
	"github.com/phoobynet/sec-sic-scraper"
)

func main () {
	sics := sec_sic_scraper.Get()
	
	fmt.Printf("%+v", sics)
}

```