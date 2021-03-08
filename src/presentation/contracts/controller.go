package contracts

type Controller interface {
	Handle() *HttpResponse
}
