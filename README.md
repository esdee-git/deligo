# deligo
experiments with go/golang web programming

i've been thinking for some more or less practical tasks to learn Go and came up with this: the ultimate goal is a web server adding links to delicious. apart from web / backedn programming it is also an exercise in a tiny bit of frontend tasks: intergrating bootstrap and some jquery into the final "product"

for simple communication with delicious api i used code from https://github.com/josegm/yummy/, since for some reason my hacks didn't work.
other than that the code used http package with Gorilla's mux. 
i had a lot of fun with the frontend, header.html, where i played somewhat with bootstrap and jquery's ajax.

eventually i also added dependency management through glide and docker image configuration in the Dockerfile file

further on, for docker's image i switched from using ubuntu image as a base (which would result in 750MB images) to alpine base (only about 15MB(!) image).
for this to work though, the Dockerfile includes a directive to add to the image file ca-certificates.crt.
it is usualy found in /etc/ssl/certs/ca-certificates.crt, needs to be present in the build folder so the docker image creation succeeds.

one more setup step is to include a file delicious.json also in the build folder
it should look like; create clientId and secret on delicious' site
{
  "clientID": "123123123123123123123123123",
  "clientSecret": "345345345345345345345345345",
  "username": "your user name",
  "password": "your password"
}

if you plan to build a docker image as prescribed in the Dockerfile, you should build the executable with:
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o deligo .
more info at http://esdeeblog.blogspot.de/2015/10/create-slim-and-lean-docker-image-for.html

