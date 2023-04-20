# Tasker - job scheduler

### Project Description
The client sends a task to the service to perform an http request to a 3rd-party services. The task is described in json format, the generated task id is returned in response and its execution starts in the background.

### Project Architecture
Two micro-services. One - receives tasks, saves to DB, returns statuses. The other - polls DB to get new tasks, concurrently runs them - makes HTTP requests to 3rd-party services and updates statuses in the DB.

* [Project flowchart](https://github.com/allensuvorov/tasker/blob/main/docs/diagram_microservices.pdf)
* [Project spec](https://github.com/allensuvorov/tasker/blob/main/docs/task_spec.pdf)


### Configuration to test/run the app:
1. First, set your postgres DB connection string (DSN) in these files:
   - location: internal/server/storage/db.go, method: NewTaskStorage()
   - location: internal/agent/storage/storage.go, method: NewStorage()

2. To test the agent (task scheduler):
   - Fill up the DB with test data by running this test: path: internal/server/storage/db_test.go, test name: "TestTaskStorage_Create30Tasks_RealDB"
   - Run the agent: cmd/agent/agent.go

3. To test/run the server:
   - POST request - run this test: path: internal/server/storage/db_test.go, test name: "TestTaskStorage_CreateTask_RealDB"
   - GET request - run this test: path: internal/server/storage/db_test.go, test name: "TestTaskStorage_GetTaskStatus_RealDB"
   - If you would like to start the server, run: cmd/main/agent.go
