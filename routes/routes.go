package routes

import (
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"
)

var (
	OCR_CLIENT *gosseract.Client
)

type Ocr_response struct {
	Ocr_data string `json:"ocr_data"`
}

type OcrRequest struct {
	Base64String string `json:"base64String"`
}

func Perform_ocr(c *fiber.Ctx) error {
	req := OcrRequest{}
	c.BodyParser(&req)

	decoded_image, err := base64.StdEncoding.DecodeString(req.Base64String)
	if err != nil {
		log.Fatalf("Failed to decode base64 string: %v", err)
	}

	OCR_CLIENT.SetImageFromBytes(decoded_image)

	ocr_text, err := OCR_CLIENT.Text()
	if err != nil {
		log.Fatalf("Failed to ocr text: %v", err)
	}

	resp := Ocr_response{Ocr_data: ocr_text}

	return c.Status(200).JSON(resp)

}

type ocr_byte struct {
	Image []byte `json:"image"`
}

func Ocr_from_byte(c *fiber.Ctx) error {

	
	req := c.Body()

	OCR_CLIENT.SetImageFromBytes(req)

	ocr_text, err := OCR_CLIENT.Text()
	if err != nil {
		log.Fatalf("Failed to ocr text: %v", err)
	}

	resp := Ocr_response{Ocr_data: ocr_text}

	return c.Status(200).JSON(resp)

}
