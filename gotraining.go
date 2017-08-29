package main

import "fmt"

// Declare a struct type to maintain information about a user
// (name, email and age)
type example struct {
  name string
  email string
  age int8
}

func main() {
  // 1. Declare a new variable of type float32 and initialize with 3.14
  pi := float32(3.14)
  fmt.Printf("%T [%v]\n", pi, pi) 

  // 2. Create a value with "example" struct type and display
  ex := example{
    name: "Spongebob",
    email: "email@example.com",
    age: 2,
  }

  fmt.Println("Name", ex.name)
  fmt.Println("Email", ex.email)
  fmt.Println("Age", ex.age)

  // 3. Declare and initialize an anonymous struct type as above
  anon := struct {
    name string
    email string
    age int8
  }{
    name: "Patrick",
    email: "foo@example.com",
    age: 3,
  }

  fmt.Println("Name", anon.name)
  fmt.Println("Email", anon.email)
  fmt.Println("Age", anon.age)

}
