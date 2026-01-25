

package antifraud_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

type codeRecorder struct {
	http.ResponseWriter
	status int
}

func (c *codeRecorder) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}











func (s *Server) handleAPIV1AuthLoginPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/auth/login"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1AuthLoginPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1AuthLoginPostOperation,
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIV1AuthLoginPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1AuthLoginPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1AuthLoginPostOperation,
			OperationSummary: "Авторизация пользователя",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *LoginRequest
			Params   = struct{}
			Response = APIV1AuthLoginPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1AuthLoginPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1AuthLoginPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1AuthLoginPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}










func (s *Server) handleAPIV1AuthRegisterPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/auth/register"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1AuthRegisterPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1AuthRegisterPostOperation,
			ID:   "",
		}
	)
	request, close, err := s.decodeAPIV1AuthRegisterPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1AuthRegisterPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1AuthRegisterPostOperation,
			OperationSummary: "Регистрация нового пользователя",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *RegisterRequest
			Params   = struct{}
			Response = APIV1AuthRegisterPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1AuthRegisterPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1AuthRegisterPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1AuthRegisterPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}









func (s *Server) handleAPIV1FraudRulesGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1FraudRulesGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesGetOperation,
			OperationSummary: "Список правил фрода",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1FraudRulesGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesGet(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}





















func (s *Server) handleAPIV1FraudRulesIDDeleteRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesIDDeleteOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesIDDeleteOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesIDDeleteOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1FraudRulesIDDeleteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1FraudRulesIDDeleteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesIDDeleteOperation,
			OperationSummary: "Деактивация правила фрода",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1FraudRulesIDDeleteParams
			Response = APIV1FraudRulesIDDeleteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1FraudRulesIDDeleteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesIDDelete(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesIDDelete(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesIDDeleteResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
















func (s *Server) handleAPIV1FraudRulesIDGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesIDGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesIDGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesIDGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1FraudRulesIDGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1FraudRulesIDGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesIDGetOperation,
			OperationSummary: "Получить правило фрода по ID",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1FraudRulesIDGetParams
			Response = APIV1FraudRulesIDGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1FraudRulesIDGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesIDGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesIDGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesIDGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}








func (s *Server) handleAPIV1FraudRulesIDPutRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesIDPutOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesIDPutOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesIDPutOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1FraudRulesIDPutParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeAPIV1FraudRulesIDPutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1FraudRulesIDPutRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesIDPutOperation,
			OperationSummary: "Обновить правило фрода",
			OperationID:      "",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *FraudRuleUpdateRequest
			Params   = APIV1FraudRulesIDPutParams
			Response = APIV1FraudRulesIDPutRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1FraudRulesIDPutParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesIDPut(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesIDPut(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesIDPutResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}















func (s *Server) handleAPIV1FraudRulesPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesPostOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesPostOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1FraudRulesPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1FraudRulesPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesPostOperation,
			OperationSummary: "Создать правило фрода",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *FraudRuleCreateRequest
			Params   = struct{}
			Response = APIV1FraudRulesPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}









func (s *Server) handleAPIV1FraudRulesValidatePostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/fraud-rules/validate"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1FraudRulesValidatePostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1FraudRulesValidatePostOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1FraudRulesValidatePostOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1FraudRulesValidatePostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1FraudRulesValidatePostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1FraudRulesValidatePostOperation,
			OperationSummary: "Валидация DSL выражения (без сохранения правила)",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *DslValidateRequest
			Params   = struct{}
			Response = APIV1FraudRulesValidatePostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1FraudRulesValidatePost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1FraudRulesValidatePost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1FraudRulesValidatePostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}








func (s *Server) handleAPIV1PingGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/ping"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1PingGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err error
	)

	var response *APIV1PingGetOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1PingGetOperation,
			OperationSummary: "Health check",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *APIV1PingGetOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1PingGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1PingGet(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1PingGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}










func (s *Server) handleAPIV1StatsMerchantsRiskGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/stats/merchants/risk"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1StatsMerchantsRiskGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1StatsMerchantsRiskGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1StatsMerchantsRiskGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1StatsMerchantsRiskGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1StatsMerchantsRiskGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1StatsMerchantsRiskGetOperation,
			OperationSummary: "Риск-метрики по мерчантам",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "to",
					In:   "query",
				}: params.To,
				{
					Name: "merchantCategoryCode",
					In:   "query",
				}: params.MerchantCategoryCode,
				{
					Name: "top",
					In:   "query",
				}: params.Top,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1StatsMerchantsRiskGetParams
			Response = APIV1StatsMerchantsRiskGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1StatsMerchantsRiskGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1StatsMerchantsRiskGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1StatsMerchantsRiskGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1StatsMerchantsRiskGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}


















