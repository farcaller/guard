// Code generated by go-enum
// DO NOT EDIT!

package ldap

import (
	"fmt"
	"strings"
)

const (
	// AuthChoiceSimple is a AuthChoice of type Simple
	AuthChoiceSimple AuthChoice = iota
	// AuthChoiceKerberos is a AuthChoice of type Kerberos
	AuthChoiceKerberos
)

const _AuthChoiceName = "SimpleKerberos"

var _AuthChoiceMap = map[AuthChoice]string{
	0: _AuthChoiceName[0:6],
	1: _AuthChoiceName[6:14],
}

func (i AuthChoice) String() string {
	if str, ok := _AuthChoiceMap[i]; ok {
		return str
	}
	return fmt.Sprintf("AuthChoice(%d)", i)
}

var _AuthChoiceValue = map[string]AuthChoice{
	_AuthChoiceName[0:6]:                   0,
	strings.ToLower(_AuthChoiceName[0:6]):  0,
	_AuthChoiceName[6:14]:                  1,
	strings.ToLower(_AuthChoiceName[6:14]): 1,
}

// ParseAuthChoice attempts to convert a string to a AuthChoice
func ParseAuthChoice(name string) (AuthChoice, error) {
	if x, ok := _AuthChoiceValue[name]; ok {
		return AuthChoice(x), nil
	}
	return AuthChoice(0), fmt.Errorf("%s is not a valid AuthChoice", name)
}

// Set implements the Golang flag.Value interface func
func (x *AuthChoice) Set(val string) error {
	v, err := ParseAuthChoice(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func
func (x *AuthChoice) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface
func (x *AuthChoice) Type() string {
	return "AuthChoice"
}
