# Tasker - job scheduler

Description: Write HTTP server for a service that would make http requests to 3rd-party services.

### To test/run the app:
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

### Basic algorithm.
The client sends a task to the service to perform an http request to a 3rd-party services. The task is described in json format, the generated task id is returned in response and its execution starts in the background.

### Code development progress:

Completed:
- increment1 - Project diagram, code with interfaces.
- increment2 - Handlers with saving to DB with tests.
- increment3 - Task agent with goroutines.