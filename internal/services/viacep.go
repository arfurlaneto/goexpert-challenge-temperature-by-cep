package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ViaCepResponse struct {
	Erro        bool   `json:"erro"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type ViaCepService interface {
	QueryCep(ctx context.Context, cep string) (*ViaCepResponse, error)
}

type ViaCepServiceImpl struct {
	client *http.Client
}

func NewViaCepService() ViaCepService {
	return &ViaCepServiceImpl{
		client: &http.Client{},
	}
}

func (s *ViaCepServiceImpl) QueryCep(ctx context.Context, cep string) (*ViaCepResponse, error) {

	cep = strings.ReplaceAll(cep, "-", "")
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", strings.ReplaceAll(cep, "-", ""))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("ViaCEP API error")
	}

	viaCepResponse := ViaCepResponse{}
	err = json.Unmarshal([]byte(body), &viaCepResponse)
	if err != nil {
		return nil, errors.New("invalid ViaCEP API response")
	}

	if viaCepResponse.Erro {
		return nil, errors.New("can not found zipcode")
	}

	return &viaCepResponse, nil
}
