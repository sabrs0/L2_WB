package main

import "fmt"

type Rights map[string]string

type Role interface {
	rights() Rights
}

type User struct {
}

func (role User) rights() Rights {
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

type UserMaker struct {
}

func (roleMaker UserMaker) makeRole() Role {
	fmt.Println("User Created")
	return User{}
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

func main() {
	adminMaker := AdminMaker{}
	userMaker := UserMaker{}

	configuer := RoleConfiguer{}

	configuer.roleMaker = adminMaker
	configuer.configRole()

	configuer.roleMaker = userMaker
	configuer.configRole()
}
