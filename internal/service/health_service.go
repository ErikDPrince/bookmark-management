package service

type Health interface {
	Check() Response
}

type Response struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

type healthService struct {
	serviceName string
	instanceID  string
}

func NewHealthService(serviceName, instanceID string) Health {
	return &healthService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (h *healthService) Check() Response {
	return Response{
		Message:     "ok",
		ServiceName: h.serviceName,
		InstanceID:  h.instanceID,
	}
}
