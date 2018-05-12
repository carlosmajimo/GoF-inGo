package decorator

import (
	"time"
)

// Retrier reintenta varias veces
type Retrier struct {
	RetryCount   int
	WaitInterval time.Duration
	Fetcher      IFetcher
}

func (r *Retrier) Fetch(args Args) (Data, string) {
	for retry := 1; retry <= r.RetryCount; retry++ {
		//fmt.Printf("Retrier busqueda para  %d\n", retry)
		if data, err := r.Fetcher.Fetch(args); err == "" {
			//fmt.Printf("Se ha obtenido para %d\n", retry)
			return data, ""
		} else if retry == r.RetryCount {
			//fmt.Printf("Retrier falló la busqueda %d veces\n", retry)
			return Data{}, err
		}
		//fmt.Printf("Retrier en el intervalo de espera %v\n", r.WaitInterval)
		time.Sleep(r.WaitInterval)
	}

	return Data{}, ""
}
