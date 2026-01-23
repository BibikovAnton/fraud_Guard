
package antifraud_v1

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		return "", false
	}
	return strings.TrimPrefix(path, prefix), true
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/api/v1/"

			if l := len("/api/v1/"); len(elem) >= l && elem[0:l] == "/api/v1/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "auth/"

				if l := len("auth/"); len(elem) >= l && elem[0:l] == "auth/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "login"

					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleAPIV1AuthLoginPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

				case 'r': // Prefix: "register"

					if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleAPIV1AuthRegisterPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

				}

			case 'f': // Prefix: "fraud-rules"

				if l := len("fraud-rules"); len(elem) >= l && elem[0:l] == "fraud-rules" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleAPIV1FraudRulesGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleAPIV1FraudRulesPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'v': // Prefix: "validate"
						origElem := elem
						if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleAPIV1FraudRulesValidatePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleAPIV1FraudRulesIDDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleAPIV1FraudRulesIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleAPIV1FraudRulesIDPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}

				}

			case 'p': // Prefix: "ping"

				if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleAPIV1PingGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

			case 's': // Prefix: "stats/"

				if l := len("stats/"); len(elem) >= l && elem[0:l] == "stats/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "merchants/risk"

					if l := len("merchants/risk"); len(elem) >= l && elem[0:l] == "merchants/risk" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleAPIV1StatsMerchantsRiskGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 'o': // Prefix: "overview"

					if l := len("overview"); len(elem) >= l && elem[0:l] == "overview" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleAPIV1StatsOverviewGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 'r': // Prefix: "rules/matches"

					if l := len("rules/matches"); len(elem) >= l && elem[0:l] == "rules/matches" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleAPIV1StatsRulesMatchesGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 't': // Prefix: "transactions/timeseries"

					if l := len("transactions/timeseries"); len(elem) >= l && elem[0:l] == "transactions/timeseries" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleAPIV1StatsTransactionsTimeseriesGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 'u': // Prefix: "users/"

					if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
						elem = elem[l:]
					} else {
						break
					}

					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/risk-profile"

						if l := len("/risk-profile"); len(elem) >= l && elem[0:l] == "/risk-profile" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleAPIV1StatsUsersIDRiskProfileGetRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

					}

				}

			case 't': // Prefix: "transactions"

				if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleAPIV1TransactionsGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleAPIV1TransactionsPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "batch"
						origElem := elem
						if l := len("batch"); len(elem) >= l && elem[0:l] == "batch" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleAPIV1TransactionsBatchPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleAPIV1TransactionsIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				}

			case 'u': // Prefix: "users"

				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleAPIV1UsersGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleAPIV1UsersPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'm': // Prefix: "me"
						origElem := elem
						if l := len("me"); len(elem) >= l && elem[0:l] == "me" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "GET":
								s.handleAPIV1UsersMeGetRequest([0]string{}, elemIsEscaped, w, r)
							case "PUT":
								s.handleAPIV1UsersMePutRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET,PUT")
							}

							return
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleAPIV1UsersIDDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleAPIV1UsersIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleAPIV1UsersIDPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}

				}

			}

		}
	}
	s.notFound(w, r)
}

type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

func (r Route) Name() string {
	return r.name
}

func (r Route) Summary() string {
	return r.summary
}

func (r Route) OperationID() string {
	return r.operationID
}

func (r Route) PathPattern() string {
	return r.pathPattern
}

func (r Route) Args() []string {
	return r.args[:r.count]
}

