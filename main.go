package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	Title  string `json:"Title"`
	IsDone bool   `json:"IsDone"`
}

var tasks []Task

func main() {

	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		j, err := json.Marshal(tasks)
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile("tasks.json", j, 0644)
		if err != nil {
			fmt.Println("")
		}
	}

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &tasks)

	app := &cli.App{
		Name:  "Console todo",
		Usage: "Manage your tasks",

		After: func(c *cli.Context) error {
			SaveTask(tasks)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Show all tasks",
				Action: func(c *cli.Context) error {
					listTasks(tasks)
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "add new tasks",
				Action: func(c *cli.Context) error {
					title := strings.Join(c.Args().Slice(), " ") //c.Args().First()
					if title == "" {
						title = TitleTask()
					}
					tasks = NewTask(title, tasks)
					return nil
				},
			},
			{
				Name:  "switch",
				Usage: "Switch task status",
				Action: func(c *cli.Context) error {
					input := c.Args().First()
					input = strings.TrimSpace(input)
					num, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Ошибка при вводе. Кажется вы ввели что-то не то...")
						return nil
					}
					if num > 0 && num <= len(tasks) {
						tasks = SwitchTaskDone(num, tasks)
					} else {
						fmt.Println("Неверный ввод номера задачи")
					}
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete task",
				Action: func(c *cli.Context) error {
					input := c.Args().First()
					input = strings.TrimSpace(input)
					num, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Ошибка при вводе. Кажется вы ввели что-то не то...")
						return nil
					}
					if num > 0 && num <= len(tasks) {
						tasks = DeleteTask(num, tasks)
					} else {
						fmt.Println("Неверный ввод номера задачи")
					}
					return nil
				},
			},
			{
				Name:  "change",
				Usage: "Change task",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						fmt.Println("Нужно указать номер задачи и название")
						return nil
					}
					input := c.Args().Get(0)
					input = strings.TrimSpace(input)
					num, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Ошибка при вводе. Кажется вы ввели что-то не то...")
						return nil
					}
					if num < 1 || num > len(tasks) {
						fmt.Println("Неверный ввод номера задачи")
						return nil
					}
					title := strings.Join(c.Args().Slice()[1:], " ")
					if title == "" {
						fmt.Println("Новое название не может быть пустым")
						return nil
					}
					tasks = ChangeTask(num, title, tasks)
					return nil
				},
			},
			{
				Name:  "sortTitle",
				Usage: "Sort task by title",
				Action: func(c *cli.Context) error {
					sort.Sort(ByTitle{tasks})
					fmt.Println("Задачи отсортированы по названию")
					return nil
				},
			},
			{
				Name:  "sortDone",
				Usage: "Sort task by done",
				Action: func(c *cli.Context) error {
					sort.Sort(ByDone{tasks})
					fmt.Println("Задачи отсортирорваны по выполнению")
					return nil
				},
			},
			{
				Name:   "help",
				Usage:  "Show all commands",
				Action: cli.ShowAppHelp,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
