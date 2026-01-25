

package antifraud_v1

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

type APIV1AuthLoginPostBadRequest ApiError

func (*APIV1AuthLoginPostBadRequest) aPIV1AuthLoginPostRes() {}

type APIV1AuthLoginPostLocked ApiError

func (*APIV1AuthLoginPostLocked) aPIV1AuthLoginPostRes() {}

type APIV1AuthLoginPostUnauthorized ApiError

func (*APIV1AuthLoginPostUnauthorized) aPIV1AuthLoginPostRes() {}

type APIV1AuthRegisterPostBadRequest ApiError

func (*APIV1AuthRegisterPostBadRequest) aPIV1AuthRegisterPostRes() {}

type APIV1AuthRegisterPostConflict ApiError

func (*APIV1AuthRegisterPostConflict) aPIV1AuthRegisterPostRes() {}

type APIV1FraudRulesGetForbidden ApiError

func (*APIV1FraudRulesGetForbidden) aPIV1FraudRulesGetRes() {}

type APIV1FraudRulesGetOKApplicationJSON []FraudRule

func (*APIV1FraudRulesGetOKApplicationJSON) aPIV1FraudRulesGetRes() {}

type APIV1FraudRulesGetUnauthorized ApiError

func (*APIV1FraudRulesGetUnauthorized) aPIV1FraudRulesGetRes() {}

type APIV1FraudRulesIDDeleteForbidden ApiError

func (*APIV1FraudRulesIDDeleteForbidden) aPIV1FraudRulesIDDeleteRes() {}


type APIV1FraudRulesIDDeleteNoContent struct{}

func (*APIV1FraudRulesIDDeleteNoContent) aPIV1FraudRulesIDDeleteRes() {}

type APIV1FraudRulesIDDeleteNotFound ApiError

func (*APIV1FraudRulesIDDeleteNotFound) aPIV1FraudRulesIDDeleteRes() {}

type APIV1FraudRulesIDDeleteUnauthorized ApiError

func (*APIV1FraudRulesIDDeleteUnauthorized) aPIV1FraudRulesIDDeleteRes() {}

type APIV1FraudRulesIDGetForbidden ApiError

func (*APIV1FraudRulesIDGetForbidden) aPIV1FraudRulesIDGetRes() {}

type APIV1FraudRulesIDGetNotFound ApiError

func (*APIV1FraudRulesIDGetNotFound) aPIV1FraudRulesIDGetRes() {}

type APIV1FraudRulesIDGetUnauthorized ApiError

func (*APIV1FraudRulesIDGetUnauthorized) aPIV1FraudRulesIDGetRes() {}

type APIV1FraudRulesIDPutConflict ApiError

func (*APIV1FraudRulesIDPutConflict) aPIV1FraudRulesIDPutRes() {}

type APIV1FraudRulesIDPutForbidden ApiError

func (*APIV1FraudRulesIDPutForbidden) aPIV1FraudRulesIDPutRes() {}

type APIV1FraudRulesIDPutNotFound ApiError

func (*APIV1FraudRulesIDPutNotFound) aPIV1FraudRulesIDPutRes() {}

type APIV1FraudRulesIDPutUnauthorized ApiError

func (*APIV1FraudRulesIDPutUnauthorized) aPIV1FraudRulesIDPutRes() {}

type APIV1FraudRulesPostConflict ApiError

func (*APIV1FraudRulesPostConflict) aPIV1FraudRulesPostRes() {}

type APIV1FraudRulesPostForbidden ApiError

func (*APIV1FraudRulesPostForbidden) aPIV1FraudRulesPostRes() {}

type APIV1FraudRulesPostUnauthorized ApiError

func (*APIV1FraudRulesPostUnauthorized) aPIV1FraudRulesPostRes() {}

type APIV1FraudRulesValidatePostForbidden ApiError

func (*APIV1FraudRulesValidatePostForbidden) aPIV1FraudRulesValidatePostRes() {}

type APIV1FraudRulesValidatePostUnauthorized ApiError

func (*APIV1FraudRulesValidatePostUnauthorized) aPIV1FraudRulesValidatePostRes() {}

type APIV1PingGetOK struct {
	Status OptString `json:"status"`
}


func (s *APIV1PingGetOK) GetStatus() OptString {
	return s.Status
}


func (s *APIV1PingGetOK) SetStatus(val OptString) {
	s.Status = val
}

type APIV1StatsMerchantsRiskGetForbidden ApiError

func (*APIV1StatsMerchantsRiskGetForbidden) aPIV1StatsMerchantsRiskGetRes() {}

type APIV1StatsMerchantsRiskGetUnauthorized ApiError

func (*APIV1StatsMerchantsRiskGetUnauthorized) aPIV1StatsMerchantsRiskGetRes() {}

type APIV1StatsOverviewGetForbidden ApiError

func (*APIV1StatsOverviewGetForbidden) aPIV1StatsOverviewGetRes() {}

type APIV1StatsOverviewGetUnauthorized ApiError

func (*APIV1StatsOverviewGetUnauthorized) aPIV1StatsOverviewGetRes() {}

type APIV1StatsRulesMatchesGetForbidden ApiError

func (*APIV1StatsRulesMatchesGetForbidden) aPIV1StatsRulesMatchesGetRes() {}

type APIV1StatsRulesMatchesGetUnauthorized ApiError

func (*APIV1StatsRulesMatchesGetUnauthorized) aPIV1StatsRulesMatchesGetRes() {}

type APIV1StatsTransactionsTimeseriesGetForbidden ApiError

func (*APIV1StatsTransactionsTimeseriesGetForbidden) aPIV1StatsTransactionsTimeseriesGetRes() {}

type APIV1StatsTransactionsTimeseriesGetUnauthorized ApiError

func (*APIV1StatsTransactionsTimeseriesGetUnauthorized) aPIV1StatsTransactionsTimeseriesGetRes() {}

type APIV1StatsUsersIDRiskProfileGetForbidden ApiError

func (*APIV1StatsUsersIDRiskProfileGetForbidden) aPIV1StatsUsersIDRiskProfileGetRes() {}

type APIV1StatsUsersIDRiskProfileGetNotFound ApiError

func (*APIV1StatsUsersIDRiskProfileGetNotFound) aPIV1StatsUsersIDRiskProfileGetRes() {}

type APIV1StatsUsersIDRiskProfileGetUnauthorized ApiError

func (*APIV1StatsUsersIDRiskProfileGetUnauthorized) aPIV1StatsUsersIDRiskProfileGetRes() {}

type APIV1TransactionsBatchPostCreated TransactionBatchResult

func (*APIV1TransactionsBatchPostCreated) aPIV1TransactionsBatchPostRes() {}

type APIV1TransactionsBatchPostMultiStatus TransactionBatchResult

func (*APIV1TransactionsBatchPostMultiStatus) aPIV1TransactionsBatchPostRes() {}

type APIV1TransactionsIDGetForbidden ApiError

func (*APIV1TransactionsIDGetForbidden) aPIV1TransactionsIDGetRes() {}

type APIV1TransactionsIDGetNotFound ApiError

func (*APIV1TransactionsIDGetNotFound) aPIV1TransactionsIDGetRes() {}

type APIV1TransactionsIDGetUnauthorized ApiError

func (*APIV1TransactionsIDGetUnauthorized) aPIV1TransactionsIDGetRes() {}

type APIV1TransactionsPostBadRequest ApiError

func (*APIV1TransactionsPostBadRequest) aPIV1TransactionsPostRes() {}

type APIV1TransactionsPostForbidden ApiError

func (*APIV1TransactionsPostForbidden) aPIV1TransactionsPostRes() {}

type APIV1TransactionsPostNotFound ApiError

func (*APIV1TransactionsPostNotFound) aPIV1TransactionsPostRes() {}

type APIV1TransactionsPostUnauthorized ApiError

func (*APIV1TransactionsPostUnauthorized) aPIV1TransactionsPostRes() {}

type APIV1UsersGetForbidden ApiError

func (*APIV1UsersGetForbidden) aPIV1UsersGetRes() {}

type APIV1UsersGetUnauthorized ApiError

func (*APIV1UsersGetUnauthorized) aPIV1UsersGetRes() {}

type APIV1UsersIDDeleteForbidden ApiError

func (*APIV1UsersIDDeleteForbidden) aPIV1UsersIDDeleteRes() {}


type APIV1UsersIDDeleteNoContent struct{}

func (*APIV1UsersIDDeleteNoContent) aPIV1UsersIDDeleteRes() {}

type APIV1UsersIDDeleteNotFound ApiError

func (*APIV1UsersIDDeleteNotFound) aPIV1UsersIDDeleteRes() {}

type APIV1UsersIDDeleteUnauthorized ApiError

func (*APIV1UsersIDDeleteUnauthorized) aPIV1UsersIDDeleteRes() {}

type APIV1UsersIDGetForbidden ApiError

func (*APIV1UsersIDGetForbidden) aPIV1UsersIDGetRes() {}

type APIV1UsersIDGetNotFound ApiError

func (*APIV1UsersIDGetNotFound) aPIV1UsersIDGetRes() {}

type APIV1UsersIDGetUnauthorized ApiError

func (*APIV1UsersIDGetUnauthorized) aPIV1UsersIDGetRes() {}

type APIV1UsersIDPutForbidden ApiError

