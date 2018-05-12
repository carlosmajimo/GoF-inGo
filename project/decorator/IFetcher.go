package decorator

// Args son los argumentos de la funcion Fetch
type Args map[string]string

// Data es lo que retorna Fetch
type Data map[string]string

// IFetcher obtiene informacion
type IFetcher interface {
	Fetch(args Args) (Data, string)
}
