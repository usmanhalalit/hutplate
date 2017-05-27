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

## Setup

The only mandatory configuration for HutPlate is that you let it know where to find your user 
(any database/datastore is fine). 
It will give you email or whatever you log your user in with, 
you just return that user's id and hashed password. 
For example, a GORM example would look like this:
```go
hutplate.Config.GetUserWithCred = func(credential interface{}) (interface{}, string) {
    // credential will be email, username or whatever you log in with
    user := models.User{}
    db.Orm.Find(&user, "email='" + credential.(string) + "'")
    
    return user.ID, user.Password
}
```

**Another highly recommended config** is that you set a session secret key:
```go
hutplate.Config.SessionSecretKey = "a_random_secret_key"
```

## Authentication

```go
success, err := hp.Auth.Login(email, password)
```
 
## Session