func (*APIV1UsersIDPutForbidden) aPIV1UsersIDPutRes() {}

type APIV1UsersIDPutNotFound ApiError

func (*APIV1UsersIDPutNotFound) aPIV1UsersIDPutRes() {}

type APIV1UsersIDPutUnauthorized ApiError

func (*APIV1UsersIDPutUnauthorized) aPIV1UsersIDPutRes() {}

type APIV1UsersMePutForbidden ApiError

func (*APIV1UsersMePutForbidden) aPIV1UsersMePutRes() {}

type APIV1UsersMePutUnauthorized ApiError

func (*APIV1UsersMePutUnauthorized) aPIV1UsersMePutRes() {}

type APIV1UsersPostConflict ApiError

func (*APIV1UsersPostConflict) aPIV1UsersPostRes() {}

type APIV1UsersPostForbidden ApiError

func (*APIV1UsersPostForbidden) aPIV1UsersPostRes() {}

type APIV1UsersPostUnauthorized ApiError

func (*APIV1UsersPostUnauthorized) aPIV1UsersPostRes() {}




type ApiError struct {
	Code ErrorCode `json:"code"`
	
	
	Message string `json:"message"`
	
	
	TraceId uuid.UUID `json:"traceId"`
	
	Timestamp time.Time `json:"timestamp"`
	
	Path string `json:"path"`
	
	Details OptApiErrorDetails `json:"details"`
}


func (s *ApiError) GetCode() ErrorCode {
	return s.Code
}


func (s *ApiError) GetMessage() string {
	return s.Message
}


func (s *ApiError) GetTraceId() uuid.UUID {
	return s.TraceId
}


func (s *ApiError) GetTimestamp() time.Time {
	return s.Timestamp
}


func (s *ApiError) GetPath() string {
	return s.Path
}


func (s *ApiError) GetDetails() OptApiErrorDetails {
	return s.Details
}


func (s *ApiError) SetCode(val ErrorCode) {
	s.Code = val
}


func (s *ApiError) SetMessage(val string) {
	s.Message = val
}


func (s *ApiError) SetTraceId(val uuid.UUID) {
	s.TraceId = val
}


func (s *ApiError) SetTimestamp(val time.Time) {
	s.Timestamp = val
}


func (s *ApiError) SetPath(val string) {
	s.Path = val
}


func (s *ApiError) SetDetails(val OptApiErrorDetails) {
	s.Details = val
}

func (*ApiError) aPIV1TransactionsBatchPostRes() {}
func (*ApiError) aPIV1TransactionsGetRes()       {}
func (*ApiError) aPIV1UsersMeGetRes()            {}


type ApiErrorDetails map[string]jx.Raw

func (s *ApiErrorDetails) init() ApiErrorDetails {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}





type AuthResponse struct {
	
	
	AccessToken string `json:"accessToken"`
	
	
	
	ExpiresIn int  `json:"expiresIn"`
	User      User `json:"user"`
}


func (s *AuthResponse) GetAccessToken() string {
	return s.AccessToken
}


func (s *AuthResponse) GetExpiresIn() int {
	return s.ExpiresIn
}


func (s *AuthResponse) GetUser() User {
	return s.User
}


func (s *AuthResponse) SetAccessToken(val string) {
	s.AccessToken = val
}


func (s *AuthResponse) SetExpiresIn(val int) {
	s.ExpiresIn = val
}


func (s *AuthResponse) SetUser(val User) {
	s.User = val
}

func (*AuthResponse) aPIV1AuthLoginPostRes()    {}
func (*AuthResponse) aPIV1AuthRegisterPostRes() {}

type BearerAuth struct {
	Token string
}


func (s *BearerAuth) GetToken() string {
	return s.Token
}


func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

type CurrencyCode string


type DslError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	
	Position OptNilInt    `json:"position"`
	Near     OptNilString `json:"near"`
}


func (s *DslError) GetCode() string {
	return s.Code
}


func (s *DslError) GetMessage() string {
	return s.Message
}


func (s *DslError) GetPosition() OptNilInt {
	return s.Position
}


func (s *DslError) GetNear() OptNilString {
	return s.Near
}


func (s *DslError) SetCode(val string) {
	s.Code = val
}


func (s *DslError) SetMessage(val string) {
	s.Message = val
}


func (s *DslError) SetPosition(val OptNilInt) {
	s.Position = val
}


func (s *DslError) SetNear(val OptNilString) {
	s.Near = val
}


type DslValidateRequest struct {
	DslExpression string `json:"dslExpression"`
}


func (s *DslValidateRequest) GetDslExpression() string {
	return s.DslExpression
}


func (s *DslValidateRequest) SetDslExpression(val string) {
	s.DslExpression = val
}


type DslValidateResponse struct {
	IsValid bool `json:"isValid"`
	
	NormalizedExpression OptNilString `json:"normalizedExpression"`
	Errors               []DslError   `json:"errors"`
}


func (s *DslValidateResponse) GetIsValid() bool {
	return s.IsValid
}


func (s *DslValidateResponse) GetNormalizedExpression() OptNilString {
	return s.NormalizedExpression
}


func (s *DslValidateResponse) GetErrors() []DslError {
	return s.Errors
}


func (s *DslValidateResponse) SetIsValid(val bool) {
	s.IsValid = val
}


func (s *DslValidateResponse) SetNormalizedExpression(val OptNilString) {
	s.NormalizedExpression = val
}


func (s *DslValidateResponse) SetErrors(val []DslError) {
	s.Errors = val
}

func (*DslValidateResponse) aPIV1FraudRulesValidatePostRes() {}






















type ErrorCode string

const (
	ErrorCodeBADREQUEST            ErrorCode = "BAD_REQUEST"
	ErrorCodeVALIDATIONFAILED      ErrorCode = "VALIDATION_FAILED"
	ErrorCodeUNAUTHORIZED          ErrorCode = "UNAUTHORIZED"
	ErrorCodeFORBIDDEN             ErrorCode = "FORBIDDEN"
	ErrorCodeNOTFOUND              ErrorCode = "NOT_FOUND"
	ErrorCodeEMAILALREADYEXISTS    ErrorCode = "EMAIL_ALREADY_EXISTS"
	ErrorCodeUSERINACTIVE          ErrorCode = "USER_INACTIVE"
	ErrorCodeRULENAMEALREADYEXISTS ErrorCode = "RULE_NAME_ALREADY_EXISTS"
	ErrorCodeDSLPARSEERROR         ErrorCode = "DSL_PARSE_ERROR"
	ErrorCodeDSLINVALIDFIELD       ErrorCode = "DSL_INVALID_FIELD"
	ErrorCodeDSLINVALIDOPERATOR    ErrorCode = "DSL_INVALID_OPERATOR"
	ErrorCodeINTERNALSERVERERROR   ErrorCode = "INTERNAL_SERVER_ERROR"
)


func (ErrorCode) AllValues() []ErrorCode {
	return []ErrorCode{
		ErrorCodeBADREQUEST,
		ErrorCodeVALIDATIONFAILED,
		ErrorCodeUNAUTHORIZED,
		ErrorCodeFORBIDDEN,
		ErrorCodeNOTFOUND,
		ErrorCodeEMAILALREADYEXISTS,
		ErrorCodeUSERINACTIVE,
		ErrorCodeRULENAMEALREADYEXISTS,
		ErrorCodeDSLPARSEERROR,
		ErrorCodeDSLINVALIDFIELD,
		ErrorCodeDSLINVALIDOPERATOR,
		ErrorCodeINTERNALSERVERERROR,
	}
}


