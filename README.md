## Hi. This is a totaly useless API written in Go.

There are only 2 methods you can use:
- **getUsers**, which gives you a list of all users in database
- **getComments** will provide comments of an actual user (requires user ID)

``` sh
https://example.com/getComments?id=1
```

All responses come in JSON format.

[Try it](http://uselesss-api.heroku.com).