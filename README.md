## RideShare:

Golang implementation of a ride sharing app.

Problem Statement: [link](https://github.com/getsimplifin/simplifin_interview_questions/wiki/Ride-Sharing)

#### Build code:
```
make build
```

#### Run Code:
```
make run
```

To test: Added the postman collection rideshare.postman_collection.json file

Sample API:

```
curl --location 'localhost:3000/user' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Rahul",
    "gender": "M",
    "age": "35"
}'
```