func (s ErrorCode) MarshalText() ([]byte, error) {
	switch s {
	case ErrorCodeBADREQUEST:
		return []byte(s), nil
	case ErrorCodeVALIDATIONFAILED:
		return []byte(s), nil
	case ErrorCodeUNAUTHORIZED:
		return []byte(s), nil
	case ErrorCodeFORBIDDEN:
		return []byte(s), nil
	case ErrorCodeNOTFOUND:
		return []byte(s), nil
	case ErrorCodeEMAILALREADYEXISTS:
		return []byte(s), nil
	case ErrorCodeUSERINACTIVE:
		return []byte(s), nil
	case ErrorCodeRULENAMEALREADYEXISTS:
		return []byte(s), nil
	case ErrorCodeDSLPARSEERROR:
		return []byte(s), nil
	case ErrorCodeDSLINVALIDFIELD:
		return []byte(s), nil
	case ErrorCodeDSLINVALIDOPERATOR:
		return []byte(s), nil
	case ErrorCodeINTERNALSERVERERROR:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *ErrorCode) UnmarshalText(data []byte) error {
	switch ErrorCode(data) {
	case ErrorCodeBADREQUEST:
		*s = ErrorCodeBADREQUEST
		return nil
	case ErrorCodeVALIDATIONFAILED:
		*s = ErrorCodeVALIDATIONFAILED
		return nil
	case ErrorCodeUNAUTHORIZED:
		*s = ErrorCodeUNAUTHORIZED
		return nil
	case ErrorCodeFORBIDDEN:
		*s = ErrorCodeFORBIDDEN
		return nil
	case ErrorCodeNOTFOUND:
		*s = ErrorCodeNOTFOUND
		return nil
	case ErrorCodeEMAILALREADYEXISTS:
		*s = ErrorCodeEMAILALREADYEXISTS
		return nil
	case ErrorCodeUSERINACTIVE:
		*s = ErrorCodeUSERINACTIVE
		return nil
	case ErrorCodeRULENAMEALREADYEXISTS:
		*s = ErrorCodeRULENAMEALREADYEXISTS
		return nil
	case ErrorCodeDSLPARSEERROR:
		*s = ErrorCodeDSLPARSEERROR
		return nil
	case ErrorCodeDSLINVALIDFIELD:
		*s = ErrorCodeDSLINVALIDFIELD
		return nil
	case ErrorCodeDSLINVALIDOPERATOR:
		*s = ErrorCodeDSLINVALIDOPERATOR
		return nil
	case ErrorCodeINTERNALSERVERERROR:
		*s = ErrorCodeINTERNALSERVERERROR
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}


type FieldError struct {
	
	
	Field string `json:"field"`
	
	Issue string `json:"issue"`
	
	RejectedValue jx.Raw `json:"rejectedValue"`
}


func (s *FieldError) GetField() string {
	return s.Field
}


func (s *FieldError) GetIssue() string {
	return s.Issue
}


func (s *FieldError) GetRejectedValue() jx.Raw {
	return s.RejectedValue
}


func (s *FieldError) SetField(val string) {
	s.Field = val
}


func (s *FieldError) SetIssue(val string) {
	s.Issue = val
}


func (s *FieldError) SetRejectedValue(val jx.Raw) {
	s.RejectedValue = val
}


type FraudRule struct {
	ID uuid.UUID `json:"id"`
	
	Name string `json:"name"`
	
	Description OptString `json:"description"`
	
	
	
	
	
	
	
	
	
	
	
	DslExpression string `json:"dslExpression"`
	
	Enabled bool `json:"enabled"`
	
	
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


func (s *FraudRule) GetID() uuid.UUID {
	return s.ID
}


func (s *FraudRule) GetName() string {
	return s.Name
}


func (s *FraudRule) GetDescription() OptString {
	return s.Description
}


func (s *FraudRule) GetDslExpression() string {
	return s.DslExpression
}


func (s *FraudRule) GetEnabled() bool {
	return s.Enabled
}


func (s *FraudRule) GetPriority() int {
	return s.Priority
}


func (s *FraudRule) GetCreatedAt() time.Time {
	return s.CreatedAt
}


func (s *FraudRule) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}


func (s *FraudRule) SetID(val uuid.UUID) {
	s.ID = val
}


func (s *FraudRule) SetName(val string) {
	s.Name = val
}


func (s *FraudRule) SetDescription(val OptString) {
	s.Description = val
}


func (s *FraudRule) SetDslExpression(val string) {
	s.DslExpression = val
}


func (s *FraudRule) SetEnabled(val bool) {
	s.Enabled = val
}


func (s *FraudRule) SetPriority(val int) {
	s.Priority = val
}


func (s *FraudRule) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}


func (s *FraudRule) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*FraudRule) aPIV1FraudRulesIDGetRes() {}
func (*FraudRule) aPIV1FraudRulesIDPutRes() {}
func (*FraudRule) aPIV1FraudRulesPostRes()  {}


type FraudRuleCreateRequest struct {
	Name          string    `json:"name"`
	Description   OptString `json:"description"`
	DslExpression string    `json:"dslExpression"`
	Enabled       OptBool   `json:"enabled"`
	Priority      OptInt    `json:"priority"`
}


func (s *FraudRuleCreateRequest) GetName() string {
	return s.Name
}


func (s *FraudRuleCreateRequest) GetDescription() OptString {
	return s.Description
}


func (s *FraudRuleCreateRequest) GetDslExpression() string {
	return s.DslExpression
}


func (s *FraudRuleCreateRequest) GetEnabled() OptBool {
	return s.Enabled
}


func (s *FraudRuleCreateRequest) GetPriority() OptInt {
	return s.Priority
}


func (s *FraudRuleCreateRequest) SetName(val string) {
	s.Name = val
}


func (s *FraudRuleCreateRequest) SetDescription(val OptString) {
	s.Description = val
}


func (s *FraudRuleCreateRequest) SetDslExpression(val string) {
	s.DslExpression = val
}


func (s *FraudRuleCreateRequest) SetEnabled(val OptBool) {
	s.Enabled = val
}


func (s *FraudRuleCreateRequest) SetPriority(val OptInt) {
	s.Priority = val
}




type FraudRuleEvaluationResult struct {
	RuleId   uuid.UUID `json:"ruleId"`
	RuleName string    `json:"ruleName"`
	Priority int       `json:"priority"`
	
	Matched bool `json:"matched"`
	
	Description string `json:"description"`
}


func (s *FraudRuleEvaluationResult) GetRuleId() uuid.UUID {
	return s.RuleId
}


func (s *FraudRuleEvaluationResult) GetRuleName() string {
	return s.RuleName
}


func (s *FraudRuleEvaluationResult) GetPriority() int {
	return s.Priority
}


func (s *FraudRuleEvaluationResult) GetMatched() bool {
	return s.Matched
}


func (s *FraudRuleEvaluationResult) GetDescription() string {
	return s.Description
}


func (s *FraudRuleEvaluationResult) SetRuleId(val uuid.UUID) {
	s.RuleId = val
}


func (s *FraudRuleEvaluationResult) SetRuleName(val string) {
	s.RuleName = val
}


func (s *FraudRuleEvaluationResult) SetPriority(val int) {
	s.Priority = val
}


func (s *FraudRuleEvaluationResult) SetMatched(val bool) {
	s.Matched = val
}


func (s *FraudRuleEvaluationResult) SetDescription(val string) {
	s.Description = val
}




type FraudRuleUpdateRequest struct {
	Name          string    `json:"name"`
	Description   OptString `json:"description"`
	DslExpression string    `json:"dslExpression"`
	Enabled       bool      `json:"enabled"`
	Priority      int       `json:"priority"`
}


func (s *FraudRuleUpdateRequest) GetName() string {
	return s.Name
}


func (s *FraudRuleUpdateRequest) GetDescription() OptString {
	return s.Description
}


func (s *FraudRuleUpdateRequest) GetDslExpression() string {
	return s.DslExpression
}


func (s *FraudRuleUpdateRequest) GetEnabled() bool {
	return s.Enabled
}


func (s *FraudRuleUpdateRequest) GetPriority() int {
	return s.Priority
}


func (s *FraudRuleUpdateRequest) SetName(val string) {
	s.Name = val
}


func (s *FraudRuleUpdateRequest) SetDescription(val OptString) {
	s.Description = val
}


func (s *FraudRuleUpdateRequest) SetDslExpression(val string) {
	s.DslExpression = val
}


func (s *FraudRuleUpdateRequest) SetEnabled(val bool) {
	s.Enabled = val
}


func (s *FraudRuleUpdateRequest) SetPriority(val int) {
	s.Priority = val
}


type Gender string

const (
	GenderMALE   Gender = "MALE"
	GenderFEMALE Gender = "FEMALE"
)


func (Gender) AllValues() []Gender {
	return []Gender{
		GenderMALE,
		GenderFEMALE,
	}
}


