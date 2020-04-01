# REST API Management for Employee's

**GET   /manage/employee**   
To get the list of all employees

```
curl -sX GET http:localhost:8080/manage/employee | jq
```

**GET manage/employee/{name/id}**    
To get details of a particular employee either by name or ID

```
curl -sX GET http:localhost:8080/manage/employee/Bob | jq
curl -sX GET http:localhost:8080/manage/employee/1 | jq
```

**GET /manage/employees/{ids}**   
To get details of multiple employees by id's
```
curl -sX GET http:localhost:8080/manage/employee/1,2,4 | jq
```

**POST /manage/employee**   
To add an employee
```
curl -sX POST http://localhost:8080/manage/employee -d '{"name": "Bob","age":32,"gender":"M"}' | jq
```

**PUT /manage/employee/{id}**   
To update an employee details
```
curl -sX PUT http://localhost:8080/manage/employee/{id} -d '{"name": "Bob2","age":34}' | jq
```
**DELETE /manage/employee/{id}**   
To delete an employee 
```
curl -sX DELETE http://localhost:8080/manage/employee/{id} | jq
```