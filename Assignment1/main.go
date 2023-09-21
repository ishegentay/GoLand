package main

import (
	"awesomeProject1/Assistant"
	"awesomeProject1/Bookkeeper"
	"awesomeProject1/Developer"
	"awesomeProject1/Engineer"
	"awesomeProject1/Manager"
	"fmt"
)

func main() {
	manager := Manager.NewManager("Manager", 260000, "ToleBi, 32")

	developer := Developer.NewDeveloper("Developer", 370000, "AbylayKhan, 43")

	bookkeeper := Bookkeeper.NewBookkeeper("Bookkeeper", 280000, "Panfilova, 54")

	engineer := Engineer.NewEngineer("Engineer", 300000, "KazbekBi, 65")

	assistant := Assistant.NewAssistant("Assistant", 170000, "Islam Karimova, 76")

	fmt.Println("Manager Position:", manager.GetPosition())
	manager.SetPosition("Senior Manager")
	fmt.Println("Manager Position (Updated):", manager.GetPosition())

	fmt.Println("Developer Salary:", developer.GetSalary())
	developer.SetSalary(385000)
	fmt.Println("Developer Salary (Updated):", developer.GetSalary())

	fmt.Println("Bookkeeper Position:", bookkeeper.GetPosition())

	fmt.Println("Engineer Position:", engineer.GetPosition())

	fmt.Println("Assistant Address:", assistant.GetAddress())

}
