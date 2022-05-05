package types

const (
	// ModuleName defines the module name
	ModuleName = "fortress"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fortress"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	FortressKey      = "Fortress-value-"
	FortressCountKey = "Fortress-count-"
)

const (
	PostKey      = "Post-value-"
	PostCountKey = "Post-count-"
)
