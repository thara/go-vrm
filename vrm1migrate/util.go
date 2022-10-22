package vrm1migrate

import (
	"github.com/qmuntal/gltf"
	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

func clamp[T number](v, min, max T) T {
	if v < min {
		return min
	} else if max < v {
		return max
	}
	return v
}

type toRGBA interface {
	RGBA() (r, g, b, a uint32)
}

func toFloat4[T constraints.Float](src toRGBA) [4]T {
	r, g, b, a := src.RGBA()
	return [4]T{T(r), T(g), T(b), T(a)}
}

func toFloat3[T constraints.Float](src toRGBA) [3]T {
	r, g, b, _ := src.RGBA()
	return [3]T{T(r), T(g), T(b)}
}

func getShadingRange0x(toony, shift float32) (min, max float32) {
	min = shift
	max = leap(1, shift, toony)
	return
}

func leap(a, b float32, t float32) float32 {
	r := float32(b - a)
	return a + r*t
}

func addExtensionsUsed(doc *gltf.Document, extName string) {
	for _, name := range doc.ExtensionsUsed {
		if name == extName {
			return
		}
	}
	doc.ExtensionsUsed = append(doc.ExtensionsUsed, extName)
}
