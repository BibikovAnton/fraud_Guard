
package antifraud_v1

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

func (s *APIV1AuthLoginPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1AuthLoginPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1AuthLoginPostBadRequest to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1AuthLoginPostBadRequest(unwrapped)
	return nil
}

func (s *APIV1AuthLoginPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1AuthLoginPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1AuthLoginPostLocked) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1AuthLoginPostLocked) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1AuthLoginPostLocked to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1AuthLoginPostLocked(unwrapped)
	return nil
}

func (s *APIV1AuthLoginPostLocked) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1AuthLoginPostLocked) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1AuthLoginPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1AuthLoginPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1AuthLoginPostUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1AuthLoginPostUnauthorized(unwrapped)
	return nil
}

func (s *APIV1AuthLoginPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1AuthLoginPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1AuthRegisterPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1AuthRegisterPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1AuthRegisterPostBadRequest to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1AuthRegisterPostBadRequest(unwrapped)
	return nil
}

func (s *APIV1AuthRegisterPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1AuthRegisterPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1AuthRegisterPostConflict) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1AuthRegisterPostConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1AuthRegisterPostConflict to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1AuthRegisterPostConflict(unwrapped)
	return nil
}

func (s *APIV1AuthRegisterPostConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1AuthRegisterPostConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesGetForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s APIV1FraudRulesGetOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []FraudRule(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

func (s *APIV1FraudRulesGetOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesGetOKApplicationJSON to nil")
	}
	var unwrapped []FraudRule
	if err := func() error {
		unwrapped = make([]FraudRule, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem FraudRule
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesGetOKApplicationJSON(unwrapped)
	return nil
}

func (s APIV1FraudRulesGetOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesGetOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDDeleteForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDDeleteForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDDeleteForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDDeleteForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDDeleteForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDDeleteForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDDeleteNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDDeleteNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDDeleteNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDDeleteNotFound(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDDeleteNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDDeleteNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDDeleteUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDDeleteUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDDeleteUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDDeleteUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDDeleteUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDDeleteUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDGetForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDGetNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDGetNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDGetNotFound(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDPutForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDPutForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDPutForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDPutForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDPutForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDPutForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDPutNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDPutNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDPutNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDPutNotFound(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDPutNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDPutNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesIDPutUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesIDPutUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesIDPutUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesIDPutUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesIDPutUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesIDPutUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesPostConflict) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesPostConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesPostConflict to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesPostConflict(unwrapped)
	return nil
}

func (s *APIV1FraudRulesPostConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesPostConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesPostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesPostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesPostForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesPostForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesPostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesPostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesPostUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesPostUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesValidatePostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesValidatePostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesValidatePostForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesValidatePostForbidden(unwrapped)
	return nil
}

func (s *APIV1FraudRulesValidatePostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesValidatePostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1FraudRulesValidatePostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1FraudRulesValidatePostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1FraudRulesValidatePostUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1FraudRulesValidatePostUnauthorized(unwrapped)
	return nil
}

func (s *APIV1FraudRulesValidatePostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1FraudRulesValidatePostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1PingGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *APIV1PingGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
}

var jsonFieldsNameOfAPIV1PingGetOK = [1]string{
	0: "status",
}

func (s *APIV1PingGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1PingGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode APIV1PingGetOK")
	}

	return nil
}

func (s *APIV1PingGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1PingGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsMerchantsRiskGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsMerchantsRiskGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsMerchantsRiskGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsMerchantsRiskGetForbidden(unwrapped)
	return nil
}

func (s *APIV1StatsMerchantsRiskGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsMerchantsRiskGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsMerchantsRiskGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsMerchantsRiskGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsMerchantsRiskGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsMerchantsRiskGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1StatsMerchantsRiskGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsMerchantsRiskGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsOverviewGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsOverviewGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsOverviewGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsOverviewGetForbidden(unwrapped)
	return nil
}

func (s *APIV1StatsOverviewGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsOverviewGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsOverviewGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsOverviewGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsOverviewGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsOverviewGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1StatsOverviewGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsOverviewGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsRulesMatchesGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsRulesMatchesGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsRulesMatchesGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsRulesMatchesGetForbidden(unwrapped)
	return nil
}

func (s *APIV1StatsRulesMatchesGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsRulesMatchesGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsRulesMatchesGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsRulesMatchesGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsRulesMatchesGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsRulesMatchesGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1StatsRulesMatchesGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsRulesMatchesGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsTransactionsTimeseriesGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsTransactionsTimeseriesGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsTransactionsTimeseriesGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsTransactionsTimeseriesGetForbidden(unwrapped)
	return nil
}

func (s *APIV1StatsTransactionsTimeseriesGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsTransactionsTimeseriesGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsTransactionsTimeseriesGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsTransactionsTimeseriesGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsTransactionsTimeseriesGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsTransactionsTimeseriesGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1StatsTransactionsTimeseriesGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsTransactionsTimeseriesGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsUsersIDRiskProfileGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsUsersIDRiskProfileGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsUsersIDRiskProfileGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsUsersIDRiskProfileGetForbidden(unwrapped)
	return nil
}

func (s *APIV1StatsUsersIDRiskProfileGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsUsersIDRiskProfileGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsUsersIDRiskProfileGetNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsUsersIDRiskProfileGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsUsersIDRiskProfileGetNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsUsersIDRiskProfileGetNotFound(unwrapped)
	return nil
}

func (s *APIV1StatsUsersIDRiskProfileGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsUsersIDRiskProfileGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1StatsUsersIDRiskProfileGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1StatsUsersIDRiskProfileGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1StatsUsersIDRiskProfileGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1StatsUsersIDRiskProfileGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1StatsUsersIDRiskProfileGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1StatsUsersIDRiskProfileGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsBatchPostCreated) Encode(e *jx.Encoder) {
	unwrapped := (*TransactionBatchResult)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsBatchPostCreated) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsBatchPostCreated to nil")
	}
	var unwrapped TransactionBatchResult
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsBatchPostCreated(unwrapped)
	return nil
}

func (s *APIV1TransactionsBatchPostCreated) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsBatchPostCreated) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsBatchPostMultiStatus) Encode(e *jx.Encoder) {
	unwrapped := (*TransactionBatchResult)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsBatchPostMultiStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsBatchPostMultiStatus to nil")
	}
	var unwrapped TransactionBatchResult
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsBatchPostMultiStatus(unwrapped)
	return nil
}

func (s *APIV1TransactionsBatchPostMultiStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsBatchPostMultiStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsIDGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsIDGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsIDGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsIDGetForbidden(unwrapped)
	return nil
}

func (s *APIV1TransactionsIDGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsIDGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsIDGetNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsIDGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsIDGetNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsIDGetNotFound(unwrapped)
	return nil
}

func (s *APIV1TransactionsIDGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsIDGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsIDGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsIDGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsIDGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsIDGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1TransactionsIDGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsIDGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsPostBadRequest to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsPostBadRequest(unwrapped)
	return nil
}

