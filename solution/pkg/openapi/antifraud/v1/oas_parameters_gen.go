

package antifraud_v1

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)


type APIV1FraudRulesIDDeleteParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1FraudRulesIDDeleteParams(packed middleware.Parameters) (params APIV1FraudRulesIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1FraudRulesIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1FraudRulesIDDeleteParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1FraudRulesIDGetParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1FraudRulesIDGetParams(packed middleware.Parameters) (params APIV1FraudRulesIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1FraudRulesIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1FraudRulesIDGetParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1FraudRulesIDPutParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1FraudRulesIDPutParams(packed middleware.Parameters) (params APIV1FraudRulesIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1FraudRulesIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1FraudRulesIDPutParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1StatsMerchantsRiskGetParams struct {
	
	
	
	
	
	
	
	
	
	
	From OptDateTime
	
	
	
	
	
	To                   OptDateTime
	MerchantCategoryCode OptString
	Top                  OptInt
}

func unpackAPIV1StatsMerchantsRiskGetParams(packed middleware.Parameters) (params APIV1StatsMerchantsRiskGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.To = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "merchantCategoryCode",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.MerchantCategoryCode = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "top",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Top = v.(OptInt)
		}
	}
	return params
}

func decodeAPIV1StatsMerchantsRiskGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1StatsMerchantsRiskGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.To.SetTo(paramsDotToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "merchantCategoryCode",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotMerchantCategoryCodeVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotMerchantCategoryCodeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.MerchantCategoryCode.SetTo(paramsDotMerchantCategoryCodeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.MerchantCategoryCode.Get(); ok {
					if err := func() error {
						if err := (validate.String{
							MinLength:    0,
							MinLengthSet: false,
							MaxLength:    0,
							MaxLengthSet: false,
							Email:        false,
							Hostname:     false,
							Regex:        regexMap["^\\d{4}$"],
						}).Validate(string(value)); err != nil {
							return errors.Wrap(err, "string")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "merchantCategoryCode",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := int(50)
		params.Top.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "top",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTopVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotTopVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Top.SetTo(paramsDotTopVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Top.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           200,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "top",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1StatsOverviewGetParams struct {
	
	
	
	
	
	
	
	
	
	
	From OptDateTime
	
	
	
	
	
	To OptDateTime
}

func unpackAPIV1StatsOverviewGetParams(packed middleware.Parameters) (params APIV1StatsOverviewGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.To = v.(OptDateTime)
		}
	}
	return params
}

func decodeAPIV1StatsOverviewGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1StatsOverviewGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.To.SetTo(paramsDotToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1StatsRulesMatchesGetParams struct {
	
	
	
	
	
	
	
	
	
	
	From OptDateTime
	
	
	
	
	
	To OptDateTime
	
	Top OptInt
}

func unpackAPIV1StatsRulesMatchesGetParams(packed middleware.Parameters) (params APIV1StatsRulesMatchesGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.To = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "top",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Top = v.(OptInt)
		}
	}
	return params
}

func decodeAPIV1StatsRulesMatchesGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1StatsRulesMatchesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.To.SetTo(paramsDotToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := int(20)
		params.Top.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "top",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTopVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotTopVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Top.SetTo(paramsDotTopVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Top.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           100,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "top",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1StatsTransactionsTimeseriesGetParams struct {
	
	
	
	
	
	
	
	
	
	
	From OptDateTime
	
	
	
	
	
	To      OptDateTime
	GroupBy OptGroupBy
	
	Timezone OptString
	
	Channel OptTransactionChannel
}

func unpackAPIV1StatsTransactionsTimeseriesGetParams(packed middleware.Parameters) (params APIV1StatsTransactionsTimeseriesGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.To = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "groupBy",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.GroupBy = v.(OptGroupBy)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "timezone",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Timezone = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "channel",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Channel = v.(OptTransactionChannel)
		}
	}
	return params
}

func decodeAPIV1StatsTransactionsTimeseriesGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1StatsTransactionsTimeseriesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.To.SetTo(paramsDotToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := GroupBy("day")
		params.GroupBy.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "groupBy",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotGroupByVal GroupBy
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotGroupByVal = GroupBy(c)
					return nil
				}(); err != nil {
					return err
				}
				params.GroupBy.SetTo(paramsDotGroupByVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.GroupBy.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "groupBy",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := string("UTC")
		params.Timezone.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "timezone",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTimezoneVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotTimezoneVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Timezone.SetTo(paramsDotTimezoneVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "timezone",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "channel",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChannelVal TransactionChannel
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotChannelVal = TransactionChannel(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Channel.SetTo(paramsDotChannelVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Channel.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "channel",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1StatsUsersIDRiskProfileGetParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1StatsUsersIDRiskProfileGetParams(packed middleware.Parameters) (params APIV1StatsUsersIDRiskProfileGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1StatsUsersIDRiskProfileGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1StatsUsersIDRiskProfileGetParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1TransactionsGetParams struct {
	
	
	UserId  OptUUID
	Status  OptTransactionStatus
	IsFraud OptBool
	
	
	
	
	
	
	
	
	
	
	From OptDateTime
	
	
	
	
	
	To OptDateTime
	
	Page OptInt
	Size OptInt
}

func unpackAPIV1TransactionsGetParams(packed middleware.Parameters) (params APIV1TransactionsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UserId = v.(OptUUID)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "status",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Status = v.(OptTransactionStatus)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "isFraud",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.IsFraud = v.(OptBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.From = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.To = v.(OptDateTime)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "size",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Size = v.(OptInt)
		}
	}
	return params
}

func decodeAPIV1TransactionsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1TransactionsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "userId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserIdVal uuid.UUID
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToUUID(val)
					if err != nil {
						return err
					}

					paramsDotUserIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserId.SetTo(paramsDotUserIdVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStatusVal TransactionStatus
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStatusVal = TransactionStatus(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Status.SetTo(paramsDotStatusVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Status.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "status",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "isFraud",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIsFraudVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotIsFraudVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IsFraud.SetTo(paramsDotIsFraudVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "isFraud",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFromVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.From.SetTo(paramsDotFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "from",
			In:   "query",
			Err:  err,
		}
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotToVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDateTime(val)
					if err != nil {
						return err
					}

					paramsDotToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.To.SetTo(paramsDotToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "to",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := int(20)
		params.Size.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSizeVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotSizeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Size.SetTo(paramsDotSizeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Size.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           100,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "size",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1TransactionsIDGetParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1TransactionsIDGetParams(packed middleware.Parameters) (params APIV1TransactionsIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1TransactionsIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1TransactionsIDGetParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1UsersGetParams struct {
	
	Page OptInt
	Size OptInt
}

func unpackAPIV1UsersGetParams(packed middleware.Parameters) (params APIV1UsersGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "size",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Size = v.(OptInt)
		}
	}
	return params
}

func decodeAPIV1UsersGetParams(args [0]string, argsEscaped bool, r *http.Request) (params APIV1UsersGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	
	{
		val := int(0)
		params.Page.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	
	{
		val := int(20)
		params.Size.SetTo(val)
	}
	
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSizeVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotSizeVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Size.SetTo(paramsDotSizeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Size.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           100,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "size",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1UsersIDDeleteParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1UsersIDDeleteParams(packed middleware.Parameters) (params APIV1UsersIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1UsersIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1UsersIDDeleteParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1UsersIDGetParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1UsersIDGetParams(packed middleware.Parameters) (params APIV1UsersIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1UsersIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1UsersIDGetParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}


type APIV1UsersIDPutParams struct {
	
	ID uuid.UUID
}

func unpackAPIV1UsersIDPutParams(packed middleware.Parameters) (params APIV1UsersIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeAPIV1UsersIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params APIV1UsersIDPutParams, _ error) {
	
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
