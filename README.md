# wifi-radar-go

### API endpoints

- GET api/map/search?city={value}
- GET api/map/search?city={value}&radius={value}
- GET api/map/coordinates?city={value}&street={value}&buildingNumber={value}
- POST api/user/{userId}/favorites (wifi object)
- GET api/user/{userId}/favorites
- DELETE api/user/{userId}/favorites/{wifiId}
- POST api/wifi/reviews (per body perduot WifiReviewDto)
- GET api/wifi/reviews?city={value}&street={value}&buildingNumber={value}
- POST /api/wifi/passwords (body - PasswordDto)
- GET api/wifi/passwords?city={value}&street={value}&buildingNumber={value}
- POST api/user/auth/login (body - username, email)
- POST api/user/auth/register (body - username, email)
