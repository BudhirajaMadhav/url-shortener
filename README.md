# url-shortener
Just another url shortener written in Go

Steps:
- [x] Connect to mongoDB
- [ ] Make a router and controller which will take the params from the url, for path and URL and push that to MongoDB.
- [ ] Check if that path still exists, if yes, return error status code(path not available).
- [ ] If user does not provide a path, then generate it with crypto.
- [ ] Make a http handler which will lookup the path and redirect it to the respective path, if there's no path then redirect it to default page.
- [ ] Add YAML functionality too.
