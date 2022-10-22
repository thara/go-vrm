package vrm1migrate

func migrateVec(x, y, z float64) []float64 {
	return []float64{-x, y, z}
}

type vec2 struct {
	x, y float64
}
