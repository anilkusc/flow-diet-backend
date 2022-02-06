# TODOS
* General
	- Convert array to json functions for working for pointer so they will not take parameter and return value they just change object within.
	- Add prepared diets like katot , mediterrian , etc.
	- Optimize Recommendation steps. Get only non-prohibited and related recipes from db.For now it is getting all recipes and filtering them.
	- Make everyting(port etc) variable and take that from environment variables.
	- Versioning information should be added to logs and also URL path
	- Control for sent json objects
	- Add language option to all of the returned backend object
* Api
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
* Shopping
* Recommendation
	- Add prepared diet recommendations
	- Add balance based diet(sufficent and balanced diet)
	- Handle blank tags etc. situations. 
	- Add Calori Based Pointing to recommendation.
	- Add Cousine Based Pointing in Recommendation
	- Add Favorite Based Recommendation. It will be recommend favorite based recipes
* Recipe
	- Tag collector from materials(Map creator from recipe lists.)
	- Set Diet Level for Ingrediens to make recommendation.It will take a look at materials and set a diet level .(1:Vegan,2Vegaterian ...)
	- Recipe should understand recipes diet level from materials. For now it will be done manually. Same about tags. Tags will be add to materials instead of recipe
* Observability
	- pprof
	- health probes