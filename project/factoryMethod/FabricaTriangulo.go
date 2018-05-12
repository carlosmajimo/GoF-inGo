package factorymethod

func FabricaTriangulo(l1, l2, l3 int) (output ITriangulo) {
	if l1 == l2 && l2 == l3 {
		output = &Equilatero{
			triangulo{
				L1: l1,
				L2: l2,
				L3: l3,
			},
		}
	} else if l1 == l2 || l1 == l3 || l2 == l3 {
		output = &Isoceles{
			triangulo{
				L1: l1,
				L2: l2,
				L3: l3,
			},
		}
	} else {
		if l1 != l2 || l1 != l3 || l3 != l2 {
			output = &Escaleno{
				triangulo{
					L1: l1,
					L2: l2,
					L3: l3,
				},
			}
		}
	}

	return
}
