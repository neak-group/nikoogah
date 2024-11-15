package auth

import "github.com/neak-group/nikoogah/internal/services/core/security/session"

func NewAuthenticator(sessionSerivce *session.SessionService) *Authenticator {
	return &Authenticator{
		sessionService:  *sessionSerivce,
		anonymousRoutes: make(map[string][]string),
	}
}

func (r *Authenticator) AddAnonymousRoute(method, route string) {
	if r.anonymousRoutes[route] == nil {
		r.anonymousRoutes[route] = make([]string, 0)
	}

	r.anonymousRoutes[route] = append(r.anonymousRoutes[route], method)
}
