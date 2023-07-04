package glbat

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// Get data for all batteries.
func GetAll() ([]Battery, error) {
	bats := []Battery{}

	ids, err := GetDetected()
	if err != nil {
		return nil, err
	}

	for _, i := range ids {
		b, err := GetBat(i)
		if err != nil {
			continue
		}

		bats = append(bats, b)
	}

	return bats, nil
}

// Get data for battery with a certain id.
func GetBat(id int) (Battery, error) {
	path := fmt.Sprintf("%s/BAT%d", sysfsPrefix, id)
	bat := Battery{
		Id: id,
	}

	if !isBattery(path) {
		return bat, errors.New("battery does not exist")
	}

	cl, err := readString(path + "/capacity_level")
	if err != nil {
		bat.CapacityLevel = "Unknown"
	}
	bat.CapacityLevel = cl

	cc, err := readInt(path + "/cycle_count")
	if err != nil {
		bat.CycleCount = 0
	}
	bat.CycleCount = cc

	ef, err := readInt(path + "/energy_full")
	if err != nil {
		return bat, errors.New("cannot get energy_full")
	}
	bat.EnergyFull = ef

	efd, err := readInt(path + "/energy_full_design")
	if err != nil {
		bat.EnergyFullDesign = bat.EnergyFull
	} else {
		bat.EnergyFullDesign = efd
	}

	en, err := readInt(path + "/energy_now")
	if err != nil {
		return bat, errors.New("cannot get energy_now")
	} else {
		bat.EnergyNow = en
	}

	c, err := readInt(path + "/capacity")
	if err != nil {
		bat.Capacity = int(bat.EnergyNow / bat.EnergyFull * 100)
	} else {
		bat.Capacity = c
	}

	m, err := readString(path + "/manufacturer")
	if err != nil {
		bat.Manufacturer = "Unknown"
	} else {
		bat.Manufacturer = m
	}

	mn, err := readString(path + "/model_name")
	if err != nil {
		bat.ModelName = "Unknown"
	} else {
		bat.ModelName = mn
	}

	pn, err := readInt(path + "/power_now")
	if err != nil {
		return bat, errors.New("cannot get power_now")
	} else {
		bat.PowerNow = pn
	}

	sn, err := readString(path + "/serial_number")
	if err != nil {
		bat.SerialNumber = "Unknown"
	} else {
		bat.SerialNumber = sn
	}

	st, err := readString(path + "/status")
	if err != nil {
		return bat, errors.New("cannot get status")
	} else {
		bat.Status = st
	}

	te, err := readString(path + "/technology")
	if err != nil {
		bat.Technology = "Unknown"
	} else {
		bat.Technology = te
	}

	pr, err := readInt(path + "/present")
	if err != nil {
		bat.Present = false
	} else {
		bat.Present = pr == 1
	}

	vn, err := readInt(path + "/voltage_now")
	if err != nil {
		return bat, errors.New("cannot get voltage_now")
	} else {
		bat.VoltageNow = vn
	}

	vm, err := readInt(path + "/voltage_min_design")
	if err != nil {
		bat.VoltageMinDesign = 0
	} else {
		bat.VoltageMinDesign = vm
	}

	return bat, nil
}

// Returns a list of detected battery ids.
func GetDetected() ([]int, error) {
	files, err := ioutil.ReadDir(sysfsPrefix)
	if err != nil {
		return nil, err
	}

	ids := []int{}
	for _, f := range files {
		if !isBattery(filepath.Join(sysfsPrefix, f.Name())) {
			continue
		}
		id, err := strconv.Atoi(strings.TrimPrefix(f.Name(), "BAT"))
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	return ids, nil
}
