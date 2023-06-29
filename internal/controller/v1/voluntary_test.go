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

func TestVoluntaryController_CreateVoluntary(t *testing.T) {

	testCases := []struct {
		name                    string
		voluntaryRequest        *dto.CreateVoluntaryRequest
		urlRequested            string
		expectedHttpStatusCode  int
		facadeVoluntaryResponse *dto.CreateVoluntaryVO
		facadeVoluntaryError    error
		expectedError           error
	}{
		{
			name: "OK",
			voluntaryRequest: &dto.CreateVoluntaryRequest{
				FirstName:    "Miguel",
				LastName:     "Ferreira",
				Neighborhood: "Centro",
				City:         "Rio de Janeiro",
			},
			urlRequested:           "/atados/v1/voluntary",
			expectedHttpStatusCode: 201,
			facadeVoluntaryResponse: &dto.CreateVoluntaryVO{
				ID: 1,
			},
			facadeVoluntaryError: nil,
			expectedError:        nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.VoluntaryControllerMock.On("CreateVoluntary", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeVoluntaryResponse,
					tc.facadeVoluntaryError,
				)

			data, err := json.Marshal(tc.voluntaryRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.CreateVoluntaryVO{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeVoluntaryResponse.ID, getParsingJson.ID)
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

func TestVoluntaryController_GetVoluntaryByID(t *testing.T) {
	testCases := []struct {
		name                    string
		voluntaryID             uint64
		urlRequested            string
		expectedHttpStatusCode  int
		facadeVoluntaryResponse *dto.GetVoluntaryByIDResponse
		facadeVoluntaryError    error
		expectedError           error
	}{
		{
			name:                   "OK",
			voluntaryID:            uint64(1),
			urlRequested:           "/atados/v1/voluntary/1",
			expectedHttpStatusCode: 200,
			facadeVoluntaryResponse: &dto.GetVoluntaryByIDResponse{
				ID:           1,
				FirstName:    "Miguel",
				LastName:     "Ferreira",
				Neighborhood: "Centro",
				City:         "Rio de Janeiro",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			facadeVoluntaryError: nil,
			expectedError:        nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.VoluntaryControllerMock.On("GetVoluntaryByID", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeVoluntaryResponse,
					tc.facadeVoluntaryError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetVoluntaryByIDResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeVoluntaryResponse.ID, getParsingJson.ID)
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

func TestVoluntaryController_GetAllVoluntarys(t *testing.T) {
	testCases := []struct {
		name                    string
		urlRequested            string
		expectedHttpStatusCode  int
		facadeVoluntaryResponse *dto.GetAllVoluntariesResponse
		facadeVoluntaryError    error
		expectedError           error
	}{
		{
			name:                   "OK",
			urlRequested:           "/atados/v1/voluntary",
			expectedHttpStatusCode: 200,
			facadeVoluntaryResponse: &dto.GetAllVoluntariesResponse{
				Voluntaries: []dto.Voluntary{
					{
						ID:           uint64(1),
						FirstName:    "Miguel",
						LastName:     "Ferreira",
						Neighborhood: "Centro",
						City:         "Rio de Janeiro",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
					{
						ID:           uint64(2),
						FirstName:    "Caio",
						LastName:     "Matos",
						Neighborhood: "Centro",
						City:         "Rio de Janeiro",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				},
				Pagination: dto.VoluntaryPagination{
					Limit:  10,
					Offset: 0,
					Total:  2,
				},
			},
			facadeVoluntaryError: nil,
			expectedError:        nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.VoluntaryControllerMock.On("GetAllVoluntaries", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeVoluntaryResponse,
					tc.facadeVoluntaryError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsingJson := &dto.GetAllVoluntariesResponse{}
				err := json.Unmarshal([]byte(responseString), getParsingJson)
				assert.NoError(t, err)
				assert.Equal(t, tc.facadeVoluntaryResponse.Voluntaries[0].ID, getParsingJson.Voluntaries[0].ID)
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

func TestVoluntaryController_UpdateVoluntary(t *testing.T) {
	testCases := []struct {
		name                   string
		voluntaryUpdateRequest *dto.UpdateVoluntaryRequest
		voluntaryID            uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeVoluntaryError   error
		expectedError          error
	}{
		{
			name: "OK",
			voluntaryUpdateRequest: &dto.UpdateVoluntaryRequest{
				FirstName:    "Caio",
				LastName:     "Matos",
				Neighborhood: "Centro",
				City:         "Rio de Janeiro",
			},
			voluntaryID:            1,
			urlRequested:           "/atados/v1/voluntary/1",
			expectedHttpStatusCode: 200,
			facadeVoluntaryError:   nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.VoluntaryControllerMock.On("UpdateVoluntary", mock.AnythingOfType("*context.emptyCtx"), mock.Anything, mock.Anything).
				Return(
					tc.facadeVoluntaryError,
				)

			data, err := json.Marshal(tc.voluntaryUpdateRequest)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Voluntary updated successfully\"\n")
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

func TestVoluntaryController_DeleteVoluntary(t *testing.T) {
	testCases := []struct {
		name                   string
		voluntaryID            uint64
		urlRequested           string
		expectedHttpStatusCode int
		facadeVoluntaryError   error
		expectedError          error
	}{
		{
			name:                   "OK",
			voluntaryID:            1,
			urlRequested:           "/atados/v1/voluntary/1",
			expectedHttpStatusCode: 200,
			facadeVoluntaryError:   nil,
			expectedError:          nil,
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.VoluntaryControllerMock.On("DeleteVoluntary", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).
				Return(
					tc.facadeVoluntaryError,
				)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.urlRequested, nil)
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				assert.Equal(t, responseString, "\"Voluntary deleted successfully\"\n")
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
