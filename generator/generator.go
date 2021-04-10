// Package generator provides an interface and implementations of generators to create a group of points.
package generator

import "github.com/RH12503/Triangula/normgeom"

// A Generator is used to generate a normgeom.NormPointGroup
type Generator interface {
	// Generate returns a generated normgeom.NormPointGroup with a specified number of generator
	Generate(numPoints int) normgeom.NormPointGroup
}
