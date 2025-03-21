package configs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	typeString   = "string"
	typeInt      = "int"
	typeFloat    = "float"
	typeBool     = "bool"
	typeDuration = "duration"
)

type Variables map[string]Variable

func NewVariablesInstance() *Variables {
	return &Variables{}
}

func (v Variables) Range() {
	for k, v := range v {
		fmt.Printf("%s: %+v\n", k, v)
	}
}

type Variable struct {
	Type     string      `toml:"type"`
	Name     string      `toml:"name"`
	Value    interface{} `toml:"value"`
	EnvName  string      `toml:"envName"`
	Required bool        `toml:"required"`
}

func (v Variables) Get(name string) Variable {
	val, ok := v[name]
	if !ok {
		name = strings.ToUpper(name)
		val, ok = v[name]
		if !ok {
			return Variable{}
		}
	}
	return val
}

func (v Variables) Collect(envs map[string]string) error {
	for name, val := range envs {
		val, err := parse(name, val)
		if err != nil {
			return err
		}
		if _, ok := v[name]; !ok {
			v[name] = val
			continue
		}
	}
	return nil
}

func (v Variable) GetString() string {
	if v.Type != typeString {
		return ""
	}
	val, ok := v.Value.(string)
	if !ok {
		return ""
	}
	return val
}

func (v Variable) GetFloat() float64 {
	if v.Type != typeFloat {
		return -1.0
	}
	val, ok := v.Value.(float64)
	if !ok {
		return 0.0
	}
	return val
}

func (v Variable) GetBool() bool {
	if v.Type != typeBool {
		return false
	}
	val, ok := v.Value.(bool)
	if !ok {
		return false
	}
	return val
}

func (v Variable) GetInt() int {
	if v.Type != typeInt {
		return -1
	}
	val, ok := v.Value.(int)
	if !ok {
		return 0
	}
	return val
}

func (v Variable) GetDuration() time.Duration {
	if v.Type != typeDuration {
		return 0
	}
	val, ok := v.Value.(string)
	if !ok {
		return 0
	}
	dur, _ := time.ParseDuration(val)
	return dur
}

func (v Variables) Validate() error {
	for name, val := range v {
		if err := val.Validate(); err != nil {
			return fmt.Errorf("variable %s: %w", name, err)
		}
	}
	return nil
}

func (v *Variable) Validate() error {
	var err error
	switch v.Type {
	case typeString:
		if _, ok := v.Value.(string); !ok {
			return fmt.Errorf("value for %s is not string", v.Name)
		}
	case typeInt:
		v.Value, err = strconv.Atoi(v.Value.(string))
		if err != nil {
			return fmt.Errorf("value for %s is not int", v.Name)
		}
	case typeFloat:
		v.Value, err = strconv.ParseFloat(v.Value.(string), 64)
		if err != nil {
			return fmt.Errorf("value for %s is not float", v.Name)
		}
	case typeBool:
		v.Value, err = strconv.ParseBool(v.Value.(string))
		if err != nil {
			return fmt.Errorf("value for %s is not bool", v.Name)
		}
	case typeDuration:
		durStr, ok := v.Value.(string)
		if !ok {
			return fmt.Errorf("value for %s is not duration", v.Name)
		}
		_, err := time.ParseDuration(durStr)
		if err != nil {
			return fmt.Errorf("value for %s is not duration", v.Name)
		}
	default:
		return fmt.Errorf("unknown type %s", v.Type)
	}
	return nil
}

// nolint: unparam
func parse(name string, val any) (Variable, error) {
	str, ok := val.(string)
	if !ok {
		str = ""
	}
	numVal, err := strconv.Atoi(str)
	if err == nil {
		return Variable{
			Type:    typeInt,
			Name:    name,
			Value:   numVal,
			EnvName: name,
		}, nil
	}
	numFloat, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return Variable{
			Type:    typeFloat,
			Name:    name,
			Value:   numFloat,
			EnvName: name,
		}, nil
	}
	boolVal, err := strconv.ParseBool(str)
	if err == nil {
		return Variable{
			Type:    typeBool,
			Name:    name,
			Value:   boolVal,
			EnvName: name,
		}, nil
	}
	_, err = time.ParseDuration(str)
	if err == nil {
		return Variable{
			Type:    typeDuration,
			Name:    name,
			Value:   str,
			EnvName: name,
		}, nil
	}
	return Variable{
		Type:    typeString,
		Name:    name,
		Value:   str,
		EnvName: name,
	}, nil
}
