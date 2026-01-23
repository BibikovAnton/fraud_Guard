
package antifraud_v1

func (s *FraudRule) setDefaults() {
	{
		val := bool(true)
		s.Enabled = val
	}
}

func (s *FraudRuleCreateRequest) setDefaults() {
	{
		val := bool(true)
		s.Enabled.SetTo(val)
	}
	{
		val := int(100)
		s.Priority.SetTo(val)
	}
}

func (s *User) setDefaults() {
	{
		val := bool(true)
		s.IsActive = val
	}
}
