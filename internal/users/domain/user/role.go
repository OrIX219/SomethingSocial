package users

import "fmt"

const (
	UserRoleUnknown   UserRole = iota
	UserRoleUser      UserRole = iota
	UserRoleModerator UserRole = iota
	UserRoleAdmin     UserRole = iota
)

type UserRole int

func (role UserRole) String() string {
	return [...]string{
		"unknown",
		"user",
		"moderator",
		"admin",
	}[role]
}

func ParseRole(role string) (UserRole, error) {
	switch role {
	case "user":
		return UserRoleUser, nil
	case "moderator":
		return UserRoleModerator, nil
	case "admin":
		return UserRoleAdmin, nil
	default:
		return UserRoleUnknown, fmt.Errorf("Invalid user role: %s", role)
	}
}

type NotEnoughRightsError struct {
	UserId   int64
	UserRole UserRole
	Action   string
}

func (e NotEnoughRightsError) Error() string {
	return fmt.Sprintf("User %d (%s) has not enough rights for: %s",
		e.UserId, e.UserRole, e.Action)
}
