package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Name string
	Tasks []Task
}

type Task struct {
	Id int
	Name string
	Status string
}

func (td *Todo) add (task Task) {
	td.Tasks = append(td.Tasks, task)
}

func (td *Todo) delete (id int) {
	tasks := []Task{}

	for _, v := range td.Tasks {
		if v.Id != id {
			tasks = append(tasks, v)
		}
	}

	td.Tasks = tasks
}

func create (name string) Todo {
	td := Todo{Name: strings.ToLower(strings.TrimSpace(name))}
	return td
}

func (td *Todo) save () {
	todos := fetch()

	update := []Todo{}
	if len(todos) > 0 {
		for _,v := range todos {
			if v.Name != td.Name {
				update = append(update, v)
			}
		}
	}

	update = append(update, *td)

	jsonFile, _ := json.MarshalIndent(update, "", "\t")
	os.WriteFile("todo.json", jsonFile, 0644);

	fmt.Println("Todo successfully saved!!")
}

func fetch () []Todo {
	jsonData, _ := os.ReadFile("todo.json")

	todos := []Todo{}
	json.Unmarshal(jsonData, &todos)

	return todos
}

func selectTodo (name string) (Todo, error) {
	todos := fetch()

	for _,v := range todos {
		if v.Name == strings.ToLower(strings.TrimSpace(name)) {
			return v, nil
		}
	}

	return  Todo{}, errors.New("Todo not found")
}