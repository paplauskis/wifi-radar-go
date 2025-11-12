# wifi-radar-go

### How to run

- ```git clone https://github.com/paplauskis/wifi-radar-go.git```
- set up ```.env``` file in ```internal/config/```
- run postgres docker container
- ```cd wifi-radar-go```
- run ```go mod tidy``` in root folder to download all dependencies
- run the app ```go run cmd/server/main.go```

### API endpoints

- GET api/map/search?city={value}
- GET api/map/search?city={value}&radius={value}
- GET api/map/coordinates?city={value}&street={value}&buildingNumber={value}
- POST api/User/{userId}/favorites (wifi object)
- GET api/User/{userId}/favorites
- DELETE api/User/{userId}/favorites/{wifiId}
- POST api/wifi/reviews (per body perduot WifiReviewDto)
- GET api/wifi/reviews?city={value}&street={value}&buildingNumber={value}
- POST /api/wifi/passwords (body - PasswordDto)
- GET api/wifi/passwords?city={value}&street={value}&buildingNumber={value}
- POST api/User/auth/login (body - username, email)
- POST api/User/auth/register (body - username, email)
