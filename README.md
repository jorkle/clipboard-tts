## clipboard-tts
Uses OpenAI text-to-speech API to speak the contents of your clipboard through your speakers/headphones.


## Requiremenets
1. Linux
2. `ffplay` command
3. Ability to configure a keyboard shortcut in your desktop environment

## Installation
```bash
go install https://github.com/jorkle/clipboard-tts@latest
```

## Instructions
1. Install clipboard-tts
2. Configure your desktop environment to run `clipboard-tts -apikey <openai-api-key>` when a keyboard shortcut is pressed


## Usage
1. Any text that you would like read to you outloud, simply copy it to your clipboard.
2. Press the "clipboard-tts" hotkey that you configuration in your desktop environment.
3. Success! You should hear the contents of your clipboard read to you through your default playback device.
