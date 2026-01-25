

package antifraud_v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)


type SecurityHandler interface {
	
	HandleBearerAuth(ctx context.Context, operationName OperationName, t BearerAuth) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityBearerAuth(ctx context.Context, operationName OperationName, req *http.Request) (context.Context, bool, error) {
	var t BearerAuth
	token, ok := findAuthorization(req.Header, "Bearer")
	if !ok {
		return ctx, false, nil
	}
	t.Token = token
	rctx, err := s.sec.HandleBearerAuth(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}


type SecuritySource interface {
	
	BearerAuth(ctx context.Context, operationName OperationName) (BearerAuth, error)
}

func (s *Client) securityBearerAuth(ctx context.Context, operationName OperationName, req *http.Request) error {
	t, err := s.sec.BearerAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"BearerAuth\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
