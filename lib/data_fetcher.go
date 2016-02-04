package lib

import (
	"snow_api/models"
	"strconv"
	"strings"
)

func FetchListOfSnowUsers(sysparm_limit string) models.Userslice {
	var snowUsers models.Userslice
	snowUserdata := new(models.SNOWUserData)
	username := ParseJsonConfig().User
	password := ParseJsonConfig().Password
	instance := ParseJsonConfig().Instance

	defaultSysparmLimit := "10"
	if sysparm_limit != "" {
		GetHttp(
			"https://"+instance+".service-now.com/api/now/table/sys_user?sysparm_limit="+sysparm_limit,
			username,
			password,
			snowUserdata,
		)
	} else {
		GetHttp(
			"https://"+instance+".service-now.com/api/now/table/sys_user?sysparm_limit="+defaultSysparmLimit,
			username,
			password,
			snowUserdata,
		)
	}
	sysparm_limitInt, err := strconv.Atoi(sysparm_limit)
	if err != nil {
		sysparm_limitInt, _ = strconv.Atoi(defaultSysparmLimit)
	}
	for i := 0; i < sysparm_limitInt; i++ {
		snowUsers = append(snowUsers, models.User{
			Sys_Id:    snowUserdata.Result[i].SysID,
			FirstName: snowUserdata.Result[i].FirstName,
			LastName:  snowUserdata.Result[i].LastName,
			Email:     snowUserdata.Result[i].Email,
		})
	}
	return snowUsers
}

func FetchOneUser(email string) models.Userslice {
	// Parse email and replace @ with ASCII encoded value, %40
	urlEncodedEmail := strings.Replace(email, "@", "%40", -1)
	snowUserData := new(models.SNOWUserData)
	GetHttp(
		"https://dev15006.service-now.com/api/now/table/sys_user?email="+urlEncodedEmail,
		"admin",
		"Professi0n@l",
		snowUserData,
	)
	var snowUser models.Userslice
	snowUser = append(snowUser, models.User{
		Sys_Id:    snowUserData.Result[0].SysID,
		FirstName: snowUserData.Result[0].FirstName,
		LastName:  snowUserData.Result[0].LastName,
		Email:     snowUserData.Result[0].Email,
	})
	return snowUser
}
