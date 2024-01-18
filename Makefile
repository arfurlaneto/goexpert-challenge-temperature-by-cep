generate-mocks:
	mockgen -source=./internal/usecases/get_temperature_by_cep.go -destination ./internal/usecases/mocks/get_temperature_by_cep.go -package mocks
	mockgen -source=./internal/services/viacep.go -destination ./internal/services/mocks/viacep.go -package mocks
	mockgen -source=./internal/services/weatherapi.go -destination ./internal/services/mocks/weatherapi.go -package mocks

test:
	go test ./... -v
