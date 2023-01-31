# Gardenly
vegetable gardening app

https://capstone-project-production-4648.up.railway.app/

## Summary
The projected outlined here will be to create an web application that, given user information (zipcode), returns information about what vegetables can be planted at the time of the request depending on their location. The project will involve creating a DB of vegetable types and their days to maturity, and creating logic to determine the vegetable types that will be returned to the user. The application is exposed to the internet through an API that allows interaction with application functions and a postgres database.

I will demonstrate the following:

1.Building a REST API in Go.

2.Creating and interacting with a postgres DataBase.

3.Ability to perform CRUD actions on resources through an API.

4.Acting as a client to APIs in order to gather location and weather data by zipcode.

5.Testing of HTTP Clients.

6.Use html templates to display response results in browser.

7.Utilizing railway cloud to host Gardenly.

## User Stories

# As a user I would like to fetch the types of vegetables I am able to plant at this time.

https://capstone-project-production-4648.up.railway.app/

<img width="650" alt="Screen Shot 2023-01-31 at 2 17 18 PM" src="https://user-images.githubusercontent.com/105764001/215873646-6a88c237-28ca-4ec7-a893-3678a0d9627d.png">

# By submitting my zip code I am redirected to this page:

https://capstone-project-production-4648.up.railway.app/93301

<img width="782" alt="Screen Shot 2023-01-31 at 2 23 56 PM" src="https://user-images.githubusercontent.com/105764001/215874341-07e01fef-8a3c-4ee4-9b06-87bfa212e447.png">

# It is possible that my zip code is not supported by gardenly. 

<img width="645" alt="Screen Shot 2023-01-31 at 2 27 27 PM" src="https://user-images.githubusercontent.com/105764001/215875192-855acdcc-b00e-423a-8a2b-88d72cabd585.png">


# A postgres database acts as a resource for Gardenly and and can be interacted with using CRUD functions.

curl https://capstone-project-production-4648.up.railway.app/vegetables/all

<img width="422" alt="Screen Shot 2023-01-31 at 2 43 45 PM" src="https://user-images.githubusercontent.com/105764001/215878634-eeb19958-80fd-4f37-b6a2-767c2b7f305a.png">



    


    
    
       

