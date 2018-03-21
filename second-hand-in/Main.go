package main

import (
	"log"
	"fmt"
)
func makeTest( values[] string,  tree *Tree){

	myRoot := GetRoot(tree)

	println("Q: "+myRoot.Data.question + " - Node Id "+ myRoot.Value)
	println("User input: " +  values[0])
	level2 := NavigateNode(myRoot, values[0])
	println("Q: "+level2.Data.question + " Your previous answer was : " + level2.Data.answer + " - Node Id "+ level2.Value)
	println("User input: " +  values[1])
	level3 := NavigateNode(level2, values[1])
	println("Q: "+level3.Data.question  + " Your previous answer was : " + level3.Data.answer + " - Node Id "+ level3.Value )
	println("User input: " +  values[2])
	level4 := NavigateNode(level3, values[2])
	//print(level4)
	println("Q: "+level4.Data.question  + " Your previous answer was : " + level4.Data.answer + " - Node Id "+ level4.Value )
	println("User input: " +  values[3])
	level5 := NavigateNode(level4, values[3])
	//print(level4)
	println("Final answer: "+level5.Data.question  + " Your previous answer was : " + level4.Data.answer + " - Node Id "+ level4.Value )
}
func main() {

	/*Answers1 := []bool{true, true, true, true}
	println(StudentStatus(Answers1))

	Answers2 := []bool{true, true, false, true}
	println(StudentStatus(Answers2))

	Answers3 := []bool{true, true, true, false}
	println(StudentStatus(Answers3))

	Answers4 := []bool{true, false, true, true}
	println(StudentStatus(Answers4))*/






	data := []Data{
	Data{"H","true","Read textbook?"},

	Data{"D","true - ","Hand ins made in time?"},
	Data{"L","false - ","Hand ins made in time?"},

	Data{"B","true","Attend lectures?"},
	Data{"F","false","Attend lectures?"},
	Data{"J","true","Attend lectures?"},
	Data{"N","false","Attend lectures?"},

	Data{"A4","true","Make exercises?"},
	Data{"C4","false","Make exercises?"},
	Data{"E4","true","Make exercises?"},
	Data{"G4","false","Make exercises?"},
	Data{"I4","true","Make exercises?"},
	Data{"K4","false","Make exercises?"},
	Data{"M4","true","Make exercises?"},
	Data{"O4","false","Make exercises?"},


		Data{"A2","true","You should be able to pass the exam"},
		Data{"A6","false","You should be able to pass the exam"},
		Data{"C2","true","You should be able to pass the exam"},
		Data{"C6","false","You could easily fail the exam"},
		Data{"E2","true","You could easily fail the exam"},
		Data{"E6","false","You could easily fail the exam"},
		Data{"G2","true","You could easily fail the exam"},
		Data{"G6","false","You could easily fail the exam"},
		Data{"I2","true","You should be able to pass the exam"},
		Data{"I6","false","You should be able to pass the exam"},
		Data{"K2","true","You could easily fail the exam"},
		Data{"K6","false","You could easily fail the exam"},
		Data{"M2","true","You could easily fail the exam"},
		Data{"M6","false","You could easily fail the exam"},
		Data{"O2","true","You could easily fail the exam"},
		Data{"O6","false","You could easily fail the exam"},
	}



	//Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(data); i++ {
		err := tree.Insert(data[i].id, data[i])
		if err != nil {
			log.Fatal("Error inserting value '", data[i].id, "': ", err)
		}
	}

	values := []string{"yes","yes","yes","yes"}
	makeTest(values, tree)
	println("\nHand ins made in time = no")
	values = []string{"yes","no","yes","yes"}
	makeTest(values, tree)
	println("\nAttend lectures” and “Read textbook” = no and no")
	values = []string{"no","yes","no","yes"}
	makeTest(values, tree)
	println("\nAttend lectures” and “Make exercises” = no and no")
	values = []string{"yes","yes","no","no"}
	makeTest(values, tree)
	//myRoot := GetRoot(tree)
	//println("Q: "+myRoot.Data.question + " A: " + myRoot.Data.answer + "Id "+ myRoot.Value)



	//Print the sorted values.
	//fmt.Print("Sorted values: | ")
	//tree.Traverse(tree.Root, func(n *Node) { fmt.Println(n.Value, ": ", n.Data, " | ") })
	fmt.Println()



	//Find values.
	/*s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}

	fmt.Println("Found " + s + ": '" + d.answer + "'")
	//Delete a value.
	err := tree.Delete(s)

	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}

	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()
	//Special case: A single-node tree. (See Tree.Delete about why this is a special case.)
	fmt.Println("Single-node tree")
	tree = &Tree{}


	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()*/


}
