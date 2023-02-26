#!/bin/bash
URL=$(curl -s https://api.github.com/repos/dev47apps/droidcam/releases/latest \
| jq '.assets | .[0].browser_download_url' \
| tr -d \")

DroidCamZIP=droidcam_latest.zip

cd /tmp/
if [ -f "$DroidCamZIP" ] ; then
    rm "$DroidCamZIP"
fi
wget --quiet "$URL" --output-document "$DroidCamZIP"

echo -n "d1038e6d62cac6f60b0dd8caa8d5849c79065a7b $DroidCamZIP" | sha1sum -c --status -
retVal=$?
if [ $retVal -ne 0 ]; then
    echo "Error the sha1sum does not match"
	exit $retVal
fi

unzip "$DroidCamZIP" -d droidcam
cd droidcam
sudo ./install-client
sudo ./install-video
sudo rmmod v4l2loopback_dc
sudo insmod /lib/modules/`uname -r`/kernel/drivers/media/video/v4l2loopback-dc.ko width=1280 height=720
cd ..
rm -rf droidcam "$DroidCamZIP"