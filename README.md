CIN114 - Freelancing Project

## Getting Started

### Prerequisites

1. Docker
Docker is the heart and soul of this application and the entire web application is readily deployable using docker.

```bash
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

## Install Docker
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-compose
```

2. SSL 

Please refer to the [SSL](./ssl/README.md) folder for more information.

3. DNS

DNS should be already conifgured for the domain names you are using. If you run into any issues with domains and DNS please refer to any documentation through GoDaddy or Cloudflare.

### Installation

1. Clone the repo
```bash
git clone https://github.com/Identityofsine/cin114
```

2. Change directory to the project root
```bash
cd cin114
```
3. Run the docker-compose command
```bash
docker-compose up -d
```

4. Access the website
```bash
https://domain.com/
```
