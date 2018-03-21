# Algirithm Hand-in
By Emmely Lundberg

For the Binary Tree implementation I have taken inspiration from this article [Binary Tree](https://appliedgo.net/bintree/).


### The implementation in summary


Go language has been used as the programming language.
The program consists of two files DecisionTree.go and Main.go.

Decision tree contains all the logic regarding the decision tree
The Main.go is used to run the program and make the tests.

The implementation of this solutions consist of a balanced Binary Tree where I have carefully inserted nodes so it would be shaped to solve the problem.

I used [this site](https://www.cs.usfca.edu/~galles/visualization/BST.html) to verify that the tree were build up as I wanted. The graphical representation can be seen in the picture below.

This is what the tree looks like. Positive answers "yes" - goes to the left. Negative answers "no" goes to the right.

[![https://gyazo.com/87f2e99e4f83c356477ba9f985c13d5a](https://i.gyazo.com/87f2e99e4f83c356477ba9f985c13d5a.png)](https://gyazo.com/87f2e99e4f83c356477ba9f985c13d5a)



The Algorithm or method I use to get to the solution

First I ask the tree for the node. And ass it to print the question that is stored in the node data.

the I use the following method that takes the answer "yes" or "no" and return the child (left or right node) that matches the answer.

Then I iterate over this method until I reached the last question.


### Implementation in details

DecisionTree.go responsibilities
- Implementation of the Binary Tree

Highlighted code:


Main.go responsibilities
- Initialize Binary Tree with key and data
- Run the tests.

Highlighted code:






```

// Special function that depending on the answer retrun the node to the left or to the right.
func NavigateNode(n *Node, b string) *Node {
	if (b == "yes") {
		return n.Left
	} else {
		return n.Right
	}
}

```



```
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



```
### Results

```
GOROOT=C:\Go #gosetup
GOPATH=C:\Users\eacl\go #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\eacl\AppData\Local\Temp\___go_build_Main_go.exe . #gosetup
"C:\Program Files\JetBrains\GoLand 2017.3.3\bin\runnerw.exe" C:\Users\eacl\AppData\Local\Temp\___go_build_Main_go.exe #gosetup

Good student
[yes yes yes yes]
Final answer: You should be able to pass the exam Your previous answer was : true - Node Id A4

Hand ins made in time = no
[yes no yes yes]
Final answer: You could easily fail the exam Your previous answer was : true - Node Id E4

Attend lectures” and “Read textbook” = no and no
[no yes no yes]
Final answer: You could easily fail the exam Your previous answer was : false - Node Id K4

Attend lectures” and “Make exercises” = no and no
[yes yes no no]
Final answer: You could easily fail the exam Your previous answer was : false - Node Id C4

``` 
### Results n details


```
GOROOT=C:\Go #gosetup
GOPATH=C:\Users\eacl\go #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\eacl\AppData\Local\Temp\___go_build_Main_go.exe . #gosetup
"C:\Program Files\JetBrains\GoLand 2017.3.3\bin\runnerw.exe" C:\Users\eacl\AppData\Local\Temp\___go_build_Main_go.exe #gosetup

Good student
[yes yes yes yes]
Q: Read textbook? - Node Id H
User input: yes
Q: Hand ins made in time? Your previous answer was : true -  - Node Id D
User input: yes
Q: Attend lectures? Your previous answer was : true - Node Id B
User input: yes
Q: Make exercises? Your previous answer was : true - Node Id A4
User input: yes
Final answer: You should be able to pass the exam Your previous answer was : true - Node Id A4

Hand ins made in time = no
[yes no yes yes]
Q: Read textbook? - Node Id H
User input: yes
Q: Hand ins made in time? Your previous answer was : true -  - Node Id D
User input: no
Q: Attend lectures? Your previous answer was : false - Node Id F
User input: yes
Q: Make exercises? Your previous answer was : true - Node Id E4
User input: yes
Final answer: You could easily fail the exam Your previous answer was : true - Node Id E4

Attend lectures” and “Read textbook” = no and no
[no yes no yes]
Q: Read textbook? - Node Id H
User input: no
Q: Hand ins made in time? Your previous answer was : false -  - Node Id L
User input: yes
Q: Attend lectures? Your previous answer was : true - Node Id J
User input: no
Q: Make exercises? Your previous answer was : false - Node Id K4
User input: yes
Final answer: You could easily fail the exam Your previous answer was : false - Node Id K4

Attend lectures” and “Make exercises” = no and no
[yes yes no no]
Q: Read textbook? - Node Id H
User input: yes
Q: Hand ins made in time? Your previous answer was : true -  - Node Id D
User input: yes
Q: Attend lectures? Your previous answer was : true - Node Id B
User input: no
Q: Make exercises? Your previous answer was : false - Node Id C4
User input: no
Final answer: You could easily fail the exam Your previous answer was : false - Node Id C4
```