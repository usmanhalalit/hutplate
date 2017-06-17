# HutPlate
[![Build Status](https://travis-ci.org/usmanhalalit/hutplate.svg?branch=master)](https://travis-ci.org/usmanhalalit/hutplate)

HutPlate is a Library over your standard net/http library which makes your life easier.
 
 It handles Authentication, Session, HTTP Error Handling, Flash Messages and Redirection. All these 
 things are a breeze to do with HutPlate which usually takes a lot of your time.
 
 Some examples:
```go
// Use any router, also using hutplate.Handler is optional
router.Handle("/", hutplate.Handler(CreatePost))

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

### Configuration

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

**Another highly recommended configuration** is that you set a session secret key:
```go
hutplate.Config.SessionSecretKey = "a_random_secret_key"
```

### Creating a HutPlate Instance

HutPlate comes with an HTTP handler, which gives you some extra power. 
But it's optional to use. If you want to keep using the default handler 
then you just call `hutplate.NewHttp` and give it your `Request` and `ResponseWriter`. 
Here's how:
```go
func MyHandler(w http.ResponseWriter, r *http.Request,) {
	hp := hutplate.NewHttp(w, r)
}
```

If you use the HutPlate handler, then it even easier:
```go
router.Handle("/", hutplate.Handler(CreatePost))

func CreatePost(hp hutplate.Http) interface{} {
	// hutplate.Http extends http.Request, so everything is still there
	// like hp.FormValue("content")

	// The http.ResponseWriter is also there
	hp.Response.Write([]byte("Success!"))
}
```


## Authentication

### Login
```go
success, _ := hp.Auth.Login(email, password)
```

It will let you know if user logging in succeeded or not by returning true or false
 and if there is an error (in rare case) it'll be in the second return value.
 
Login requires you to store bcrypt hashed password when you register/save your user.
Here is an example of how you hash you password using bcrypt

```go
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
```

### Check if Logged In
```go
hp.Auth.Check()
```
Will return true or false.

### Get Logged in User ID
```go
hp.Auth.UserId()
```

### Get the Logged in User

It gives you the whole user object, not just id. To use this feature,
 you need to set a configuration to let HutPlate know how to 
 get the user from your database/datastore.
```go
user, err = hp.Auth.User()
```

**Configuration**
```go
hutplate.Config.GetUserWithId = func(userId interface{}) interface {} {
    // You will receive the userId just return the user with corresponding id  
}
```

A GORM example:
```go
hutplate.Config.GetUserWithId = func(userId interface{}) interface {} {
    user := models.User{}
    if userId == nil {
        return user
    }
    db.Orm.Find(&user, userId)
    return user
}
```

### Logout
```go
hp.Auth.Logout()
```

## HutPlate HTTP Handler

HutPlate also comes with an HTTP handler, using it is optional but you get so much
power by using it. 

 - It automatically creates an HutPlate instance for you.
 - You can return an error from the handler, so you do much less `if err != nil`.
 - You can configure how to handle all errors or group of errors. There is a sensible default.
 - You can return a `Redirect` or a plain string response (text, HTML, etc.).

```go
// Use any router
router.Handle("/", hutplate.Handler(CreatePost))

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

Configure the error handler

```go
hutplate.Config.ErrorHandler = func(err error, hut hutplate.Http) {
        // TODO log the error
		http.Error(hut.Response, err.Error(), 404)
}
```

## Redirect

```go
// Plain redirect
hp.Response.Redirect("/admin")

// Redirect with a flash message
hp.Response.Redirect("/login").With("error", "Please do login!")

// Redirect with a different status code
hp.Response.Redirect("/login", 301)
```
 
## Session

Set a session value and it will persist.
```go
err := hp.Session.Set("test_key", "test_value")
```

Get the session value
```go
err := hp.Session.Get("test_key")
```

### Flash Messages
```go
err = hut.Session.SetFlash("test_key", "value")
```

Getting a flash message is same as getting normal session value:
```go
err := hp.Session.Get("test_key")
```

[As mentioned above](#Redirect) you can also easily set a flash message while you redirect.  


### Session Config

#### Storage Path

By default Session uses file system storage. You can optionally 
configure the storage directory by setting:
```go
hutplate.Config.SessionDirectory = "path/to/dir"
```
By default it uses operating system's temp directory.

#### Secret Key
It is recommended that you set a session secret key by setting:
```go
hutplate.Config.SessionSecretKey = "your_key"
````

### Custom Session Store

As noted above, by default Session uses file system storage. But
you can use an entire different session storage system. HutPlate
uses [Gorilla Sessions](https://github.com/gorilla/sessions).
So you can use any session storage supported by Gorilla Sessions, 
i.e. the store that implements Gorilla's `sessions.Store` interface.

[Here is a list](https://github.com/gorilla/sessions#store-implementations)
 of session stores supported by Gorilla Sessions.

To use your own session store, use the SessionStore config,
for example:
```
import "github.com/gorilla/sessions"
hutplate.Config.SessionStore = sessions.NewCookieStore(...)
```

### Clear Context

Important Note: If you aren't using gorilla/mux router, you need to wrap 
your handlers with context.ClearHandler as or else you will leak memory! 
An easy way to do this is to wrap the top-level mux when calling 
http.ListenAndServe:
```go
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
```
The ClearHandler function is provided by the gorilla/context package.

More examples are available 
[on the Gorilla website](http://www.gorillatoolkit.org/pkg/sessions).

## List of Config

```go
hutplate.Config.GetUserWithId = func(userId interface{}) interface {} {

}

hutplate.Config.GetUserWithCred = func(credential interface{}) (interface{}, string) {
    
}

hutplate.Config.ErrorHandler = func(err error, hp hutplate.Http) {

}

hutplate.Config.SessionSecretKey = "a_secret_key"
hutplate.Config.SessionDirectory = "path/to/dir"
```


___
&copy; 2017 [Muhammad Usman](http://usman.it/). Licensed under MIT license.