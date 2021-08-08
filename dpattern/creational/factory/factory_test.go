package factory

import "testing"

func TestAndroid(t *testing.T) {
	mobile, _ := getMobile("android")
	if mobile.getManufacturer() != "Google" {
		t.Errorf("Expected Google, got %s", mobile.getManufacturer())
	}
}

func TestIOS(t *testing.T) {
	mobile, _ := getMobile("ios")
	if mobile.getManufacturer() != "Apple" {
		t.Errorf("Expected Apple, got %s", mobile.getManufacturer())
	}
}

func TestSetManfucturer(t *testing.T) {
	mobile, _ := getMobile("android")
	mobile.setManufacturer("HTC")
	if mobile.getManufacturer() != "HTC" {
		t.Errorf("Expected HTC, got %s", mobile.getManufacturer())
	}
}