func (s *APIV1TransactionsPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsPostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsPostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsPostForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsPostForbidden(unwrapped)
	return nil
}

func (s *APIV1TransactionsPostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsPostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsPostNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsPostNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsPostNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsPostNotFound(unwrapped)
	return nil
}

func (s *APIV1TransactionsPostNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsPostNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1TransactionsPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1TransactionsPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1TransactionsPostUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1TransactionsPostUnauthorized(unwrapped)
	return nil
}

func (s *APIV1TransactionsPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1TransactionsPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersGetForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDDeleteForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDDeleteForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDDeleteForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDDeleteForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersIDDeleteForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDDeleteForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDDeleteNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDDeleteNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDDeleteNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDDeleteNotFound(unwrapped)
	return nil
}

func (s *APIV1UsersIDDeleteNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDDeleteNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDDeleteUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDDeleteUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDDeleteUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDDeleteUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersIDDeleteUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDDeleteUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDGetForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDGetForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDGetForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDGetForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersIDGetForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDGetForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDGetNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDGetNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDGetNotFound(unwrapped)
	return nil
}

func (s *APIV1UsersIDGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDGetUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDGetUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersIDGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDPutForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDPutForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDPutForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDPutForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersIDPutForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDPutForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDPutNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDPutNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDPutNotFound to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDPutNotFound(unwrapped)
	return nil
}

func (s *APIV1UsersIDPutNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDPutNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersIDPutUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersIDPutUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersIDPutUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersIDPutUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersIDPutUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersIDPutUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersMePutForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersMePutForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersMePutForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersMePutForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersMePutForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersMePutForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersMePutUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersMePutUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersMePutUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersMePutUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersMePutUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersMePutUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersPostConflict) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersPostConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersPostConflict to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersPostConflict(unwrapped)
	return nil
}

func (s *APIV1UsersPostConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersPostConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersPostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersPostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersPostForbidden to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersPostForbidden(unwrapped)
	return nil
}

func (s *APIV1UsersPostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersPostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *APIV1UsersPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*ApiError)(s)

	unwrapped.Encode(e)
}

func (s *APIV1UsersPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode APIV1UsersPostUnauthorized to nil")
	}
	var unwrapped ApiError
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = APIV1UsersPostUnauthorized(unwrapped)
	return nil
}

func (s *APIV1UsersPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *APIV1UsersPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *ApiError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *ApiError) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		s.Code.Encode(e)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		e.FieldStart("traceId")
		json.EncodeUUID(e, s.TraceId)
	}
	{
		e.FieldStart("timestamp")
		json.EncodeDateTime(e, s.Timestamp)
	}
	{
		e.FieldStart("path")
		e.Str(s.Path)
	}
	{
		if s.Details.Set {
			e.FieldStart("details")
			s.Details.Encode(e)
		}
	}
}

var jsonFieldsNameOfApiError = [6]string{
	0: "code",
	1: "message",
	2: "traceId",
	3: "timestamp",
	4: "path",
	5: "details",
}

func (s *ApiError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ApiError to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Code.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "traceId":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.TraceId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"traceId\"")
			}
		case "timestamp":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.Timestamp = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		case "path":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Path = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"path\"")
			}
		case "details":
			if err := func() error {
				s.Details.Reset()
				if err := s.Details.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"details\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ApiError")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfApiError) {
					name = jsonFieldsNameOfApiError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *ApiError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *ApiError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s ApiErrorDetails) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s ApiErrorDetails) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

func (s *ApiErrorDetails) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ApiErrorDetails to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ApiErrorDetails")
	}

	return nil
}

func (s ApiErrorDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *ApiErrorDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *AuthResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *AuthResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("accessToken")
		e.Str(s.AccessToken)
	}
	{
		e.FieldStart("expiresIn")
		e.Int(s.ExpiresIn)
	}
	{
		e.FieldStart("user")
		s.User.Encode(e)
	}
}

var jsonFieldsNameOfAuthResponse = [3]string{
	0: "accessToken",
	1: "expiresIn",
	2: "user",
}

func (s *AuthResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AuthResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "accessToken":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.AccessToken = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"accessToken\"")
			}
		case "expiresIn":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.ExpiresIn = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expiresIn\"")
			}
		case "user":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.User.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"user\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AuthResponse")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAuthResponse) {
					name = jsonFieldsNameOfAuthResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *AuthResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *AuthResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s CurrencyCode) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

func (s *CurrencyCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CurrencyCode to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = CurrencyCode(unwrapped)
	return nil
}

func (s CurrencyCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *CurrencyCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *DslError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *DslError) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		e.Str(s.Code)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		if s.Position.Set {
			e.FieldStart("position")
			s.Position.Encode(e)
		}
	}
	{
		if s.Near.Set {
			e.FieldStart("near")
			s.Near.Encode(e)
		}
	}
}

var jsonFieldsNameOfDslError = [4]string{
	0: "code",
	1: "message",
	2: "position",
	3: "near",
}

func (s *DslError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DslError to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Code = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "position":
			if err := func() error {
				s.Position.Reset()
				if err := s.Position.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"position\"")
			}
		case "near":
			if err := func() error {
				s.Near.Reset()
				if err := s.Near.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"near\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DslError")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDslError) {
					name = jsonFieldsNameOfDslError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *DslError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *DslError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *DslValidateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *DslValidateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("dslExpression")
		e.Str(s.DslExpression)
	}
}

var jsonFieldsNameOfDslValidateRequest = [1]string{
	0: "dslExpression",
}

