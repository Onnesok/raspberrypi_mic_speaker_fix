# Raspberry pi microphone and sound problem fixed
</body>
</html>


[![Facebook URL](https://img.shields.io/static/v1?color=red&label=Facebook&logo=Facebook&logoColor=white&style=for-the-badge&message=Connect)](https://facebook.com/Onnesok.94)
[![Youtube URL](https://img.shields.io/static/v1?color=red&label=Youtube&logo=Youtube&logoColor=white&style=for-the-badge&message=subscribe)](https://www.youtube.com/Onnesok)
[![Instagram URL](https://img.shields.io/static/v1?color=red&label=Instagram&logo=Instagram&logoColor=white&style=for-the-badge&message=follow)](https://www.instagram.com/Onnesok/)
[![Pinterest URL](https://img.shields.io/static/v1?color=red&label=Pinterest&logo=pinterest&logoColor=white&style=for-the-badge&message=Follow)](https://www.pinterest.com/ratulhasan94/)

**In freshly installed raspbian os sometimes we need to specify cards to work with microphone. Sometimes it may require to create .asoundrc file to specify speaker and microphone card seperately**

<p>update and upgrade your pi</p>

```
sudo apt update -y
sudo apt upgrade -y
```

## Now find device ids. 
<p> For speakers type </p>

```
aplay -l 

```
<p> Output </p>

```
**** List of PLAYBACK Hardware Devices ****
card 0: ALSA [bcm2835 ALSA], device 0: bcm2835 ALSA [bcm2835 ALSA]
  Subdevices: 7/7
  Subdevice #0: subdevice #0
  Subdevice #1: subdevice #1
  Subdevice #2: subdevice #2
  Subdevice #3: subdevice #3
  Subdevice #4: subdevice #4
  Subdevice #5: subdevice #5
  Subdevice #6: subdevice #6
card 0: ALSA [bcm2835 ALSA], device 1: bcm2835 IEC958/HDMI [bcm2835 IEC958/HDMI]
  Subdevices: 1/1
  Subdevice #0: subdevice #0
card 1: Microphone [USB Microphone], device 0: USB Audio [USB Audio]
  Subdevices: 1/1
  Subdevice #0: subdevice #0
  
  ```
  <p> 
  Based on the above output we have three devices:

The onboard 3.5mm plug listed as BCM2835 ALSA: Card 0, Device 0 (“hw:0,0”)
The onboard HDMI connection listed as bcm2835 IEC958/HDMI: Card 0, Device 1 (“hw:0,1”)
The USB microphone listed as USB Audio: : Card 1, Device 0 (“hw:1,0”)
In this example we want to use the 3.5mm audio jack, so we’ll use Card 0, Device 0 as the way to locate our speaker device. The USB microphone doesn’t have audio output and is not valid as a speaker setting. It is listed however, which can cause confusion.
  </p>
  
  # Now for microphone
  
  ```
  arecord -l
  
  ```
  
  <p> Output </p>
  
  ```
  
  **** List of CAPTURE Hardware Devices ****
card 1: Microphone [USB Microphone], device 0: USB Audio [USB Audio]
  Subdevices: 1/1
  Subdevice #0: subdevice #0
  
  ```
  
  <p> Based on the output usb microphone is there as Card 1 and Device 0 (“hw:1,0”) </p>

<hr/>
  