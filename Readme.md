# TODOS
* General
	- Convert array to json functions for working for pointer so they will not take parameter and return value they just change object within.
	- Add cousines like italian, mexican etc.
	- Add prepared diets like katot , mediterrian , etc.
	- Optimize Recommendation steps. Get only non-prohibited and related recipes from db.For now it is getting all recipes and filtering them.
* Api
	- DRY for session code while getting session values.(Everytime calling session, err := app.SessionStore.Get(r, "session") ...)
* User
	- mail or phone based authentication
	- role based access
	- user role control function in user object
* Shopping
* Recommendation
	- Add prepared diet recommendations
	- Add balance based diet(sufficent and balanced diet)
	- User pointer in funcitons for getting uint array
	- Handle blank tags etc. situations. 
* Recipe
	- Tag collector from materials(Map creator from recipe lists.)
	- Set Diet Level for Ingrediens to make recommendation.It will take a look at materials and set a diet level .(1:Vegan,2Vegaterian ...)
	- Recipe should understand recipes diet level from materials. For now it will be done manually. Same about tags. Tags will be add to materials instead of recipe