// Code generated by "enumer -type=messageType -json"; DO NOT EDIT

package main

import "fmt"
import "encoding/json"

const _messageType_name = "typeDevicestypeSharetypeDeviceUpdatetypeError"

var _messageType_index = [...]uint8{0, 11, 20, 36, 45}

func (i messageType) String() string {
	if i < 0 || i >= messageType(len(_messageType_index)-1) {
		return fmt.Sprintf("messageType(%d)", i)
	}
	return _messageType_name[_messageType_index[i]:_messageType_index[i+1]]
}

var _messageTypeNameToValue_map = map[string]messageType{
	_messageType_name[0:11]:  0,
	_messageType_name[11:20]: 1,
	_messageType_name[20:36]: 2,
	_messageType_name[36:45]: 3,
}

func messageTypeString(s string) (messageType, error) {
	if val, ok := _messageTypeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to messageType values", s)
}

func (i messageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *messageType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("messageType should be a string, got %s", data)
	}

	var err error
	*i, err = messageTypeString(s)
	return err
}
