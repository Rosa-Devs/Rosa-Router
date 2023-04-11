<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Mihalic2040/Rosa-Router">
    <img src="images/Logo.svg" alt="Logo" width="200" height="200">
  </a>

# Rosa Router
</div>


## The project was inspired by the architectural disadvantages of the most popular messengers such as:

### Client - Server architecture:
<br />
    <div align="center">
        <a href="https://github.com/Mihalic2040/Rosa-Router">
            <img src="images/client-server.png" alt="Logo">
        </a>
    </div>

### No source code:
<br />
    <div align="center">
        <a href="https://github.com/Mihalic2040/Rosa-Router">
            <img src="images/code.png" alt="Logo">
        </a>
    </div>

## In my opinion, this program can solve these problems and provide people with the most reliable messaging service.

# What is this?
Router this is base of CLI. He provide a local web api to communicate with P2P network.
This will allow me to write a client to a dear person who knows how to interact with the JSON HTTP API.

## Hot potato protocol
<br />
    <div align="center">
        <a href="https://github.com/Mihalic2040/Rosa-Router">
            <img src="images/potato.png" alt="Logo">
        </a>
    </div>
Also, the router is responsible for supporting the hot potato protocol as it runs as a service.
The essence of the problem was that if 2 nodes are not in the network, it is not realistic to deliver a message to it. This protocol allows a message to live for 60 minutes before being marked as undelivered. If the receiving node cannot receive this message within 60 minutes, it will be deleted and marked as not delivered to the sender.

# How to build and run
    
    git clone https://github.com/Mihalic2040/Rosa-Router
    cd Rosa-Router
    make build  #Generate launch codes for nuclear missiles

<!-- CONTRIBUTING -->
# Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

