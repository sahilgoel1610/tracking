package aerospike
import "fmt"
import ."github.com/aerospike/aerospike-client-go"
import "time"
// import "../memory_store_interface"


type AeroStore struct{
	client *Client
	defaultNamespace string 
}

func GetAeroClient()(aeroStore AeroStore) {
	client,err := NewClient("172.28.128.3", 3000)
	if (err != nil ) {

	}
	return AeroStore { client, "test"}
}

func (as *AeroStore) AddPair(key string, value string) bool {
	keyObj, err := NewKey(as.defaultNamespace, as.defaultNamespace + time.Now().String(), key)
	bins := BinMap {"first": value}
	err = as.client.Add(nil, keyObj, bins)
	if (err != nil ) {
		fmt.Println(err)
	}
	return err == nil
}

func (as *AeroStore) GetValue (key string) string {
	keyObj, err := NewKey(as.defaultNamespace, as.defaultNamespace, key)
	value, err := as.client.Get(nil, keyObj)
	if (err != nil) {
		return value.Bins["first"].(string)
	} else {
		return "Error"
	}
}


func (as *AeroStore) RemoveKey(key string) bool{
	keyObj, err := NewKey(as.defaultNamespace, as.defaultNamespace, key)
	existed, err := as.client.Delete(nil, keyObj)
	return existed && err == nil
}