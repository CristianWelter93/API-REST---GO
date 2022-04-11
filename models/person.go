package main

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

func (p Person) New(firstName string, lastName string, address *Address) {
    p.Firstname = firstName
    p.Lastname = lastName
    p.Address = address
}