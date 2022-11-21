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

const managementBaseURL = "https://devapi.lrinternal.com/identity/v2/manage"
const createAccountURL = managementBaseURL + "/account"

type EmailStruct struct {
	Type  string `validate:"required"`
	Value string `validate:"required,email"`
}
type Account struct {
	Email         []EmailStruct `validate:"required"`
	Password      string        `validate:"required,min=6,max=32"`
	FirstName     string
	LastName      string
	Gender        string
	EmailVerified bool
}

func CreateAccountHandler(c *fiber.Ctx) error {
	fmt.Println("inside create profile handler")
	account := new(Account)
	if err := c.BodyParser(account); err != nil {
		return requestMalformed(c, err)
	}
	account.EmailVerified = true
	fmt.Println(account)
	validate := validator.New()
	if err := validate.Struct(account); err != nil {
		return err
	}

	base, _ := url.Parse(createAccountURL)
	fmt.Println(base)
	params := url.Values{}
	apikey, apisecret := c.Query("apikey"), c.Query("apisecret")
	if apikey == "" || apisecret == "" {
		return ApiKeyOrSecretKeyMissing(c)
	}
	params.Add("apikey", apikey)
	params.Add("apisecret", apisecret)

	base.RawQuery = params.Encode()
	body, err := json.Marshal(account)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Request not formatted"})
	}
	fmt.Println(string(body))
	req, _ := http.NewRequest(http.MethodPost, base.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	result, _ := io.ReadAll(resp.Body)
	fmt.Println(resp)
	defer resp.Body.Close()
	c.Set("content-type", "application/json")
	return okResponse(c, result)
}
