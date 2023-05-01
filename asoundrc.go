//written by Ratul Hasan
pcm.!default {
	type asym
	capture.pcm "mic"
	playback.pcm "speaker"
  }
  
  pcm.mic {
	type plug
	slave {
	  pcm "hw:1,0"     //modify this according to your card
	}
  }
  
  pcm.speaker {
	type plug
	slave {
	  pcm "hw:0,0"  //modify this according to your card
	}
  }
  
  