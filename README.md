## Setup

_temporarily_

`export PB_ACCESS_TOKEN='<YOUR_PUSHBULLET_ACCESS_TOKEN>'`

_permanently_ (change .bashrc to your profile file)

`echo "PB_ACCESS_TOKEN='<YOUR_PUSHBULLET_ACCESS_TOKEN>'" >> ~/.bashrc`

`echo -e 'export PB_ACCESS_TOKEN\nexport PATH=$PB_ACCESS_TOKEN:$PATH' >> ~/.bashrc`

## Usages

./pushbullet-cli note title body