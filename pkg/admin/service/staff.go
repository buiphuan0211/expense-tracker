package service

type StaffInterface interface {
}

type staffImpl struct{}

func Staff() StaffInterface {
	return staffImpl{}
}
