Usage
=====

Local Build
-----------

    make all
    
Docker Build
------------

    make
    docker build -t mdota .
    docker run -p 5000:5000 -e PORT=5000 -it mdota

Heroku Build
------------

    # Do Once
    # Install Heroku Toolbelt - https://toolbelt.heroku.com/
    heroku login    # Login to Heroku
    heroku create   # Create new app with random name
    heroku plugins:install heroku-container-registry    # Use docker plugin
    heroku container:login  # Login to heroku docker registry
    
    # For Each Release
    make heroku     # Build docker image and push to heroku
    heroku open     # Open URL
