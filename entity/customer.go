package entity

import "fmt"

type Customer struct {
	Name string
	PhoneNo string
	TotalFee float64
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Customer name cannot be empty")
	}
	if len(c.Name) > 100 {
		return fmt.Errorf("Customer name cannot exceed 100 characters")
	}
	if len(c.PhoneNo) < 11  {
		fmt.Println("Phone length is: ", len(c.PhoneNo))
		return fmt.Errorf("phone number must be about 12 number. Example: +6282169672989 ")
	}
	return nil
}