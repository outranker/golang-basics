20. how web application works

- nothing new i guess?

21. making a server

- we use net/http pkg to create and handle requests
- http.HandleFunc(a1, a2). a1 is route name, a2 is handler func. a2 takes 2 args, first one is response writer second one is request
- we use http.ListenAndServer() to serve our app

22. functions and handlers

- express like server and its handlers, nothing special so far
