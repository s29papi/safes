package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ProposalResponse struct {
	Type         string       `json:"type"`
	Transaction  *Transaction `json:"transaction,omitempty"`  // Only present when type is "TRANSACTION"
	Timestamp    int64        `json:"timestamp,omitempty"`    // Only present when type is "DATE_LABEL"
	ConflictType string       `json:"conflictType,omitempty"` // Only present when type is "TRANSACTION"
}

type Transaction struct {
	TxInfo        TxInfo        `json:"txInfo"`
	ID            string        `json:"id"`
	TxHash        string        `json:"txHash"`
	Timestamp     int64         `json:"timestamp"`
	TxStatus      string        `json:"txStatus"`
	ExecutionInfo ExecutionInfo `json:"executionInfo"`
	SafeAppInfo   SafeAppInfo   `json:"safeAppInfo"`
}

type TxInfo struct {
	Type             string `json:"type"`
	HumanDescription string `json:"humanDescription"`
	Creator          Party  `json:"creator"`
	TransactionHash  string `json:"transactionHash"`
	Implementation   Party  `json:"implementation"`
	Factory          Party  `json:"factory"`
	SaltNonce        string `json:"saltNonce"`
}

type Party struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	LogoUri string `json:"logoUri"`
}

type ExecutionInfo struct {
	Type                   string  `json:"type"`
	Nonce                  int     `json:"nonce"`
	ConfirmationsRequired  int     `json:"confirmationsRequired"`
	ConfirmationsSubmitted int     `json:"confirmationsSubmitted"`
	MissingSigners         []Party `json:"missingSigners"`
}

type SafeAppInfo struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	LogoUri string `json:"logoUri"`
}

func GetProposals(apiURL string) ([]ProposalResponse, error) {
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}
	fmt.Println(baseURL.String())
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response struct {
		Count    int                `json:"count"`
		Next     *string            `json:"next"`
		Previous *string            `json:"previous"`
		Results  []ProposalResponse `json:"results"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}
	return response.Results, nil
}
