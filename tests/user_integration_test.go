package tests

//-------------------------
// IMPORTANT: To run this you should enable the database destruction on the env vars.
//-------------------------

/*func TestUsersCRUDIntegrationTest(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// prepare
	config := config.New("../.env")
	config.Database.Purge = true
	database := repository.NewDatabase(config.Config, nopLogger{})
	repository := repository.New(database)
	service := service.New(repository, nopAuth{}, config)
	endpoints := transport.New(service, transport.NewErrorsMapper(nopLogger{}))

	// happy run :)
	testSignup(t, endpoints)
	testLogin(t, endpoints, "test")
	testGetUser(t, endpoints, http.StatusOK)
	testUpdateUser(t, endpoints)
	testLogin(t, endpoints, "test2")
	testDeleteUser(t, endpoints)
	testGetUser(t, endpoints, http.StatusNotFound)

	// admin run :o
	testCreateUser(t, endpoints)
	testLogin(t, endpoints, "admin")
}

func testSignup(t *testing.T, endpoints transport.TransportLayer) {
	c := makeTestContextWithHTTPRequest(map[string]string{
		"username":        "test",
		"email":           "test@email.com",
		"password":        "password",
		"repeat_password": "password",
	})
	transport.HandleRequest(endpoints, c, requests.MakeSignupRequest, endpoints.Service().Signup)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func testLogin(t *testing.T, endpoints transport.TransportLayer, username string) {
	c := makeTestContextWithHTTPRequest(map[string]string{
		"username_or_email": username,
		"password":          "password",
	})
	transport.HandleRequest(endpoints, c, requests.MakeLoginRequest, endpoints.Service().Login)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func testCreateUser(t *testing.T, endpoints transport.TransportLayer) {
	c := makeTestContextWithHTTPRequest(map[string]any{
		"email":    "admin@email.com",
		"username": "admin",
		"password": "password",
		"is_admin": true,
	})
	transport.HandleRequest(endpoints, c, requests.MakeCreateUserRequest, endpoints.Service().CreateUser)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func testGetUser(t *testing.T, endpoints transport.TransportLayer, status int) {
	c := makeTestContextWithHTTPRequest(map[string]string{})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeGetUserRequest, endpoints.Service().GetUser)
	assert.Equal(t, status, c.Writer.Status())
}

func testUpdateUser(t *testing.T, endpoints transport.TransportLayer) {
	c := makeTestContextWithHTTPRequest(map[string]string{
		"username": "test2",
		"email":    "test2@email.com",
	})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeUpdateUserRequest, endpoints.Service().UpdateUser)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func testDeleteUser(t *testing.T, endpoints transport.TransportLayer) {
	c := makeTestContextWithHTTPRequest(map[string]string{})
	addValueAndParamToContext(c, "ID", 1, "user_id", "1")
	transport.HandleRequest(endpoints, c, requests.MakeDeleteUserRequest, endpoints.Service().DeleteUser)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

//--------------------
//      HELPERS
//--------------------

func makeTestHTTPRequest(body []byte) *http.Request {
	req, _ := http.NewRequest("", "", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func addRequestToContext(request *http.Request) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	return c
}

func makeTestContextWithHTTPRequest(body any) *gin.Context {
	jsonData, _ := json.Marshal(body)
	httpReq := makeTestHTTPRequest(jsonData)
	return addRequestToContext(httpReq)
}

func addValueAndParamToContext(context *gin.Context, ctxKey string, ctxValue any, paramKey, paramValue string) {
	context.Set(ctxKey, ctxValue)
	context.Params = append(context.Params, gin.Param{Key: paramKey, Value: paramValue})
}

//--------------------
//       NOPS
//--------------------

type nopLogger struct{}

func (l nopLogger) Info(args ...interface{})                  {}
func (l nopLogger) Warn(args ...interface{})                  {}
func (l nopLogger) Error(args ...interface{})                 {}
func (l nopLogger) Fatalf(format string, args ...interface{}) {}

type nopAuth struct{}

func (a nopAuth) GenerateToken(user auth.User, role auth.Role) (string, error)         { return "", nil }
func (a nopAuth) ValidateToken(role auth.Role, shouldMatchUserID bool) gin.HandlerFunc { return nil }
*/