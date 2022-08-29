## Instalación de Docker

```bash
sudo su
pacman -S docker docker-compose
systemctl enable docker
systemctl restart docker
groupadd -r -g 82 www-data
useradd -M -r -u 82 -g 82 -c "User HTTP files" -s /usr/bin/nologin www-data
exit
usermod -aG docker,www-data $(whoami)
```