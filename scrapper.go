package main

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/hemreari/goldwatcher/price"
	log "github.com/sirupsen/logrus"
)

func GetPrices() *price.Price {
	price := price.Price{}

	c := colly.NewCollector()

	c.OnHTML("div.kurlar", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			text := strings.Split(strings.TrimSpace(el.Text), "\n")
			name := strings.Trim(text[0], " ")
			priceStr := strings.Trim(text[2], " ")
			if strings.Contains(priceStr, ","); true {
				priceStr = strings.TrimSuffix(priceStr, ",")
			}
			priceInt, _ := strconv.Atoi(priceStr)
			switch name {
			case "22 Ayar Altın":
				price.Ayar22Altin = priceInt
			case "Çeyrek Ziynet":
				price.Ceyrek = priceInt
			case "Yarım Ziynet":
				price.Yarim = priceInt
			case "Tam Ziynet":
				price.Tam = priceInt
			case "Cumhuriyet":
				price.Cumhuriyet = priceInt
			case "IAB Kapanış":
				price.IabKapanis = priceInt
			default:
				log.Errorf("'%s' name not recognized.", name)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit("http://akod.org.tr")

	return &price
}
