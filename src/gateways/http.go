package gateways

import (
	service "go-fiber-template/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	UserService service.IUsersService
	IPService   service.IIpService
	StripeService service.IStripeService
}

func NewHTTPGateway(app *fiber.App, users service.IUsersService, ip service.IIpService,stripe service.IStripeService) {
	gateway := &HTTPGateway{
		UserService: users,
		IPService:   ip,
		StripeService: stripe,
	}

	RouteUsers(*gateway, app)
	RouteIP(*gateway, app)
}