func (s Gender) MarshalText() ([]byte, error) {
	switch s {
	case GenderMALE:
		return []byte(s), nil
	case GenderFEMALE:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *Gender) UnmarshalText(data []byte) error {
	switch Gender(data) {
	case GenderMALE:
		*s = GenderMALE
		return nil
	case GenderFEMALE:
		*s = GenderFEMALE
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type GroupBy string

const (
	GroupByHour GroupBy = "hour"
	GroupByDay  GroupBy = "day"
	GroupByWeek GroupBy = "week"
)


func (GroupBy) AllValues() []GroupBy {
	return []GroupBy{
		GroupByHour,
		GroupByDay,
		GroupByWeek,
	}
}


func (s GroupBy) MarshalText() ([]byte, error) {
	switch s {
	case GroupByHour:
		return []byte(s), nil
	case GroupByDay:
		return []byte(s), nil
	case GroupByWeek:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *GroupBy) UnmarshalText(data []byte) error {
	switch GroupBy(data) {
	case GroupByHour:
		*s = GroupByHour
		return nil
	case GroupByDay:
		*s = GroupByDay
		return nil
	case GroupByWeek:
		*s = GroupByWeek
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


func (s *LoginRequest) GetEmail() string {
	return s.Email
}


func (s *LoginRequest) GetPassword() string {
	return s.Password
}


func (s *LoginRequest) SetEmail(val string) {
	s.Email = val
}


func (s *LoginRequest) SetPassword(val string) {
	s.Password = val
}


type MaritalStatus string

const (
	MaritalStatusSINGLE   MaritalStatus = "SINGLE"
	MaritalStatusMARRIED  MaritalStatus = "MARRIED"
	MaritalStatusDIVORCED MaritalStatus = "DIVORCED"
	MaritalStatusWIDOWED  MaritalStatus = "WIDOWED"
)


func (MaritalStatus) AllValues() []MaritalStatus {
	return []MaritalStatus{
		MaritalStatusSINGLE,
		MaritalStatusMARRIED,
		MaritalStatusDIVORCED,
		MaritalStatusWIDOWED,
	}
}


func (s MaritalStatus) MarshalText() ([]byte, error) {
	switch s {
	case MaritalStatusSINGLE:
		return []byte(s), nil
	case MaritalStatusMARRIED:
		return []byte(s), nil
	case MaritalStatusDIVORCED:
		return []byte(s), nil
	case MaritalStatusWIDOWED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *MaritalStatus) UnmarshalText(data []byte) error {
	switch MaritalStatus(data) {
	case MaritalStatusSINGLE:
		*s = MaritalStatusSINGLE
		return nil
	case MaritalStatusMARRIED:
		*s = MaritalStatusMARRIED
		return nil
	case MaritalStatusDIVORCED:
		*s = MaritalStatusDIVORCED
		return nil
	case MaritalStatusWIDOWED:
		*s = MaritalStatusWIDOWED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type MccCode string


type MerchantRiskRow struct {
	MerchantId           string     `json:"merchantId"`
	MerchantCategoryCode OptMccCode `json:"merchantCategoryCode"`
	
	TxCount int `json:"txCount"`
	
	
	
	Gmv float64 `json:"gmv"`
	
	DeclineRate float64 `json:"declineRate"`
}


func (s *MerchantRiskRow) GetMerchantId() string {
	return s.MerchantId
}


func (s *MerchantRiskRow) GetMerchantCategoryCode() OptMccCode {
	return s.MerchantCategoryCode
}


func (s *MerchantRiskRow) GetTxCount() int {
	return s.TxCount
}


func (s *MerchantRiskRow) GetGmv() float64 {
	return s.Gmv
}


func (s *MerchantRiskRow) GetDeclineRate() float64 {
	return s.DeclineRate
}


func (s *MerchantRiskRow) SetMerchantId(val string) {
	s.MerchantId = val
}


func (s *MerchantRiskRow) SetMerchantCategoryCode(val OptMccCode) {
	s.MerchantCategoryCode = val
}


func (s *MerchantRiskRow) SetTxCount(val int) {
	s.TxCount = val
}


func (s *MerchantRiskRow) SetGmv(val float64) {
	s.Gmv = val
}


func (s *MerchantRiskRow) SetDeclineRate(val float64) {
	s.DeclineRate = val
}


type MerchantRiskStats struct {
	Items []MerchantRiskRow `json:"items"`
}


func (s *MerchantRiskStats) GetItems() []MerchantRiskRow {
	return s.Items
}


func (s *MerchantRiskStats) SetItems(val []MerchantRiskRow) {
	s.Items = val
}

func (*MerchantRiskStats) aPIV1StatsMerchantsRiskGetRes() {}


func NewNilGender(v Gender) NilGender {
	return NilGender{
		Value: v,
	}
}


type NilGender struct {
	Value Gender
	Null  bool
}


func (o *NilGender) SetTo(v Gender) {
	o.Null = false
	o.Value = v
}


func (o NilGender) IsNull() bool { return o.Null }


func (o *NilGender) SetToNull() {
	o.Null = true
	var v Gender
	o.Value = v
}


func (o NilGender) Get() (v Gender, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}


func (o NilGender) Or(d Gender) Gender {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}


type NilInt struct {
	Value int
	Null  bool
}


func (o *NilInt) SetTo(v int) {
	o.Null = false
	o.Value = v
}


func (o NilInt) IsNull() bool { return o.Null }


func (o *NilInt) SetToNull() {
	o.Null = true
	var v int
	o.Value = v
}


func (o NilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}


func (o NilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewNilMaritalStatus(v MaritalStatus) NilMaritalStatus {
	return NilMaritalStatus{
		Value: v,
	}
}


type NilMaritalStatus struct {
	Value MaritalStatus
	Null  bool
}


func (o *NilMaritalStatus) SetTo(v MaritalStatus) {
	o.Null = false
	o.Value = v
}


func (o NilMaritalStatus) IsNull() bool { return o.Null }


func (o *NilMaritalStatus) SetToNull() {
	o.Null = true
	var v MaritalStatus
	o.Value = v
}


func (o NilMaritalStatus) Get() (v MaritalStatus, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}


func (o NilMaritalStatus) Or(d MaritalStatus) MaritalStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}


type NilString struct {
	Value string
	Null  bool
}


func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}


func (o NilString) IsNull() bool { return o.Null }


func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}


func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}


func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptApiError(v ApiError) OptApiError {
	return OptApiError{
		Value: v,
		Set:   true,
	}
}


type OptApiError struct {
	Value ApiError
	Set   bool
}


func (o OptApiError) IsSet() bool { return o.Set }


func (o *OptApiError) Reset() {
	var v ApiError
	o.Value = v
	o.Set = false
}


func (o *OptApiError) SetTo(v ApiError) {
	o.Set = true
	o.Value = v
}


func (o OptApiError) Get() (v ApiError, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptApiError) Or(d ApiError) ApiError {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptApiErrorDetails(v ApiErrorDetails) OptApiErrorDetails {
	return OptApiErrorDetails{
		Value: v,
		Set:   true,
	}
}


type OptApiErrorDetails struct {
	Value ApiErrorDetails
	Set   bool
}


func (o OptApiErrorDetails) IsSet() bool { return o.Set }


func (o *OptApiErrorDetails) Reset() {
	var v ApiErrorDetails
	o.Value = v
	o.Set = false
}


func (o *OptApiErrorDetails) SetTo(v ApiErrorDetails) {
	o.Set = true
	o.Value = v
}


func (o OptApiErrorDetails) Get() (v ApiErrorDetails, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptApiErrorDetails) Or(d ApiErrorDetails) ApiErrorDetails {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}


type OptBool struct {
	Value bool
	Set   bool
}


func (o OptBool) IsSet() bool { return o.Set }


func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}


func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}


func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}


type OptDateTime struct {
	Value time.Time
	Set   bool
}


func (o OptDateTime) IsSet() bool { return o.Set }


func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}


func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}


func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}


type OptFloat64 struct {
	Value float64
	Set   bool
}


func (o OptFloat64) IsSet() bool { return o.Set }


func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}


func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}


func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptGender(v Gender) OptGender {
	return OptGender{
		Value: v,
		Set:   true,
	}
}


type OptGender struct {
	Value Gender
	Set   bool
}


func (o OptGender) IsSet() bool { return o.Set }


func (o *OptGender) Reset() {
	var v Gender
	o.Value = v
	o.Set = false
}


func (o *OptGender) SetTo(v Gender) {
	o.Set = true
	o.Value = v
}


func (o OptGender) Get() (v Gender, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptGender) Or(d Gender) Gender {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptGroupBy(v GroupBy) OptGroupBy {
	return OptGroupBy{
		Value: v,
		Set:   true,
	}
}


type OptGroupBy struct {
	Value GroupBy
	Set   bool
}


func (o OptGroupBy) IsSet() bool { return o.Set }


func (o *OptGroupBy) Reset() {
	var v GroupBy
	o.Value = v
	o.Set = false
}


func (o *OptGroupBy) SetTo(v GroupBy) {
	o.Set = true
	o.Value = v
}


func (o OptGroupBy) Get() (v GroupBy, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptGroupBy) Or(d GroupBy) GroupBy {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}


type OptInt struct {
	Value int
	Set   bool
}


func (o OptInt) IsSet() bool { return o.Set }


func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}


func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}


func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptMaritalStatus(v MaritalStatus) OptMaritalStatus {
	return OptMaritalStatus{
		Value: v,
		Set:   true,
	}
}


type OptMaritalStatus struct {
	Value MaritalStatus
	Set   bool
}


func (o OptMaritalStatus) IsSet() bool { return o.Set }


func (o *OptMaritalStatus) Reset() {
	var v MaritalStatus
	o.Value = v
	o.Set = false
}


func (o *OptMaritalStatus) SetTo(v MaritalStatus) {
	o.Set = true
	o.Value = v
}


func (o OptMaritalStatus) Get() (v MaritalStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptMaritalStatus) Or(d MaritalStatus) MaritalStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptMccCode(v MccCode) OptMccCode {
	return OptMccCode{
		Value: v,
		Set:   true,
	}
}


type OptMccCode struct {
	Value MccCode
	Set   bool
}


func (o OptMccCode) IsSet() bool { return o.Set }


func (o *OptMccCode) Reset() {
	var v MccCode
	o.Value = v
	o.Set = false
}


func (o *OptMccCode) SetTo(v MccCode) {
	o.Set = true
	o.Value = v
}


func (o OptMccCode) Get() (v MccCode, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptMccCode) Or(d MccCode) MccCode {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptNilDateTime(v time.Time) OptNilDateTime {
	return OptNilDateTime{
		Value: v,
		Set:   true,
	}
}


type OptNilDateTime struct {
	Value time.Time
	Set   bool
	Null  bool
}


func (o OptNilDateTime) IsSet() bool { return o.Set }


func (o *OptNilDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
	o.Null = false
}


func (o *OptNilDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Null = false
	o.Value = v
}


func (o OptNilDateTime) IsNull() bool { return o.Null }


func (o *OptNilDateTime) SetToNull() {
	o.Set = true
	o.Null = true
	var v time.Time
	o.Value = v
}


func (o OptNilDateTime) Get() (v time.Time, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptNilDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptNilInt(v int) OptNilInt {
	return OptNilInt{
		Value: v,
		Set:   true,
	}
}


type OptNilInt struct {
	Value int
	Set   bool
	Null  bool
}


func (o OptNilInt) IsSet() bool { return o.Set }


func (o *OptNilInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
	o.Null = false
}


func (o *OptNilInt) SetTo(v int) {
	o.Set = true
	o.Null = false
	o.Value = v
}


func (o OptNilInt) IsNull() bool { return o.Null }


func (o *OptNilInt) SetToNull() {
	o.Set = true
	o.Null = true
	var v int
	o.Value = v
}


func (o OptNilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptNilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptNilString(v string) OptNilString {
	return OptNilString{
		Value: v,
		Set:   true,
	}
}


type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}


func (o OptNilString) IsSet() bool { return o.Set }


func (o *OptNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Null = false
}


func (o *OptNilString) SetTo(v string) {
	o.Set = true
	o.Null = false
	o.Value = v
}


func (o OptNilString) IsNull() bool { return o.Null }


func (o *OptNilString) SetToNull() {
	o.Set = true
	o.Null = true
	var v string
	o.Value = v
}


func (o OptNilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptNilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}


type OptString struct {
	Value string
	Set   bool
}


func (o OptString) IsSet() bool { return o.Set }


func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}


func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}


