package usecases

import (
	"context"
	"testing"

	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/services"
	"github.com/arfurlaneto/goexpert-challenge-temperature-by-cep/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type GetTemperatureByCepUseCaseTestSuite struct {
	suite.Suite
	controller            *gomock.Controller
	viaCepServiceMock     *mocks.MockViaCepService
	weatherApiServiceMock *mocks.MockWeatherApiService
}

func TestGetTemperatureByCepUseCase(t *testing.T) {
	suite.Run(t, new(GetTemperatureByCepUseCaseTestSuite))
}

func (s *GetTemperatureByCepUseCaseTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.viaCepServiceMock = mocks.NewMockViaCepService(s.controller)
	s.weatherApiServiceMock = mocks.NewMockWeatherApiService(s.controller)
}

func (s *GetTemperatureByCepUseCaseTestSuite) TestNewGetTemperatureByCepHandler() {
	usecase := NewGetTemperatureByCepUseCase(s.viaCepServiceMock, s.weatherApiServiceMock)
	assert.NotNil(s.T(), usecase)
}

func (s *GetTemperatureByCepUseCaseTestSuite) TestExecuteSuccess() {

	ctx := context.Background()
	validCep := "01451-000"

	s.viaCepServiceMock.EXPECT().QueryCep(ctx, validCep).Return(
		&services.ViaCepResponse{
			Cep:         "01451-000",
			Logradouro:  "Avenida Brigadeiro Faria Lima",
			Complemento: "de 1884 a 3250 - lado par",
			Bairro:      "Jardim Paulistano",
			Localidade:  "São Paulo",
			UF:          "SP",
			IBGE:        "3550308",
			GIA:         "1004",
			DDD:         "11",
			SIAFI:       "7107",
		},
		nil,
	)

	s.weatherApiServiceMock.EXPECT().QueryWeather(ctx, "Brazil - São Paulo - São Paulo").AnyTimes().Return(
		&services.WeatherApiResponse{
			Location: services.WeatherApiResponseLocation{
				Name:    "São Paulo",
				Region:  "São Paulo",
				Country: "Brazil",
			},
			Current: services.WeatherApiResponseCurrent{
				TemperatureCelsius: 24,
			},
		},
		nil,
	)

	usecase := NewGetTemperatureByCepUseCase(s.viaCepServiceMock, s.weatherApiServiceMock)
	output, err := usecase.Execute(ctx, &TemperatureByCepInput{Cep: validCep})

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 24.0, output.TemperatureCelsius)
	assert.Equal(s.T(), 75.2, output.TemperatureFahrenheit)
	assert.Equal(s.T(), 297.15, output.TemperatureKelvin)
}
