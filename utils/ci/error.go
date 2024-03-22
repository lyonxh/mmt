package ci

import "errors"

var (
	HarborRequestError        = errors.New("请求harbor错误")
	HarborProjectAlreadyExist = errors.New("The Project Named Already Exists")
	HarborUnauthorized        = errors.New("Invide Authorized")
	HarborServerError         = errors.New("Internal Server Error")
	HarborProjectNotFound     = errors.New("Project Name Not Found")

	GiteaRequestError = errors.New("请求gitea错误")
	GiteaUnauthorized = errors.New("Token Is Required")
	GiteaHookNotFound = errors.New("Hook Not Found")
	GiteaUserNotAdmin = errors.New("user is not admin")

	BuildRequestError  = errors.New("请求build错误")
	TektonRequestError = errors.New("请求tekton错误")

	GitlabRequestError = errors.New("请求gitlab错误")
	GitlabUnauthorized = errors.New("Invide Authorized")
	GitlabHookNotFound = errors.New("Hook Not Found")
)
