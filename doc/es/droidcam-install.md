## Instalación de DroidCam

```bash
sudo pacman -S --needed linux-headers gcc make libappindicator-gtk3 android-tools android-udev

# Descarga la última versión de droidcam desde su repositorio de GitHub en la carpeta /tmp con el nombre droidcam_latest.zip
# https://github.com/dev47apps/droidcam/releases

# Verifica el sha1sum
sha1sum droidcam_latest.zip

# Instala
unzip droidcam_latest.zip -d droidcam
cd droidcam
sudo ./install-client
sudo ./install-video
sudo rmmod v4l2loopback_dc
sudo insmod /lib/modules/`uname -r`/kernel/drivers/media/video/v4l2loopback-dc.ko width=1280 height=720
cd ..
rm -r droidcam droidcam_latest.zip

# Si usa adb, puede ejecutarlo con el siguiente comando
adb devices
droidcam-cli -v -vflip adb 4747
```