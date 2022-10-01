package vrm1migrate

import "errors"

func migrateVec(x, y, z *float64) ([]float64, error) {
	if x == nil || y == nil || z == nil {
		return []float64{}, errors.New("invalid vec")
	}
	return []float64{-(*x), *y, *z}, nil
}
