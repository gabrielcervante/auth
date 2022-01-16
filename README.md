# auth

This is a simple signup api

You can access the db.go file and change the database credentials to your local postgres credentials.

To run it properly create a database named "auth" if you doesn't changed the database name in db.go, create a table called "users" and 3 columns called "email", "password" and "id", so your setup will be completed and now you can run the application using docker.

It has a docker config file, you can create a container.

Use it to create a container:

sudo docker build --tag container-name-here .

Use it to run your container and be able to access localhost:

sudo docker run --network=host container-name-here

