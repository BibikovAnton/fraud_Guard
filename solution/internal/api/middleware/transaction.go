package middleware

import (
	"context"
	"net/http"
	"solution/internal/api/antifraud/v1"
	"solution/internal/service"
)

type TransactionMiddleware struct {
	rawHandler *v1.RawTransactionHandler
}

func NewTransactionMiddleware(transactionService service.TransactionService) *TransactionMiddleware {
	return &TransactionMiddleware{
		rawHandler: v1.NewRawTransactionHandler(transactionService),
	}
}

func (m *TransactionMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "POST" && r.URL.Path == "/api/v1/transactions") ||
			(r.Method == "POST" && r.URL.Path == "/api/v1/transactions/batch") {
			
			ctx := r.Context()
			
			ctx = context.WithValue(ctx, "user_role", "USER")
			ctx = context.WithValue(ctx, "user_id", "demo-user-id")
			
			r = r.WithContext(ctx)
			
			if r.URL.Path == "/api/v1/transactions" {
				m.rawHandler.CreateTransactionRaw(w, r)
				return
			}
			if r.URL.Path == "/api/v1/transactions/batch" {
				m.rawHandler.CreateBatchTransactionRaw(w, r)
				return
			}
		}
		
		next.ServeHTTP(w, r)
	})
}
