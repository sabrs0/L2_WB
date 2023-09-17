package pattern

import "fmt"

type Rights map[string]string

type Role interface {
	rights() Rights
}

type DefaultUser struct {
}

func (role DefaultUser) rights() Rights {
	rights := make(Rights)
	rights["db"] = "nothing"
	return rights
}

type Analytic struct {
}

func (role Analytic) rights() Rights {
	rights := make(Rights)
	rights["db"] = "read only"
	return rights
}

type Admin struct {
}

func (role Admin) rights() Rights {
	rights := make(Rights)
	rights["db"] = "read write add"
	return rights
}

type RoleMaker interface {
	makeRole() Role
}

type AdminMaker struct {
}

func (roleMaker AdminMaker) makeRole() Role {
	fmt.Println("Admin Created")

	return Admin{}
}

type DefaultUserMaker struct {
}

func (roleMaker DefaultUserMaker) makeRole() Role {
	fmt.Println("User Created")
	return DefaultUser{}
}

type AnalyticMaker struct {
}

func (roleMaker AnalyticMaker) makeRole() Role {
	fmt.Println("Analytic Created")
	return Analytic{}
}

type RoleConfiguer struct {
	roleMaker RoleMaker
}

func (configuer RoleConfiguer) configRole() {
	role := configuer.roleMaker.makeRole()
	fmt.Println(role.rights())
}

func factoryMethodPattern() {
	adminMaker := AdminMaker{}
	userMaker := DefaultUserMaker{}

	configuer := RoleConfiguer{}

	configuer.roleMaker = adminMaker
	configuer.configRole()

	configuer.roleMaker = userMaker
	configuer.configRole()
}
