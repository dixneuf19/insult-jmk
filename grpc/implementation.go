package insultJMK

import (
	"log"
	"strings"

	"github.com/dixneuf19/insult-jmk/insulter"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
	Insulter *insulter.Insults
}

// SayHello generates response to a Ping request
func (s *Server) GetInsult(ctx context.Context, in *InsultRequest) (*InsultResponse, error) {
	message := in.GetName()
	if message != "" {
		message += ", "
	}
	insult := strings.Trim(s.Insulter.GenerateInsult(), " ")
	message += insult
	log.Printf("Generate '%s'", message)

	return &InsultResponse{Message: message}, nil
}
