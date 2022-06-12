package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains	:", strings.Contains("test", "es"))
	fmt.Println("Count		:", strings.Count("test", "t"))
	fmt.Println("HasPrefix	:", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix	:", strings.HasSuffix("test", "st"))
	fmt.Println("Index		:", strings.Index("test", "e"))
	fmt.Println("Join		:", strings.Join([]string{"a", "b"}, "-")) // ["a,"b].join("-")
	fmt.Println("Repeat		:", strings.Repeat("a", 5))
	fmt.Println("Replace 1	:", strings.Replace("aaaa", "a", "b", 2))
	fmt.Println("Replace 2	:", strings.Replace("aaaa", "a", "b", -1))
	fmt.Println("Split		:", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower	:", strings.ToLower("TEST"))
	fmt.Println("ToUpper	:", strings.ToUpper("test"))
	fmt.Println("Trim		:", strings.Trim("  abc  ", " "))
	fmt.Println("TrimLeft	:", strings.TrimLeft("  abc  ", " "))
	fmt.Println("TrimRight	:", strings.TrimRight("  abc  ", " "))
	fmt.Println("TrimSpace	:", strings.TrimSpace("  abc  "))
	fmt.Println("TrimPrefix	:", strings.TrimPrefix("  abc  ", " "))
	fmt.Println("TrimSuffix	:", strings.TrimSuffix("  abc  ", " "))
	fmt.Println("Fields		:", strings.Fields("  abc  "))
}
