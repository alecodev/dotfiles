## Install DroidCam

```bash
sudo pacman -S --needed linux-headers gcc make libappindicator-gtk3 android-tools android-udev

# Download the latest version of droidcam from its GitHub repository in the /tmp folder with the name droidcam_latest.zip
# https://github.com/dev47apps/droidcam/releases

# Check the sha1sum
sha1sum droidcam_latest.zip

# Install
unzip droidcam_latest.zip -d droidcam
cd droidcam
sudo ./install-client
sudo ./install-video
sudo rmmod v4l2loopback_dc
sudo insmod /lib/modules/`uname -r`/kernel/drivers/media/video/v4l2loopback-dc.ko width=1280 height=720
cd ..
rm -r droidcam droidcam_latest.zip

# If you use adb you can run it with the following command
adb devices
droidcam-cli -v -vflip adb 4747
```