func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/api/v1/"

			if l := len("/api/v1/"); len(elem) >= l && elem[0:l] == "/api/v1/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "auth/"

				if l := len("auth/"); len(elem) >= l && elem[0:l] == "auth/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'l': // Prefix: "login"

					if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = APIV1AuthLoginPostOperation
							r.summary = "Авторизация пользователя"
							r.operationID = ""
							r.pathPattern = "/api/v1/auth/login"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'r': // Prefix: "register"

					if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = APIV1AuthRegisterPostOperation
							r.summary = "Регистрация нового пользователя"
							r.operationID = ""
							r.pathPattern = "/api/v1/auth/register"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				}

			case 'f': // Prefix: "fraud-rules"

				if l := len("fraud-rules"); len(elem) >= l && elem[0:l] == "fraud-rules" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = APIV1FraudRulesGetOperation
						r.summary = "Список правил фрода"
						r.operationID = ""
						r.pathPattern = "/api/v1/fraud-rules"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = APIV1FraudRulesPostOperation
						r.summary = "Создать правило фрода"
						r.operationID = ""
						r.pathPattern = "/api/v1/fraud-rules"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'v': // Prefix: "validate"
						origElem := elem
						if l := len("validate"); len(elem) >= l && elem[0:l] == "validate" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = APIV1FraudRulesValidatePostOperation
								r.summary = "Валидация DSL выражения (без сохранения правила)"
								r.operationID = ""
								r.pathPattern = "/api/v1/fraud-rules/validate"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = APIV1FraudRulesIDDeleteOperation
							r.summary = "Деактивация правила фрода"
							r.operationID = ""
							r.pathPattern = "/api/v1/fraud-rules/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = APIV1FraudRulesIDGetOperation
							r.summary = "Получить правило фрода по ID"
							r.operationID = ""
							r.pathPattern = "/api/v1/fraud-rules/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							r.name = APIV1FraudRulesIDPutOperation
							r.summary = "Обновить правило фрода"
							r.operationID = ""
							r.pathPattern = "/api/v1/fraud-rules/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				}

			case 'p': // Prefix: "ping"

				if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = APIV1PingGetOperation
						r.summary = "Health check"
						r.operationID = ""
						r.pathPattern = "/api/v1/ping"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

			case 's': // Prefix: "stats/"

				if l := len("stats/"); len(elem) >= l && elem[0:l] == "stats/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "merchants/risk"

					if l := len("merchants/risk"); len(elem) >= l && elem[0:l] == "merchants/risk" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = APIV1StatsMerchantsRiskGetOperation
							r.summary = "Риск-метрики по мерчантам"
							r.operationID = ""
							r.pathPattern = "/api/v1/stats/merchants/risk"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'o': // Prefix: "overview"

					if l := len("overview"); len(elem) >= l && elem[0:l] == "overview" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = APIV1StatsOverviewGetOperation
							r.summary = "Обзор ключевых метрик (дашборд)"
							r.operationID = ""
							r.pathPattern = "/api/v1/stats/overview"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'r': // Prefix: "rules/matches"

					if l := len("rules/matches"); len(elem) >= l && elem[0:l] == "rules/matches" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = APIV1StatsRulesMatchesGetOperation
							r.summary = "Статистика срабатываний правил"
							r.operationID = ""
							r.pathPattern = "/api/v1/stats/rules/matches"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 't': // Prefix: "transactions/timeseries"

					if l := len("transactions/timeseries"); len(elem) >= l && elem[0:l] == "transactions/timeseries" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = APIV1StatsTransactionsTimeseriesGetOperation
							r.summary = "Таймсерия по транзакциям"
							r.operationID = ""
							r.pathPattern = "/api/v1/stats/transactions/timeseries"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'u': // Prefix: "users/"

					if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
						elem = elem[l:]
					} else {
						break
					}

					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/risk-profile"

						if l := len("/risk-profile"); len(elem) >= l && elem[0:l] == "/risk-profile" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = APIV1StatsUsersIDRiskProfileGetOperation
								r.summary = "Риск-профиль пользователя (в т.ч. для USER — только свой)"
								r.operationID = ""
								r.pathPattern = "/api/v1/stats/users/{id}/risk-profile"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

					}

				}

			case 't': // Prefix: "transactions"

				if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = APIV1TransactionsGetOperation
						r.summary = "Список транзакций"
						r.operationID = ""
						r.pathPattern = "/api/v1/transactions"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = APIV1TransactionsPostOperation
						r.summary = "Создать транзакцию и выполнить антифрод-проверку"
						r.operationID = ""
						r.pathPattern = "/api/v1/transactions"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "batch"
						origElem := elem
						if l := len("batch"); len(elem) >= l && elem[0:l] == "batch" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = APIV1TransactionsBatchPostOperation
								r.summary = "Создать батч транзакций"
								r.operationID = ""
								r.pathPattern = "/api/v1/transactions/batch"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = APIV1TransactionsIDGetOperation
							r.summary = "Получить транзакцию"
							r.operationID = ""
							r.pathPattern = "/api/v1/transactions/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				}

			case 'u': // Prefix: "users"

				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = APIV1UsersGetOperation
						r.summary = "Список пользователей"
						r.operationID = ""
						r.pathPattern = "/api/v1/users"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = APIV1UsersPostOperation
						r.summary = "Создание пользователя админом"
						r.operationID = ""
						r.pathPattern = "/api/v1/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'm': // Prefix: "me"
						origElem := elem
						if l := len("me"); len(elem) >= l && elem[0:l] == "me" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								r.name = APIV1UsersMeGetOperation
								r.summary = "Получить профиль текущего пользователя"
								r.operationID = ""
								r.pathPattern = "/api/v1/users/me"
								r.args = args
								r.count = 0
								return r, true
							case "PUT":
								r.name = APIV1UsersMePutOperation
								r.summary = "Обновить профиль текущего пользователя"
								r.operationID = ""
								r.pathPattern = "/api/v1/users/me"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = APIV1UsersIDDeleteOperation
							r.summary = "Деактивация пользователя"
							r.operationID = ""
							r.pathPattern = "/api/v1/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = APIV1UsersIDGetOperation
							r.summary = "Получение профиля пользователя"
							r.operationID = ""
							r.pathPattern = "/api/v1/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							r.name = APIV1UsersIDPutOperation
							r.summary = "Полное обновление профиля пользователя"
							r.operationID = ""
							r.pathPattern = "/api/v1/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				}

			}

		}
	}
	return r, false
}
