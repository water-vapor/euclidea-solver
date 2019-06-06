package configs

const (
	// HashPrecision is the inverse of Tolerance
	HashPrecision    = 100000000
	// Prime is the coefficient used in hash function
	Prime            = 19260817
	// Tolerance is the value below which two values are considered the same
	Tolerance        = 1e-8
	// EarlyStop determines if we check at early levels for solution
	EarlyStop        = true
	// ImageSize is the resolution of the output image
	ImageSize        = 1000
	// MaxPointCoord is maximum coordinate of a point to be displayed on the image
	MaxPointCoord    = 1e3
	// RandomPointRange is the initiating coordinate range of a random point
	RandomPointRange = 10
)
