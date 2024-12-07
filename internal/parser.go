package internal

//Takes in arr bytes and return arr byte
var CommandParser = map[string] func (buff []byte) ([]byte, error) {
    "$": aggregate,
    "#": flush,
    "&": simplestring,
    "%": integer,
}

func aggregate(buff []byte) ([]byte, error)  {
    return buff, nil
}

func flush(buff []byte) ([]byte, error) {
    return buff, nil
}

func simplestring(buff []byte) ([]byte, error) {
    return buff, nil
}

func integer(buff []byte) ([]byte, error) {
    return buff, nil
}