func (s *Server) handleAPIV1StatsOverviewGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/stats/overview"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1StatsOverviewGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1StatsOverviewGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1StatsOverviewGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1StatsOverviewGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1StatsOverviewGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1StatsOverviewGetOperation,
			OperationSummary: "Обзор ключевых метрик (дашборд)",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "to",
					In:   "query",
				}: params.To,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1StatsOverviewGetParams
			Response = APIV1StatsOverviewGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1StatsOverviewGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1StatsOverviewGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1StatsOverviewGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1StatsOverviewGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}



















func (s *Server) handleAPIV1StatsRulesMatchesGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/stats/rules/matches"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1StatsRulesMatchesGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1StatsRulesMatchesGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1StatsRulesMatchesGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1StatsRulesMatchesGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1StatsRulesMatchesGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1StatsRulesMatchesGetOperation,
			OperationSummary: "Статистика срабатываний правил",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "to",
					In:   "query",
				}: params.To,
				{
					Name: "top",
					In:   "query",
				}: params.Top,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1StatsRulesMatchesGetParams
			Response = APIV1StatsRulesMatchesGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1StatsRulesMatchesGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1StatsRulesMatchesGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1StatsRulesMatchesGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1StatsRulesMatchesGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}




















func (s *Server) handleAPIV1StatsTransactionsTimeseriesGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/stats/transactions/timeseries"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1StatsTransactionsTimeseriesGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1StatsTransactionsTimeseriesGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1StatsTransactionsTimeseriesGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1StatsTransactionsTimeseriesGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1StatsTransactionsTimeseriesGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1StatsTransactionsTimeseriesGetOperation,
			OperationSummary: "Таймсерия по транзакциям",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "to",
					In:   "query",
				}: params.To,
				{
					Name: "groupBy",
					In:   "query",
				}: params.GroupBy,
				{
					Name: "timezone",
					In:   "query",
				}: params.Timezone,
				{
					Name: "channel",
					In:   "query",
				}: params.Channel,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1StatsTransactionsTimeseriesGetParams
			Response = APIV1StatsTransactionsTimeseriesGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1StatsTransactionsTimeseriesGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1StatsTransactionsTimeseriesGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1StatsTransactionsTimeseriesGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1StatsTransactionsTimeseriesGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}













func (s *Server) handleAPIV1StatsUsersIDRiskProfileGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/stats/users/{id}/risk-profile"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1StatsUsersIDRiskProfileGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1StatsUsersIDRiskProfileGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1StatsUsersIDRiskProfileGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1StatsUsersIDRiskProfileGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1StatsUsersIDRiskProfileGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1StatsUsersIDRiskProfileGetOperation,
			OperationSummary: "Риск-профиль пользователя (в т.ч. для USER — только свой)",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1StatsUsersIDRiskProfileGetParams
			Response = APIV1StatsUsersIDRiskProfileGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1StatsUsersIDRiskProfileGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1StatsUsersIDRiskProfileGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1StatsUsersIDRiskProfileGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1StatsUsersIDRiskProfileGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}












func (s *Server) handleAPIV1TransactionsBatchPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/transactions/batch"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1TransactionsBatchPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1TransactionsBatchPostOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1TransactionsBatchPostOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1TransactionsBatchPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1TransactionsBatchPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1TransactionsBatchPostOperation,
			OperationSummary: "Создать батч транзакций",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *TransactionBatchCreateRequest
			Params   = struct{}
			Response = APIV1TransactionsBatchPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1TransactionsBatchPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1TransactionsBatchPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1TransactionsBatchPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}













func (s *Server) handleAPIV1TransactionsGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/transactions"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1TransactionsGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1TransactionsGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1TransactionsGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1TransactionsGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1TransactionsGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1TransactionsGetOperation,
			OperationSummary: "Список транзакций",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "userId",
					In:   "query",
				}: params.UserId,
				{
					Name: "status",
					In:   "query",
				}: params.Status,
				{
					Name: "isFraud",
					In:   "query",
				}: params.IsFraud,
				{
					Name: "from",
					In:   "query",
				}: params.From,
				{
					Name: "to",
					In:   "query",
				}: params.To,
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "size",
					In:   "query",
				}: params.Size,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1TransactionsGetParams
			Response = APIV1TransactionsGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1TransactionsGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1TransactionsGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1TransactionsGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1TransactionsGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}







func (s *Server) handleAPIV1TransactionsIDGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/transactions/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1TransactionsIDGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1TransactionsIDGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1TransactionsIDGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1TransactionsIDGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1TransactionsIDGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1TransactionsIDGetOperation,
			OperationSummary: "Получить транзакцию",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1TransactionsIDGetParams
			Response = APIV1TransactionsIDGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1TransactionsIDGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1TransactionsIDGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1TransactionsIDGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1TransactionsIDGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
































