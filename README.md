# HutPlate

HutPlate is a Library over your standard net/http library which makes your life easier.
 
 It handles Authentication, Session, HTTP Error Handling, Flash Messages and Redirection. All these 
 things are a breeze to do with HutPlate which usually takes a lot of your time.
 
 Some examples:
```go
// Use any router, also using hutplate.Handler is optional
router.Handle("/", hutplate.Handler(myHandler))

func CreatePost(hp hutplate.Http) interface{} {
	if ! hp.Auth.Check() {
	    // Will redirect and set flash message in session
	    return hp.Response.Redirect("/login").With("notice", "You have to login first!")
	}
	
	// hutplate.Http extends http.Request, so everything is still there
	if err := createPost(hp.FormValue("content")); err =! nil {
	    // Will show a generic error message to user with 500 status.
	    //  You can also customize the behaviour
	    return err
	}
	
	// The http.ResponseWriter is also there
	hp.Response.Write([]byte("Success!"))
}
```
 
 