package src

func ExamplePrintKb() {
	PrintKb(2 * BytesPerKB)
	// Output:
	// 2048 KB
}

func ExamplePrintKb_2() {
	PrintKb(5 * BytesPerKB)
	// Output:
	// 5120 KB
}

func ExamplePrintKb_3() {
	PrintKb(BytesPerKB)
	// Output:
	// n * 1024 KB
}
