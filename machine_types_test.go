package gcpcomputepricing

import (
	"testing"
)

func TestGetE2standard2OnDemand(t *testing.T) {
	hourlyRate := getter(t, getE2standard2)

	if hourlyRate != 67011420 {
		t.Errorf("Hourly rate should be 67011420, instead of %d", hourlyRate)
	}
}

func TestGetE2standard4OnDemand(t *testing.T) {
	hourlyRate := getter(t, getE2standard4)

	if hourlyRate != 134022840 {
		t.Errorf("Hourly rate should be 134022840, instead of %d", hourlyRate)
	}
}

func TestGetE2standard8OnDemand(t *testing.T) {
	hourlyRate := getter(t, getE2standard8)

	if hourlyRate != 268045680 {
		t.Errorf("Hourly rate should be 268045680, instead of %d", hourlyRate)
	}
}

func TestGetE2standard16OnDemand(t *testing.T) {
	hourlyRate := getter(t, getE2standard16)

	if hourlyRate != 536091360 {
		t.Errorf("Hourly rate should be 536091360, instead of %d", hourlyRate)
	}
}

func TestGetE2standard32OnDemand(t *testing.T) {
	hourlyRate := getter(t, getE2standard32)

	if hourlyRate != 1072182720 {
		t.Errorf("Hourly rate should be 1072182720, instead of %d", hourlyRate)
	}
}

func getter(t *testing.T, fn machineTypeGetter) uint64 {
	p, err := Fetch()
	if err != nil {
		t.Fatal(err)
	}
	hourlyRate, err := fn(p, Opts{
		Type:   OnDemand,
		Region: "us-west1",
	})
	if err != nil {
		t.Fatal(err)
	}

	return hourlyRate
}
