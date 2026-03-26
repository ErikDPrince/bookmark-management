package service

type Health interface {
	Check() Response
}

type Response struct {
	Message     string `json:"message" example:"ok"`
	ServiceName string `json:"service_name" example:"bookmark-management"`
	InstanceID  string `json:"instance_id" example:"123e4567-e89b-12d3-a456-426614174000"`
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
