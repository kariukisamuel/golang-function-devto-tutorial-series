package main

import "fmt"

type VoterRegister struct {
	name           string
	pollingStation string
	id             int
}

func main() {

	v := NewVoter("samuel", "ruiru", 23)
	p := &v
	fmt.Println(v)
	p.UpdateVoter()
	fmt.Println(v)
}

func NewVoter(name string, pollingStation string, id int) VoterRegister {
	return VoterRegister{
		name:           name,
		pollingStation: pollingStation,
		id:             id,
	}
}
func (v *VoterRegister) UpdateVoter() {
	v.name = "Njoki"
	v.pollingStation = "Ruaka"
	v.id = 2
}

func (v VoterRegister) Publish() VoterRegister {
	return VoterRegister{v.name, v.pollingStation, v.id}
}