func (s *DslValidateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DslValidateRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "dslExpression":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.DslExpression = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dslExpression\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DslValidateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDslValidateRequest) {
					name = jsonFieldsNameOfDslValidateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *DslValidateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *DslValidateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *DslValidateResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *DslValidateResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("isValid")
		e.Bool(s.IsValid)
	}
	{
		if s.NormalizedExpression.Set {
			e.FieldStart("normalizedExpression")
			s.NormalizedExpression.Encode(e)
		}
	}
	{
		e.FieldStart("errors")
		e.ArrStart()
		for _, elem := range s.Errors {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfDslValidateResponse = [3]string{
	0: "isValid",
	1: "normalizedExpression",
	2: "errors",
}

func (s *DslValidateResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DslValidateResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "isValid":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Bool()
				s.IsValid = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isValid\"")
			}
		case "normalizedExpression":
			if err := func() error {
				s.NormalizedExpression.Reset()
				if err := s.NormalizedExpression.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"normalizedExpression\"")
			}
		case "errors":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Errors = make([]DslError, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem DslError
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Errors = append(s.Errors, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"errors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DslValidateResponse")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDslValidateResponse) {
					name = jsonFieldsNameOfDslValidateResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *DslValidateResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *DslValidateResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s ErrorCode) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *ErrorCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorCode to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch ErrorCode(v) {
	case ErrorCodeBADREQUEST:
		*s = ErrorCodeBADREQUEST
	case ErrorCodeVALIDATIONFAILED:
		*s = ErrorCodeVALIDATIONFAILED
	case ErrorCodeUNAUTHORIZED:
		*s = ErrorCodeUNAUTHORIZED
	case ErrorCodeFORBIDDEN:
		*s = ErrorCodeFORBIDDEN
	case ErrorCodeNOTFOUND:
		*s = ErrorCodeNOTFOUND
	case ErrorCodeUSERNOTFOUND:
		*s = ErrorCodeUSERNOTFOUND
	case ErrorCodeEMAILALREADYEXISTS:
		*s = ErrorCodeEMAILALREADYEXISTS
	case ErrorCodeUSERINACTIVE:
		*s = ErrorCodeUSERINACTIVE
	case ErrorCodeDSLPARSEERROR:
		*s = ErrorCodeDSLPARSEERROR
	case ErrorCodeDSLINVALIDFIELD:
		*s = ErrorCodeDSLINVALIDFIELD
	case ErrorCodeDSLINVALIDOPERATOR:
		*s = ErrorCodeDSLINVALIDOPERATOR
	case ErrorCodeRULENAMEALREADYEXISTS:
		*s = ErrorCodeRULENAMEALREADYEXISTS
	case ErrorCodeINTERNALSERVERERROR:
		*s = ErrorCodeINTERNALSERVERERROR
	default:
		*s = ErrorCode(v)
	}

	return nil
}

func (s ErrorCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *ErrorCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *FieldError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *FieldError) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("field")
		e.Str(s.Field)
	}
	{
		e.FieldStart("issue")
		e.Str(s.Issue)
	}
	{
		if len(s.RejectedValue) != 0 {
			e.FieldStart("rejectedValue")
			e.Raw(s.RejectedValue)
		}
	}
}

var jsonFieldsNameOfFieldError = [3]string{
	0: "field",
	1: "issue",
	2: "rejectedValue",
}

func (s *FieldError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FieldError to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "field":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Field = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"field\"")
			}
		case "issue":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Issue = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"issue\"")
			}
		case "rejectedValue":
			if err := func() error {
				v, err := d.RawAppend(nil)
				s.RejectedValue = jx.Raw(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"rejectedValue\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FieldError")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFieldError) {
					name = jsonFieldsNameOfFieldError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *FieldError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *FieldError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *FraudRule) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *FraudRule) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		e.FieldStart("dslExpression")
		e.Str(s.DslExpression)
	}
	{
		e.FieldStart("enabled")
		e.Bool(s.Enabled)
	}
	{
		e.FieldStart("priority")
		e.Int(s.Priority)
	}
	{
		e.FieldStart("createdAt")
		json.EncodeDateTime(e, s.CreatedAt)
	}
	{
		e.FieldStart("updatedAt")
		json.EncodeDateTime(e, s.UpdatedAt)
	}
}

var jsonFieldsNameOfFraudRule = [8]string{
	0: "id",
	1: "name",
	2: "description",
	3: "dslExpression",
	4: "enabled",
	5: "priority",
	6: "createdAt",
	7: "updatedAt",
}

func (s *FraudRule) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FraudRule to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "dslExpression":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.DslExpression = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dslExpression\"")
			}
		case "enabled":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Bool()
				s.Enabled = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "priority":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Int()
				s.Priority = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		case "createdAt":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"createdAt\"")
			}
		case "updatedAt":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.UpdatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updatedAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FraudRule")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b11111011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFraudRule) {
					name = jsonFieldsNameOfFraudRule[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *FraudRule) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *FraudRule) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *FraudRuleCreateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *FraudRuleCreateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		e.FieldStart("dslExpression")
		e.Str(s.DslExpression)
	}
	{
		if s.Enabled.Set {
			e.FieldStart("enabled")
			s.Enabled.Encode(e)
		}
	}
	{
		if s.Priority.Set {
			e.FieldStart("priority")
			s.Priority.Encode(e)
		}
	}
}

var jsonFieldsNameOfFraudRuleCreateRequest = [5]string{
	0: "name",
	1: "description",
	2: "dslExpression",
	3: "enabled",
	4: "priority",
}

func (s *FraudRuleCreateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FraudRuleCreateRequest to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "dslExpression":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.DslExpression = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dslExpression\"")
			}
		case "enabled":
			if err := func() error {
				s.Enabled.Reset()
				if err := s.Enabled.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "priority":
			if err := func() error {
				s.Priority.Reset()
				if err := s.Priority.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FraudRuleCreateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFraudRuleCreateRequest) {
					name = jsonFieldsNameOfFraudRuleCreateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *FraudRuleCreateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *FraudRuleCreateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *FraudRuleEvaluationResult) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *FraudRuleEvaluationResult) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ruleId")
		json.EncodeUUID(e, s.RuleId)
	}
	{
		e.FieldStart("ruleName")
		e.Str(s.RuleName)
	}
	{
		e.FieldStart("priority")
		e.Int(s.Priority)
	}
	{
		e.FieldStart("enabled")
		e.Bool(s.Enabled)
	}
	{
		e.FieldStart("matched")
		e.Bool(s.Matched)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
}

var jsonFieldsNameOfFraudRuleEvaluationResult = [6]string{
	0: "ruleId",
	1: "ruleName",
	2: "priority",
	3: "enabled",
	4: "matched",
	5: "description",
}

func (s *FraudRuleEvaluationResult) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FraudRuleEvaluationResult to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ruleId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.RuleId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ruleId\"")
			}
		case "ruleName":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.RuleName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ruleName\"")
			}
		case "priority":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Priority = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		case "enabled":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Bool()
				s.Enabled = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "matched":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Bool()
				s.Matched = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"matched\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FraudRuleEvaluationResult")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFraudRuleEvaluationResult) {
					name = jsonFieldsNameOfFraudRuleEvaluationResult[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *FraudRuleEvaluationResult) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *FraudRuleEvaluationResult) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *FraudRuleUpdateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *FraudRuleUpdateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		e.FieldStart("dslExpression")
		e.Str(s.DslExpression)
	}
	{
		e.FieldStart("enabled")
		e.Bool(s.Enabled)
	}
	{
		e.FieldStart("priority")
		e.Int(s.Priority)
	}
}

var jsonFieldsNameOfFraudRuleUpdateRequest = [5]string{
	0: "name",
	1: "description",
	2: "dslExpression",
	3: "enabled",
	4: "priority",
}

func (s *FraudRuleUpdateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FraudRuleUpdateRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "dslExpression":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.DslExpression = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dslExpression\"")
			}
		case "enabled":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Bool()
				s.Enabled = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"enabled\"")
			}
		case "priority":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int()
				s.Priority = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"priority\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FraudRuleUpdateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFraudRuleUpdateRequest) {
					name = jsonFieldsNameOfFraudRuleUpdateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *FraudRuleUpdateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *FraudRuleUpdateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s Gender) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *Gender) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Gender to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch Gender(v) {
	case GenderMALE:
		*s = GenderMALE
	case GenderFEMALE:
		*s = GenderFEMALE
	default:
		*s = Gender(v)
	}

	return nil
}

