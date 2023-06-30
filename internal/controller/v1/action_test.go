package v1_test

import (
	"atados/challenger/internal/dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestActionController_CreateAction(t *testing.T) {

	testCases := []struct {
		name                   string
		actionRequest          *dto.CreateActionRequest
		urlRequested           string
		expectedHttpStatusCode int
		facadeActionResponse   *dto.CreateActionResponse
		facadeActionError      error
		expectedError          error
	}{
		{
			name: "OK",
			actionRequest: &dto.CreateActionRequest{
				Name:         "Action test",
				Institution:  "Institution fake",
				City:         "São Paulo",
				Neighborhood: "Limoeiro",
				Address:      "Rua Palmeira, 25",
				Description:  "Reuniao as 15 horas",
			},
			urlRequested:           "/atados/v1/action",
			expectedHttpStatusCode: 201,
			facadeActionResponse: &dto.CreateActionResponse{
				ID: 1,
			},
			facadeActionError: nil,
			expectedError:     nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ActionControllerMock.On("CreateAction", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeActionResponse,
					tc.facadeActionError,
				)

			data, err := json.Marshal(tc.actionRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.CreateActionResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeActionResponse.ID, getParsingJson.ID)
				return
			}

			errorResponse := &dto.Error{}
			err = json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}

}

func TestActionController_GetActionByID(t *testing.T) {
	testCases := []struct {
		name                   string
		actionID               uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeActionResponse   *dto.GetActionByIDResponse
		facadeActionError      error
		expectedError          error
	}{
		{
			name:                   "OK",
			actionID:               uint64(1),
			urlRequested:           "/atados/v1/action/1",
			expectedHttpStatusCode: 200,
			facadeActionResponse: &dto.GetActionByIDResponse{
				ID:           1,
				Name:         "Action test",
				Institution:  "Institution fake",
				City:         "São Paulo",
				Neighborhood: "Limoeiro",
				Address:      "Rua Palmeira, 25",
				Description:  "Reuniao as 15 horas",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			facadeActionError: nil,
			expectedError:     nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ActionControllerMock.On("GetActionByID", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeActionResponse,
					tc.facadeActionError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetActionByIDResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeActionResponse.ID, getParsingJson.ID)
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}

}

func TestActionController_GetAllActions(t *testing.T) {
	testCases := []struct {
		name                   string
		urlRequested           string
		expectedHttpStatusCode int
		facadeActionResponse   *dto.GetAllActionsResponse
		facadeActionError      error
		expectedError          error
	}{
		{
			name:                   "OK",
			urlRequested:           "/atados/v1/action",
			expectedHttpStatusCode: 200,
			facadeActionResponse: &dto.GetAllActionsResponse{
				Actions: []dto.Action{
					{
						ID:           uint64(1),
						Name:         "Action test",
						Institution:  "Institution fake",
						City:         "São Paulo",
						Neighborhood: "Limoeiro",
						Address:      "Rua Palmeira, 25",
						Description:  "Reuniao as 15 horas",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
					{
						ID:           uint64(2),
						Name:         "Action test 2",
						Institution:  "Institution fake 2",
						City:         "Rio de Janeiro",
						Neighborhood: "Bragança",
						Address:      "Rua das flores, 25",
						Description:  "Reuniao as 17 horas",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				},
				Pagination: dto.ActionPagination{
					Limit:  10,
					Offset: 0,
					Total:  2,
				},
			},
			facadeActionError: nil,
			expectedError:     nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ActionControllerMock.On("GetAllActions", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeActionResponse,
					tc.facadeActionError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetAllActionsResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeActionResponse.Actions[0].ID, getParsingJson.Actions[0].ID)
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}

func TestActionController_UpdateAction(t *testing.T) {
	testCases := []struct {
		name                   string
		actionUpdateRequest    *dto.UpdateActionRequest
		actionID               uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeActionError      error
		expectedError          error
	}{
		{
			name: "OK",
			actionUpdateRequest: &dto.UpdateActionRequest{
				Name:         "Action test updated",
				Institution:  "Institution fake updated",
				City:         "Rio de Janeiro",
				Neighborhood: "Bragança",
				Address:      "Rua das flores, 25",
				Description:  "Reuniao as 17 horas",
			},
			actionID:               1,
			urlRequested:           "/atados/v1/action/1",
			expectedHttpStatusCode: 200,
			facadeActionError:      nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ActionControllerMock.On("UpdateAction", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeActionError,
				)

			data, err := json.Marshal(tc.actionUpdateRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Action updated successfully\"\n")
				return
			}

			errorResponse := &dto.Error{}
			err = json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}

func TestActionController_DeleteAction(t *testing.T) {
	testCases := []struct {
		name                   string
		actionID               uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeActionError      error
		expectedError          error
	}{
		{
			name:                   "OK",
			actionID:               1,
			urlRequested:           "/atados/v1/action/1",
			expectedHttpStatusCode: 200,
			facadeActionError:      nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ActionControllerMock.On("DeleteAction", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeActionError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Action deleted successfully\"\n")
				return
			}

			errorResponse := &dto.Error{}
			err := json.Unmarshal([]byte(responseString), errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}
}
