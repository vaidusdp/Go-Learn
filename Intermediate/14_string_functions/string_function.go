package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go"
	fmt.Println(len(str))

	fmt.Println(str + " " + "World!")

	fmt.Println(str[0]) //Returns ASCII Values
	fmt.Printf("%c\n",str[0]) //Returns ASCII Values
	fmt.Println(str[0:3])


	num := 18
	str2 := strconv.Itoa(num) // Int to decimal string
	fmt.Println(str2)

	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ", ")
	fmt.Println(parts)

	countries := []string{"Germany", "India"}
	joined := strings.Join(countries, " ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))

	replaced := strings.Replace(str, "Go", "World", 1)
	replaced = strings.ReplaceAll(str, "Go", "World")
	fmt.Println(replaced)

	str3 := "   Vaidus   "
	str3 = strings.TrimSpace(str3)
	fmt.Println(str3)

	fmt.Println(strings.ToLower(str3))
	fmt.Println(strings.ToUpper(str3))

	fmt.Println(strings.Repeat("foo ", 3))

	fmt.Println(strings.Count("Hello", "l"))

	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str4 := "Hello, 123 Go! 11"
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(str4, -1)
	fmt.Println(nums)

	str5 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str5))

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World!")

	res := builder.String()
	fmt.Println(res)

	builder.WriteRune('!')
	res = builder.String()
	fmt.Println(res)

	builder.Reset()
	builder.WriteString("Fresh")
	res1 := builder.String()
	fmt.Println(res1)

}