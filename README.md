# What is this?

This a service that returns temperatures (celsius, fahrenheit and kelvin) for a Brazilian CEP.

It is also one of the challenges of https://goexpert.fullcycle.com.br/pos-goexpert/.

# How to run it?

Run with docker compose:

```bash
docker compose up
```

And do a request, something like this:

```bash
curl --request GET --url 'http://localhost:8080/?cep=69400970'
```

# How to use in production?

Edit the `.env` file and add you Weather API Key (https://www.weatherapi.com/).

The docker image require no external dependencies, so it runs anywhere docker runs.

# Google Cloud Run

This service is also available online (only until challenge validation). Some URLs:

```
# invalid zip code
https://goexpert-challenge-temperature-by-cep-npf4263taq-uc.a.run.app/?cep=1

# zip code not found
https://goexpert-challenge-temperature-by-cep-npf4263taq-uc.a.run.app/?cep=11111111

# success
https://goexpert-challenge-temperature-by-cep-npf4263taq-uc.a.run.app/?cep=69400970
```
