package factory

import "fmt"

type Mobile interface {
	setModel(model string)
	getModel() string
	setManufacturer(model string)
	getManufacturer() string
}

type mobile struct {
	model        string
	manufacturer string
}

func (m *mobile) setModel(model string) {
	m.model = model
}

func (m *mobile) getModel() string {
	return m.model
}

func (m *mobile) setManufacturer(manufacturer string) {
	m.manufacturer = manufacturer
}

func (m *mobile) getManufacturer() string {
	return m.manufacturer
}

type android struct {
	mobile
}

func newAndroid() Mobile {
	return &android{
		mobile{
			model:        "type X",
			manufacturer: "Google",
		},
	}
}

type ios struct {
	mobile
}

func newIOS() Mobile {
	return &ios{
		mobile{
			model:        "type Y",
			manufacturer: "Apple",
		},
	}
}

func getMobile(mobType string) (Mobile, error) {
	if mobType == "android" {
		return newAndroid(), nil
	}
	if mobType == "ios" {
		return newIOS(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
