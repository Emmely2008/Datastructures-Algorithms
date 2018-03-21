package main

import (
	"log"
	"fmt"
)

func makeTest(answers [] string, tree *Tree, verbose bool) {
	fmt.Printf("%v", answers)
	fmt.Println("")
	myRoot := GetRoot(tree)
	if (verbose) {
		fmt.Println("Q: " + myRoot.Data.question + " - Node Id " + myRoot.Value)
		fmt.Println("User input: " + answers[0])
	}
	level2 := NavigateNode(myRoot, answers[0])
	if (verbose) {
		fmt.Println("Q: " + level2.Data.question + " Your previous answer was : " + level2.Data.answer + " - Node Id " + level2.Value)
		fmt.Println("User input: " + answers[1])
	}
	level3 := NavigateNode(level2, answers[1])
	if (verbose) {
		fmt.Println("Q: " + level3.Data.question + " Your previous answer was : " + level3.Data.answer + " - Node Id " + level3.Value)
		fmt.Println("User input: " + answers[2])
	}
	level4 := NavigateNode(level3, answers[2])
	//print(level4)
	if (verbose) {
		fmt.Println("Q: " + level4.Data.question + " Your previous answer was : " + level4.Data.answer + " - Node Id " + level4.Value)
		fmt.Println("User input: " + answers[3])
	}
	level5 := NavigateNode(level4, answers[3])
	//print(level4)
	fmt.Println("Final answer: " + level5.Data.question + " Your previous answer was : " + level4.Data.answer + " - Node Id " + level4.Value)
}

func main() {

	// Data for tree
	data := []Data{
		Data{"H", "true", "Read textbook?"},

		Data{"D", "true - ", "Hand ins made in time?"},
		Data{"L", "false - ", "Hand ins made in time?"},

		Data{"B", "true", "Attend lectures?"},
		Data{"F", "false", "Attend lectures?"},
		Data{"J", "true", "Attend lectures?"},
		Data{"N", "false", "Attend lectures?"},

		Data{"A4", "true", "Make exercises?"},
		Data{"C4", "false", "Make exercises?"},
		Data{"E4", "true", "Make exercises?"},
		Data{"G4", "false", "Make exercises?"},
		Data{"I4", "true", "Make exercises?"},
		Data{"K4", "false", "Make exercises?"},
		Data{"M4", "true", "Make exercises?"},
		Data{"O4", "false", "Make exercises?"},


		Data{"A2", "true", "You should be able to pass the exam"},
		Data{"A6", "false", "You should be able to pass the exam"},
		Data{"C2", "true", "You should be able to pass the exam"},
		Data{"C6", "false", "You could easily fail the exam"},
		Data{"E2", "true", "You could easily fail the exam"},
		Data{"E6", "false", "You could easily fail the exam"},
		Data{"G2", "true", "You could easily fail the exam"},
		Data{"G6", "false", "You could easily fail the exam"},
		Data{"I2", "true", "You should be able to pass the exam"},
		Data{"I6", "false", "You should be able to pass the exam"},
		Data{"K2", "true", "You could easily fail the exam"},
		Data{"K6", "false", "You could easily fail the exam"},
		Data{"M2", "true", "You could easily fail the exam"},
		Data{"M6", "false", "You could easily fail the exam"},
		Data{"O2", "true", "You could easily fail the exam"},
		Data{"O6", "false", "You could easily fail the exam"},
	}

	//Create a tree and fill it from the data values.
	tree := &Tree{}
	for i := 0; i < len(data); i++ {
		err := tree.Insert(data[i].id, data[i])
		if err != nil {
			log.Fatal("Error inserting value '", data[i].id, "': ", err)
		}
	}


	// Student tests
	verbose := true;
	fmt.Println("\nGood student")
	values := []string{"yes", "yes", "yes", "yes"}
	makeTest(values, tree, verbose)

	fmt.Println("\nHand ins made in time = no")
	values = []string{"yes", "no", "yes", "yes"}
	makeTest(values, tree, verbose)

	fmt.Println("\nAttend lectures” and “Read textbook” = no and no")
	values = []string{"no", "yes", "no", "yes"}
	makeTest(values, tree, verbose)

	fmt.Println("\nAttend lectures” and “Make exercises” = no and no")
	values = []string{"yes", "yes", "no", "no"}
	makeTest(values, tree, verbose)

}
