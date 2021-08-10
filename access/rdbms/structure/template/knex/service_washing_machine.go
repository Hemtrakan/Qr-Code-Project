package knex

type serviceWashingMachine struct {
	history []washingMachineInfo
	ops     washingMachineOps
}
