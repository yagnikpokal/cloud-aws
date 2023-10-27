# Amazon SNS
Simple notification service
It is a message publishing and processing services (pubsub)
### Features
- Fully managed service so no need to worry about the harddisk, infrastructur and host
- Durable service means
  - it is gaurentee that message will not lost
  - It is not a gaurentee that they will deliver right away may be they can deliver after few time
- Autoscaling

### Topics
- a logical access point and communication channel through which message sent to multiple services

### Publisher/producer
- Publisher will publish the messages to topic
- Event publishers will only sends messages to one sns topic

### Subscriber/consumer
- Subscriber will subscribe the messages from topic
- SNS subscriber can be a http/https, emails, SMS messages, mobile notifications, SQS queue, Lambda functions
- Application to person -->> text message, email etc
- Application to application -->> Lamnda, SQS, http end points
- As many as subscriber as we want to listen for topic notification

<img width="644" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/e56853b7-07e7-43e7-b928-9369ab11f4d5">

### Can send a message to one person only 
- using the application filtering even if your topic is made for multiple customers
<img width="644" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/56216f2c-5215-494f-8253-7d8f47873904">

### Send a message application to application
<img width="606" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/17414996-21f6-43fa-8d37-feb4bffaf7ad">

### Hands on in the console
- Click on serachbar type sns
- Click on create
- Create subscription based on need

 ## Questions
 What if we can bypass topic? Means we can directly send messages to lambda or any other app. Why we need SNS topic?
 Answer: topic knows lambda, SQS internally and our application dont have idea so if we bypass the topic then we have to know all the pplication like lambda, sqs etc
 
 <img width="606" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/42ef6648-ccc9-496b-87da-a6451223ebab">


