In this project, I am playing with vars.
I found a nice pongo based jinja2 templating app in golang, and it works super neat.
But when I cross-compile in windows, I get a bunch of errors because the env var KEY names and some VALUES include exotic characters.

Here I try to print the env vars in windows, and sanitize them unset them before passing env vars to the pongo templating engine.


Here are some example fail vars:

```
CommonProgramFiles(x86)=C:\Program Files (x86)\Common Files
ProgramFiles(x86)=C:\Program Files (x86)
```