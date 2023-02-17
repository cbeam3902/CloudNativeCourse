# Lab 6
This lab introduces accessing an external service through their public REST API
The API uses information from this [website](api.openweathermap.org), so make an account to follow along.

The task for this lab is to extend the information collected to include the city's pressure, humidity, and wind speed.

To execute the code, you'll need to export your API key

    export OPENWEATHERMAP_API_KEY=KEY
    
and then run the code with the city of your choice

    go run main.go Charlotte

Not sure if case matters for the location.