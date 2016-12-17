package storage

import (
	"github.com/golang/protobuf/proto"
	"log"
)

type Applicant struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (a *Applicant) MarshalBinary() ([]byte, error) {
	// if a.ID == nil {
	// 	a.ID = randId(20)
	// }
	return proto.Marshal(&ApSignup{
		ID:        a.ID,
		Firstname: a.Firstname,
		Lastname:  a.Lastname,
		Email:     a.Email,
		Password:  a.Password,
	})
}

func (a *Applicant) UnmarshalBinary(data []byte) error {
	var pb ApSignup

	if err := proto.Unmarshal(data, &pb); err != nil {
		log.Println(err)
		return err
	}

	a.ID = pb.GetID()
	a.Firstname = pb.GetFirstname()
	a.Lastname = pb.GetLastname()
	a.Email = pb.GetEmail()
	a.Password = pb.GetPassword()
	return nil
}
