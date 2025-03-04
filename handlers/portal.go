package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func (h *ApiHandler) IncAndGetRefreshed(ctx context.Context) (int64, error) {
	return h.redisClient.Incr(ctx, "refreshes").Result()
}

func (h *ApiHandler) HandleIndex(c *fiber.Ctx) error {
	refreshes, _ := h.IncAndGetRefreshed(c.Context())

	return c.Render("index", fiber.Map{
		"refreshed": refreshes,
	})
}