func (s *Server) handleAPIV1TransactionsPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/transactions"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1TransactionsPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1TransactionsPostOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1TransactionsPostOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1TransactionsPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1TransactionsPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1TransactionsPostOperation,
			OperationSummary: "Создать транзакцию и выполнить антифрод-проверку",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *TransactionCreateRequest
			Params   = struct{}
			Response = APIV1TransactionsPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1TransactionsPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1TransactionsPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1TransactionsPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}









func (s *Server) handleAPIV1UsersGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/users"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1UsersGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1UsersGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersGetOperation,
			OperationSummary: "Список пользователей",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "size",
					In:   "query",
				}: params.Size,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1UsersGetParams
			Response = APIV1UsersGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1UsersGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}




















func (s *Server) handleAPIV1UsersIDDeleteRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/api/v1/users/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersIDDeleteOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersIDDeleteOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersIDDeleteOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1UsersIDDeleteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1UsersIDDeleteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersIDDeleteOperation,
			OperationSummary: "Деактивация пользователя",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1UsersIDDeleteParams
			Response = APIV1UsersIDDeleteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1UsersIDDeleteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersIDDelete(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersIDDelete(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersIDDeleteResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}












func (s *Server) handleAPIV1UsersIDGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/users/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersIDGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersIDGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersIDGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1UsersIDGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response APIV1UsersIDGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersIDGetOperation,
			OperationSummary: "Получение профиля пользователя",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = APIV1UsersIDGetParams
			Response = APIV1UsersIDGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1UsersIDGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersIDGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersIDGet(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersIDGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}























func (s *Server) handleAPIV1UsersIDPutRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/users/{id}"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersIDPutOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersIDPutOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersIDPutOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeAPIV1UsersIDPutParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeAPIV1UsersIDPutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1UsersIDPutRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersIDPutOperation,
			OperationSummary: "Полное обновление профиля пользователя",
			OperationID:      "",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "path",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = *UserUpdateRequest
			Params   = APIV1UsersIDPutParams
			Response = APIV1UsersIDPutRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackAPIV1UsersIDPutParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersIDPut(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersIDPut(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersIDPutResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}










func (s *Server) handleAPIV1UsersMeGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/users/me"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersMeGetOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersMeGetOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersMeGetOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response APIV1UsersMeGetRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersMeGetOperation,
			OperationSummary: "Получить профиль текущего пользователя",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = APIV1UsersMeGetRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersMeGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersMeGet(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersMeGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
















func (s *Server) handleAPIV1UsersMePutRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/users/me"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersMePutOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersMePutOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersMePutOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1UsersMePutRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1UsersMePutRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersMePutOperation,
			OperationSummary: "Обновить профиль текущего пользователя",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UserUpdateRequest
			Params   = struct{}
			Response = APIV1UsersMePutRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersMePut(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersMePut(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersMePutResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}











func (s *Server) handleAPIV1UsersPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	statusWriter := &codeRecorder{ResponseWriter: w}
	w = statusWriter
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/users"),
	}

	
	ctx, span := s.cfg.Tracer.Start(r.Context(), APIV1UsersPostOperation,
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)

		attrSet := labeler.AttributeSet()
		attrs := attrSet.ToSlice()
		code := statusWriter.status
		if code != 0 {
			codeAttr := semconv.HTTPResponseStatusCode(code)
			attrs = append(attrs, codeAttr)
			span.SetAttributes(codeAttr)
		}
		attrOpt := metric.WithAttributes(attrs...)

		
		s.requests.Add(ctx, 1, attrOpt)

		
		s.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)

			
			
			
			
			code := statusWriter.status
			if code >= 100 && code < 500 {
				span.SetStatus(codes.Error, stage)
			}

			attrSet := labeler.AttributeSet()
			attrs := attrSet.ToSlice()
			if code != 0 {
				attrs = append(attrs, semconv.HTTPResponseStatusCode(code))
			}

			s.errors.Add(ctx, 1, metric.WithAttributes(attrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: APIV1UsersPostOperation,
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearerAuth(ctx, APIV1UsersPostOperation, r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerAuth",
					Err:              err,
				}
				defer recordError("Security:BearerAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			defer recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeAPIV1UsersPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response APIV1UsersPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    APIV1UsersPostOperation,
			OperationSummary: "Создание пользователя админом",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *UserCreateRequest
			Params   = struct{}
			Response = APIV1UsersPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.APIV1UsersPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.APIV1UsersPost(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIV1UsersPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
