package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/usecases"
	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/utils"
)

type GetTemperatureByCepResponse struct {
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

type GetTemperatureByCepHandler struct {
	getTemperatureByCepUseCase usecases.GetTemperatureByCepUseCase
}

func NewGetTemperatureByCepHandler(getTemperatureByCepUseCase usecases.GetTemperatureByCepUseCase) *GetTemperatureByCepHandler {
	return &GetTemperatureByCepHandler{
		getTemperatureByCepUseCase: getTemperatureByCepUseCase,
	}
}

func (h *GetTemperatureByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep, ok := h.getCepFromRequest(r)
	if !ok {
		w.WriteHeader(422)
		w.Write([]byte("invalid zipcode"))
		return
	}

	input := &usecases.TemperatureByCepInput{Cep: cep}
	output, err := h.getTemperatureByCepUseCase.Execute(r.Context(), input)

	if err != nil {
		if err.Error() == "can not found zipcode" {
			w.WriteHeader(404)
			w.Write([]byte("can not found zipcode"))
		} else {
			fmt.Printf("ERROR: %s\n", err.Error())
			w.WriteHeader(500)
			w.Write([]byte("internal server error"))
		}
		return
	}

	response := GetTemperatureByCepResponse{
		TemperatureCelsius:    utils.RoundFloat(output.TemperatureCelsius, 1),
		TemperatureFahrenheit: utils.RoundFloat(output.TemperatureFahrenheit, 1),
		TemperatureKelvin:     utils.RoundFloat(output.TemperatureKelvin, 1),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *GetTemperatureByCepHandler) getCepFromRequest(r *http.Request) (string, bool) {
	cep := r.URL.Query().Get("cep")

	if cep == "" {
		return "", false
	}

	cepRegex := regexp.MustCompile(`^\d{5}-{0,1}\d{3}$`)
	if !cepRegex.Match([]byte(cep)) {
		return "", false
	}

	return cep, true
}
