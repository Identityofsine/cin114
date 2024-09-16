# ~/ssl 

The ssl folder is meant only for docker, please keep this private and do not give this up and handle with caution...

## Generation 

> Please "cd" into the directory of the project and run the command:
```bash
PROJECT_ROOT=$PWD
```

OR

> Please replace the variable with the path to the root directory of the project:
```bash
PROJECT_ROOT=/path/to/project/
```

### 1. Install the required packages to use certbot.
```bash
sudo apt update
sudo apt install python3 python3-venv libaugeas0
```

### 2. Setup a virtual python3 environment.
```bash
sudo python3 -m venv /opt/certbot/
sudo /opt/certbot/bin/pip install --upgrade pip
```

### 3. Install certbot.
```bash
sudo /opt/certbot/bin/pip install certbot
```
### 4. Prepare the Certbot command.
```bash
sudo ln -s /opt/certbot/bin/certbot /usr/bin/certbot
```

### 5. Generate the SSL certificate.
```bash
sudo certbot certonly --standalone -d www.cin114.net -d cin114.net -d cin114films.com -d www.cin114films.com
```

### 6. Copy the generated SSL certificate to the ssl folder.
```bash
sudo cp /etc/letsencrypt/live/cin114.net/fullchain.pem $PROJECT_ROOT/ssl/cin114.net.crt
sudo cp /etc/letsencrypt/live/cin114.net/privkey.pem $PROJECT_ROOT/ssl/cin114.net.key
sudo cp /etc/letsencrypt/live/cin114films.com/fullchain.pem $PROJECT_ROOT/ssl/cin114films.com.crt
sudo cp /etc/letsencrypt/live/cin114films.com/privkey.pem $PROJECT_ROOT/ssl/cin114films.com.key
```

## Errors

Any errors can be solved with a google search, however you should know that you shouldn't run this on your main computer. These commands are strictly for the server (the endpoint of those domains.)
