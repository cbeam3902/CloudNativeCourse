# Lab 7
This lab introduces Docker and how to containerize the server from lab 5. The code here is the same as the one from lab 5 with no changes.

First install docker

    $ sudo apt-get update

    $ sudo apt install docker.io

    // Verify   

    $ sudo docker run hello-world

Then execute the following line:

    $ sudo docker image build -t movieserver .
    $ sudo docker container run -p 50051:50051 movieserver

If you have docker compose, then you can do `sudo docker compose up -d` and check the logs using `sudo docker compose logs`. Once you're done you can use `docker compose down` to stop it.