func (s Gender) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *Gender) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *LoginRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *LoginRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("password")
		e.Str(s.Password)
	}
}

var jsonFieldsNameOfLoginRequest = [2]string{
	0: "email",
	1: "password",
}

func (s *LoginRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "email":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLoginRequest) {
					name = jsonFieldsNameOfLoginRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *LoginRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *LoginRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s MaritalStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *MaritalStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MaritalStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch MaritalStatus(v) {
	case MaritalStatusSINGLE:
		*s = MaritalStatusSINGLE
	case MaritalStatusMARRIED:
		*s = MaritalStatusMARRIED
	case MaritalStatusDIVORCED:
		*s = MaritalStatusDIVORCED
	case MaritalStatusWIDOWED:
		*s = MaritalStatusWIDOWED
	default:
		*s = MaritalStatus(v)
	}

	return nil
}

func (s MaritalStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *MaritalStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s MccCode) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

func (s *MccCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MccCode to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = MccCode(unwrapped)
	return nil
}

func (s MccCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *MccCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *MerchantRiskRow) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *MerchantRiskRow) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("merchantId")
		e.Str(s.MerchantId)
	}
	{
		if s.MerchantCategoryCode.Set {
			e.FieldStart("merchantCategoryCode")
			s.MerchantCategoryCode.Encode(e)
		}
	}
	{
		e.FieldStart("txCount")
		e.Int(s.TxCount)
	}
	{
		if s.Gmv.Set {
			e.FieldStart("gmv")
			s.Gmv.Encode(e)
		}
	}
	{
		e.FieldStart("declineRate")
		e.Float64(s.DeclineRate)
	}
}

var jsonFieldsNameOfMerchantRiskRow = [5]string{
	0: "merchantId",
	1: "merchantCategoryCode",
	2: "txCount",
	3: "gmv",
	4: "declineRate",
}

func (s *MerchantRiskRow) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MerchantRiskRow to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "merchantId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.MerchantId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantId\"")
			}
		case "merchantCategoryCode":
			if err := func() error {
				s.MerchantCategoryCode.Reset()
				if err := s.MerchantCategoryCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantCategoryCode\"")
			}
		case "txCount":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.TxCount = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"txCount\"")
			}
		case "gmv":
			if err := func() error {
				s.Gmv.Reset()
				if err := s.Gmv.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gmv\"")
			}
		case "declineRate":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Float64()
				s.DeclineRate = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"declineRate\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MerchantRiskRow")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00010101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMerchantRiskRow) {
					name = jsonFieldsNameOfMerchantRiskRow[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *MerchantRiskRow) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *MerchantRiskRow) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *MerchantRiskStats) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *MerchantRiskStats) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfMerchantRiskStats = [1]string{
	0: "items",
}

func (s *MerchantRiskStats) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MerchantRiskStats to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]MerchantRiskRow, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem MerchantRiskRow
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MerchantRiskStats")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMerchantRiskStats) {
					name = jsonFieldsNameOfMerchantRiskStats[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *MerchantRiskStats) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *MerchantRiskStats) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o NilGender) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

func (o *NilGender) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilGender to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v Gender
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s NilGender) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *NilGender) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o NilInt) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Int(int(o.Value))
}

func (o *NilInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilInt to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

func (s NilInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *NilInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o NilMaritalStatus) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

func (o *NilMaritalStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilMaritalStatus to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v MaritalStatus
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s NilMaritalStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *NilMaritalStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o NilString) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

func (o *NilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode NilString to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v string
		o.Value = v
		o.Null = true
		return nil
	}
	o.Null = false
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

func (s NilString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *NilString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptApiError) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptApiError) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptApiError to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptApiError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptApiError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptApiErrorDetails) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptApiErrorDetails) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptApiErrorDetails to nil")
	}
	o.Set = true
	o.Value = make(ApiErrorDetails)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptApiErrorDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptApiErrorDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptFloat64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Float64(float64(o.Value))
}

func (o *OptFloat64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptFloat64 to nil")
	}
	o.Set = true
	v, err := d.Float64()
	if err != nil {
		return err
	}
	o.Value = float64(v)
	return nil
}

func (s OptFloat64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptFloat64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptGender) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

func (o *OptGender) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptGender to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptGender) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptGender) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptMaritalStatus) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

func (o *OptMaritalStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptMaritalStatus to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptMaritalStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptMaritalStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptMccCode) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptMccCode) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptMccCode to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptMccCode) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptMccCode) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptNilDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	format(e, o.Value)
}

func (o *OptNilDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilDateTime to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v time.Time
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

func (s OptNilDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

func (s *OptNilDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

func (o OptNilInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Int(int(o.Value))
}

func (o *OptNilInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilInt to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v int
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

func (s OptNilInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptNilInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptNilString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

func (o *OptNilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptNilString to nil")
	}
	if d.Next() == jx.Null {
		if err := d.Null(); err != nil {
			return err
		}

		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	}
	o.Set = true
	o.Null = false
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

func (s OptNilString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptNilString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptTransactionChannel) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

func (o *OptTransactionChannel) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionChannel to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptTransactionChannel) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptTransactionChannel) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptTransactionCreateRequestMetadata) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptTransactionCreateRequestMetadata) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionCreateRequestMetadata to nil")
	}
	o.Set = true
	o.Value = make(TransactionCreateRequestMetadata)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptTransactionCreateRequestMetadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptTransactionCreateRequestMetadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptTransactionDecision) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptTransactionDecision) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionDecision to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptTransactionDecision) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptTransactionDecision) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptTransactionLocation) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptTransactionLocation) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionLocation to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptTransactionLocation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptTransactionLocation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptTransactionMetadata) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

func (o *OptTransactionMetadata) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTransactionMetadata to nil")
	}
	o.Set = true
	o.Value = make(TransactionMetadata)
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptTransactionMetadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptTransactionMetadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (o OptUserRole) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

func (o *OptUserRole) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUserRole to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

func (s OptUserRole) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *OptUserRole) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *PagedTransactions) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *PagedTransactions) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("total")
		e.Int(s.Total)
	}
	{
		e.FieldStart("page")
		e.Int(s.Page)
	}
	{
		e.FieldStart("size")
		e.Int(s.Size)
	}
}

var jsonFieldsNameOfPagedTransactions = [4]string{
	0: "items",
	1: "total",
	2: "page",
	3: "size",
}

func (s *PagedTransactions) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PagedTransactions to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]Transaction, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Transaction
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		case "total":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Total = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total\"")
			}
		case "page":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Page = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"page\"")
			}
		case "size":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int()
				s.Size = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"size\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PagedTransactions")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfPagedTransactions) {
					name = jsonFieldsNameOfPagedTransactions[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *PagedTransactions) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *PagedTransactions) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *PagedUsers) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *PagedUsers) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("total")
		e.Int(s.Total)
	}
	{
		e.FieldStart("page")
		e.Int(s.Page)
	}
	{
		e.FieldStart("size")
		e.Int(s.Size)
	}
}

