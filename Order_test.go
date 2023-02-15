package mockuptest

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name   string `valid:"required~name not blank"`
	Reason string
	Limit  uint
}

func Test_Name_Notblank(t *testing.T) {
	g := NewGomegaWithT(t)

	Order := Order{
		Name:   "",
		Reason: "ssss",
		Limit:  20,
	}

	ok, err := govalidator.ValidateStruct(Order)

	g.Expect(ok).NotTo(BeTrue())

	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("name not blank"))

}