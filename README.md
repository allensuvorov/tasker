# tasker - job scheduler

Write HTTP server for a service that would make http requests to 3rd-party services.

### Work algorithm.
The client sends a task to the service to perform an http request to a 3rd-party services. The task is described in json format, the generated task id is returned in response and its execution starts in the background.

Code development progress:

- increment1 - Project diagram, code with interfaces.
- increment2 - Handlers with saving to DB and tests.
- increment3 - Task agents with goroutines and tests.
- Add Limit db connection pool
- increment4 - Docker, and other cool stuff.

Plan:
1. First - do what I already know 
2. Next - do what needs study:
- mode DB to Docker
https://habr.com/ru/post/578744/
- second container for storage methods tests