package handler

import (
	"context"

	"github.com/VladRomanciuc/CSProject/auth/facebook"
	"github.com/VladRomanciuc/CSProject/auth/google"
	"github.com/gofiber/fiber/v2"
)

// Auth fiber handler
func GoogleAuth(c *fiber.Ctx) error {
	path := gauth.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}
func FacebookAuth(c *fiber.Ctx) error {
	path := fauth.ConfigFacebook()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

// Callback to receive google's response
func GoogleCallback(c *fiber.Ctx) error {
	token, error := gauth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	user := gauth.GetGoogleUser(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{
		"id": user.ID,
		"name": user.Name,
		"given_name": user.GivenName,
		"family_name": user.FamilyName,
		"gender": user.Gender,
		"email": user.Email,
		"picture": user.Picture})
}

// Callback to receive google's response
func FacebookCallback(c *fiber.Ctx) error {
	token, error := fauth.ConfigFacebook().Exchange(context.Background(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	user := fauth.GetFacebookUser(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{
		"id": user.ID,
		"name": user.Name,
		"given_name": user.GivenName,
		"family_name": user.FamilyName,
		"gender": user.Gender,
		"email": user.Email,
		"picture": user.Picture})
}