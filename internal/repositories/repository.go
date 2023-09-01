package repositories

import (
	"context"
	"github.com/grading-system-golang/internal/models"
	"github.com/jackc/pgx/v5"
	"time"
)

type Repository interface {
	CreateMark(mark models.Mark) (int, error)
	GetMarkByID(markID int) (models.Mark, error)
	GetAllMarks() ([]models.Mark, error)
	DeleteMark(markID int) error
	UpdateMark(mark models.Mark) error
	AddStudentToLesson(studentID int, lessonID int) (int, error)
	RemoveStudentFromLesson(studentID int, lessonID int) error
	GetStudentsForLesson(lessonID int) ([]models.User, error)
	GetLessonsForStudent(studentID int) ([]models.Lesson, error)
	GetStudentLesson(studentID int, lessonID int) (models.StudentLesson, error)
	AddLesson(lesson models.Lesson) (int, error)
	DeleteLesson(id int) error
	UpdateLesson(lesson models.Lesson) error
	AllLessons() ([]models.Lesson, error)
	GetLessonByID(id int) (models.Lesson, error)
	AddRole(role models.Role) (int, error)
	DeleteRole(id int) error
	UpdateRole(role models.Role) error
	AllRoles() ([]models.Role, error)
	GetRoleByID(id int) (models.Role, error)
	AddUser(user models.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user models.User) error
	AllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	GetTopRating(period time.Duration, limit int) ([]models.Rating, error)
	GetTopRatingByLesson(lessonID int, period time.Duration, limit int) ([]models.Rating, error)
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
