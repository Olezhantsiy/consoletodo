package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewTask(title string, tasks []Task) []Task {
	tasks = append(tasks, Task{Title: title, IsDone: false})
	fmt.Printf("Задача \"%s\" добавлена", title)
	return tasks
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
	}
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		fmt.Println("Пустая строка...")
	}
	return title
}

func listTasks(tasks []Task) {
	if len(tasks) < 1 {
		fmt.Println("Задач нет")
		return
	}
	for i, task := range tasks {
		done := "Не выполнено"
		if task.IsDone != false {
			done = "Выполнено"
		}
		fmt.Println(i+1, task.Title, done)
	}
}

func SwitchTaskDone(NumTask int, tasks []Task) []Task {
	tasks[NumTask-1].IsDone = !tasks[NumTask-1].IsDone
	task := tasks[NumTask-1]
	done := "Не выполнена"
	if task.IsDone != false {
		done = "Выполнена"
	}
	fmt.Printf("Задача \"%s\" %s\n", task.Title, done)
	return tasks
}

func ChangeTask(NumTask int, Title string, tasks []Task) []Task {
	if NumTask != 0 {
		tasks[NumTask-1].Title = Title
		title := tasks[NumTask-1].Title
		fmt.Printf("Задача изменена на \"%s\"", title)
	}
	return tasks
}

func DeleteTask(NumTask int, tasks []Task) []Task {
	if NumTask != 0 {
		tasks = append(tasks[:NumTask-1], tasks[NumTask:]...)
		fmt.Println("Задача удалена")
	}
	return tasks
}

func SaveTask(tasks []Task) {
	j, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Ошибка сохранения", err)
		return
	}
	err = os.WriteFile("tasks.json", j, 0644)
	if err != nil {
		fmt.Println("Ошибка сохранения", err)
	}
}

type Tasks []Task

func (t Tasks) Len() int      { return len(t) }
func (t Tasks) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type ByTitle struct{ Tasks }

func (t ByTitle) Less(i, j int) bool { return t.Tasks[i].Title < t.Tasks[j].Title }

type ByDone struct{ Tasks }

func (t ByDone) Less(i, j int) bool { return t.Tasks[i].IsDone }
