# TODOS
* General
	- Add Dont read from database if object soft deleted feature
	- Convert array to json functions for working for pointer so they will not take parameter and return value they just change object within.
* Api
	- DRY for session code while getting session values.(Everytime calling session, err := app.SessionStore.Get(r, "session") ...)
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
* Shopping
	- Add list shoppinglist by date interval
* Recommendation
	- Add prepared diet recommendations
	- Add balance based diet(sufficent and balanced diet)