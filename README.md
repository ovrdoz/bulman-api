# bulman-api

### to run local
export CONNECTION_STRING="mongodb+srv://admin:admin123@cluster0.s2ibu.mongodb.net/bulman?retryWrites=true&w=majority"
go run main.go 

### to build
git add --all && git commit -m "add first commit" && git push &&  heroku logs -a bulman-api
heroku logs --tail -a bulman-api 