var jsonFieldsNameOfPagedUsers = [4]string{
	0: "items",
	1: "total",
	2: "page",
	3: "size",
}

func (s *PagedUsers) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PagedUsers to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]User, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem User
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		case "total":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Total = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"total\"")
			}
		case "page":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Page = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"page\"")
			}
		case "size":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int()
				s.Size = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"size\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PagedUsers")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfPagedUsers) {
					name = jsonFieldsNameOfPagedUsers[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *PagedUsers) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *PagedUsers) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *RegisterRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *RegisterRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("password")
		e.Str(s.Password)
	}
	{
		e.FieldStart("fullName")
		e.Str(s.FullName)
	}
	{
		if s.Region.Set {
			e.FieldStart("region")
			s.Region.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
	{
		if s.MaritalStatus.Set {
			e.FieldStart("maritalStatus")
			s.MaritalStatus.Encode(e)
		}
	}
}

var jsonFieldsNameOfRegisterRequest = [7]string{
	0: "email",
	1: "password",
	2: "fullName",
	3: "region",
	4: "gender",
	5: "age",
	6: "maritalStatus",
}

func (s *RegisterRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RegisterRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "email":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		case "fullName":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.FullName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fullName\"")
			}
		case "region":
			if err := func() error {
				s.Region.Reset()
				if err := s.Region.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"region\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "maritalStatus":
			if err := func() error {
				s.MaritalStatus.Reset()
				if err := s.MaritalStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"maritalStatus\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RegisterRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfRegisterRequest) {
					name = jsonFieldsNameOfRegisterRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *RegisterRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *RegisterRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *RuleMatchRow) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *RuleMatchRow) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("ruleId")
		json.EncodeUUID(e, s.RuleId)
	}
	{
		e.FieldStart("ruleName")
		e.Str(s.RuleName)
	}
	{
		e.FieldStart("matches")
		e.Int(s.Matches)
	}
	{
		e.FieldStart("uniqueUsers")
		e.Int(s.UniqueUsers)
	}
	{
		if s.UniqueMerchants.Set {
			e.FieldStart("uniqueMerchants")
			s.UniqueMerchants.Encode(e)
		}
	}
	{
		e.FieldStart("shareOfDeclines")
		e.Float64(s.ShareOfDeclines)
	}
}

var jsonFieldsNameOfRuleMatchRow = [6]string{
	0: "ruleId",
	1: "ruleName",
	2: "matches",
	3: "uniqueUsers",
	4: "uniqueMerchants",
	5: "shareOfDeclines",
}

func (s *RuleMatchRow) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RuleMatchRow to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "ruleId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.RuleId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ruleId\"")
			}
		case "ruleName":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.RuleName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ruleName\"")
			}
		case "matches":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Matches = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"matches\"")
			}
		case "uniqueUsers":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int()
				s.UniqueUsers = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uniqueUsers\"")
			}
		case "uniqueMerchants":
			if err := func() error {
				s.UniqueMerchants.Reset()
				if err := s.UniqueMerchants.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"uniqueMerchants\"")
			}
		case "shareOfDeclines":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Float64()
				s.ShareOfDeclines = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"shareOfDeclines\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RuleMatchRow")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00101111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfRuleMatchRow) {
					name = jsonFieldsNameOfRuleMatchRow[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *RuleMatchRow) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *RuleMatchRow) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *RuleMatchStats) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *RuleMatchStats) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfRuleMatchStats = [1]string{
	0: "items",
}

func (s *RuleMatchStats) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RuleMatchStats to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]RuleMatchRow, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem RuleMatchRow
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RuleMatchStats")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfRuleMatchStats) {
					name = jsonFieldsNameOfRuleMatchStats[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *RuleMatchStats) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *RuleMatchStats) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *StatsOverview) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *StatsOverview) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("from")
		json.EncodeDateTime(e, s.From)
	}
	{
		e.FieldStart("to")
		json.EncodeDateTime(e, s.To)
	}
	{
		e.FieldStart("volume")
		e.Int(s.Volume)
	}
	{
		e.FieldStart("gmv")
		e.Float64(s.Gmv)
	}
	{
		e.FieldStart("approvalRate")
		e.Float64(s.ApprovalRate)
	}
	{
		e.FieldStart("declineRate")
		e.Float64(s.DeclineRate)
	}
	{
		e.FieldStart("topRiskMerchants")
		e.ArrStart()
		for _, elem := range s.TopRiskMerchants {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfStatsOverview = [7]string{
	0: "from",
	1: "to",
	2: "volume",
	3: "gmv",
	4: "approvalRate",
	5: "declineRate",
	6: "topRiskMerchants",
}

func (s *StatsOverview) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StatsOverview to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "from":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.From = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from\"")
			}
		case "to":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.To = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"to\"")
			}
		case "volume":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Volume = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"volume\"")
			}
		case "gmv":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Float64()
				s.Gmv = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gmv\"")
			}
		case "approvalRate":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Float64()
				s.ApprovalRate = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"approvalRate\"")
			}
		case "declineRate":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Float64()
				s.DeclineRate = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"declineRate\"")
			}
		case "topRiskMerchants":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				s.TopRiskMerchants = make([]MerchantRiskRow, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem MerchantRiskRow
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.TopRiskMerchants = append(s.TopRiskMerchants, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"topRiskMerchants\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StatsOverview")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b01111111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStatsOverview) {
					name = jsonFieldsNameOfStatsOverview[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *StatsOverview) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *StatsOverview) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *Transaction) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *Transaction) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("userId")
		json.EncodeUUID(e, s.UserId)
	}
	{
		e.FieldStart("amount")
		e.Float64(s.Amount)
	}
	{
		e.FieldStart("currency")
		s.Currency.Encode(e)
	}
	{
		e.FieldStart("status")
		s.Status.Encode(e)
	}
	{
		if s.MerchantId.Set {
			e.FieldStart("merchantId")
			s.MerchantId.Encode(e)
		}
	}
	{
		if s.MerchantCategoryCode.Set {
			e.FieldStart("merchantCategoryCode")
			s.MerchantCategoryCode.Encode(e)
		}
	}
	{
		e.FieldStart("timestamp")
		json.EncodeDateTime(e, s.Timestamp)
	}
	{
		if s.IpAddress.Set {
			e.FieldStart("ipAddress")
			s.IpAddress.Encode(e)
		}
	}
	{
		if s.DeviceId.Set {
			e.FieldStart("deviceId")
			s.DeviceId.Encode(e)
		}
	}
	{
		if s.Channel.Set {
			e.FieldStart("channel")
			s.Channel.Encode(e)
		}
	}
	{
		if s.Location.Set {
			e.FieldStart("location")
			s.Location.Encode(e)
		}
	}
	{
		e.FieldStart("isFraud")
		e.Bool(s.IsFraud)
	}
	{
		if s.Metadata.Set {
			e.FieldStart("metadata")
			s.Metadata.Encode(e)
		}
	}
	{
		e.FieldStart("createdAt")
		json.EncodeDateTime(e, s.CreatedAt)
	}
}

