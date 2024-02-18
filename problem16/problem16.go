package problem16

import (
	"errors"
)

func Power(base float64, exp int) (float64, error) {
	if base == 0.0 && exp < 0 {
		return 0.0, errors.New("base == 0 && exp < 0")
	}
	unsignedExp := exp
	if exp < 0 {
		unsignedExp = -exp
	}
	res := PowerWithUnsignedExp(base, unsignedExp)
	if exp < 0 {
		res = 1.0 / res
	}
	return res, nil
}

func PowerWithUnsignedExp(base float64, exp int) float64 {
	if exp == 0 {
		return 1.0
	}
	if exp == 1 {
		return base
	}
	res, fin := base, 1.0
	if exp&0x1 == 1 {
		fin = base
	}
	for exp >= 2 {
		res *= res
		exp >>= 1
	}
	res *= fin
	return res
}
