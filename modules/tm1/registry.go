package tm1

func Tm1Registry() Service {
	tm1Repository := NewRepository()
	tm1Service := NewService(tm1Repository)

	return tm1Service
}