var jsonFieldsNameOfTransaction = [15]string{
	0:  "id",
	1:  "userId",
	2:  "amount",
	3:  "currency",
	4:  "status",
	5:  "merchantId",
	6:  "merchantCategoryCode",
	7:  "timestamp",
	8:  "ipAddress",
	9:  "deviceId",
	10: "channel",
	11: "location",
	12: "isFraud",
	13: "metadata",
	14: "createdAt",
}

func (s *Transaction) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Transaction to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "userId":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.UserId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"userId\"")
			}
		case "amount":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Float64()
				s.Amount = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"amount\"")
			}
		case "currency":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "merchantId":
			if err := func() error {
				s.MerchantId.Reset()
				if err := s.MerchantId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantId\"")
			}
		case "merchantCategoryCode":
			if err := func() error {
				s.MerchantCategoryCode.Reset()
				if err := s.MerchantCategoryCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantCategoryCode\"")
			}
		case "timestamp":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.Timestamp = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		case "ipAddress":
			if err := func() error {
				s.IpAddress.Reset()
				if err := s.IpAddress.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ipAddress\"")
			}
		case "deviceId":
			if err := func() error {
				s.DeviceId.Reset()
				if err := s.DeviceId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"deviceId\"")
			}
		case "channel":
			if err := func() error {
				s.Channel.Reset()
				if err := s.Channel.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"channel\"")
			}
		case "location":
			if err := func() error {
				s.Location.Reset()
				if err := s.Location.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"location\"")
			}
		case "isFraud":
			requiredBitSet[1] |= 1 << 4
			if err := func() error {
				v, err := d.Bool()
				s.IsFraud = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isFraud\"")
			}
		case "metadata":
			if err := func() error {
				s.Metadata.Reset()
				if err := s.Metadata.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metadata\"")
			}
		case "createdAt":
			requiredBitSet[1] |= 1 << 6
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"createdAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Transaction")
	}
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b10011111,
		0b01010000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransaction) {
					name = jsonFieldsNameOfTransaction[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *Transaction) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *Transaction) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionBatchCreateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionBatchCreateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfTransactionBatchCreateRequest = [1]string{
	0: "items",
}

func (s *TransactionBatchCreateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionBatchCreateRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]TransactionCreateRequest, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem TransactionCreateRequest
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionBatchCreateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionBatchCreateRequest) {
					name = jsonFieldsNameOfTransactionBatchCreateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionBatchCreateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionBatchCreateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionBatchResult) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionBatchResult) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("items")
		e.ArrStart()
		for _, elem := range s.Items {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfTransactionBatchResult = [1]string{
	0: "items",
}

func (s *TransactionBatchResult) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionBatchResult to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Items = make([]TransactionBatchResultItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem TransactionBatchResultItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionBatchResult")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionBatchResult) {
					name = jsonFieldsNameOfTransactionBatchResult[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionBatchResult) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionBatchResult) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionBatchResultItem) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionBatchResultItem) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("index")
		e.Int(s.Index)
	}
	{
		if s.Decision.Set {
			e.FieldStart("decision")
			s.Decision.Encode(e)
		}
	}
	{
		if s.Error.Set {
			e.FieldStart("error")
			s.Error.Encode(e)
		}
	}
}

var jsonFieldsNameOfTransactionBatchResultItem = [3]string{
	0: "index",
	1: "decision",
	2: "error",
}

func (s *TransactionBatchResultItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionBatchResultItem to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "index":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Index = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"index\"")
			}
		case "decision":
			if err := func() error {
				s.Decision.Reset()
				if err := s.Decision.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"decision\"")
			}
		case "error":
			if err := func() error {
				s.Error.Reset()
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionBatchResultItem")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionBatchResultItem) {
					name = jsonFieldsNameOfTransactionBatchResultItem[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionBatchResultItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionBatchResultItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s TransactionChannel) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *TransactionChannel) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionChannel to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch TransactionChannel(v) {
	case TransactionChannelWEB:
		*s = TransactionChannelWEB
	case TransactionChannelMOBILE:
		*s = TransactionChannelMOBILE
	case TransactionChannelPOS:
		*s = TransactionChannelPOS
	case TransactionChannelOTHER:
		*s = TransactionChannelOTHER
	default:
		*s = TransactionChannel(v)
	}

	return nil
}

func (s TransactionChannel) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionChannel) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionCreateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionCreateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("userId")
		json.EncodeUUID(e, s.UserId)
	}
	{
		e.FieldStart("amount")
		e.Float64(s.Amount)
	}
	{
		e.FieldStart("currency")
		s.Currency.Encode(e)
	}
	{
		if s.MerchantId.Set {
			e.FieldStart("merchantId")
			s.MerchantId.Encode(e)
		}
	}
	{
		if s.MerchantCategoryCode.Set {
			e.FieldStart("merchantCategoryCode")
			s.MerchantCategoryCode.Encode(e)
		}
	}
	{
		e.FieldStart("timestamp")
		json.EncodeDateTime(e, s.Timestamp)
	}
	{
		if s.IpAddress.Set {
			e.FieldStart("ipAddress")
			s.IpAddress.Encode(e)
		}
	}
	{
		if s.DeviceId.Set {
			e.FieldStart("deviceId")
			s.DeviceId.Encode(e)
		}
	}
	{
		if s.Channel.Set {
			e.FieldStart("channel")
			s.Channel.Encode(e)
		}
	}
	{
		if s.Location.Set {
			e.FieldStart("location")
			s.Location.Encode(e)
		}
	}
	{
		if s.Metadata.Set {
			e.FieldStart("metadata")
			s.Metadata.Encode(e)
		}
	}
}

var jsonFieldsNameOfTransactionCreateRequest = [11]string{
	0:  "userId",
	1:  "amount",
	2:  "currency",
	3:  "merchantId",
	4:  "merchantCategoryCode",
	5:  "timestamp",
	6:  "ipAddress",
	7:  "deviceId",
	8:  "channel",
	9:  "location",
	10: "metadata",
}

