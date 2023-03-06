package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/e-commerce/api/params"
	mockdb "project/e-commerce/db/mock"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/global"
	"project/e-commerce/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	user, password := randomUser(t)

	testCase := []struct {
		name          string
		body          gin.H
		buidStubs     func(store *mockdb.MockStore)
		response      params.UserResponse
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder, user params.UserResponse)
	}{
		{
			name: "ok",
			body: gin.H{
				"user_name": user.UserName,
				"password":  password,
				"email":     user.Email,
			},
			buidStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).Times(1).
					Return(user, nil)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).Times(1)
			},
			response: params.UserResponse{
				UserName: user.UserName,
				Email:    user.Email,
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder, user params.UserResponse) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotEmpty(t, user)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mock store object
			store := mockdb.NewMockStore(ctrl)
			// build stubs
			tc.buidStubs(store)

			// start test server and send request
			server := newTestServer(global.CONFIG, store)
			recorder := httptest.NewRecorder()

			// Marshal body to JSON data
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// request
			url := "/login"
			requset, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			//
			server.Router.ServeHTTP(recorder, requset)

			// verfy response
			tc.checkResponse(t, recorder, tc.response)

		})
	}
}

func randomUser(t *testing.T) (user sqlc.User, password string) {
	password = utils.RandomString(6)
	hashedPassword, err := utils.HashedPassword(password)
	require.NoError(t, err)

	user = sqlc.User{
		UserName:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	return
}
