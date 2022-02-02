# TODOS
* Api
	- DRY for session code while getting session values.(Everytime calling session, err := app.SessionStore.Get(r, "session") ...)
	- DRY for getting cookie on test codes.(Everytime calling sign in)
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
	- phone uniqueness control