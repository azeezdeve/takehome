package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"trustwallet/internal/model"
)

type (
	Eth struct {
	}

	req struct {
		JsonRPC string `json:"jsonrpc"`
		Method  string `json:"method"`
		ID      int    `json:"id"`
	}

	resp struct {
		JsonRPC string `json:"json"`
		Result  string `json:"result"`
		ID      int    `json:"id"`
	}
)

func New() Eth {
	return Eth{}
}

// GetCurrentBlock: This is for test ask
func (e Eth) GetCurrentBlock(ctx context.Context) (int, error) {
	var result resp
	body := req{
		JsonRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      83,
	}

	bodyInbyte, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	url := "https://cloudflare-eth.com"
	reqBody, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyInbyte))
	client := &http.Client{}
	respBody, err := client.Do(reqBody)
	if err != nil {
		return 0, err
	}
	defer respBody.Body.Close()

	ioReader, err := ioutil.ReadAll(respBody.Body)
	if err != nil {
		return 0, err
	}

	if respBody.StatusCode >= 400 {
		return 0, errors.New(string(ioReader))
	}

	err = json.Unmarshal(ioReader, &result)
	if err != nil {
		return 0, err
	}

	response, err := strconv.ParseInt(result.Result, 0, 0)
	if err != nil {
		return 0, err
	}

	return int(response), nil
}

var _ model.IBlock = Eth{}
