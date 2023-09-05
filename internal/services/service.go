package services

import (
	"context"
	"fmt"
	"github.com/grading-system-golang/internal/models"
	"github.com/grading-system-golang/internal/repositories"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	AddUser(user models.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	AddRole(role models.Role) (int, error)
	DeleteRole(id int) error
	UpdateRole(role models.Role) error
	GetAllRoles() ([]models.Role, error)
	GetRoleByID(id int) (models.Role, error)
	AddStudentToLesson(studentID int, lessonID int) (int, error)
	RemoveStudentFromLesson(studentID int, lessonID int) error
	GetStudentsForLesson(lessonID int) ([]models.User, error)
	GetLessonsForStudent(studentID int) ([]models.Lesson, error)
	GetStudentLesson(studentID int, lessonID int) (models.HomeWork, error)
	CreateMark(mark models.Mark) (int, error)
	GetMarkByID(markID int) (models.Mark, error)
	GetAllMarks() ([]models.Mark, error)
	DeleteMark(markID int) error
	UpdateMark(mark models.Mark) error
	AddLesson(lesson models.Lesson) (int, error)
	DeleteLesson(id int) error
	UpdateLesson(lesson models.Lesson) error
	GetAllLessons() ([]models.Lesson, error)
	GetLessonByID(id int) (models.Lesson, error)
	AddRoleUser(roleID int, userID int) (int, error)
	RemoveRoleUser(roleID int, userID int) error
	GetUsersForRole(roleID int) ([]models.User, error)
	GetRolesForUser(userID int) ([]models.Role, error)
	GetUserRole(userID int) (models.Role, error)

	RegisterUser(user models.User) (int, error)
	GetTokenFromUser(username string, password string) (string, error)
	ParseToken(accessToken string) (*Claims, error)
	GetTopRatingFromCache(period string, limit int) ([]models.Rating, error)
	GetTopRatingByLessonFromCache(lessonID int, period time.Duration, limit int) ([]models.Rating, error)

	getCacheKey(baseKey string, args ...interface{}) string
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

func (s *ServiceV1) getCacheKey(baseKey string, args ...interface{}) string {
	keyParts := []string{baseKey}
	for _, arg := range args {
		keyParts = append(keyParts, argToString(arg))
	}
	return strings.Join(keyParts, ":")
}

func argToString(arg interface{}) string {
	switch v := arg.(type) {
	case time.Duration:
		return v.String()
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
