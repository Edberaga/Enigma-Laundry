package entity

import "fmt"

type Service struct {
	ServiceName string
	Satuan string
	Price float64
}

func (s *Service) Validate() error {
	if s.ServiceName == "" {
		return fmt.Errorf("Service name cannot be empty")
	}
	if len(s.ServiceName) > 100 {
		return fmt.Errorf("Service name cannot exceed 100 characters")
	}
	if s.Satuan == "" {
		return fmt.Errorf("Satuan cannot be empty")
	}
	if s.Price <= 0 {
		return fmt.Errorf("Price must be positive")
	}
	return nil
}