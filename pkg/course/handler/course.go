package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/amchicas/go-api-gateway/pkg/course/domain"
	"github.com/amchicas/go-course-srv/pkg/pb"
	fiber "github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	courseService pb.CourseServiceClient
}

func NewCourseHandler(courseService pb.CourseServiceClient) CourseHandler {

	return CourseHandler{
		courseService: courseService,
	}

}
func (h *CourseHandler) Add(c *fiber.Ctx) error {

	var input domain.Course

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	userId := c.Locals("userID").(uint64)
	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument Id",
		})
	}
	role := c.Locals("role").(int64)

	if role != 1 && userId != uid {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not Authoritzed",
		})

	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})

	}
	course := &pb.Course{
		Title:       input.Title,
		Subtitle:    input.Subtitle,
		Description: input.Description,
		Status:      &pb.Course_CourseStatus(input.Status),
	}
	req := &pb.CourseReq{Course: course}
	res, err := h.courseService.CreateCourse(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on add request", "data": err.Error()})
	}

	return c.JSON(&fiber.Map{"status": fmt.Sprint(res.Status), "Error": fmt.Sprint(res.Error), "msg": fmt.Sprint(res.Msg)})
}
func (h *CourseHandler) Update(c *fiber.Ctx) error {

	var input domain.Course
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userId := c.Locals("userID").(uint64)
	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument Id",
		})
	}
	role := c.Locals("role").(int64)

	if role != 1 && userId != uid {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not Authoritzed",
		})

	}
	if err := c.BodyParser(&input); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})

	}

	course := &pb.Course{
		Name:        input.Name,
		Lastname:    input.Lastname,
		Description: input.Description,
		Votes:       input.Votes,
		Students:    input.Students,
		Website:     input.Website,
		Youtube:     input.Youtube,
		Linkedin:    input.Linkedin,
		Twitter:     input.Twitter,
		Facebook:    input.Facebook,
	}
	req := &pb.UpdateReq{Uid: uid, Course: course}
	res, err := h.courseService.Update(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": "Updated"})
}
func (h *CourseHandler) Find(c *fiber.Ctx) error {

	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument Id",
		})
	}
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &pb.FindCourseReq{Uid: uid}
	res, err := h.courseService.FindCourse(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on find request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": res.Course})
}
func (h *CourseHandler) Delete(c *fiber.Ctx) error {

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userId := c.Locals("userID").(uint64)
	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument Id",
		})
	}
	role := c.Locals("role").(int64)

	if role != 1 && userId != uid {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Not Authoritzed",
		})

	}
	req := &pb.DeleteReq{Uid: uid}
	res, err := h.courseService.Remove(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error: delete request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": res.Msg})
}