func (s *TransactionCreateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionCreateRequest to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "userId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.UserId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"userId\"")
			}
		case "amount":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Float64()
				s.Amount = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"amount\"")
			}
		case "currency":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Currency.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"currency\"")
			}
		case "merchantId":
			if err := func() error {
				s.MerchantId.Reset()
				if err := s.MerchantId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantId\"")
			}
		case "merchantCategoryCode":
			if err := func() error {
				s.MerchantCategoryCode.Reset()
				if err := s.MerchantCategoryCode.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"merchantCategoryCode\"")
			}
		case "timestamp":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.Timestamp = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		case "ipAddress":
			if err := func() error {
				s.IpAddress.Reset()
				if err := s.IpAddress.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ipAddress\"")
			}
		case "deviceId":
			if err := func() error {
				s.DeviceId.Reset()
				if err := s.DeviceId.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"deviceId\"")
			}
		case "channel":
			if err := func() error {
				s.Channel.Reset()
				if err := s.Channel.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"channel\"")
			}
		case "location":
			if err := func() error {
				s.Location.Reset()
				if err := s.Location.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"location\"")
			}
		case "metadata":
			if err := func() error {
				s.Metadata.Reset()
				if err := s.Metadata.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"metadata\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionCreateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b00100111,
		0b00000000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionCreateRequest) {
					name = jsonFieldsNameOfTransactionCreateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionCreateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionCreateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s TransactionCreateRequestMetadata) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s TransactionCreateRequestMetadata) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

func (s *TransactionCreateRequestMetadata) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionCreateRequestMetadata to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionCreateRequestMetadata")
	}

	return nil
}

func (s TransactionCreateRequestMetadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionCreateRequestMetadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionDecision) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionDecision) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("transaction")
		s.Transaction.Encode(e)
	}
	{
		e.FieldStart("ruleResults")
		e.ArrStart()
		for _, elem := range s.RuleResults {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfTransactionDecision = [2]string{
	0: "transaction",
	1: "ruleResults",
}

func (s *TransactionDecision) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionDecision to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "transaction":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Transaction.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"transaction\"")
			}
		case "ruleResults":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.RuleResults = make([]FraudRuleEvaluationResult, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem FraudRuleEvaluationResult
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.RuleResults = append(s.RuleResults, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ruleResults\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionDecision")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionDecision) {
					name = jsonFieldsNameOfTransactionDecision[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionDecision) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionDecision) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionLocation) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionLocation) encodeFields(e *jx.Encoder) {
	{
		if s.Country.Set {
			e.FieldStart("country")
			s.Country.Encode(e)
		}
	}
	{
		if s.City.Set {
			e.FieldStart("city")
			s.City.Encode(e)
		}
	}
	{
		if s.Latitude.Set {
			e.FieldStart("latitude")
			s.Latitude.Encode(e)
		}
	}
	{
		if s.Longitude.Set {
			e.FieldStart("longitude")
			s.Longitude.Encode(e)
		}
	}
}

var jsonFieldsNameOfTransactionLocation = [4]string{
	0: "country",
	1: "city",
	2: "latitude",
	3: "longitude",
}

func (s *TransactionLocation) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionLocation to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "country":
			if err := func() error {
				s.Country.Reset()
				if err := s.Country.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"country\"")
			}
		case "city":
			if err := func() error {
				s.City.Reset()
				if err := s.City.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"city\"")
			}
		case "latitude":
			if err := func() error {
				s.Latitude.Reset()
				if err := s.Latitude.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latitude\"")
			}
		case "longitude":
			if err := func() error {
				s.Longitude.Reset()
				if err := s.Longitude.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"longitude\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionLocation")
	}

	return nil
}

func (s *TransactionLocation) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionLocation) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s TransactionMetadata) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s TransactionMetadata) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

func (s *TransactionMetadata) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionMetadata to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionMetadata")
	}

	return nil
}

func (s TransactionMetadata) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionMetadata) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s TransactionStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *TransactionStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch TransactionStatus(v) {
	case TransactionStatusAPPROVED:
		*s = TransactionStatusAPPROVED
	case TransactionStatusDECLINED:
		*s = TransactionStatusDECLINED
	default:
		*s = TransactionStatus(v)
	}

	return nil
}

func (s TransactionStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionsTimePoint) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionsTimePoint) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("bucketStart")
		json.EncodeDateTime(e, s.BucketStart)
	}
	{
		e.FieldStart("txCount")
		e.Int(s.TxCount)
	}
	{
		e.FieldStart("gmv")
		e.Float64(s.Gmv)
	}
	{
		e.FieldStart("approvalRate")
		e.Float64(s.ApprovalRate)
	}
	{
		e.FieldStart("declineRate")
		e.Float64(s.DeclineRate)
	}
}

var jsonFieldsNameOfTransactionsTimePoint = [5]string{
	0: "bucketStart",
	1: "txCount",
	2: "gmv",
	3: "approvalRate",
	4: "declineRate",
}

