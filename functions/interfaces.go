package functions

import "fmt"

type architect struct {
	person
	specialization string
	buildingsDone int
}

type dentist struct {
	person
	specialization string
	surgeries int
}

func (a architect) greet() {
	fmt.Println("Morning! Im called", a.name, "and ive done", a.buildingsDone, "buildings.")
}

func (d dentist) greet() {
	fmt.Println("Morning! Im called", d.name, "and ive made", d.surgeries, "surgeries.")
}

type human interface {
	greet()
}

func beHuman(h human) {

	switch profession := h.(type) {
	case architect:
		fmt.Println("This architect is specialized in", profession.specialization, "and made", profession.buildingsDone, " buildings")
	case dentist:
		fmt.Println("This dentist is specialized in", profession.specialization, "and made", profession.surgeries, "surgeries")
	}

	h.greet()
}

func interfaceMethodCall() {

	paul := architect{
		person{
			"Paul",
			"Constructor",
			40,
		},
		"Concrete",
		3,
	}

	henry := dentist{
		person{
			"Henry",
			"Tooth",
			65,
		},
		"Hair Transplant",
		1,
	}

	beHuman(paul)
	beHuman(henry)
}