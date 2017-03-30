# goApi
Demo REST API in go lang with mysql database and AWS SQS messaging service to build scalable and fast REST API

*This project spins up:*
  1. An queue on sqs 
  2. Setup a database with dummy videos list
  3. API listening to:
    
    a. videos: 
       Returns list of videos
   
    b. feedback:
       a.Pushes feedback message to sqs queue 
       b. Receiver polls the queue for feedback message to record feedback
       c. Accepts feedback as {"videoId": 1, "like": true}
    
 
