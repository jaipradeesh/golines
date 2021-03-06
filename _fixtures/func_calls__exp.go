package fixtures

import "fmt"

func testFunc() {
	fmt.Printf(
		"This is a really long line that can be broken up twice %s %s",
		fmt.Sprintf(
			"This is a really long sub-line that should be broken up more because %s %s",
			argument1,
			argument2,
		),
		fmt.Sprintf("A short one %d", 3),
	)
	fmt.Printf(
		"A short prefix",
		"This is a really long line that can be broken up twice %s %s",
		fmt.Sprintf(
			"This is a really long sub-line that should be broken up more because %s %s",
			argument1,
			argument2,
		),
		fmt.Sprintf("A short one %d", 3),
	)
	fmt.Print(
		"This is a function with a really long single argument. We want to see if it's properly split",
	)
}
