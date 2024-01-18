package usecases

import (
	"context"
	"fmt"

	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/services"
	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/utils"
)

type TemperatureByCepInput struct {
	Cep string
}

type TemperatureByCepOutput struct {
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

type GetTemperatureByCepUseCase interface {
	Execute(ctx context.Context, input *TemperatureByCepInput) (*TemperatureByCepOutput, error)
}

type GetTemperatureByCepUseCaseImpl struct {
	cepService     services.ViaCepService
	weatherService services.WeatherApiService
}

func NewGetTemperatureByCepUseCase(cepService services.ViaCepService, weatherService services.WeatherApiService) GetTemperatureByCepUseCase {
	return &GetTemperatureByCepUseCaseImpl{
		cepService:     cepService,
		weatherService: weatherService,
	}
}

func (u *GetTemperatureByCepUseCaseImpl) Execute(ctx context.Context, input *TemperatureByCepInput) (*TemperatureByCepOutput, error) {

	cepResponse, err := u.cepService.QueryCep(ctx, input.Cep)
	if err != nil {
		return nil, err
	}

	location := fmt.Sprintf("Brazil - %s - %s", utils.UfToStateNameMap[cepResponse.UF], cepResponse.Localidade)

	weatherResponse, err := u.weatherService.QueryWeather(ctx, location)
	if err != nil {
		return nil, err
	}

	temperatureCelsius := weatherResponse.Current.TemperatureCelsius
	temperatureFahrenheit := weatherResponse.Current.TemperatureCelsius*1.8 + 32
	temperatureKelvin := weatherResponse.Current.TemperatureCelsius + 273.15

	weatherLocation := fmt.Sprintf("%s, %s, %s", weatherResponse.Location.Name, weatherResponse.Location.Region, weatherResponse.Location.Country)
	fmt.Printf("Asked for %s => Got %s C=%f F=%f K=%f\n", location, weatherLocation, temperatureCelsius, temperatureFahrenheit, temperatureKelvin)

	return &TemperatureByCepOutput{
		TemperatureCelsius:    temperatureCelsius,
		TemperatureFahrenheit: temperatureFahrenheit,
		TemperatureKelvin:     temperatureKelvin,
	}, nil
}