func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionChannel(v TransactionChannel) OptTransactionChannel {
	return OptTransactionChannel{
		Value: v,
		Set:   true,
	}
}


type OptTransactionChannel struct {
	Value TransactionChannel
	Set   bool
}


func (o OptTransactionChannel) IsSet() bool { return o.Set }


func (o *OptTransactionChannel) Reset() {
	var v TransactionChannel
	o.Value = v
	o.Set = false
}


func (o *OptTransactionChannel) SetTo(v TransactionChannel) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionChannel) Get() (v TransactionChannel, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionChannel) Or(d TransactionChannel) TransactionChannel {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionCreateRequestMetadata(v TransactionCreateRequestMetadata) OptTransactionCreateRequestMetadata {
	return OptTransactionCreateRequestMetadata{
		Value: v,
		Set:   true,
	}
}


type OptTransactionCreateRequestMetadata struct {
	Value TransactionCreateRequestMetadata
	Set   bool
}


func (o OptTransactionCreateRequestMetadata) IsSet() bool { return o.Set }


func (o *OptTransactionCreateRequestMetadata) Reset() {
	var v TransactionCreateRequestMetadata
	o.Value = v
	o.Set = false
}


func (o *OptTransactionCreateRequestMetadata) SetTo(v TransactionCreateRequestMetadata) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionCreateRequestMetadata) Get() (v TransactionCreateRequestMetadata, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionCreateRequestMetadata) Or(d TransactionCreateRequestMetadata) TransactionCreateRequestMetadata {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionDecision(v TransactionDecision) OptTransactionDecision {
	return OptTransactionDecision{
		Value: v,
		Set:   true,
	}
}


type OptTransactionDecision struct {
	Value TransactionDecision
	Set   bool
}


func (o OptTransactionDecision) IsSet() bool { return o.Set }


func (o *OptTransactionDecision) Reset() {
	var v TransactionDecision
	o.Value = v
	o.Set = false
}


func (o *OptTransactionDecision) SetTo(v TransactionDecision) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionDecision) Get() (v TransactionDecision, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionDecision) Or(d TransactionDecision) TransactionDecision {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionLocation(v TransactionLocation) OptTransactionLocation {
	return OptTransactionLocation{
		Value: v,
		Set:   true,
	}
}


type OptTransactionLocation struct {
	Value TransactionLocation
	Set   bool
}


func (o OptTransactionLocation) IsSet() bool { return o.Set }


func (o *OptTransactionLocation) Reset() {
	var v TransactionLocation
	o.Value = v
	o.Set = false
}


func (o *OptTransactionLocation) SetTo(v TransactionLocation) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionLocation) Get() (v TransactionLocation, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionLocation) Or(d TransactionLocation) TransactionLocation {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionMetadata(v TransactionMetadata) OptTransactionMetadata {
	return OptTransactionMetadata{
		Value: v,
		Set:   true,
	}
}


type OptTransactionMetadata struct {
	Value TransactionMetadata
	Set   bool
}


func (o OptTransactionMetadata) IsSet() bool { return o.Set }


func (o *OptTransactionMetadata) Reset() {
	var v TransactionMetadata
	o.Value = v
	o.Set = false
}


func (o *OptTransactionMetadata) SetTo(v TransactionMetadata) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionMetadata) Get() (v TransactionMetadata, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionMetadata) Or(d TransactionMetadata) TransactionMetadata {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptTransactionStatus(v TransactionStatus) OptTransactionStatus {
	return OptTransactionStatus{
		Value: v,
		Set:   true,
	}
}


type OptTransactionStatus struct {
	Value TransactionStatus
	Set   bool
}


func (o OptTransactionStatus) IsSet() bool { return o.Set }


func (o *OptTransactionStatus) Reset() {
	var v TransactionStatus
	o.Value = v
	o.Set = false
}


func (o *OptTransactionStatus) SetTo(v TransactionStatus) {
	o.Set = true
	o.Value = v
}


func (o OptTransactionStatus) Get() (v TransactionStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptTransactionStatus) Or(d TransactionStatus) TransactionStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}


type OptUUID struct {
	Value uuid.UUID
	Set   bool
}


func (o OptUUID) IsSet() bool { return o.Set }


func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}


func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}


func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptUUID) Or(d uuid.UUID) uuid.UUID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


func NewOptUserRole(v UserRole) OptUserRole {
	return OptUserRole{
		Value: v,
		Set:   true,
	}
}


type OptUserRole struct {
	Value UserRole
	Set   bool
}


func (o OptUserRole) IsSet() bool { return o.Set }


func (o *OptUserRole) Reset() {
	var v UserRole
	o.Value = v
	o.Set = false
}


func (o *OptUserRole) SetTo(v UserRole) {
	o.Set = true
	o.Value = v
}


func (o OptUserRole) Get() (v UserRole, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}


