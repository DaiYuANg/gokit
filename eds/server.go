package eds

import (
	"codeberg.org/miekg/dns"
)

type Server struct {
	port      int16
	add       string
	dnsServer *dns.Server
}

func (s *Server) Start() error {
	return s.dnsServer.ListenAndServe()
}
