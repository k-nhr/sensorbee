package tuple

type Map map[string]Value

func (m Map) Type() TypeID {
	return TypeMap
}

func (m Map) AsBool() (Bool, error) {
	return false, castError(m.Type(), TypeBool)
}

func (m Map) AsInt() (Int, error) {
	return 0, castError(m.Type(), TypeInt)
}

func (m Map) AsFloat() (Float, error) {
	return 0, castError(m.Type(), TypeFloat)
}

func (m Map) AsString() (String, error) {
	return "", castError(m.Type(), TypeString)
}

func (m Map) AsBlob() (Blob, error) {
	return nil, castError(m.Type(), TypeBlob)
}

func (m Map) AsTimestamp() (Timestamp, error) {
	return Timestamp{}, castError(m.Type(), TypeTimestamp)
}

func (m Map) AsArray() (Array, error) {
	return nil, castError(m.Type(), TypeArray)
}

func (m Map) AsMap() (Map, error) {
	return m, nil
}

func (m Map) clone() Value {
	return m.Copy()
}

func (m Map) Copy() Map {
	out := make(map[string]Value, len(m))
	for key, val := range m {
		out[key] = val.clone()
	}
	return Map(out)
}

// Get returns value(s) from a structured Map as addressed by the
// given path expression. Returns an error when the path expression
// is invalid or the path is not found in the Map.
//
// Type conversion can be done for each type using the Value
// interface's methods.
//
// Example:
//  v, err := map.Get("path")
//  s, err := v.AsString() // cast to String
//
// Path Expression Example:
// Given the following Map structure
//  Map{
//  	"store": Map{
//  		"name": String("store name"),
//  		"book": Array([]Value{
//  			Map{
//  				"title": String("book name"),
//  			},
//  		}),
//  	},
//  }
// To get values, access the following path expressions
//  `store`              -> get store's Map
//  `store.name`         -> get "store name"
//  `store.book[0].title -> get "book name"
// or
//  `["store"]`                     -> get store's Map
//  `["store"]["name"]`             -> get "store name"
//  `["store"]["book"][0]["title"]` -> get "book name"
//
func (m Map) Get(path string) (Value, error) {
	var v Value
	err := scanMap(m, path, &v)
	return v, err
}
