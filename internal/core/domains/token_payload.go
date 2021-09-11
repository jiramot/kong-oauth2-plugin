package domains

type TokenPayload struct {
    Issue     string    `json:"iss"`
    Subject   string    `json:"sub"`
    Scope     string    `json:"scope"`
    Amr       [1]string `json:"amr"`
    IssuedAt  int64     `json:"iat"`
    ExpiredAt int64     `json:"exp"`
    ClientId  string    `json:"aud"`
}
