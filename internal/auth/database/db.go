package db

/// REPOSITORY LAYER

type Repo struct {
	_ map[string]uint32
}

// Connect to DB
func Connect() (Repo, error) {
	return Repo{make(map[string]uint32)}, nil
}
