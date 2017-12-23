package security

import (
	"testing"


	"net/http"
	"net/http/httptest"
)

//func TestMain(m *testing.M){
//	logger.Init()
//	logger.Info.Println("Logger ready")
//
//	os.Exit(m.Run())
//}

func mockHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do nothing
	})
}


func TestNoSession(t *testing.T) {

	secureHandler := Secure(mockHandler())
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	secureHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Secutity return %v", w.Code)
	}

}

func TestAuthentication(t *testing.T) {

	secureHandler := Secure(mockHandler())
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Add("Authorization", "Basic Y2hhcmxlc0BjYmlsbGV0dGUuY29tOnRvdG8=")

	w := httptest.NewRecorder()

	secureHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expexted 200 received  %v ", w.Code)
	}

	cookieString := w.HeaderMap.Get("Set-Cookie")
	expected := "session=A session id; Path=/"

	if cookieString != expected {

		t.Errorf("Expexted %v received %v ", expected, w.Code)
	}

}
