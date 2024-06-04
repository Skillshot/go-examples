package server

import (
	"fmt"
)

type Server struct {
	logger Logger
}

func New(l Logger) Server {
	return Server{
		logger: l,
	}
}

func (s Server) RenderAdminPage(user string) string {
	if user != "admin" {
		s.logger.Error(fmt.Sprintf("User %s tried to access the admin page", user))
		return "Access Denied"
	}
	s.logger.Info(fmt.Sprintf("User %s accessed the admin page", user))
	return "Admin Page"
}
