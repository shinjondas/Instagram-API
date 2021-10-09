# Instagram-API

This API has been solely programmed on [Go](https://golang.org/), and for database storage [MongoDB](https://www.mongodb.com/) has been used. This is originally done as a part of a technical task for Appointy. 

## Cloning repository

Use git to clone the repository as follows.

```bash
 git clone https://github.com/shinjondas/Instagram-API/
```
## Functionalities
- Adding new user to platform![New User](https://github.com/shinjondas/Instagram-API/blob/main/output/PostUser.PNG)
- Retrieving data regarding a user ![Get User](https://github.com/shinjondas/Instagram-API/blob/main/output/GetUser.PNG)
- Adding new post to platform![Create Post](https://github.com/shinjondas/Instagram-API/blob/main/output/PostPost.PNG)
- Retrieving all data related to that post![Fetch Post](https://github.com/shinjondas/Instagram-API/blob/main/output/GetPost.PNG)
- Getting all posts posted from a given userID![Fetch All Posts of a user](https://github.com/shinjondas/Instagram-API/blob/main/output/AllPostsOfUser.PNG)
- Encrypted paswords for added security using [md5](https://en.wikipedia.org/wiki/MD5)

## Usage

```go
go mod init instagram-api
go run index.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to add and update tests as appropriate.
