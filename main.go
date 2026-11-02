package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	ID         int
	Name       string
	Position   string
	Department string
}

func (e *Employee) save() error {
	filename := fmt.Sprintf("%d.txt", e.ID)
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0600)
}

func loadEmployee(id int) (*Employee, error) {
	filename := fmt.Sprintf("%d.txt", id)
	info, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	e := &Employee{}
	err = json.Unmarshal(info, e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func main() {
	p := &Employee{ID: 1, Name: "Hoshina Subaru", Position: "Wife", Department: "Uy's heart"}
	p.save()
	e, err := loadEmployee(p.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(e)
}
