package main

import (
	"fmt"
	"log"
	"strings"

	"net/url"

	"regexp"

	"github.com/cdipaolo/goml/linear"
)

func main() {
	urlOrginal := "http://sampiy10.gazetevatan.com/"

	u, err := url.Parse(urlOrginal)
	if err != nil {
		panic(err)
	}
	urlRest := u.Path
	urlQuery := u.RawQuery

	// fmt.Println(urlRest)
	// fmt.Println(urlQuery)
	isQuary := 0
	if len(urlQuery) > 0 {
		isQuary = 1
		// fmt.Println(isQuary)
	}

	idsReg := regexp.MustCompile(`\d{5,}`)
	idArray := idsReg.FindAllString(urlRest, -1)
	isId := 0
	if len(idArray) > 0 {
		isId = 1
		// fmt.Println(isId)
	}
	// fmt.Printf("%q\n", idArray)

	datareg1 := regexp.MustCompile(`(\d+/\d+/\d+)`)
	datareg2 := regexp.MustCompile(`(\d+-\d+-\d+)`)
	date := datareg1.FindAllString(urlRest, -1)
	date2 := datareg2.FindAllString(urlRest, -1)
	// fmt.Printf("%q\n", date)
	// fmt.Printf("%q\n", date2)

	isDate := 0
	if len(date) > 0 || len(date2) > 0 {
		isDate = 1
		// fmt.Println(isDate)
	}

	urlRestSpliteArray := strings.FieldsFunc(urlRest+urlQuery, Split)
	// fmt.Println(urlRestSpliteArray)
	NowordsAfterBase := len(urlRestSpliteArray)
	// fmt.Println(NowordsAfterBase)

	nothingAfterBase := 0
	if NowordsAfterBase == 0 {
		nothingAfterBase = 1
		// fmt.Println(nothingAfterBase)
	}
	//  how to load the model

	model := linear.Logistic{}
	if err := model.RestoreFromFile("UrlModel.json"); err != nil {
		log.Fatal(err)
	}
	//'isId', 'isDate', 'isQuary', 'NowordsAfterBase','nothingAfterBase'
	myLinkFeatures := []float64{float64(isId), float64(isDate), float64(isQuary), float64(NowordsAfterBase), float64(nothingAfterBase)}
	fmt.Println(urlOrginal)
	fmt.Println(myLinkFeatures)

	prediction, err2 := model.Predict(myLinkFeatures)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("loaded")
	fmt.Println(prediction)

}
func Split(r rune) bool {
	return r == '/' || r == '-'
}
