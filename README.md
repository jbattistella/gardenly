# GardenAppProject
vegetable gardening app

## Summary
The projected outlined here will be to create an api (gardenapp) that, given user information (zipcode), returns information about what vegetables can be planted at the time of the request got their given growing zone. The project will involve creating a DB of vegetable types and their growing times in relation to the hardiness zones in the United States, and an api that can take in user information and return a list of vegetables.

I will demonstrate the following:

1.Building a REST API in Go.
2.Creating and interacting with a postgres DataBase.
3.Acting as a client to api's in order to gather hardiness zone by zipcode.
4.Testing of HTTP handlers
5.Testing of Http Client

## User Stories

### As a user I would like to fetch information about what vegetables I am able to plant at this time.

Example:

    curl http://localhost:8080/api/v1/users -d { "name" : "joseph", "zipcode" :35235}

    You can plant:
    carrot
    beats
    turnips
    lettuce
    ...
    
If the postal code is not found then the user will revieve an error message.

    curl http://localhost:8080/api/v1/users -d { "name" : "joseph", "zipcode" :35235}
    
    Error: "xyzip" is not found, please provide a valid zipcode.
    
    
Notes:

I would like to add functionality to this project as i complete different parts. 
Added functionality:
1.App is hosted remotely
2.Send an email to a user containing the growing information.
3.Store user information and send emails on a regular basis that contain more personalized information.

    
    
       

