#!/bin/bash
#signup
curl -X 'POST' \
  'http://localhost:8080/user/signup' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": "myadress 123121",
  "age": 25,
  "diet_level": 1,
  "dislikes": [
    "onion"
  ],
  "email": "test@test.com",
  "favorite_cousines": [
    "italian"
  ],
  "favorite_recipes": [
    0
  ],
  "gender": "male",
  "height": 170,
  "likes": [
    "kebap",
    "pizza"
  ],
  "name": "test user",
  "password": "admin",
  "phone": "+905355353535",
  "preferred_meals": [
    "breakfast"
  ],
  "prohibits": [
    "sugar"
  ],
  "role": "admin",
  "username": "admin",
  "wants": "gain",
  "weight": 70
}'
#signin
curl -X 'POST' \
  'http://localhost:8080/user/signin' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "admin",
  "username": "admin"
}'
#create recipe 1
curl -X 'POST' \
  'http://localhost:8080/recipes/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "appropriate_meals": [
    "noon",
    "afternoon",
    "evening",
  ],
  "calori": 552,
  "cooking_time_minute": 30,
  "cousines": [
    "japan","asia"
  ],
  "for_how_many_people": 2,
  "ingredients": [
    {
      "isexist": false,
      "isoptional": true,
      "material": {
        "material_photo_urls": [
          "S3URL1",
          "S3URL2"
        ],
        "name": "banana"
      },
      "measurement": {
        "quantity": "gram",
        "size": 2
      }
    }
  ],
  "photo_urls": [
    "S3URL1",
    "S3URL2"
  ],
  "preperation": "bla bla bla",
  "preperation_time": 15,
  "recipe_diet_level": 1,
  "tags": [
    "vegan"
  ],
  "title": "Sushi With Wassabi",
  "video_urls": [
    "S3URL1",
    "S3URL2"
  ]
}'
#create recipe 2
curl -X 'POST' \
  'http://localhost:8080/recipes/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "appropriate_meals": [
    "evening"
  ],
  "calori": 507,
  "cooking_time_minute": 15,
  "cousines": [
    "turkey","mediterrian"
  ],
  "for_how_many_people": 2,
  "ingredients": [
    {
      "isexist": false,
      "isoptional": true,
      "material": {
        "material_photo_urls": [
          "S3URL1",
          "S3URL2"
        ],
        "name": "banana"
      },
      "measurement": {
        "quantity": "gram",
        "size": 2
      }
    }
  ],
  "photo_urls": [
    "S3URL1",
    "S3URL2"
  ],
  "preperation": "bla blaasda bla",
  "preperation_time": 25,
  "recipe_diet_level": 3,
  "tags": [
    "kebap", "meat" , "owen"
  ],
  "title": "Adana Kebap",
  "video_urls": [
    "S3URL1",
    "S3URL2"
  ]
}'

#create recipe 3
curl -X 'POST' \
  'http://localhost:8080/recipes/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "appropriate_meals": [
   "noon","evening"
  ],
  "calori": 700,
  "cooking_time_minute": 25,
  "cousines": [
    "italian","mediterrian"
  ],
  "for_how_many_people": 2,
  "ingredients": [
    {
      "isexist": false,
      "isoptional": true,
      "material": {
        "material_photo_urls": [
          "S3URL1",
          "S3URL2"
        ],
        "name": "banana"
      },
      "measurement": {
        "quantity": "gram",
        "size": 2
      }
    }
  ],
  "photo_urls": [
    "S3URL1",
    "S3URL2"
  ],
  "preperation": "bla blaasda bla",
  "preperation_time": 25,
  "recipe_diet_level": 3,
  "tags": [
    "pizza", "fastfood" , "owen"
  ],
  "title": "Italian Pizza",
  "video_urls": [
    "S3URL1",
    "S3URL2"
  ]
}'