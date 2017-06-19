package builder

import (
	"strconv"
)

// Product
type computer struct {
	model string
}

// Builder
type builder interface {
	buildCPU(param float64) string
	buildRAM(param int) string
	buildDisk(param int) string
	getComputer(model string) *computer
}

// ConcreteBuilder
type dumbBuilder struct {
}

func (db *dumbBuilder) getComputer(model string) *computer {
	return &computer{model}
}

func (db *dumbBuilder) buildCPU(param float64) string {
	return "CPU:" + strconv.FormatFloat(param, 'f', 2, 64) + "GHz"
}

func (db *dumbBuilder) buildRAM(param int) string {
	return "RAM:" + strconv.Itoa(param) + "GB"
}

func (db *dumbBuilder) buildDisk(param int) string {
	return "Disk:" + strconv.Itoa(param) + "GB"
}

// Director
type director struct {
	builder builder
}

func (d *director) Constructor() string {
	model := d.builder.buildCPU(3.30) + "-"
	model += d.builder.buildRAM(8) + "-"
	model += d.builder.buildDisk(500)

	return model
}
