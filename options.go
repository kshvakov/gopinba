package gopinba

// Options for construct new Pinba access object (see New())
type Options struct {
	PinbaHost  string // default 127.0.0.1
	PinbaPort  int    // default 30002
	ServerName string // default unknown
}
