package memStoreFactory
import "../redis"
import "../aerospike"
import "../memory_store_interface"

func Get(typeOfStore string) (memStore memoryStoreInterface.MemStore) {
	var memStore1 memoryStoreInterface.MemStore
	if (typeOfStore == "redis") {
		redisStore := redis.GetRedisClient()
		memStore1 = &redisStore
	} else if (typeOfStore == "aerospike") {
		aeroStore := aerospike.GetAeroClient()
		memStore1 = &aeroStore
	} else {
		return nil
	}

	return memStore1;
}