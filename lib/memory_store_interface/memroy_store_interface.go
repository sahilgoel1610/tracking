package memoryStoreInterface

type MemStore interface {
	AddPair(key string, value string) bool
	GetValue(key string) string
	RemoveKey(key string) bool
}