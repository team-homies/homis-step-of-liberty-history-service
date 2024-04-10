package fiberkit

import "github.com/gofiber/fiber/v2"

type FiberKit struct {
	C *fiber.Ctx
}

func (f *FiberKit) GetLocalsString(key string) string {
	return f.C.Locals(key).(string)
}

func (f *FiberKit) GetLocalsInt(key string) int {
	return f.C.Locals(key).(int)
}

func (f *FiberKit) HttpOK(data any) error {
	return f.C.JSON(data) // 제이슨으로 변환 및 response
}

func (f *FiberKit) HttpFail(data any, errCode int) error {
	return f.C.Status(errCode).JSON(data)
}
