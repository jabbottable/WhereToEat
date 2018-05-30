package store

type JwtToken struct {
    Token string `json:"token"`
}

type Exception struct {
    Message string `json:"message"`
}

type Location struct {
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}
