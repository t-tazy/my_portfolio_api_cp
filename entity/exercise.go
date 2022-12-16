package entity

import "time"

type ExerciseID int

// 誤代入を防ぐため、IDフィールドは独自型を使う
type Exercise struct {
	ID          ExerciseID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Created     time.Time  `json:"created"`
}

type Exercises []*Exercise
