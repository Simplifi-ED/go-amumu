package usecase

type DeviceCase struct {
}

func NewDeviceCase() *DeviceCase {
	return &DeviceCase{}
}

func (usecase *DeviceCase) SendToSMTPSERVER() error {
	return nil
}
