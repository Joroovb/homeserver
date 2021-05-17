package sonarr

type sonarr struct {
	host string
	port int
	key  string
	protocol string
}

func New(host string, port int, key string, protocol string) sonarr {
	return sonarr{
		host,
		port,
		key,
		protocol,
	}
}

func (s sonarr) GetSeries() {

}

