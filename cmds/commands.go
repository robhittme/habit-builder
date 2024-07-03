package commands

import (
	"encoding/json"
	"log"
	"os"
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

const Dir = "./habits"

func LoadHabits() ([]Habit, error) {
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

func AddHabit(h Habit) bool {
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
