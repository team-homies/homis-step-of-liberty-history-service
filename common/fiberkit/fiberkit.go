package fiberkit

import "github.com/gofiber/fiber/v2"

type FiberKit struct {
	C *fiber.Ctx
}

// c.Locals( , ), c.Next()에서 넘어온 string키값을 string으로 반환
func (f *FiberKit) GetLocalsString(key string) string {
	return f.C.Locals(key).(string)
}

// c.Locals( , ), c.Next()에서 넘어온 string키값을 int로 반환
func (f *FiberKit) GetLocalsInt(key string) int {
	return f.C.Locals(key).(int)
}

// json으로 변환 및 response OK
func (f *FiberKit) HttpOK(data any) error {
	return f.C.JSON(data)
}

// json으로 변환 및 response Fail
func (f *FiberKit) HttpFail(data any, errCode int) error {
	return f.C.Status(errCode).JSON(data)
}
