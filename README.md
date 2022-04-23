# ascii-image-generator

### How to use 
This is program is really simple to use. All you need to do is download the source code and compile it using the command: 

```bash
go build
```

After that, you're gonna generate your executable, called __ascii-image-generator__.

Finally, just run your program and provide a file as input: 
```bash
./ascii-image-generator <image png file>
```

You're gonna see that, the converted image in your terminal output. Depending on the resolution of your input image, it can be hard to see the image clearly, in such a case, you can save the image in a text file, like this:


```bash
./ascii-image-generator <image png file> > file.txt
```
