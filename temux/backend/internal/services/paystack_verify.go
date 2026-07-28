package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"temux/internal/config"
)

const PaystackVerifyURL = "https://api.paystack.co/transaction/verify/"

type VerifyData struct {
	Status    string `json:"status"`
	Reference string `json:"reference"`
	Amount    int64  `json:"amount"`
}

type VerifyResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    VerifyData `json:"data"`
}

func VerifyPayment(reference string) (*VerifyResponse, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		PaystackVerifyURL+reference,
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+config.PaystackSecretKey(),
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result VerifyResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if !result.Status {
		return nil, fmt.Errorf(result.Message)
	}

	return &result, nil
}
