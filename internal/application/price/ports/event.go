package ports

type PriceEventyIF interface {
	SendPriceEvent(message []byte)
	Close()
}
