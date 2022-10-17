package sec_sic_scraper

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

type SIC struct {
	Code          int    `json:"code"`
	IndustryTitle string `json:"industryTitle"`
	Office        string `json:"office"`
}

var SourceURL = "https://www.sec.gov/corpfin/division-of-corporation-finance-standard-industrial-classification-sic-code-list"

func Get(sourceURL *string) ([]SIC, error) {
	if sourceURL == nil {
		sourceURL = &SourceURL
	}

	sics := make([]SIC, 0)

	c := colly.NewCollector()

	var err error

	c.OnHTML("table.sic", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if err != nil {
				return
			}

			if !strings.Contains(el.Text, "SIC Code") {
				sic := SIC{}

				el.ForEach("td", func(i int, td *colly.HTMLElement) {
					switch i {
					case 0:
						sicCode, sicCodeErr := strconv.Atoi(td.Text)
						if sicCodeErr != nil {
							err = sicCodeErr
							return
						}
						sic.Code = sicCode
					case 1:
						sic.Office = td.Text
					case 2:
						sic.IndustryTitle = td.Text
					}
				})

				sics = append(sics, sic)
			}
		})
	})

	if err != nil {
		return nil, err
	}

	visitErr := c.Visit(*sourceURL)
	if visitErr != nil {
		return nil, err
	}

	return sics, nil
}
