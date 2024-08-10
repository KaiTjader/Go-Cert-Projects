package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

		"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
//  Display()
// }

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}
	fmt.Println("saving the node succeeded!")
	return nil
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	todoText := getUserInput("Todo text:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func printSomething(value any) {
	typedValue, ok := value.(int)
	if ok {
		fmt.Println("typedValue: ", typedValue)
	}
	// switch value.(type) {
	// case int:
	//  //...
	//  fmt.Print("int: ")
	// case float64:
	//  fmt.Print("float64: ")
	// default:
	//  fmt.Print("other: ");
	// }
	// fmt.Println(value)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

// func outputData(data saver) {
//  data.Display()
//  saveData(data)
// }

// func add(a, b interface{}) interface{} {
//  aInt, aIsInt := a.(int)
//  bInt, bIsInt := b.(int)

//  if aIsInt && bIsInt {
//      return aInt + bInt
//  }

//  aFloat, aIsFloat := a.(float64)
//  bFloat, bIsFloat := b.(float64)

//  if aIsFloat && bIsFloat {
//      return bFloat + aFloat
//  }
// }

func add[T int | float64 | string](a, b T) T {
	return a + b
}
