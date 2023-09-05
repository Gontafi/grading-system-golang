package models

import "time"

type Role struct {
	ID       int    `json:"ID"`
	Name     string `json:"role_name"`
	StatusID int    `json:"status_id"`
}
type User struct {
	ID           int       `json:"id" required:"-"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RoleUser struct {
	ID     int `json:"id"`
	RoleID int `json:"role_Id"`
	UserID int `json:"userID"`
}
type Lesson struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
}
type StudentLesson struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	LessonID  int `json:"lesson_id"`
}

type HomeWork struct {
	ID               int    `json:"id"`
	StudentID        int    `json:"student_id"`
	LessonID         int    `json:"lesson_id"`
	TeacherID        int    `json:"teacher_id"`
	StatusID         int    `json:"status_id"`
	HomeTaskResponse string `json:"home_task_response"`
}
type Mark struct {
	ID              int    `json:"id"`
	HomeWorkID      int    `json:"home_work_id"`
	HomeWorkGrade   int    `json:"home_work_grade"`
	AttendanceGrade int    `json:"attendance_grade"`
	Date            string `json:"date"`
}

type Rating struct {
	StudentID   int     `json:"student_id"`
	StudentName string  `json:"student_name"`
	Score       float64 `json:"score"`
}
