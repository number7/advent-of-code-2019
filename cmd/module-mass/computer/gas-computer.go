package computer

func ComputeGas(mass int) int {
	return compute(mass, new(int))
}

func compute(mass int, acc *int) int {
	if mass <= 0 {
		return *acc
	}
	gas := (mass / 3) - 2
	if gas > 0 {
		*acc += gas
	}
	return compute(gas, acc)
}