func (o OptUserRole) Or(d UserRole) UserRole {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


type PagedTransactions struct {
	Items []Transaction `json:"items"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
}


func (s *PagedTransactions) GetItems() []Transaction {
	return s.Items
}


func (s *PagedTransactions) GetTotal() int {
	return s.Total
}


func (s *PagedTransactions) GetPage() int {
	return s.Page
}


func (s *PagedTransactions) GetSize() int {
	return s.Size
}


func (s *PagedTransactions) SetItems(val []Transaction) {
	s.Items = val
}


func (s *PagedTransactions) SetTotal(val int) {
	s.Total = val
}


func (s *PagedTransactions) SetPage(val int) {
	s.Page = val
}


func (s *PagedTransactions) SetSize(val int) {
	s.Size = val
}

func (*PagedTransactions) aPIV1TransactionsGetRes() {}


type PagedUsers struct {
	Items []User `json:"items"`
	
	Total int `json:"total"`
	
	Page int `json:"page"`
	
	Size int `json:"size"`
}


func (s *PagedUsers) GetItems() []User {
	return s.Items
}


func (s *PagedUsers) GetTotal() int {
	return s.Total
}


func (s *PagedUsers) GetPage() int {
	return s.Page
}


func (s *PagedUsers) GetSize() int {
	return s.Size
}


func (s *PagedUsers) SetItems(val []User) {
	s.Items = val
}


func (s *PagedUsers) SetTotal(val int) {
	s.Total = val
}


func (s *PagedUsers) SetPage(val int) {
	s.Page = val
}


func (s *PagedUsers) SetSize(val int) {
	s.Size = val
}

func (*PagedUsers) aPIV1UsersGetRes() {}


type RegisterRequest struct {
	
	Email string `json:"email"`
	
	
	
	
	Password string `json:"password"`
	
	FullName string `json:"fullName"`
	
	
	Region OptString `json:"region"`
	Gender OptGender `json:"gender"`
	
	Age           OptInt           `json:"age"`
	MaritalStatus OptMaritalStatus `json:"maritalStatus"`
}


func (s *RegisterRequest) GetEmail() string {
	return s.Email
}


func (s *RegisterRequest) GetPassword() string {
	return s.Password
}


func (s *RegisterRequest) GetFullName() string {
	return s.FullName
}


func (s *RegisterRequest) GetRegion() OptString {
	return s.Region
}


func (s *RegisterRequest) GetGender() OptGender {
	return s.Gender
}


func (s *RegisterRequest) GetAge() OptInt {
	return s.Age
}


func (s *RegisterRequest) GetMaritalStatus() OptMaritalStatus {
	return s.MaritalStatus
}


func (s *RegisterRequest) SetEmail(val string) {
	s.Email = val
}


func (s *RegisterRequest) SetPassword(val string) {
	s.Password = val
}


func (s *RegisterRequest) SetFullName(val string) {
	s.FullName = val
}


func (s *RegisterRequest) SetRegion(val OptString) {
	s.Region = val
}


func (s *RegisterRequest) SetGender(val OptGender) {
	s.Gender = val
}


func (s *RegisterRequest) SetAge(val OptInt) {
	s.Age = val
}


func (s *RegisterRequest) SetMaritalStatus(val OptMaritalStatus) {
	s.MaritalStatus = val
}


type RuleMatchRow struct {
	RuleId   uuid.UUID `json:"ruleId"`
	RuleName string    `json:"ruleName"`
	
	Matches int `json:"matches"`
	
	UniqueUsers int `json:"uniqueUsers"`
	
	UniqueMerchants OptInt `json:"uniqueMerchants"`
	
	ShareOfDeclines float64 `json:"shareOfDeclines"`
}


func (s *RuleMatchRow) GetRuleId() uuid.UUID {
	return s.RuleId
}


func (s *RuleMatchRow) GetRuleName() string {
	return s.RuleName
}


func (s *RuleMatchRow) GetMatches() int {
	return s.Matches
}


func (s *RuleMatchRow) GetUniqueUsers() int {
	return s.UniqueUsers
}


func (s *RuleMatchRow) GetUniqueMerchants() OptInt {
	return s.UniqueMerchants
}


func (s *RuleMatchRow) GetShareOfDeclines() float64 {
	return s.ShareOfDeclines
}


func (s *RuleMatchRow) SetRuleId(val uuid.UUID) {
	s.RuleId = val
}


func (s *RuleMatchRow) SetRuleName(val string) {
	s.RuleName = val
}


func (s *RuleMatchRow) SetMatches(val int) {
	s.Matches = val
}


func (s *RuleMatchRow) SetUniqueUsers(val int) {
	s.UniqueUsers = val
}


func (s *RuleMatchRow) SetUniqueMerchants(val OptInt) {
	s.UniqueMerchants = val
}


func (s *RuleMatchRow) SetShareOfDeclines(val float64) {
	s.ShareOfDeclines = val
}


type RuleMatchStats struct {
	Items []RuleMatchRow `json:"items"`
}


func (s *RuleMatchStats) GetItems() []RuleMatchRow {
	return s.Items
}


func (s *RuleMatchStats) SetItems(val []RuleMatchRow) {
	s.Items = val
}

func (*RuleMatchStats) aPIV1StatsRulesMatchesGetRes() {}




type StatsOverview struct {
	
	From time.Time `json:"from"`
	
	To time.Time `json:"to"`
	
	Volume int `json:"volume"`
	
	
	
	Gmv float64 `json:"gmv"`
	
	
	
	ApprovalRate float64 `json:"approvalRate"`
	
	
	
	DeclineRate float64 `json:"declineRate"`
	
	TopRiskMerchants []MerchantRiskRow `json:"topRiskMerchants"`
}


func (s *StatsOverview) GetFrom() time.Time {
	return s.From
}


func (s *StatsOverview) GetTo() time.Time {
	return s.To
}


func (s *StatsOverview) GetVolume() int {
	return s.Volume
}


func (s *StatsOverview) GetGmv() float64 {
	return s.Gmv
}


func (s *StatsOverview) GetApprovalRate() float64 {
	return s.ApprovalRate
}


func (s *StatsOverview) GetDeclineRate() float64 {
	return s.DeclineRate
}


func (s *StatsOverview) GetTopRiskMerchants() []MerchantRiskRow {
	return s.TopRiskMerchants
}


func (s *StatsOverview) SetFrom(val time.Time) {
	s.From = val
}


func (s *StatsOverview) SetTo(val time.Time) {
	s.To = val
}


func (s *StatsOverview) SetVolume(val int) {
	s.Volume = val
}


func (s *StatsOverview) SetGmv(val float64) {
	s.Gmv = val
}


func (s *StatsOverview) SetApprovalRate(val float64) {
	s.ApprovalRate = val
}


func (s *StatsOverview) SetDeclineRate(val float64) {
	s.DeclineRate = val
}


func (s *StatsOverview) SetTopRiskMerchants(val []MerchantRiskRow) {
	s.TopRiskMerchants = val
}

func (*StatsOverview) aPIV1StatsOverviewGetRes() {}


type Transaction struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	
	
	
	Amount   float64           `json:"amount"`
	Currency CurrencyCode      `json:"currency"`
	Status   TransactionStatus `json:"status"`
	
	MerchantId           OptString  `json:"merchantId"`
	MerchantCategoryCode OptMccCode `json:"merchantCategoryCode"`
	
	Timestamp time.Time `json:"timestamp"`
	
	IpAddress OptString `json:"ipAddress"`
	
	DeviceId OptString              `json:"deviceId"`
	Channel  OptTransactionChannel  `json:"channel"`
	Location OptTransactionLocation `json:"location"`
	
	
	IsFraud bool `json:"isFraud"`
	
	Metadata OptTransactionMetadata `json:"metadata"`
	
	CreatedAt time.Time `json:"createdAt"`
}


func (s *Transaction) GetID() uuid.UUID {
	return s.ID
}


func (s *Transaction) GetUserId() uuid.UUID {
	return s.UserId
}


func (s *Transaction) GetAmount() float64 {
	return s.Amount
}


func (s *Transaction) GetCurrency() CurrencyCode {
	return s.Currency
}


func (s *Transaction) GetStatus() TransactionStatus {
	return s.Status
}


func (s *Transaction) GetMerchantId() OptString {
	return s.MerchantId
}


func (s *Transaction) GetMerchantCategoryCode() OptMccCode {
	return s.MerchantCategoryCode
}


func (s *Transaction) GetTimestamp() time.Time {
	return s.Timestamp
}


func (s *Transaction) GetIpAddress() OptString {
	return s.IpAddress
}


func (s *Transaction) GetDeviceId() OptString {
	return s.DeviceId
}


func (s *Transaction) GetChannel() OptTransactionChannel {
	return s.Channel
}


func (s *Transaction) GetLocation() OptTransactionLocation {
	return s.Location
}


func (s *Transaction) GetIsFraud() bool {
	return s.IsFraud
}


func (s *Transaction) GetMetadata() OptTransactionMetadata {
	return s.Metadata
}


func (s *Transaction) GetCreatedAt() time.Time {
	return s.CreatedAt
}


func (s *Transaction) SetID(val uuid.UUID) {
	s.ID = val
}


func (s *Transaction) SetUserId(val uuid.UUID) {
	s.UserId = val
}


func (s *Transaction) SetAmount(val float64) {
	s.Amount = val
}


func (s *Transaction) SetCurrency(val CurrencyCode) {
	s.Currency = val
}


func (s *Transaction) SetStatus(val TransactionStatus) {
	s.Status = val
}


func (s *Transaction) SetMerchantId(val OptString) {
	s.MerchantId = val
}


func (s *Transaction) SetMerchantCategoryCode(val OptMccCode) {
	s.MerchantCategoryCode = val
}


func (s *Transaction) SetTimestamp(val time.Time) {
	s.Timestamp = val
}


func (s *Transaction) SetIpAddress(val OptString) {
	s.IpAddress = val
}


func (s *Transaction) SetDeviceId(val OptString) {
	s.DeviceId = val
}


func (s *Transaction) SetChannel(val OptTransactionChannel) {
	s.Channel = val
}


func (s *Transaction) SetLocation(val OptTransactionLocation) {
	s.Location = val
}


func (s *Transaction) SetIsFraud(val bool) {
	s.IsFraud = val
}


func (s *Transaction) SetMetadata(val OptTransactionMetadata) {
	s.Metadata = val
}


func (s *Transaction) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}


type TransactionBatchCreateRequest struct {
	
	Items []TransactionCreateRequest `json:"items"`
}


func (s *TransactionBatchCreateRequest) GetItems() []TransactionCreateRequest {
	return s.Items
}


func (s *TransactionBatchCreateRequest) SetItems(val []TransactionCreateRequest) {
	s.Items = val
}


type TransactionBatchResult struct {
	Items []TransactionBatchResultItem `json:"items"`
}


func (s *TransactionBatchResult) GetItems() []TransactionBatchResultItem {
	return s.Items
}


func (s *TransactionBatchResult) SetItems(val []TransactionBatchResultItem) {
	s.Items = val
}



type TransactionBatchResultItem struct {
	
	Index int `json:"index"`
	
	Decision OptTransactionDecision `json:"decision"`
	
	
	Error OptApiError `json:"error"`
}


func (s *TransactionBatchResultItem) GetIndex() int {
	return s.Index
}


func (s *TransactionBatchResultItem) GetDecision() OptTransactionDecision {
	return s.Decision
}


func (s *TransactionBatchResultItem) GetError() OptApiError {
	return s.Error
}


func (s *TransactionBatchResultItem) SetIndex(val int) {
	s.Index = val
}


func (s *TransactionBatchResultItem) SetDecision(val OptTransactionDecision) {
	s.Decision = val
}


func (s *TransactionBatchResultItem) SetError(val OptApiError) {
	s.Error = val
}



type TransactionChannel string

const (
	TransactionChannelWEB    TransactionChannel = "WEB"
	TransactionChannelMOBILE TransactionChannel = "MOBILE"
	TransactionChannelPOS    TransactionChannel = "POS"
	TransactionChannelOTHER  TransactionChannel = "OTHER"
)


func (TransactionChannel) AllValues() []TransactionChannel {
	return []TransactionChannel{
		TransactionChannelWEB,
		TransactionChannelMOBILE,
		TransactionChannelPOS,
		TransactionChannelOTHER,
	}
}


func (s TransactionChannel) MarshalText() ([]byte, error) {
	switch s {
	case TransactionChannelWEB:
		return []byte(s), nil
	case TransactionChannelMOBILE:
		return []byte(s), nil
	case TransactionChannelPOS:
		return []byte(s), nil
	case TransactionChannelOTHER:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *TransactionChannel) UnmarshalText(data []byte) error {
	switch TransactionChannel(data) {
	case TransactionChannelWEB:
		*s = TransactionChannelWEB
		return nil
	case TransactionChannelMOBILE:
		*s = TransactionChannelMOBILE
		return nil
	case TransactionChannelPOS:
		*s = TransactionChannelPOS
		return nil
	case TransactionChannelOTHER:
		*s = TransactionChannelOTHER
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}


type TransactionCreateRequest struct {
	
	
	
	UserId uuid.UUID `json:"userId"`
	
	
	Amount               float64      `json:"amount"`
	Currency             CurrencyCode `json:"currency"`
	MerchantId           OptString    `json:"merchantId"`
	MerchantCategoryCode OptMccCode   `json:"merchantCategoryCode"`
	
	
	Timestamp time.Time                           `json:"timestamp"`
	IpAddress OptString                           `json:"ipAddress"`
	DeviceId  OptString                           `json:"deviceId"`
	Channel   OptTransactionChannel               `json:"channel"`
	Location  OptTransactionLocation              `json:"location"`
	Metadata  OptTransactionCreateRequestMetadata `json:"metadata"`
}


func (s *TransactionCreateRequest) GetUserId() uuid.UUID {
	return s.UserId
}


func (s *TransactionCreateRequest) GetAmount() float64 {
	return s.Amount
}


func (s *TransactionCreateRequest) GetCurrency() CurrencyCode {
	return s.Currency
}


func (s *TransactionCreateRequest) GetMerchantId() OptString {
	return s.MerchantId
}


func (s *TransactionCreateRequest) GetMerchantCategoryCode() OptMccCode {
	return s.MerchantCategoryCode
}


func (s *TransactionCreateRequest) GetTimestamp() time.Time {
	return s.Timestamp
}


func (s *TransactionCreateRequest) GetIpAddress() OptString {
	return s.IpAddress
}


func (s *TransactionCreateRequest) GetDeviceId() OptString {
	return s.DeviceId
}


func (s *TransactionCreateRequest) GetChannel() OptTransactionChannel {
	return s.Channel
}


func (s *TransactionCreateRequest) GetLocation() OptTransactionLocation {
	return s.Location
}


func (s *TransactionCreateRequest) GetMetadata() OptTransactionCreateRequestMetadata {
	return s.Metadata
}


func (s *TransactionCreateRequest) SetUserId(val uuid.UUID) {
	s.UserId = val
}


func (s *TransactionCreateRequest) SetAmount(val float64) {
	s.Amount = val
}


func (s *TransactionCreateRequest) SetCurrency(val CurrencyCode) {
	s.Currency = val
}


func (s *TransactionCreateRequest) SetMerchantId(val OptString) {
	s.MerchantId = val
}


func (s *TransactionCreateRequest) SetMerchantCategoryCode(val OptMccCode) {
	s.MerchantCategoryCode = val
}


func (s *TransactionCreateRequest) SetTimestamp(val time.Time) {
	s.Timestamp = val
}


func (s *TransactionCreateRequest) SetIpAddress(val OptString) {
	s.IpAddress = val
}


func (s *TransactionCreateRequest) SetDeviceId(val OptString) {
	s.DeviceId = val
}


func (s *TransactionCreateRequest) SetChannel(val OptTransactionChannel) {
	s.Channel = val
}


func (s *TransactionCreateRequest) SetLocation(val OptTransactionLocation) {
	s.Location = val
}


func (s *TransactionCreateRequest) SetMetadata(val OptTransactionCreateRequestMetadata) {
	s.Metadata = val
}

type TransactionCreateRequestMetadata map[string]jx.Raw

func (s *TransactionCreateRequestMetadata) init() TransactionCreateRequestMetadata {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}




type TransactionDecision struct {
	Transaction Transaction `json:"transaction"`
	
	
	
	
	RuleResults []FraudRuleEvaluationResult `json:"ruleResults"`
}


func (s *TransactionDecision) GetTransaction() Transaction {
	return s.Transaction
}


func (s *TransactionDecision) GetRuleResults() []FraudRuleEvaluationResult {
	return s.RuleResults
}


func (s *TransactionDecision) SetTransaction(val Transaction) {
	s.Transaction = val
}


func (s *TransactionDecision) SetRuleResults(val []FraudRuleEvaluationResult) {
	s.RuleResults = val
}

func (*TransactionDecision) aPIV1TransactionsIDGetRes() {}
func (*TransactionDecision) aPIV1TransactionsPostRes()  {}










type TransactionLocation struct {
	
	Country OptString `json:"country"`
	
	City OptString `json:"city"`
	
	Latitude OptFloat64 `json:"latitude"`
	
	Longitude OptFloat64 `json:"longitude"`
}


func (s *TransactionLocation) GetCountry() OptString {
	return s.Country
}


func (s *TransactionLocation) GetCity() OptString {
	return s.City
}


func (s *TransactionLocation) GetLatitude() OptFloat64 {
	return s.Latitude
}


func (s *TransactionLocation) GetLongitude() OptFloat64 {
	return s.Longitude
}


func (s *TransactionLocation) SetCountry(val OptString) {
	s.Country = val
}


func (s *TransactionLocation) SetCity(val OptString) {
	s.City = val
}


func (s *TransactionLocation) SetLatitude(val OptFloat64) {
	s.Latitude = val
}


func (s *TransactionLocation) SetLongitude(val OptFloat64) {
	s.Longitude = val
}


type TransactionMetadata map[string]jx.Raw

func (s *TransactionMetadata) init() TransactionMetadata {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}





type TransactionStatus string

const (
	TransactionStatusAPPROVED TransactionStatus = "APPROVED"
	TransactionStatusDECLINED TransactionStatus = "DECLINED"
)


func (TransactionStatus) AllValues() []TransactionStatus {
	return []TransactionStatus{
		TransactionStatusAPPROVED,
		TransactionStatusDECLINED,
	}
}


func (s TransactionStatus) MarshalText() ([]byte, error) {
	switch s {
	case TransactionStatusAPPROVED:
		return []byte(s), nil
	case TransactionStatusDECLINED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *TransactionStatus) UnmarshalText(data []byte) error {
	switch TransactionStatus(data) {
	case TransactionStatusAPPROVED:
		*s = TransactionStatusAPPROVED
		return nil
	case TransactionStatusDECLINED:
		*s = TransactionStatusDECLINED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}



type TransactionsTimePoint struct {
	
	BucketStart time.Time `json:"bucketStart"`
	
	TxCount int `json:"txCount"`
	
	
	
	Gmv float64 `json:"gmv"`
	
	ApprovalRate float64 `json:"approvalRate"`
	
	DeclineRate float64 `json:"declineRate"`
}


func (s *TransactionsTimePoint) GetBucketStart() time.Time {
	return s.BucketStart
}


func (s *TransactionsTimePoint) GetTxCount() int {
	return s.TxCount
}


func (s *TransactionsTimePoint) GetGmv() float64 {
	return s.Gmv
}


func (s *TransactionsTimePoint) GetApprovalRate() float64 {
	return s.ApprovalRate
}


func (s *TransactionsTimePoint) GetDeclineRate() float64 {
	return s.DeclineRate
}


func (s *TransactionsTimePoint) SetBucketStart(val time.Time) {
	s.BucketStart = val
}


func (s *TransactionsTimePoint) SetTxCount(val int) {
	s.TxCount = val
}


func (s *TransactionsTimePoint) SetGmv(val float64) {
	s.Gmv = val
}


func (s *TransactionsTimePoint) SetApprovalRate(val float64) {
	s.ApprovalRate = val
}


func (s *TransactionsTimePoint) SetDeclineRate(val float64) {
	s.DeclineRate = val
}


type TransactionsTimeSeries struct {
	Points []TransactionsTimePoint `json:"points"`
}


func (s *TransactionsTimeSeries) GetPoints() []TransactionsTimePoint {
	return s.Points
}


func (s *TransactionsTimeSeries) SetPoints(val []TransactionsTimePoint) {
	s.Points = val
}

func (*TransactionsTimeSeries) aPIV1StatsTransactionsTimeseriesGetRes() {}


type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
	
	Region        OptString        `json:"region"`
	Gender        OptGender        `json:"gender"`
	Age           OptInt           `json:"age"`
	MaritalStatus OptMaritalStatus `json:"maritalStatus"`
	Role          UserRole         `json:"role"`
	
	
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


func (s *User) GetID() uuid.UUID {
	return s.ID
}


func (s *User) GetEmail() string {
	return s.Email
}


func (s *User) GetFullName() string {
	return s.FullName
}


func (s *User) GetRegion() OptString {
	return s.Region
}


func (s *User) GetGender() OptGender {
	return s.Gender
}


func (s *User) GetAge() OptInt {
	return s.Age
}


func (s *User) GetMaritalStatus() OptMaritalStatus {
	return s.MaritalStatus
}


func (s *User) GetRole() UserRole {
	return s.Role
}


func (s *User) GetIsActive() bool {
	return s.IsActive
}


func (s *User) GetCreatedAt() time.Time {
	return s.CreatedAt
}


func (s *User) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}


func (s *User) SetID(val uuid.UUID) {
	s.ID = val
}


func (s *User) SetEmail(val string) {
	s.Email = val
}


func (s *User) SetFullName(val string) {
	s.FullName = val
}


func (s *User) SetRegion(val OptString) {
	s.Region = val
}


func (s *User) SetGender(val OptGender) {
	s.Gender = val
}


func (s *User) SetAge(val OptInt) {
	s.Age = val
}


func (s *User) SetMaritalStatus(val OptMaritalStatus) {
	s.MaritalStatus = val
}


func (s *User) SetRole(val UserRole) {
	s.Role = val
}


func (s *User) SetIsActive(val bool) {
	s.IsActive = val
}


func (s *User) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}


func (s *User) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*User) aPIV1UsersIDGetRes() {}
func (*User) aPIV1UsersIDPutRes() {}
func (*User) aPIV1UsersMeGetRes() {}
func (*User) aPIV1UsersMePutRes() {}
func (*User) aPIV1UsersPostRes()  {}


type UserCreateRequest struct {
	Email         string           `json:"email"`
	Password      string           `json:"password"`
	FullName      string           `json:"fullName"`
	Region        OptString        `json:"region"`
	Gender        OptGender        `json:"gender"`
	Age           OptInt           `json:"age"`
	MaritalStatus OptMaritalStatus `json:"maritalStatus"`
	Role          UserRole         `json:"role"`
}


func (s *UserCreateRequest) GetEmail() string {
	return s.Email
}


func (s *UserCreateRequest) GetPassword() string {
	return s.Password
}


func (s *UserCreateRequest) GetFullName() string {
	return s.FullName
}


func (s *UserCreateRequest) GetRegion() OptString {
	return s.Region
}


func (s *UserCreateRequest) GetGender() OptGender {
	return s.Gender
}


func (s *UserCreateRequest) GetAge() OptInt {
	return s.Age
}


func (s *UserCreateRequest) GetMaritalStatus() OptMaritalStatus {
	return s.MaritalStatus
}


func (s *UserCreateRequest) GetRole() UserRole {
	return s.Role
}


func (s *UserCreateRequest) SetEmail(val string) {
	s.Email = val
}


func (s *UserCreateRequest) SetPassword(val string) {
	s.Password = val
}


func (s *UserCreateRequest) SetFullName(val string) {
	s.FullName = val
}


func (s *UserCreateRequest) SetRegion(val OptString) {
	s.Region = val
}


func (s *UserCreateRequest) SetGender(val OptGender) {
	s.Gender = val
}


func (s *UserCreateRequest) SetAge(val OptInt) {
	s.Age = val
}


func (s *UserCreateRequest) SetMaritalStatus(val OptMaritalStatus) {
	s.MaritalStatus = val
}


func (s *UserCreateRequest) SetRole(val UserRole) {
	s.Role = val
}


type UserRiskProfile struct {
	UserId uuid.UUID `json:"userId"`
	
	TxCount24h int `json:"txCount_24h"`
	
	
	Gmv24h float64 `json:"gmv_24h"`
	
	DistinctDevices24h int `json:"distinctDevices_24h"`
	
	DistinctIps24h int `json:"distinctIps_24h"`
	
	DistinctCities24h int `json:"distinctCities_24h"`
	
	DeclineRate30d float64 `json:"declineRate_30d"`
	
	
	
	LastSeenAt OptNilDateTime `json:"lastSeenAt"`
}


func (s *UserRiskProfile) GetUserId() uuid.UUID {
	return s.UserId
}


func (s *UserRiskProfile) GetTxCount24h() int {
	return s.TxCount24h
}


func (s *UserRiskProfile) GetGmv24h() float64 {
	return s.Gmv24h
}


func (s *UserRiskProfile) GetDistinctDevices24h() int {
	return s.DistinctDevices24h
}


func (s *UserRiskProfile) GetDistinctIps24h() int {
	return s.DistinctIps24h
}


func (s *UserRiskProfile) GetDistinctCities24h() int {
	return s.DistinctCities24h
}


func (s *UserRiskProfile) GetDeclineRate30d() float64 {
	return s.DeclineRate30d
}


func (s *UserRiskProfile) GetLastSeenAt() OptNilDateTime {
	return s.LastSeenAt
}


func (s *UserRiskProfile) SetUserId(val uuid.UUID) {
	s.UserId = val
}


func (s *UserRiskProfile) SetTxCount24h(val int) {
	s.TxCount24h = val
}


func (s *UserRiskProfile) SetGmv24h(val float64) {
	s.Gmv24h = val
}


func (s *UserRiskProfile) SetDistinctDevices24h(val int) {
	s.DistinctDevices24h = val
}


func (s *UserRiskProfile) SetDistinctIps24h(val int) {
	s.DistinctIps24h = val
}


func (s *UserRiskProfile) SetDistinctCities24h(val int) {
	s.DistinctCities24h = val
}


func (s *UserRiskProfile) SetDeclineRate30d(val float64) {
	s.DeclineRate30d = val
}


func (s *UserRiskProfile) SetLastSeenAt(val OptNilDateTime) {
	s.LastSeenAt = val
}

func (*UserRiskProfile) aPIV1StatsUsersIDRiskProfileGetRes() {}



type UserRole string

const (
	UserRoleADMIN UserRole = "ADMIN"
	UserRoleUSER  UserRole = "USER"
)


func (UserRole) AllValues() []UserRole {
	return []UserRole{
		UserRoleADMIN,
		UserRoleUSER,
	}
}


func (s UserRole) MarshalText() ([]byte, error) {
	switch s {
	case UserRoleADMIN:
		return []byte(s), nil
	case UserRoleUSER:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}


func (s *UserRole) UnmarshalText(data []byte) error {
	switch UserRole(data) {
	case UserRoleADMIN:
		*s = UserRoleADMIN
		return nil
	case UserRoleUSER:
		*s = UserRoleUSER
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}








type UserUpdateRequest struct {
	FullName      string           `json:"fullName"`
	Region        NilString        `json:"region"`
	Gender        NilGender        `json:"gender"`
	Age           NilInt           `json:"age"`
	MaritalStatus NilMaritalStatus `json:"maritalStatus"`
	
	Role OptUserRole `json:"role"`
	
	IsActive OptBool `json:"isActive"`
}


func (s *UserUpdateRequest) GetFullName() string {
	return s.FullName
}


func (s *UserUpdateRequest) GetRegion() NilString {
	return s.Region
}


func (s *UserUpdateRequest) GetGender() NilGender {
	return s.Gender
}


func (s *UserUpdateRequest) GetAge() NilInt {
	return s.Age
}


func (s *UserUpdateRequest) GetMaritalStatus() NilMaritalStatus {
	return s.MaritalStatus
}


func (s *UserUpdateRequest) GetRole() OptUserRole {
	return s.Role
}


func (s *UserUpdateRequest) GetIsActive() OptBool {
	return s.IsActive
}


func (s *UserUpdateRequest) SetFullName(val string) {
	s.FullName = val
}


func (s *UserUpdateRequest) SetRegion(val NilString) {
	s.Region = val
}


func (s *UserUpdateRequest) SetGender(val NilGender) {
	s.Gender = val
}


func (s *UserUpdateRequest) SetAge(val NilInt) {
	s.Age = val
}


func (s *UserUpdateRequest) SetMaritalStatus(val NilMaritalStatus) {
	s.MaritalStatus = val
}


func (s *UserUpdateRequest) SetRole(val OptUserRole) {
	s.Role = val
}


func (s *UserUpdateRequest) SetIsActive(val OptBool) {
	s.IsActive = val
}


type ValidationError struct {
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	TraceId     uuid.UUID    `json:"traceId"`
	Timestamp   time.Time    `json:"timestamp"`
	Path        string       `json:"path"`
	FieldErrors []FieldError `json:"fieldErrors"`
}


func (s *ValidationError) GetCode() string {
	return s.Code
}


func (s *ValidationError) GetMessage() string {
	return s.Message
}


func (s *ValidationError) GetTraceId() uuid.UUID {
	return s.TraceId
}


func (s *ValidationError) GetTimestamp() time.Time {
	return s.Timestamp
}


func (s *ValidationError) GetPath() string {
	return s.Path
}


func (s *ValidationError) GetFieldErrors() []FieldError {
	return s.FieldErrors
}


func (s *ValidationError) SetCode(val string) {
	s.Code = val
}


func (s *ValidationError) SetMessage(val string) {
	s.Message = val
}


func (s *ValidationError) SetTraceId(val uuid.UUID) {
	s.TraceId = val
}


func (s *ValidationError) SetTimestamp(val time.Time) {
	s.Timestamp = val
}


func (s *ValidationError) SetPath(val string) {
	s.Path = val
}


func (s *ValidationError) SetFieldErrors(val []FieldError) {
	s.FieldErrors = val
}

func (*ValidationError) aPIV1AuthLoginPostRes()         {}
func (*ValidationError) aPIV1AuthRegisterPostRes()      {}
func (*ValidationError) aPIV1FraudRulesIDPutRes()       {}
func (*ValidationError) aPIV1FraudRulesPostRes()        {}
func (*ValidationError) aPIV1StatsOverviewGetRes()      {}
func (*ValidationError) aPIV1TransactionsBatchPostRes() {}
func (*ValidationError) aPIV1TransactionsGetRes()       {}
func (*ValidationError) aPIV1TransactionsPostRes()      {}
func (*ValidationError) aPIV1UsersGetRes()              {}
func (*ValidationError) aPIV1UsersIDPutRes()            {}
func (*ValidationError) aPIV1UsersMePutRes()            {}
func (*ValidationError) aPIV1UsersPostRes()             {}
