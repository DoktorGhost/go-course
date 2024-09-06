package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title       string
	Description string
}

func main() {
	ads := []Ad{
		{
			Title:       "Куплю велосипед MeRiDa",
			Description: "Куплю велосипед meriDA в хорошем состоянии",
		},
		{
			Title:       "Продам ВаЗ 2101",
			Description: "Продам ваз 2101 в хорошем состоянии",
		},
		{
			Title:       "Продам БМВ",
			Description: "Продам бМв в хорошем состоянии",
		},
		{
			Title:       "Продам macBook pro",
			Description: "Продам macBook PRO в хорошем состоянии",
		},
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Description)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	if len(censor) == 0 {
		return ads
	}

	var result []Ad
	for key, value := range censor {
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(key))
		for _, ad := range ads {
			if re.MatchString(ad.Title) || re.MatchString(ad.Description) {
				textTitle := re.ReplaceAllString(ad.Title, value)
				textDescription := re.ReplaceAllString(ad.Description, value)
				result = append(result, Ad{textTitle, textDescription})
			}
		}
	}

	return result
}
