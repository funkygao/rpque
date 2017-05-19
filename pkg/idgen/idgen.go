package idgen

// IdGenerator is an interface that generate global unique numeric id.
type IdGenerator interface {

	// Next returns a new global unique int64.
	Next() int64
}
