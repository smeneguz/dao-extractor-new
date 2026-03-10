package etherscan

import "fmt"

type EtherscanResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func (r EtherscanResponse) IsOK() bool {
	return r.Status == "1" && r.Message == "OK"
}

func (r EtherscanResponse) Error() error {
	if r.Status != "1" {
		return fmt.Errorf("%s, status: %s", r.Result, r.Status)
	}

	if r.Message != "OK" {
		return fmt.Errorf("%s, message: %s", r.Result, r.Message)
	}

	return nil
}
