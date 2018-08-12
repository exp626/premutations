package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Premutations!")
}

func (as *ActionSuite) Test_V1PremutationsInit() {
	res := as.JSON("/v1/init").Post([]int{1, 2, 3, 5, 4})
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "[1,2,3,4,5]")
}

func (as *ActionSuite) Test_V1PremutationsNext() {
	as.JSON("/v1/init").Post([]int{1, 2, 3, 5, 4})
	res := as.JSON("/v1/next").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "[1,2,3,5,4]")
}
