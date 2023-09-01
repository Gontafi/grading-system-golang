package repositories

import (
	"context"
	"github.com/grading-system-golang/internal/models"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	CreateMark(mark models.Mark) (int, error)
	GetMarkByID(markID int) (models.Mark, error)
	GetAllMarks() ([]models.Mark, error)
	DeleteMark(markID int) error
	UpdateMark(mark models.Mark) error
	GetStudentLesson(studentID int, lessonID int) (models.StudentLesson, error)
	AddStudentToLesson(studentID int, lessonID int) (int, error)
	RemoveStudentFromLesson(studentID int, lessonID int) error
	GetStudentsForLesson(lessonID int) ([]models.Student, error)
	GetLessonsForStudent(studentID int) ([]models.Lesson, error)
	AddLesson(lesson models.Lesson) (int, error)
	DeleteLesson(id int) error
	UpdateLesson(lesson models.Lesson) error
	AllLessons() ([]models.Lesson, error)
	GetLessonByID(id int) (models.Lesson, error)
	AddTeacher(teacher models.Teacher) (int, error)
	DeleteTeacher(id int) error
	UpdateTeacher(teacher models.Teacher) error
	AllTeachers() ([]models.Teacher, error)
	GetTeacherByID(id int) (models.Teacher, error)
	AddStudent(student models.Student) (int, error)
	DeleteStudent(id int) error
	UpdateStudent(student models.Student) error
	AllStudents() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	AddUser(user models.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user models.User) error
	AllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	GetTopRating() ([]models.Rating, error)
	GetTopRatingByLesson(lessonID int, limit int) ([]models.Rating, error)
}

type RepositoryV1 struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewRepository(ctx context.Context, db *pgx.Conn) *RepositoryV1 {
	return &RepositoryV1{
		ctx: ctx,
		db:  db,
	}
}
