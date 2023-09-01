package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grading-system-golang/internal/services"
)

type Handler struct {
	Service *services.ServiceV1
}

func NewHandler(service *services.ServiceV1) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	{
		auth.Post("/sign-in", h.SignIn)
		auth.Post("/sign-up", h.SignUp)
	}

	api := app.Group("/api")
	api.Use(h.AuthMiddleware())

	lessons := api.Group("/lessons")
	{
		lessons.Get("/", h.GetLessons)
		lessons.Get("/:id", h.GetLesson)
		lessons.Post("/", h.CreateLesson)
		lessons.Put("/:id", h.UpdateLesson)
		lessons.Delete("/:id", h.DeleteLesson)
		lessons.Get("students/:student_id", h.GetLessonsForStudent)
	}

	students := api.Group("/students")
	{
		students.Get("/", h.GetStudents)
		students.Get("/:id", h.GetStudent)
		students.Post("/", h.CreateStudent)
		students.Put("/:id", h.UpdateStudent)
		students.Delete("/:id", h.DeleteStudent)
		students.Post("/:id/lessons/:lesson_id", h.AddStudentToLesson)
		students.Delete("/:id/lessons/:lesson_id", h.RemoveStudentFromLesson)
		students.Get("lessons/:lesson_id", h.GetStudentsForLesson)
	}

	teachers := api.Group("/teachers")
	{
		teachers.Get("/", h.GetTeachers)
		teachers.Get("/:id", h.GetTeacher)
		teachers.Post("/", h.CreateTeacher)
		teachers.Put("/:id", h.UpdateTeacher)
		teachers.Delete("/:id", h.DeleteTeacher)
	}

	users := api.Group("/users")
	{
		users.Get("/", h.GetUsers)
		users.Get("/:id", h.GetUser)
		users.Post("/", h.CreateUser)
		users.Put("/:id", h.UpdateUser)
		users.Delete("/:id", h.DeleteUser)
	}
	topRatings := api.Group("/top-ratings")
	{
		topRatings.Get("/", h.GetTopRatingFromCache)
	}

	topRatingsByLesson := api.Group("/top-ratings-by-lesson/:lessonID")
	{
		topRatingsByLesson.Get("/", h.GetTopRatingByLessonFromCache)
	}
}
