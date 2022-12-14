// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package internal

import "encoding/json"

type VRMCSpringBoneShapeSchemaJsonCapsule struct {
	// The capsule head. vector3
	Offset []float64 `json:"offset,omitempty"`

	// The capsule radius
	Radius float64 `json:"radius,omitempty"`

	// The capsule tail. vector3
	Tail []float64 `json:"tail,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VRMCSpringBoneShapeSchemaJsonCapsule) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain VRMCSpringBoneShapeSchemaJsonCapsule
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["offset"]; !ok || v == nil {
		plain.Offset = []float64{
			0,
			0,
			0,
		}
	}
	if v, ok := raw["radius"]; !ok || v == nil {
		plain.Radius = 0
	}
	if v, ok := raw["tail"]; !ok || v == nil {
		plain.Tail = []float64{
			0,
			0,
			0,
		}
	}
	*j = VRMCSpringBoneShapeSchemaJsonCapsule(plain)
	return nil
}

type VRMCSpringBoneShapeSchemaJsonSphere struct {
	// The sphere center. vector3
	Offset []float64 `json:"offset,omitempty"`

	// The sphere radius
	Radius float64 `json:"radius,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VRMCSpringBoneShapeSchemaJsonSphere) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain VRMCSpringBoneShapeSchemaJsonSphere
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["offset"]; !ok || v == nil {
		plain.Offset = []float64{
			0,
			0,
			0,
		}
	}
	if v, ok := raw["radius"]; !ok || v == nil {
		plain.Radius = 0
	}
	*j = VRMCSpringBoneShapeSchemaJsonSphere(plain)
	return nil
}

// Shape of collider. Have one of sphere and capsule
type VRMCSpringBoneShapeSchemaJson struct {
	// Capsule corresponds to the JSON schema field "capsule".
	Capsule *VRMCSpringBoneShapeSchemaJsonCapsule `json:"capsule,omitempty"`

	// Extensions corresponds to the JSON schema field "extensions".
	Extensions interface{} `json:"extensions,omitempty"`

	// Extras corresponds to the JSON schema field "extras".
	Extras interface{} `json:"extras,omitempty"`

	// Sphere corresponds to the JSON schema field "sphere".
	Sphere *VRMCSpringBoneShapeSchemaJsonSphere `json:"sphere,omitempty"`
}
