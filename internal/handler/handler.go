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

	teacherRoute := app.Group("/marks", h.TeacherRoleMiddleware())
	{
		teacherRoute.Post("/check-homework", h.CheckHomeWorkAndPutGrades)
		teacherRoute.Put("/change-mark/:id", h.UpdateMark)
		teacherRoute.Delete("/delete-mark/:id", h.DeleteMark)
	}
	topRatings := app.Group("/top-ratings")
	{
		topRatings.Get("/", h.GetTopRatingFromCache)
	}

	topRatingsByLesson := app.Group("/top-ratings-by-lesson/:lessonID")
	{
		topRatingsByLesson.Get("/", h.GetTopRatingByLessonFromCache)
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

	users := api.Group("/users")
	{
		users.Get("/", h.GetUsers)
		users.Get("/:id", h.GetUser)
		users.Post("/", h.CreateUser)
		users.Put("/:id", h.UpdateUser)
		users.Delete("/:id", h.DeleteUser)
		users.Get("/:student_id/lessons/:lesson_id", h.GetStudentLesson)
		users.Get("lessons/:lesson_id", h.GetStudentsForLesson)
	}
	enroll := api.Group("/student-lesson/", h.AdminRoleMiddleware())
	{
		enroll.Post("student/:student_id/lesson/:lesson_id", h.AddStudentToLesson)
		enroll.Delete("student/:student_id/lesson/:lesson_id", h.RemoveStudentFromLesson)
	}
	marks := api.Group("/marks")
	{
		marks.Get("/:id", h.GetMark)
		marks.Get("/", h.GetMarks)
	}
}
