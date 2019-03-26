package home_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/b-b3rn4rd/omni-task-two/resouces/home"
	"github.com/emicklei/go-restful"
	"gotest.tools/assert"
)

type r struct {
	Message string `json:"message"`
}

func TestHome(t *testing.T) {
	t.Run("TestHomeIsWelcoming", func(t *testing.T) {
		httpWriter := httptest.NewRecorder()
		bodyReader := strings.NewReader(`{}`)
		httpRequest, _ := http.NewRequest("GET", "/", bodyReader)
		httpRequest.Header.Set("Content-Type", "application/json")
		request := &restful.Request{Request: httpRequest}
		response := restful.NewResponse(httpWriter)
		response.SetRequestAccepts("application/json")
		home.Home(request, response)

		d := json.NewDecoder(httpWriter.Body)
		actualResponseBody := &r{}
		expectedResponseBody := &r{"Hello World"}

		d.Decode(&actualResponseBody)

		assert.Equal(t, "application/json", httpWriter.Header().Get("Content-Type"))
		assert.DeepEqual(t, expectedResponseBody, actualResponseBody)
	})
}
