package response

// GenericResponse is a reusable response model.
type GenericResponse interface {
	ParseXML(data []byte) error
}
