package http

import (
	"errors"
	"fmt"
	"io"
	"letsgo-mini-is/internal/adapter/handler/http/dto"
	"letsgo-mini-is/internal/adapter/handler/http/exception"
	"letsgo-mini-is/internal/core/domain"
	"letsgo-mini-is/internal/core/port"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PictureHandler struct {
	pictureService port.PictureService
}

func NewPictureHandler(pictureService port.PictureService) *PictureHandler {
	return &PictureHandler{
		pictureService: pictureService,
	}
}

func (h *PictureHandler) Create(c *fiber.Ctx) error {

	file, err := c.FormFile("picture")
	if err != nil {
		return exception.NewParsingFailedExceptino(err)
	}
	openFile, err := file.Open()
	if err != nil {
		return exception.NewInternalErrorException(err, "사용자로부터 요청받은 이미지를 열람했더니 에러가 발생하였습니다.")
	}
	defer openFile.Close()
	content, err := io.ReadAll(openFile)
	if err != nil {
		return exception.NewInternalErrorException(err, "열람한 이미지를 byte로 변환했더니 에러가 발생하였습니다.")
	}

	ext := filepath.Ext(file.Filename)
	res, err := h.pictureService.Create(c.Context(), &domain.Picture{
		ID:        uuid.New(),
		Content:   content,
		Extension: strings.Split(ext, ".")[1],
	})
	if err != nil {
		return exception.NewInternalErrorException(err)
	}

	return c.Status(fiber.StatusCreated).JSON(&dto.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success Created",
		Data:    dto.PictureCreateResponseOf(res),
	})
}

func (h *PictureHandler) Find(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return exception.NewParsingFailedExceptino(errors.New("id의 값이 제대로 넘어 오지 않았습니다."))
	}

	err := uuid.Validate(id)
	if err != nil {
		return exception.NewInvalidUUIDException(err)
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return exception.NewInternalErrorException(err, "string uuid를 google.uuid로 parse 하던 중 에레가 발생했습니다.")
	}

	res, err := h.pictureService.Find(c.Context(), uuid)
	if err != nil {
		return exception.NewInternalErrorException(err)
	}

	fmt.Println(res.Extension)
	key := fmt.Sprintf("image/%s", res.Extension)
	c.Set("Content-Type", key)
	return c.Status(fiber.StatusOK).Send(res.Content)
}

func (h *PictureHandler) Update(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return exception.NewParsingFailedExceptino(errors.New("Query Params의 값이 제대로 넘어 오지 않았습니다."))
	}
	uuid, err := uuid.FromBytes([]byte(id))
	if err != nil {
		return exception.NewInvalidUUIDException(err)
	}

	file, err := c.FormFile("picture")
	if err != nil {
		return exception.NewParsingFailedExceptino(err)
	}
	openFile, err := file.Open()
	if err != nil {
		return exception.NewInternalErrorException(err, "사용자로부터 요청받은 이미지를 열람했더니 에러가 발생하였습니다.")
	}
	defer openFile.Close()
	content, err := io.ReadAll(openFile)
	if err != nil {
		return exception.NewInternalErrorException(err, "열람한 이미지를 byte로 변환했더니 에러가 발생하였습니다.")
	}
	if err != nil {
		return exception.NewInternalErrorException(err)
	}

	ext := filepath.Ext(file.Filename)

	err = h.pictureService.Update(c.Context(), &domain.Picture{
		ID:        uuid,
		Content:   content,
		Extension: strings.Split(ext, ".")[1],
	})

	if err != nil {
		return exception.NewInternalErrorException(err, "업데이트를 함수를 실행하던 중 에러가 발생하였습니다.")
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "성공적으로 갱신하였습니다.",
	})

}

func (h *PictureHandler) Delete(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return exception.NewParsingFailedExceptino(errors.New("Query Params의 값이 제대로 넘어 오지 않았습니다."))
	}
	uuid, err := uuid.FromBytes([]byte(id))
	if err != nil {
		return exception.NewInvalidUUIDException(err)
	}

	err = h.pictureService.Delete(c.Context(), uuid)
	if err != nil {
		return exception.NewInternalErrorException(err, "삭제 함수를 실행하던 중 에러가 발행하였습니다.")
	}

	return c.SendStatus(fiber.StatusOK)
}
