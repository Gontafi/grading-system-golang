package models

import "time"

type User struct {
	ID           int       `json:"id" required:"-"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Student struct {
	UserID  int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Teacher struct {
	UserID  int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Lesson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type StudentLesson struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	LessonID  int `json:"lesson_id"`
}

type Mark struct {
	ID              int       `json:"id"`
	TeacherID       int       `json:"teacher_id"`
	StudentID       int       `json:"student_id"`
	LessonID        int       `json:"lesson_id"`
	HomeWorkGrade   int       `json:"home_work_grade"`
	AttendanceGrade int       `json:"attendance_grade"`
	Date            time.Time `json:"date"`
}

type Attendance struct {
	ID        int       `json:"id"`
	LessonID  int       `json:"lesson_id"`
	StudentID int       `json:"student_id"`
	TeacherID int       `json:"teacher_id"`
	Status    bool      `json:"status"`
	Date      time.Time `json:"date"`
}

type Rating struct {
	StudentID      int     `json:"student_id"`
	StudentName    string  `json:"student_name"`
	StudentSurname string  `json:"student_surname"`
	Score          float64 `json:"score"`
}
