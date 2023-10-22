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
	switch role {
	case UserRoleUser:
		return "user"
	case UserRoleModerator:
		return "moderator"
	case UserRoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
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
