# Hey, thanks for coming! This chat allows people to share the experience while using a chatbot for making phone calls

## The idea

The idea is to make interaction with a chatbot more social and let unlimited amount of peers use a chatbot and share the experience between each other.

I also did not add usernames to each message to bring some kind of anonymity to each particular message.

You know who is participating in the chat, but knowing who said something particular can apply pressure on the person speaking.

## Tech

 - It was decided to Go :smirk: with [Revel](https://github.com/revel/revel) framework as it has a [chat example](https://github.com/revel/examples/tree/master/chat) already implemented and it has more features than the [Gorilla WebSocket chat example](https://github.com/gorilla/websocket/tree/master/examples/chat).

 - [Chatbot](https://en.wikipedia.org/wiki/Chatbot) logic to interact with the users,
 - [WebSockets](https://en.wikipedia.org/wiki/WebSocket) for the chat between users functionality,
 - UI using some cool [CodePens](https://codepen.io/) for the better UX,
 - [Twilio REST API](https://www.twilio.com/docs/usage/api) for making phone calls 

## Deployment with Docker

revelChat is very easy to install and deploy in a Docker container.

```sh
cd revelChat
docker build -t revel-chat .
docker run -it -p 8080:8080 revel-chat
```

This will create the image and pull in the necessary dependencies.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 8080 of the host to port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -it -p 8080:8080 revel-chat
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8080
```

## Feedback

Please feel free to leave comments if you happen to visit this page and suggest what can be added or improved

## Demo

![](https://github.com/Dysar/revelChat/blob/master/public/images/screely-1553024086978.jpg)
