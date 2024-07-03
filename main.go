package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	a "habit-builder/api"
)

type Habit struct {
	Name      string `json:"name"`
	Cue       string `json:"cue"`
	Craving   string `json:"craving"`
	Action    string `json:"action"`
	Reward    string `json:"reward"`
	Streak    int    `json:"streak"`
	Count     int    `json:"count"`
	Frequency int    `json:"frequency"`
}

var habits []Habit

const Dir = "./habits"

func loadHabits() ([]Habit, error) {
	habits := []Habit{}
	files, err := os.ReadDir(Dir)
	if err != nil {
		log.Fatal("read directory error", err)
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			habit := Habit{}
			f, err := os.ReadFile(Dir + "/" + file.Name())
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
				return nil, err
			}
			err = json.Unmarshal(f, &habit)
			if err != nil {
				return nil, err
			}
			habits = append(habits, habit)
		}
	}
	return habits, nil
}

func addHabit(h Habit) bool {
	b, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(Dir+"/"+h.Name+".json", []byte(b), 0777)
	return true
}

func incrementHabitCounter(name string) {
	habit := Habit{}
	f, err := os.ReadFile(Dir + "/" + name + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(f, &habit)
	if f != nil {
		habit.Name = name
		habit.Count++
	}
	b, err := json.Marshal(habit)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(Dir+"/"+name+".json", []byte(b), 0777)
}
func listCommands() {
	//TODO: update to be a constant list to use elsewhere.
	commands := []string{
		"listHabits",
		"addHabit",
		"habitCompleted",
	}
	for _, c := range commands {
		fmt.Printf("- %v \n", c)
	}
}
func main() {
	a.InitRoutes()

	//TODO: move this to another file and keep separate

	//	if len(os.Args) < 2 {
	//		listCommands()
	//		os.Exit(1)
	//
	// }
	// switch os.Args[1] {
	// case "listHabits":
	//
	//	habits, err := loadHabits()
	//	if err != nil {
	//		fmt.Printf("error loading habits: %v", err)
	//		os.Exit(1)
	//	}
	//	for _, h := range habits {
	//		fmt.Println(h.Name)
	//	}
	//
	// case "addHabit":
	//
	//	var h Habit
	//	scanner := bufio.NewScanner(os.Stdin)
	//
	//	//Name
	//	fmt.Println("Enter habit name: ")
	//	scanner.Scan()
	//	name := scanner.Text()
	//	h.Name = name
	//
	//	addHabit(h)
	//
	// case "habitCompleted":
	//
	//	scanner := bufio.NewScanner(os.Stdin)
	//	fmt.Println("habit to be updated")
	//	scanner.Scan()
	//	habitName := scanner.Text()
	//	incrementHabitCounter(habitName)
	//
	// default:
	//
	//		listCommands()
	//	}
}
