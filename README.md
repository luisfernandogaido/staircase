# staircase
Amazon Coding Interview Question - Recursive Staircase Problem

https://www.youtube.com/watch?v=5o-kdjv7FD0

### how to use

Download de binary of your SO in bin folder. In terminal type, for example:

```
staircase -n 4 -j 1,3,5
```

You should not put very high `n` proporcionally to `j` elements, since the possible paths are huge and
you need stop de program.
You can increase accelerate de results by try something like this:

```
staircase.exe -n 20 -j 1,2,4 > out.txt
```

This solution is not as elegant as the one showed in the video. But it shows the paths and is interesting.