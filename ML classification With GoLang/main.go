package main

import (
	"fmt"
	"log"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func main() {
	xtrain, ytrain, err := base.LoadDataFromCSV("InputsGoLang.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%T \n", xtrain)
	fmt.Println("%T \n", ytrain)

	// xtest, ytest, err := base.LoadDataFromCSV("test.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xtrain, ytrain)

	err = model.Learn()
	if err != nil {
		fmt.Println("error learning")
		log.Fatal(err)
	}

	fmt.Println("trained")

	s1 := []float64{1, 0, 0, 4, 0}

	mypred, _ := model.Predict(s1)
	err = model.Learn()
	if err != nil {
		fmt.Println("error predictng")
		log.Fatal(err)
	}
	fmt.Println("test")
	fmt.Println(mypred)
	fmt.Println("save model")
	model.PersistToFile("UrlModel.json")

	// /// how to load the model
	// // model := linear.Logistic{}
	// // if err := model.RestoreFromFile("firstMlModel.json"); err != nil {
	// // 	log.Fatal(err)
	// // }
	// // s1 := []float64{32, 1, 46.3, 41.3, 17.5, 17.8, 8.5, 7.01, 4.79, 70, 16.9, 74.5}
	// // prediction, err2 := model.Predict(s1)
	// // if err2 != nil {
	// // 	log.Fatal(err2)
	// // }
	// // fmt.Println("loaded")
	// // fmt.Println(prediction)

}
