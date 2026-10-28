package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// tmpl := template.New("example")

	// tmpl, err := tmpl.Parse("Welcome {{.name}}! How are you doing?\n")

	// if err != nil {
	// 	panic(err)
	// }

	// tmpl := template.Must(template.New("example").Parse("Welcome {{.name}}! How are you doing?"))

	// data := map[string]interface{} {
	// 	"name": "Jhon",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is ur name?")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	templates := map[string]string {
		"Welcome": "Welcome, {{.name}}! We're glad to see you.",
		"Notification": "{{.name}}, you have a notification: {{.notification}}",
		"Error": "Oops! An error occured: {{.errMessage}}",
	}

	parsedTemplates := make(map[string]*template.Template)

	for name,tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["Welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["Notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Println("Enter your error message: ")
			errMessage, _ := reader.ReadString('\n')
			errMessage = strings.TrimSpace(errMessage)
			tmpl = parsedTemplates["Error"]
			data = map[string]interface{}{"name": name, "errMessage": errMessage}
		case "4":
			return
		default:
			fmt.Println("Invalid Choice!")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error Occoured")
		}
	}
}