// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldMixedString holds the string denoting the mixed_string field in the database.
	FieldMixedString = "mixed_string"
	// FieldMixedEnum holds the string denoting the mixed_enum field in the database.
	FieldMixedEnum = "mixed_enum"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBuffer holds the string denoting the buffer field in the database.
	FieldBuffer = "buffer"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldNewName holds the string denoting the new_name field in the database.
	FieldNewName = "renamed"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"

	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"

	// CarFieldID holds the string denoting the id field of the Car.
	CarFieldID = "id"
	// PetFieldID holds the string denoting the id field of the Pet.
	PetFieldID = "id"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CarTable is the table the holds the car relation/edge.
	CarTable = "cars"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "user_car"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "owner_id"
	// FriendsTable is the table the holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friends"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldMixedString,
	FieldMixedEnum,
	FieldAge,
	FieldName,
	FieldNickname,
	FieldPhone,
	FieldBuffer,
	FieldTitle,
	FieldNewName,
	FieldBlob,
	FieldState,
	FieldStatus,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user", "friend"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMixedString holds the default value on creation for the mixed_string field.
	DefaultMixedString string
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultPhone holds the default value on creation for the phone field.
	DefaultPhone string
	// DefaultTitle holds the default value on creation for the title field.
	DefaultTitle string
)

// MixedEnum defines the type for the mixed_enum enum field.
type MixedEnum string

// MixedEnumOn is the default MixedEnum.
const DefaultMixedEnum = MixedEnumOn

// MixedEnum values.
const (
	MixedEnumOn  MixedEnum = "on"
	MixedEnumOff MixedEnum = "off"
)

func (me MixedEnum) String() string {
	return string(me)
}

// MixedEnumValidator is a validator for the "mixed_enum" field enum values. It is called by the builders before save.
func MixedEnumValidator(me MixedEnum) error {
	switch me {
	case MixedEnumOn, MixedEnumOff:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for mixed_enum field: %q", me)
	}
}

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
	StateOnline    State = "online"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateLoggedIn, StateLoggedOut, StateOnline:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", s)
	}
}

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusDone    Status = "done"
	StatusPending Status = "pending"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDone, StatusPending:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}
