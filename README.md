# gardenly
vegetable gardening app

https://capstone-project-production-4648.up.railway.app/

## Summary
The projected outlined here will be to create an api (gardenly) that, given user information (zipcode), returns information about what vegetables can be planted at the time of the request got their given growing zone. The project will involve creating a DB of vegetable types and their growing times in relation to the forst dates for zip codes in most of the United States, and an api that can take in user information and return a list of vegetables.

I will demonstrate the following:

1.Building a REST API in Go.
2.Creating and interacting with a postgres DataBase.
3.Ability to perform CRUD actions on resources through an API.
4.Acting as a client to api's in order to gather location and weather data by zipcode.
5.Testing of HTTP handlers.
6.Testing of HTTP Client.
7.Use html templates to display response results in browser.

## User Stories

### As a user I would like to fetch information about what vegetables I am able to plant at this time.

<img width="650" alt="Screen Shot 2023-01-31 at 2 17 18 PM" src="https://user-images.githubusercontent.com/105764001/215873646-6a88c237-28ca-4ec7-a893-3678a0d9627d.png">

Example:

    curl https://capstone-project-production-4648.up.railway.app/
    
    <img width="650" alt="image" src="https://user-images.githubusercontent.com/105764001/215872982-88af2239-a09a-4e81-a0e8-415ac210f545.png">

    
    There are 30 days until the last frost. 

    You can plant:
    carrot
    beats
    turnips
    lettuce
    ...
    
If the postal code is not found then the user will revieve an error message.

    curl http://localhost:8080/gardenly/34647589
    
    Error: "xyzip" is not found, please provide a valid zipcode.
    
    


    
    
       

