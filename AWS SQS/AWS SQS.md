# SQS Service
- Simple queue service
- Offers async communication as opposed with the API calls
- Follows fifo order
- Oldest service
- Fully managed service
- Provides queue for storing massages to travel to different applications or microservices
<img width="660" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/eb7e4d53-1ab1-4c34-825c-8ccafebaf218">

### Core concept of SQS
<img width="484" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/fd2a1786-b5a5-42bc-8f39-112cbd4a8cf5">


- Multiple producer -->> SQS Queue -->> COnsumer 
- Scale from 1 messgage/sec to 10000message/sec
- Default retention of message 4 days to 14 days
- No limit for messages
- Message were deleted by queue once consumer will read
- Low latency <10 ms 
- Go to amazon SQS -->>Create queue-->>Select standard-->> give name-->> Create queue

### Visibility time out and deletion or reque of message
<img width="516" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/46f1a345-6545-497e-9acd-89fee4705f22">

### Why SQS needed since we can use direct api calls?
<img width="651" alt="image" src="https://github.com/yagnikpokal/cloud-aws/assets/29862439/36fceac1-3317-4286-9eed-776534aa381e">

