package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Habit struct {
	Cue     string `json:"cue"`
	Craving string `json:"craving"`
	Action  string `json:"action"`
	Reward  string `json:"reward"`
	Streak  int    `json:"streak"`
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
	os.WriteFile(Dir + "/" + 
	return true
}

func main() {
	fmt.Println("hello habit tracker")
	//connect to database
	// make call
}
