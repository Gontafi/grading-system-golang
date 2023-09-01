package services

import (
	"context"
	"github.com/grading-system-golang/internal/models"
	"github.com/grading-system-golang/internal/repositories"
	"github.com/redis/go-redis/v9"
	"time"
)

type Service interface {
	AddUser(user models.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	AddTeacher(teacher models.Teacher) (int, error)
	DeleteTeacher(id int) error
	UpdateTeacher(teacher models.Teacher) error
	GetAllTeachers() ([]models.Teacher, error)
	GetTeacherByID(id int) (models.Teacher, error)
	AddStudent(student models.Student) (int, error)
	DeleteStudent(id int) error
	UpdateStudent(student models.Student) error
	GetAllStudents() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	AddStudentToLesson(studentID int, lessonID int) (int, error)
	RemoveStudentFromLesson(studentID int, lessonID int) error
	GetStudentsForLesson(lessonID int) ([]models.Student, error)
	GetLessonsForStudent(studentID int) ([]models.Lesson, error)
	CreateMark(mark models.Mark) (int, error)
	GetMarkByID(markID int) (models.Mark, error)
	GetAllMarks() ([]models.Mark, error)
	DeleteMark(markID int) error
	UpdateMark(mark models.Mark) error
	GetStudentLesson(studentID int, lessonID int) (models.StudentLesson, error)
	AddLesson(lesson models.Lesson) (int, error)
	DeleteLesson(id int) error
	UpdateLesson(lesson models.Lesson) error
	GetAllLessons() ([]models.Lesson, error)
	GetLessonByID(id int) (models.Lesson, error)
}

type ServiceV1 struct {
	repository *repositories.RepositoryV1
	rdb        *redis.Client
	ctx        context.Context
	expiry     time.Duration
}

func NewService(
	r *repositories.RepositoryV1,
	rdb *redis.Client,
	ctx context.Context,
	expiry time.Duration,
) *ServiceV1 {
	return &ServiceV1{
		repository: r,
		rdb:        rdb,
		ctx:        ctx,
		expiry:     expiry,
	}
}
