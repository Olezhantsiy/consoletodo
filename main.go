package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title  string
	isDone bool
}

func main() {

	var tasks []Task

	tasks = append(tasks, Task{"Проснуться", true})
	tasks = append(tasks, Task{"Поесть", true})
	tasks = append(tasks, Task{"Поспать", false})

	for {
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть задачи")
		fmt.Println("3. Изменить статус задачи")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Изменить задачу")
		fmt.Println("6. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			NewTask(TitleTask(), tasks)
		case 2:
			listTasks(tasks)
		case 3:
			SwitchTaskDone(NumTask(), tasks)
		case 4:
			DeleteTask(NumTask(), tasks)
		case 5:
			ChangeTask(NumTask(), TitleTask(), tasks)
		case 6:
			fmt.Print("Выход")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func NewTask(title string, tasks []Task) {
	tasks = append(tasks, Task{Title: title, isDone: false})
	listTasks(tasks)
}

func NumTask() int {
	fmt.Println("Введите номер задачи")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 0
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Кажется вы ввели что-то не то...")
		return 0
	}
	return num
}

func TitleTask() string {
	fmt.Println("Введите название задачи")
	title, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при вводе названия", err)
		return "0"
	}
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		fmt.Println("Пустая строка...")
		return "0"
	}
	return title
}

func listTasks(tasks []Task) {
	for i, task := range tasks {
		status := "Не выполнено"
		if task.isDone != false {
			status = "Выполнено"
		}
		fmt.Println(i+1, task.Title, status)
	}
	fmt.Println("Нажмите чтобы продолжить")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func SwitchTaskDone(NumTask int, tasks []Task) {
	if NumTask != 0 {
		task := tasks[NumTask-1].isDone
		if task != true {
			tasks[NumTask-1].isDone = true
		} else {
			tasks[NumTask-1].isDone = false
		}
		listTasks(tasks)
	}
}

func ChangeTask(NumTask int, Title string, tasks []Task) {
	if NumTask != 0 {
		tasks[NumTask-1].Title = Title
		listTasks(tasks)
	}
}

func DeleteTask(NumTask int, tasks []Task) {
	if NumTask != 0 {
		tasks = append(tasks[:NumTask-1], tasks[NumTask:]...)
		listTasks(tasks)
	}
}
