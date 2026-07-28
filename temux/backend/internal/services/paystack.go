package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"temux/internal/config"
	"net/http"
	"io"
)

const (
	PaystackInitializeURL = "https://api.paystack.co/transaction/initialize"
)

type InitializeRequest struct {
	Email  string `json:"email"`
	Amount int64  `json:"amount"`
}

type InitializeData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type InitializeResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    InitializeData `json:"data"`
}

func InitializePayment(
	email string,
	amount float64,
) (*InitializeResponse, error) {

	reqBody := InitializeRequest{
		Email:  email,
		Amount: int64(amount * 100),
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		PaystackInitializeURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+config.PaystackSecretKey(),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result InitializeResponse

	err = json.Unmarshal(
		respBody,
		&result,
	)

	if err != nil {
		return nil, err
	}

	if !result.Status {
		return nil, fmt.Errorf(result.Message)
	}

	return &result, nil
}
