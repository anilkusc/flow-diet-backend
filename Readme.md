# TODOS
* General
	- Convert array to json functions for working for pointer so they will not take parameter and return value they just change object within.
	- Add prepared diets like katot , mediterrian , etc.
	- Optimize Recommendation steps. Get only non-prohibited and related recipes from db.For now it is getting all recipes and filtering them.
	- Make everyting(port etc) variable and take that from environment variables.
	- Versioning information should be added to logs and also URL path
	- Control for sent json objects
	- Add language option to all of the returned backend object
	- Add random recpie feature. It will be give random recipe which users mostly like.
* Api
	- Session should be define and get variables on only middlewares and migrate to the next handler.
* Search
	- Add upparcase and lowercase letter searching. For now it is is case sensitive. It should be case insensitive.
	- Add tag searching if editor wants to add or user search
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
* Shopping
	- Create Automatic Shopping lists from weekly recipes.
* Recommendation
	- Add prepared diet recommendations
	- Add balance based diet(sufficent and balanced diet)
	- Handle blank tags etc. situations. 
	- Add Calori Based Pointing to recommendation.
	- Add Favorite Based Recommendation. It will be recommend favorite based recipes
	- Enable query recommendation for date. Take recommendations only for specified time interval.It can be maxium 1 week.
	- Add  date based recommendations feature. It means look at the season and give recommendation for season.
* Recipe
	- Tag collector from materials(Map creator from recipe lists.)
	- Set Diet Level for Ingrediens to make recommendation.It will take a look at materials and set a diet level .(1:Vegan,2Vegaterian ...)
	- Recipe should understand recipes diet level from materials. For now it will be done manually. Same about tags. Tags will be add to materials instead of recipe
	- Add  date based recommendations feature. It means look at the season and give recommendation for season.
* Calendar
	- Add meals for recipes beside of the date.
* Observability
	- pprof
	- health probes