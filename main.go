package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func startTodo (reader *bufio.Reader) string {
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')

	return name
}

func deleteTask(reader *bufio.Reader, todo *Todo) {
	selectTasks := "Enter the ID of the task to be deleted from the options below: "

	for _,v := range todo.Tasks {
		selectTasks += "\n"+ strconv.Itoa(v.Id)
	}

	fmt.Println(selectTasks)
	taskIdString, _ := reader.ReadString('\n')
	taskId, _ := strconv.Atoi(strings.TrimSpace(taskIdString));

	todo.delete(taskId)

	nextAfterStart (reader, todo)
}

func nextAfterStart (reader *bufio.Reader, todo *Todo) {
	fmt.Println("Press \n a: To add tasks \n b: Delete Task \n c: To save")
	selected, _ := reader.ReadString('\n')
	selected = strings.TrimSpace(selected)

	if selected == "a" {
		startTaskCreation(reader, todo)
	}else if selected == "b" {
		if(len(todo.Tasks) > 0) {
			deleteTask(reader, todo)
		}else{
			nextAfterStart (reader, todo)
		}
	}else if selected == "c" {
		todo.save()
	}else {
		nextAfterStart (reader, todo)
	}
}

func startTaskCreation(reader *bufio.Reader, todo *Todo) {
	fmt.Println("What is the name of this task?")
	taskname, _ := reader.ReadString('\n')
	status := selectTaskStatus(reader)

	task := Task{
		Id: rand.Intn(100000),
		Name: taskname,
		Status: status,
	}

	todo.add(task)

	afterTaskAddition(reader, todo)
}

func selectTaskStatus(reader *bufio.Reader) string {
	fmt.Println("Select task status \n a: pending \n b: completed")
	taskoption, _ := reader.ReadString('\n')
	taskoption = strings.TrimSpace(taskoption)

	var status string

	if(taskoption == "a") {
		status = "pending"
	}else if taskoption == "b" {
		status = "completed"
	}else{
		selectTaskStatus(reader)
	}

	return status
}

func afterTaskAddition(reader *bufio.Reader, todo *Todo) {
	fmt.Println("Press \n a: To add another task \n b: To save")
	whatNext, _ := reader.ReadString('\n')
	whatNext = strings.TrimSpace(whatNext)

	if(whatNext == "a") {
		startTaskCreation(reader, todo)
	}else if whatNext == "b" {
		todo.save()
	}else{
		afterTaskAddition(reader, todo)
	}
}

func updateExistingTodo(reader *bufio.Reader) {
	existingTodos := fetch()
	selectTodos := "Enter the name of Todo from the options below: "

	for _,v := range existingTodos {
		selectTodos += "\n"+ v.Name
	}

	fmt.Println(selectTodos)
	todoName, _ := reader.ReadString('\n')
	todo,err := selectTodo(strings.TrimSpace(todoName))

	if(err != nil) {
		updateExistingTodo(reader)
	}else{
		nextAfterStart(reader, &todo)
	}
}

func startProcess(reader *bufio.Reader) {
	fmt.Println("Press \n a: To create a todo list \n b: To update todo list")

	selected, _ := reader.ReadString('\n')
	selected = strings.TrimSpace(selected)

	if selected == "a" {
		name := startTodo(reader)
		todo := create(name)
		
		nextAfterStart(reader, &todo)
	}else if selected == "b" {
		updateExistingTodo(reader)
	}else{
		startProcess(reader)
	}
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	startProcess(reader);
}