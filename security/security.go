package security
import (
	"net/http"
	"fmt"
	"dix975.com/basic/logger"
)


type Handle struct {
	ControllerFunc func(http.ResponseWriter, *http.Request)
	ForwardToLogin bool
}

func (handle Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

		cookie, error := r.Cookie("session")

		//todo validate session cookie
		logger.Debug.Println("Cookie : ", cookie)

		if (error == http.ErrNoCookie) {

			logger.Debug.Println("Cookie not found")

			user, _, ok := r.BasicAuth()
			if (ok) {
				logger.Debug.Printf("Will authenticate user [%v]", user)
				//todo query user in database and validate password
				//todo generate real session id
				SetSessionCookie(w)

			}else {


				if handle.ForwardToLogin {

					redirectTo := fmt.Sprintf("/app/loginView.do?successUrl=%v", r.RequestURI)
					r.Method = "GET"
					http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
				}else{

					logger.Info.Println("No session and no Authentication header.")
					http.Error(w, http.StatusText( http.StatusUnauthorized),  http.StatusUnauthorized)
				}

				return
			}
		}

		//logger.Trace.Println("Security pass")

		handle.ControllerFunc(w, r)

}

func SetSessionCookie(w http.ResponseWriter){
	http.SetCookie(w, &http.Cookie{Name : "session", Value: "A session id", Path:"/"})
}
