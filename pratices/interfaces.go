package pratices

//Prefer small, focused interfaces (io.Reader, io.Writer).
//Avoid over-engineering with unnecessary interfaces.

type Fetcher interface {
	Fetch(url string) ([]byte, error)
}
