package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/amchicas/go-api-gateway/pkg/profile/domain"
	"github.com/amchicas/go-profile-srv/pkg/pb"
	fiber "github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	ProfileService pb.ProfileServiceClient
}

func NewProfileHandler(profileService pb.ProfileServiceClient) ProfileHandler {

	return ProfileHandler{
		ProfileService: profileService,
	}

}
func (h *ProfileHandler) Add(c *fiber.Ctx) error {

	var input domain.Profile

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
	profile := &pb.Profile{
		Id:          uid,
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
	req := &pb.AddReq{Profile: profile}
	res, err := h.ProfileService.AddProfile(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on add request", "data": err.Error()})
	}

	return c.JSON(&fiber.Map{"status": fmt.Sprint(res.Status), "Error": fmt.Sprint(res.Error), "msg": fmt.Sprint(res.Msg)})
}
func (h *ProfileHandler) Update(c *fiber.Ctx) error {

	var input domain.Profile
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

	profile := &pb.Profile{
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
	req := &pb.UpdateReq{Uid: uid, Profile: profile}
	res, err := h.ProfileService.Update(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": "Updated"})
}
func (h *ProfileHandler) Find(c *fiber.Ctx) error {

	uid, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid argument Id",
		})
	}
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &pb.FindProfileReq{Uid: uid}
	res, err := h.ProfileService.FindProfile(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on find request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": res.Profile})
}
func (h *ProfileHandler) Delete(c *fiber.Ctx) error {

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
	res, err := h.ProfileService.Remove(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error: delete request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": res.Status, "message": res.Msg})
}
