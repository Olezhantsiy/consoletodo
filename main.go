package main

import (
	"bufio"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	Title  string
	isDone bool
}

var tasks []Task

type Tasks []Task

func (t Tasks) Len() int      { return len(t) }
func (t Tasks) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type ByTitle struct{ Tasks }

func (t ByTitle) Less(i, j int) bool { return t.Tasks[i].Title < t.Tasks[j].Title }

type ByStatus struct{ Tasks }

func (t ByStatus) Less(i, j int) bool { return t.Tasks[i].isDone }

func main() {

	tasks = append(tasks, Task{"Проснуться", true})
	tasks = append(tasks, Task{"Поесть", true})
	tasks = append(tasks, Task{"Поспать", false})

	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть задачи")
		fmt.Println("3. Изменить статус задачи")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Изменить задачу")
		fmt.Println("6. Сортировать задачи")
		fmt.Println("7. Выход")
		fmt.Print("Выберите действие: ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при вводе")
			continue
		}
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Неверный ввод, введите число")
			continue
		}

		switch choice {
		case 1:
			screen.Clear()
			screen.MoveTopLeft()
			tasks = NewTask(TitleTask(), tasks)
		case 2:
			screen.Clear()
			screen.MoveTopLeft()
			listTasks(tasks)
		case 3:
			screen.Clear()
			screen.MoveTopLeft()
			num := NumTask()
			if num > 0 && num <= len(tasks) {
				tasks = SwitchTaskDone(num, tasks)
			} else {
				fmt.Println("Неверный ввод номера задачи")
			}
		case 4:
			screen.Clear()
			screen.MoveTopLeft()
			num := NumTask()
			if num > 0 && num <= len(tasks) {
				tasks = DeleteTask(NumTask(), tasks)
			} else {
				fmt.Println("Неверный ввод номера задачи")
			}
		case 5:
			screen.Clear()
			screen.MoveTopLeft()
			num := NumTask()
			if num > 0 && num <= len(tasks) {
				tasks = ChangeTask(NumTask(), TitleTask(), tasks)
			} else {
				fmt.Println("Неверный ввод номера задачи")
			}
		case 6:
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Println("1. Сортировка по названии задачи")
			fmt.Println("2. Сортировка по статусу задачи")
			//fmt.Println("3. Показать действующие задачи")
			fmt.Println("3. Вернуться")

			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка при вводе")
				continue
			}
			input = strings.TrimSpace(input)
			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Неверный ввод, введите число")
				continue
			}

			switch choice {
			case 1:
				screen.Clear()
				screen.MoveTopLeft()
				sort.Sort(ByTitle{tasks})
				listTasks(tasks)
			case 2:
				screen.Clear()
				screen.MoveTopLeft()
				sort.Sort(ByStatus{tasks})
				listTasks(tasks)
			case 3:
				return
			default:
				fmt.Println("Неверный выбор. Попробуйте снова.")
			}
		case 7:
			fmt.Print("Выход")
			return
		default:
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Println("Неверный выбор. Попробуйте снова...")
		}
	}
}

func NewTask(title string, tasks []Task) []Task {
	tasks = append(tasks, Task{Title: title, isDone: false})
	listTasks(tasks)
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

func SwitchTaskDone(NumTask int, tasks []Task) []Task {
	tasks[NumTask-1].isDone = !tasks[NumTask-1].isDone
	listTasks(tasks)
	return tasks
}

func ChangeTask(NumTask int, Title string, tasks []Task) []Task {
	if NumTask != 0 {
		tasks[NumTask-1].Title = Title
		listTasks(tasks)
	}
	return tasks
}

func DeleteTask(NumTask int, tasks []Task) []Task {
	if NumTask != 0 {
		tasks = append(tasks[:NumTask-1], tasks[NumTask:]...)
		listTasks(tasks)
	}
	return tasks
}
