package handler

import (
	_ "bytes"
	"encoding/json"
	
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/mock"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUseCase := mock.NewMockUserUseCase(ctrl)
	cartUseCase := mock.NewMockCartUseCase(ctrl)
	userHandler := NewUserHandler(userUseCase, cartUseCase)

	
	tests := []struct {
		name          string
		input         models.UserSignUp
		beforeTest    func(mock.MockUserUseCase, models.UserSignUp)
		responseBody response.Response
	}{
		{
			name: "user signup",
			input: models.UserSignUp{
				Name:     "aarathy",
				Email:    "aarathy@gmail.com",
				PhNo:     "+919745503907",
				Password: "1234",
			},
			
			beforeTest: func(muuc mock.MockUserUseCase, signup models.UserSignUp) {
				muuc.EXPECT().UserSignUp(signup).Times(1).Return(models.TokenResponse{
					Token: "lkjjjzhbxkjggjnsuahjkll",
					UserDetails: models.UserDetails{
						ID:    1,
						Name:  "aarathy",
						Email: "aarathy@gmail.com",
						PhNo:  "+919745503907",
					},
				}, nil)
			},
			responseBody: response.Response{
				Statuscode: http.StatusCreated,
				Message: "successfully signedup",
				Data: models.TokenResponse{
					Token: "lkjjjzhbxkjggjnsuahjkll",
					UserDetails: models.UserDetails{
						ID:    1,
						Name:  "aarathy",
						Email: "aarathy@gmail.com",
						PhNo:  "+919745503907",
					},	
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*userUseCase, tt.input)
			router:=gin.Default()
			router.POST("/signup",userHandler.SignUpHandler)	

			mockRequest := httptest.NewRequest("POST", "/signup", strings.NewReader(`
			{
				"name":"aarathy",
				"email":"aarathy@gmail.com",
				"phno":"+919745503907",
				"password":"1234"
			}`))
			
			mockRequest.Header.Set("Content-Type","application/json")
			w:=httptest.NewRecorder()
			router.ServeHTTP(w,mockRequest)
			var actualResponse response.Response
			json.Unmarshal(w.Body.Bytes(),&actualResponse)
			assert.Equal(t,tt.responseBody.Statuscode,actualResponse.Statuscode)
			assert.Equal(t,tt.responseBody.Error,actualResponse.Error)
			assert.Equal(t,tt.responseBody.Message,actualResponse.Message)
			assert.Equal(t,tt.responseBody.Data.(models.TokenResponse).Token,actualResponse.Data.(map[string]interface{})["Token"])
		})
	}
}
