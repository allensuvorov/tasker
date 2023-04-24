# Tasker - job scheduler

### Project Description
The client sends a task to the service to perform an http request to a 3rd-party services. The task is described in json format, the generated task id is returned in response and its execution starts in the background.

### Project Architecture
Two microservices. First microservice receives tasks, saves them to DB, returns statuses. Second microservice polls DB to get new tasks, concurrently runs them - makes HTTP requests to 3rd-party services and updates statuses in the DB.

- [Project flowchart](https://github.com/allensuvorov/tasker/blob/main/docs/diagram_microservices.pdf)
- [Project spec](https://github.com/allensuvorov/tasker/blob/main/docs/task_spec.pdf)


### Configuration
Set your postgres DB connection string (DSN) in these files:
   - location: internal/server/storage/db.go, method: NewTaskStorage()
   - location: internal/agent/storage/storage.go, method: NewStorage()

### To test/run the app
To run and test the agent (job scheduler):
   - NOTE: agent runs independently. No need to start the server.
   - First, fill up the DB with test data by running this test: 
      - path: internal/server/storage/db_test.go
      - test name: "TestTaskStorage_Create30Tasks_RealDB"
   - Run the agent: cmd/agent/agent.go

To run and test the server:
   - To start the server run: cmd/main/main.go
   - For a POST request run this test: 
      - path "internal/server/storage/db_test.go"
      - test name: "TestTaskStorage_CreateTask_RealDB"
   - For a GET request run this test: 
      - path: internal/server/storage/db_test.go, 
      - test name: "TestTaskStorage_GetTaskStatus_RealDB"
