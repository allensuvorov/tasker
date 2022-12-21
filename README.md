# tasker - job scheduler

Write HTTP server for a service that would make http requests to 3rd-party services.

### Work algorithm.
The client sends a task to the service to perform an http request to a 3rd-party services. The task is described in json format, the generated task id is returned in response and its execution starts in the background.

### Code development progress:

Completed:
- increment1 - Project diagram, code with interfaces.
- increment2 - Handlers with saving to DB.
- increment3 - Task agent with goroutines.

Planed:
- increment4 - Docker, Tests and other cool stuff.
- move DB to Docker. How-to https://habr.com/ru/post/578744/
- second container for storage methods tests