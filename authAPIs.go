package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Credentials struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=32"`
}
type Profile struct {
	FirstName string `validate:"omitempty"`
	LastName  string `validate:"omitempty"`
	Gender    string
}

const authBaseURL = "https://devapi.lrinternal.com/identity/v2/auth"
const authURL = authBaseURL + "/login"
const updateAccountURL = authBaseURL + "/account"

func LoginHandler(c *fiber.Ctx) error {
	fmt.Println("inside login handler")
	credentials := new(Credentials)
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	validate := validator.New()
	if err := validate.Struct(credentials); err != nil {
		return err
	}

	base, _ := url.Parse(authURL)
	fmt.Println("apikey", c.Query("apikey"))
	params := url.Values{}
	params.Add("apikey", c.Query("apikey"))
	base.RawQuery = params.Encode()
	fmt.Println(base)
	body, err := json.Marshal(map[string]string{
		"email":    credentials.Email,
		"password": credentials.Password})
	if err != nil {
		return err
	}
	resp, err := http.Post(base.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	result, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	c.Set("content-type", "application/json")
	return c.Status(fiber.StatusOK).Send(result)
}

func UpdateProfileHandler(c *fiber.Ctx) error {
	fmt.Println("inside update profile handler")
	profile := new(Profile)
	if err := c.BodyParser(profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}
	fmt.Println(profile)
	validate := validator.New()
	validate.Struct(profile)
	base, _ := url.Parse(updateAccountURL)
	fmt.Println(base)
	params := url.Values{}
	params.Add("apikey", c.Query("apikey"))
	headers := c.GetReqHeaders()
	fmt.Println("header", headers)
	token := headers["Authorization"]
	base.RawQuery = params.Encode()
	body, _ := json.Marshal(profile)
	req, _ := http.NewRequest(http.MethodPut, base.String(), bytes.NewBuffer(body))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	result, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	c.Set("content-type", "application/json")
	return c.Status(fiber.StatusOK).Send(result)

}
