package util

import (
	// "errors"
	// "fmt"

	"github.com/spf13/cobra"
)

type FlagProvider struct {
	cmd     *cobra.Command
	flagSet map[string]any
}

func NewFlagProvider(cmd *cobra.Command) FlagProvider {
	return FlagProvider{cmd, make(map[string]any)}
}

func (self *FlagProvider) Register(
  name, shorthand string,
  defaultValue any,
  usage string,
) *FlagProvider {
  cmdFlags := self.cmd.PersistentFlags()
  
  switch defaultValue.(type) {
  case uint8:
    self.flagSet[name] = new(uint8)
    cmdFlags.Uint8VarP(
      self.flagSet[name].(*uint8),
      name,
      shorthand,
      defaultValue.(uint8),
      usage,
    )
  case uint16:
    self.flagSet[name] = new(uint16)
    cmdFlags.Uint16VarP(
      self.flagSet[name].(*uint16),
      name,
      shorthand,
      defaultValue.(uint16),
      usage,
    )
  case uint32:
    self.flagSet[name] = new(uint32)
    cmdFlags.Uint32VarP(
      self.flagSet[name].(*uint32),
      name,
      shorthand,
      defaultValue.(uint32),
      usage,
    )
  case uint64:
    self.flagSet[name] = new(uint64)
    cmdFlags.Uint64VarP(
      self.flagSet[name].(*uint64),
      name,
      shorthand,
      defaultValue.(uint64),
      usage,
    )
   case int8:
    self.flagSet[name] = new(int8)
    cmdFlags.Int8VarP(
      self.flagSet[name].(*int8),
      name,
      shorthand,
      defaultValue.(int8),
      usage,
    )
   case int16:
    self.flagSet[name] = new(int16)
    cmdFlags.Int16VarP(
      self.flagSet[name].(*int16),
      name,
      shorthand,
      defaultValue.(int16),
      usage,
    )
   case int32:
    self.flagSet[name] = new(int32)
    cmdFlags.Int32VarP(
      self.flagSet[name].(*int32),
      name,
      shorthand,
      defaultValue.(int32),
      usage,
    )
   case int64:
    self.flagSet[name] = new(int64)
    cmdFlags.Int64VarP(
      self.flagSet[name].(*int64),
      name,
      shorthand,
      defaultValue.(int64),
      usage,
    )
   case string:
    self.flagSet[name] = new(string)
    cmdFlags.StringVarP(
      self.flagSet[name].(*string),
      name,
      shorthand,
      defaultValue.(string),
      usage,
    )
   case bool:
    self.flagSet[name] = new(bool)
    cmdFlags.BoolVarP(
      self.flagSet[name].(*bool),
      name,
      shorthand,
      defaultValue.(bool),
      usage,
    )
  }
  return self
}

func (self *FlagProvider) Provided(name string) bool {
  return self.cmd.Flag(name) != nil
}

func (self *FlagProvider) Value(name string) any {
  flag := self.cmd.Flag(name)

  if flag.Changed {
    return self.flagSet[name]
  } else {
    return nil
  }
}

func (self *FlagProvider) ValueOrDefault(name string) any {
	return self.flagSet[name]
}
