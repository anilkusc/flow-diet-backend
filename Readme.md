# TODOS
* General
	- Convert array to json functions for working for pointer so they will not take parameter and return value they just change object within.
	- Add prepared diets like katot , mediterrian , etc.
	- Make everyting(port etc) variable and take that from environment variables.
	- Versioning information should be added to logs and also URL path
	- Control for sent json objects
	- Add language option to all of the returned backend object
	- Add random recpie feature. It will be give random recipe which users mostly like.
	- Do not return user id
	- Seperate ingredients package
	- paths are rehandled
	- PUT,PATCH,DELETE methods will be implemented
	- Return id after create 
* Api
	- Session should be define and get variables on only middlewares and migrate to the next handler.
* Search
	- Add upparcase and lowercase letter searching. For now it is is case sensitive. It should be case insensitive.
	- Add tag searching if editor wants to add or user search
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
	- Signin and signup models will be seperated from user.
	- User update will be implemented on the api
	- Json ignore for user password
* Shopping
* Recommendation
	- Add prepared diet recommendations
	- Add balance based diet(sufficent and balanced diet)
	- Handle blank tags etc. situations. 
	- Add Calori Based Pointing to recommendation.
	- Add Favorite Based Recommendation. It will be recommend favorite based recipes
	- Add  date based recommendations feature. It means look at the season and give recommendation for season.
* Recipe
	- Tag collector from materials(Map creator from recipe lists.)
	- Set Diet Level for Ingrediens to make recommendation.It will take a look at materials and set a diet level .(1:Vegan,2Vegaterian ...)
	- Recipe should understand recipes diet level from materials. For now it will be done manually. Same about tags. Tags will be add to materials instead of recipe
* Calendar
* Observability
	- pprof
	- health probes
	- add egress rule for security. If egress body is bigger than 50 Mb return error