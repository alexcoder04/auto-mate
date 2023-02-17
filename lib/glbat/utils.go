package glbat

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type Battery struct {
	// battery id
	Id int

	// current battery capacity in percent
	Capacity         int
	CapacityLevel    string
	CycleCount       int
	EnergyFull       int
	EnergyFullDesign int
	EnergyNow        int
	Manufacturer     string
	ModelName        string
	PowerNow         int
	// whether the battery is present
	Present      bool
	SerialNumber string
	// battery status, e. g. "Charging"
	Status           string
	Technology       string
	VoltageNow       int
	VoltageMinDesign int
}

const sysfsPrefix = "/sys/class/power_supply"

func isBattery(path string) bool {
	if !strings.HasPrefix(filepath.Base(path), "BAT") {
		return false
	}
	data, err := ioutil.ReadFile(filepath.Join(path, "type"))
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(data)) == "Battery"
}

func readString(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func readInt(file string) (int, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}
