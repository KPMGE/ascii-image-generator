# ascii-image-generator

### How to use 
In order to generate ascii representations from your png images, all you need to do is download this source code and compile it using the command: 

```bash
go build
```

After that, you're gonna generate the executable, called __ascii-image-generator__.

Finally, just run your program and provide a file as input: 
```bash
./ascii-image-generator <image png file>
```

You're gonna see the converted image in your terminal. Depending on the resolution of your input image, it can be hard to see it clearly, in such a case, you can either zoom out your terminal as much as possible or save the image in a text file and, like this:

```bash
./ascii-image-generator <image png file> > file.txt
```
### Examples 
Running the program and passing in a sample image, we see one clearly one of the most powerful editors ever is being displayed on our terminal! 
![vim terminal image](https://github.com/KPMGE/ascii-image-generator/blob/main/images/vim-gen.png)
