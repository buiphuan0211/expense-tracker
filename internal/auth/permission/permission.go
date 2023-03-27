package permission

const (
	category = "category"
	staff    = "staff"
)

var Scope = struct {
	Category listPermissions
	Staff    listPermissions
}{
	Category: listPermissions{
		View: getPermissionsView(category),
		Edit: getPermissionsEdit(category),
	},
}
