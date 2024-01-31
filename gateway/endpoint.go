package gateway

import (
	"context"

	"github.com/gagguilhermearaujo/two-services/hashing"
	"github.com/gofiber/fiber/v2"
)

func makeCreateHashEndpoint(s *GatewayServer) {
	s.fiberApp.Post("/CreateHash", func(c *fiber.Ctx) error {
		req := new(hashing.CreateHashRequest)
		err := c.BodyParser(req)
		if err != nil {
			return err
		}
		res, err := s.hashingService.CreateHash(context.TODO(), req)
		if err != nil {
			return err
		}

		return c.JSON(&res)
	})
}

func makeCheckHashEndpoint(s *GatewayServer) {
	s.fiberApp.Post("/CheckHash", func(c *fiber.Ctx) error {
		req := new(hashing.CheckHashRequest)
		err := c.BodyParser(req)
		if err != nil {
			return err
		}
		res, err := s.hashingService.CheckHash(context.TODO(), req)
		if err != nil {
			return err
		}

		return c.JSON(map[string]bool{"hash_exists": res.HashExists})
	})
}

func makeGetHashEndpoint(s *GatewayServer) {
	s.fiberApp.Post("/GetHash", func(c *fiber.Ctx) error {
		req := new(hashing.GetHashRequest)
		err := c.BodyParser(req)
		if err != nil {
			return err
		}
		res, err := s.hashingService.GetHash(context.TODO(), req)
		if err != nil {
			return err
		}

		return c.JSON(&res)
	})
}
