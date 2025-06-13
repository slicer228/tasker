# Summary
This small application is designed to run your tasks remotely.
# Customize
You can customize configs in ./config folder.
If you want to customize task core, go to ./internal/trasport/rest/task/createTask/createTask.go and seek for tasker.CreateTask() call.
# Launch
## Method #1(Using docker)
Clone repository to your server
```bash
git clone https://github.com/slicer228/tasker
cd tasker
```
Then, build docker image
```bash
docker build -t tasker .
```
Lauch your app, you can customize config file path
```bash
docker run -p 127.0.0.1:8080:8080 -e CONFIG_PATH=./config/prod.yaml tasker
```
Now you can access your http server on 127.0.0.1:8080
## Method #2(run manual)
Firstly, you should install go 1.24 on your server
Clone repository to your server
```bash
git clone https://github.com/slicer228/tasker
cd tasker
```
Windows
```cmd
go build -o ./build/executable/app.exe ./cmd/run.go
```
Then, you need to set env var CONFIG_PATH before launch
```cmd
set CONFIG_PATH=./config/local.yaml
start ./build/executable/app.exe
```
Linux
```bash
go build -o ./build/executable/app ./cmd/run.go
```
Then, you need to set env var CONFIG_PATH before launch
```bash
export CONFIG_PATH=./config/local.yaml
./build/executable/app
```
Now you can access your http server on 127.0.0.1:8080
# Routers(using localhost) examples
## Create task
### Method POST
### Request body
```json
{}
```
### Response body
```json
{
  "task_id": 1
}
```
### Url
127.0.0.1:8080/task
## Run task
### Method POST
### Request body
```json
{
  "task_id": 1
}
```
### Response body
```json
{}
```
### Url
127.0.0.1:8080/task/run
## Get task info
### Method GET
### Request body
```json
{}
```
### Response body
```json
{
  "status": "ready",
  "createdAt": "2025-06-10 08:25:29.21"
  "jobs": [
    {
     "startedAt": "2025-06-13 08:25:47.37",
      "completedAt": "2025-06-13 08:25:52.37",
      "timeSpent": "00:00:05.00",
      "status": "job_completed",
      "jobNumber": 1,
      "result": {
          "Value": 1
      } 
    }
  ]
}
```
### Url
127.0.0.1:8080/task?task_id=1
## Delete task
### Method DELETE
### Request body
```
{
  "task_id": 1
}
```
### Response body
```json
{}
```
### Url
127.0.0.1:8080/task
