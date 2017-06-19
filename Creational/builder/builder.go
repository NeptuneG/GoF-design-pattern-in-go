package builder

// Product
type team struct {
}

// Builder
type builder interface {
	getPointGuard()
	getShootingGuard()
	getSmallForward()
	getPowerForward()
	getCenter()
	getTeam() *team
}

type dumbBuilder struct {
}

func (db *dumbBuilder) getPointGuard() {
	
}

func (db *dumbBuilder) getShootingGuard() {
	
}

func (db *dumbBuilder) getSmallForward() {
	
}

func (db *dumbBuilder) getPowerForward() {
	
}

func (db *dumbBuilder) getCenter() {
	
}

func (db *dumbBuilder) getTeam() *team {
	return &team{}
}

// Director
type director struct {
	builder builder
}

func (d *director) Constructor() {
	d.builder.getCenter()
	d.builder.getPointGuard()
	d.builder.getPowerForward()
	d.builder.getShootingGuard()
	d.builder.getSmallForward()
}