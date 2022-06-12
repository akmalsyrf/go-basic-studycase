package helper

func SayHello(name string) string { //diawali dengan huruf besar, bisa diakses di package lain
	return "Hello, " + name
}

func sayGoodbye(name string) string { //diawali dengan huruf kecil, tidak bisa diakses di package lain
	return "Goodbye, " + name
}