func (s *TransactionsTimePoint) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionsTimePoint to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "bucketStart":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.BucketStart = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"bucketStart\"")
			}
		case "txCount":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.TxCount = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"txCount\"")
			}
		case "gmv":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Float64()
				s.Gmv = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gmv\"")
			}
		case "approvalRate":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Float64()
				s.ApprovalRate = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"approvalRate\"")
			}
		case "declineRate":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Float64()
				s.DeclineRate = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"declineRate\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionsTimePoint")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionsTimePoint) {
					name = jsonFieldsNameOfTransactionsTimePoint[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionsTimePoint) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionsTimePoint) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *TransactionsTimeSeries) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *TransactionsTimeSeries) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("points")
		e.ArrStart()
		for _, elem := range s.Points {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfTransactionsTimeSeries = [1]string{
	0: "points",
}

func (s *TransactionsTimeSeries) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TransactionsTimeSeries to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "points":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Points = make([]TransactionsTimePoint, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem TransactionsTimePoint
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Points = append(s.Points, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"points\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TransactionsTimeSeries")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfTransactionsTimeSeries) {
					name = jsonFieldsNameOfTransactionsTimeSeries[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *TransactionsTimeSeries) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *TransactionsTimeSeries) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *User) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("fullName")
		e.Str(s.FullName)
	}
	{
		if s.Region.Set {
			e.FieldStart("region")
			s.Region.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
	{
		if s.MaritalStatus.Set {
			e.FieldStart("maritalStatus")
			s.MaritalStatus.Encode(e)
		}
	}
	{
		e.FieldStart("role")
		s.Role.Encode(e)
	}
	{
		e.FieldStart("isActive")
		e.Bool(s.IsActive)
	}
	{
		e.FieldStart("createdAt")
		json.EncodeDateTime(e, s.CreatedAt)
	}
	{
		e.FieldStart("updatedAt")
		json.EncodeDateTime(e, s.UpdatedAt)
	}
}

var jsonFieldsNameOfUser = [11]string{
	0:  "id",
	1:  "email",
	2:  "fullName",
	3:  "region",
	4:  "gender",
	5:  "age",
	6:  "maritalStatus",
	7:  "role",
	8:  "isActive",
	9:  "createdAt",
	10: "updatedAt",
}

func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}
	var requiredBitSet [2]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "email":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "fullName":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.FullName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fullName\"")
			}
		case "region":
			if err := func() error {
				s.Region.Reset()
				if err := s.Region.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"region\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "maritalStatus":
			if err := func() error {
				s.MaritalStatus.Reset()
				if err := s.MaritalStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"maritalStatus\"")
			}
		case "role":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				if err := s.Role.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		case "isActive":
			requiredBitSet[1] |= 1 << 0
			if err := func() error {
				v, err := d.Bool()
				s.IsActive = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isActive\"")
			}
		case "createdAt":
			requiredBitSet[1] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"createdAt\"")
			}
		case "updatedAt":
			requiredBitSet[1] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.UpdatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updatedAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b10000111,
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUser) {
					name = jsonFieldsNameOfUser[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *UserCreateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *UserCreateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("password")
		e.Str(s.Password)
	}
	{
		e.FieldStart("fullName")
		e.Str(s.FullName)
	}
	{
		if s.Region.Set {
			e.FieldStart("region")
			s.Region.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
	{
		if s.MaritalStatus.Set {
			e.FieldStart("maritalStatus")
			s.MaritalStatus.Encode(e)
		}
	}
	{
		e.FieldStart("role")
		s.Role.Encode(e)
	}
}

var jsonFieldsNameOfUserCreateRequest = [8]string{
	0: "email",
	1: "password",
	2: "fullName",
	3: "region",
	4: "gender",
	5: "age",
	6: "maritalStatus",
	7: "role",
}

func (s *UserCreateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserCreateRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "email":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		case "fullName":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.FullName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fullName\"")
			}
		case "region":
			if err := func() error {
				s.Region.Reset()
				if err := s.Region.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"region\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "maritalStatus":
			if err := func() error {
				s.MaritalStatus.Reset()
				if err := s.MaritalStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"maritalStatus\"")
			}
		case "role":
			requiredBitSet[0] |= 1 << 7
			if err := func() error {
				if err := s.Role.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserCreateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b10000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserCreateRequest) {
					name = jsonFieldsNameOfUserCreateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *UserCreateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *UserCreateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *UserRiskProfile) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *UserRiskProfile) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("userId")
		json.EncodeUUID(e, s.UserId)
	}
	{
		e.FieldStart("txCount_24h")
		e.Int(s.TxCount24h)
	}
	{
		e.FieldStart("gmv_24h")
		e.Float64(s.Gmv24h)
	}
	{
		e.FieldStart("distinctDevices_24h")
		e.Int(s.DistinctDevices24h)
	}
	{
		e.FieldStart("distinctIps_24h")
		e.Int(s.DistinctIps24h)
	}
	{
		e.FieldStart("distinctCities_24h")
		e.Int(s.DistinctCities24h)
	}
	{
		e.FieldStart("declineRate_30d")
		e.Float64(s.DeclineRate30d)
	}
	{
		if s.LastSeenAt.Set {
			e.FieldStart("lastSeenAt")
			s.LastSeenAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfUserRiskProfile = [8]string{
	0: "userId",
	1: "txCount_24h",
	2: "gmv_24h",
	3: "distinctDevices_24h",
	4: "distinctIps_24h",
	5: "distinctCities_24h",
	6: "declineRate_30d",
	7: "lastSeenAt",
}

func (s *UserRiskProfile) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRiskProfile to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "userId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.UserId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"userId\"")
			}
		case "txCount_24h":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.TxCount24h = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"txCount_24h\"")
			}
		case "gmv_24h":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Float64()
				s.Gmv24h = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gmv_24h\"")
			}
		case "distinctDevices_24h":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int()
				s.DistinctDevices24h = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"distinctDevices_24h\"")
			}
		case "distinctIps_24h":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int()
				s.DistinctIps24h = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"distinctIps_24h\"")
			}
		case "distinctCities_24h":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Int()
				s.DistinctCities24h = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"distinctCities_24h\"")
			}
		case "declineRate_30d":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				v, err := d.Float64()
				s.DeclineRate30d = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"declineRate_30d\"")
			}
		case "lastSeenAt":
			if err := func() error {
				s.LastSeenAt.Reset()
				if err := s.LastSeenAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"lastSeenAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserRiskProfile")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b01111111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserRiskProfile) {
					name = jsonFieldsNameOfUserRiskProfile[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *UserRiskProfile) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *UserRiskProfile) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s UserRole) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

func (s *UserRole) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserRole to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	switch UserRole(v) {
	case UserRoleADMIN:
		*s = UserRoleADMIN
	case UserRoleUSER:
		*s = UserRoleUSER
	default:
		*s = UserRole(v)
	}

	return nil
}

func (s UserRole) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *UserRole) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *UserUpdateRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *UserUpdateRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("fullName")
		e.Str(s.FullName)
	}
	{
		e.FieldStart("region")
		s.Region.Encode(e)
	}
	{
		e.FieldStart("gender")
		s.Gender.Encode(e)
	}
	{
		e.FieldStart("age")
		s.Age.Encode(e)
	}
	{
		e.FieldStart("maritalStatus")
		s.MaritalStatus.Encode(e)
	}
	{
		if s.Role.Set {
			e.FieldStart("role")
			s.Role.Encode(e)
		}
	}
	{
		if s.IsActive.Set {
			e.FieldStart("isActive")
			s.IsActive.Encode(e)
		}
	}
}

var jsonFieldsNameOfUserUpdateRequest = [7]string{
	0: "fullName",
	1: "region",
	2: "gender",
	3: "age",
	4: "maritalStatus",
	5: "role",
	6: "isActive",
}

func (s *UserUpdateRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserUpdateRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "fullName":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.FullName = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fullName\"")
			}
		case "region":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Region.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"region\"")
			}
		case "gender":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "age":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "maritalStatus":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				if err := s.MaritalStatus.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"maritalStatus\"")
			}
		case "role":
			if err := func() error {
				s.Role.Reset()
				if err := s.Role.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		case "isActive":
			if err := func() error {
				s.IsActive.Reset()
				if err := s.IsActive.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isActive\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserUpdateRequest")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserUpdateRequest) {
					name = jsonFieldsNameOfUserUpdateRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *UserUpdateRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *UserUpdateRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

func (s *ValidationError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

func (s *ValidationError) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		e.Str(s.Code)
	}
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		e.FieldStart("traceId")
		json.EncodeUUID(e, s.TraceId)
	}
	{
		e.FieldStart("timestamp")
		json.EncodeDateTime(e, s.Timestamp)
	}
	{
		e.FieldStart("path")
		e.Str(s.Path)
	}
	{
		e.FieldStart("fieldErrors")
		e.ArrStart()
		for _, elem := range s.FieldErrors {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfValidationError = [6]string{
	0: "code",
	1: "message",
	2: "traceId",
	3: "timestamp",
	4: "path",
	5: "fieldErrors",
}

func (s *ValidationError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ValidationError to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Code = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "traceId":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.TraceId = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"traceId\"")
			}
		case "timestamp":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.Timestamp = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"timestamp\"")
			}
		case "path":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Path = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"path\"")
			}
		case "fieldErrors":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				s.FieldErrors = make([]FieldError, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem FieldError
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.FieldErrors = append(s.FieldErrors, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fieldErrors\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ValidationError")
	}
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00111111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfValidationError) {
					name = jsonFieldsNameOfValidationError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

func (s *ValidationError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

func (s *ValidationError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
