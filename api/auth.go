package api

import (
	"context"
	"net/http"

	"github.com/Jimbo8702/lets_get_one/db"
	"github.com/Jimbo8702/lets_get_one/types"
	"github.com/Jimbo8702/lets_get_one/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/nedpals/supabase-go"
	"github.com/sagikazarmark/slog-shim"
)

type AuthHandler struct {
	store 		*db.Store
	log 		*slog.Logger
	validator 	validator.Validator
	supa 		*supabase.Client
}

func NewAuthHandler(db *db.Store, ls *slog.Logger, vd validator.Validator, supa *supabase.Client) *AuthHandler {
	return &AuthHandler{
		store: db,
		log: ls,
		validator: vd,
		supa: supa,
	}
}

func (h *AuthHandler) HandleSignupWithEmail(c *fiber.Ctx) error {
	var params *types.SignupParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := h.validator.Validate(params); errors != nil {
		h.log.Error("validation error with params", params)
		return c.JSON(errors)
	}
	resp, err := h.supa.Auth.SignUp(context.Background(), supabase.UserCredentials{
		Email: 		params.Email,
		Password: 	params.Password,
		// Data: 		fiber.Map{"fullName": params.FullName},
	})
	if err != nil {
		h.log.Error("error with supabase signup: ", err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	acc := types.User{
		UserID: resp.ID,
	}
	_, err = h.store.User.Insert(&acc)
	if err != nil {
		h.log.Error("error inserting user: ", err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	h.log.Info("new user signup: email: ", resp.Email)
	return c.Render("home", fiber.Map{"isAuthenticated": true})
}
