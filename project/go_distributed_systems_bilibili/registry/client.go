package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	rsp, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry service "+"responded with code %v", rsp.StatusCode)
	}

	return nil
}
