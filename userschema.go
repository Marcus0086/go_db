package main

type (
	Address struct {
		Country string `json:"country"`
		City    string `json:"city"`
		State   string `json:"state"`
		Pincode int32  `json:"pincode"`
	}

	User struct {
		Name    string  `json:"name"`
		Age     int32   `json:"age"`
		Contact string  `json:"contact"`
		Company string  `json:"company"`
		Address Address `json:"address"`
	}
)
