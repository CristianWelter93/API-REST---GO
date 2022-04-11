package main

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

func (a Address) New(city string, state string) {
	a.City = city
    a.State = state
}