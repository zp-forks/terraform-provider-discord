package discord

import (
	"context"
	"log"

	"github.com/andersfylling/disgord"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Role struct {
	ServerId disgord.Snowflake
	RoleId   disgord.Snowflake
	Role     *disgord.Role
}

func insertRole(array []*disgord.Role, value *disgord.Role, index int) []*disgord.Role {
	return append(array[:index], append([]*disgord.Role{value}, array[index:]...)...)
}

func removeRole(array []*disgord.Role, index int) []*disgord.Role {
	return append(array[:index], array[index+1:]...)
}

func removeRoleById(array []disgord.Snowflake, id disgord.Snowflake) []disgord.Snowflake {
	roles := make([]disgord.Snowflake, 0, len(array))
	for _, x := range array {
		if x != id {
			roles = append(roles, x)
		}
	}

	return roles
}

func moveRole(array []*disgord.Role, srcIndex int, dstIndex int) []*disgord.Role {
	value := array[srcIndex]
	return insertRole(removeRole(array, srcIndex), value, dstIndex)
}

func findRoleIndex(array []*disgord.Role, value *disgord.Role) (int, bool) {
	for index, element := range array {
		if element.ID == value.ID {
			return index, true
		}
	}

	return -1, false
}

func findRoleById(array []*disgord.Role, id disgord.Snowflake) *disgord.Role {
	for _, element := range array {
		if element.ID == id {
			return element
		}
	}

	return nil
}

func reorderRoles(ctx context.Context, m interface{}, serverId disgord.Snowflake, role *disgord.Role, position int) (bool, diag.Diagnostics) {
	client := m.(*Context).Client

	roles, err := client.Guild(serverId).GetRoles()
	if err != nil {
		return false, diag.Errorf("Failed to fetch roles: %s", err.Error())
	}
	index, exists := findRoleIndex(roles, role)
	if !exists {
		return false, diag.Errorf("Role somehow does not exists")
	}

	moveRole(roles, index, position)

	params := make([]disgord.UpdateGuildRolePositions, 0, len(roles))
	for index, r := range roles {
		params = append(params, disgord.UpdateGuildRolePositions{ID: r.ID, Position: index})
	}

	// FIXME: ロギングいる？
	log.Print(spew.Sdump(params))

	if roles, err = client.Guild(serverId).UpdateRolePositions(params); err != nil {
		return false, diag.Errorf("Failed to re-order roles: %s", err.Error())
	}

	return true, nil
}

func getRole(ctx context.Context, client *disgord.Client, serverId disgord.Snowflake, roleId disgord.Snowflake) (*disgord.Role, error) {
	if roles, err := client.Guild(serverId).GetRoles(); err != nil {
		return nil, err
	} else {
		role := findRoleById(roles, roleId)
		return role, nil
	}
}
