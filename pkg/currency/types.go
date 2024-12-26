package currency

type Response struct {
    Date string
    Rate float64
}

type Converter struct {
    BaseURL string
}
