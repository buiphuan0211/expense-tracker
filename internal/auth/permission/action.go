package permission

import "fmt"

var action = struct {
	Admin string
	View  string
	Edit  string
}{
	Admin: "admin",
	View:  "view",
	Edit:  "edit",
}

type listPermissions struct {
	View []string
	Edit []string
}

func generatePermissions(value, permissions string) []string {
	return []string{
		fmt.Sprintf("%s_%s", value, permissions),
		//fmt.Sprintf("%s_%s", value, action.Admin),
	}
}

func getPermissionsView(value string) []string {
	return generatePermissions(value, action.View)
}

func getPermissionsEdit(value string) []string {
	return generatePermissions(value, action.Edit)
}
