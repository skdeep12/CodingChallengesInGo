package payroll

import "fmt"

type database struct {
	m            map[string]any
	unionMembers map[string]string
}

func (d *database) AddEmployee(e Identity) {
	d.m[e.GetID()] = e
}

func (d *database) GetEmployee(id string) any {
	if val, ok := d.m[id]; !ok {
		return nil
	} else {
		return val
	}
}

func (d *database) DeleteEmployee(id string) {
	delete(d.m, id)
}

func (d *database) GetUnionMember(id string) any {
	if val, ok := d.unionMembers[id]; !ok {
		return nil
	} else {
		return d.m[val]
	}
}

func (d *database) AddUnionMember(memberID, empID string) error {
	emp := d.GetEmployee(empID)
	if emp == nil {
		return fmt.Errorf("employee with id %s does not exist", empID)
	}
	d.unionMembers[memberID] = empID
	return nil
}

type Identity interface {
	GetID() string
}

var PayrollDatabase = database{
	m: make(map[string]any),
}
