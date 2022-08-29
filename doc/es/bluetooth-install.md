## Instalación de Bluetooth

```bash
sudo su
pacman -S bluez bluez-utils
systemctl enable bluetooth
sed -i 's/#AutoEnable=true/AutoEnable=true/' /etc/bluetooth/main.conf
printf "### Automatically switch to newly-connected devices\nload-module module-switch-on-connect\n" >> /etc/pulse/default.pa'
```