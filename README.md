Simple program to adjust the backlighting on my laptop.

* `igotyourbacklight` outputs the current backlight stats.
* `igotyourbacklight <int>` where int is a number between -100 and 100 adjusts the brightness up or down by %.

*Example:*

`igotyourbacklight -10` makes the screen 10% dimmer.


## i3wm Setup

Add the following to your i3wm config:
```
bindsym XF86MonBrightnessUp exec igotyourbacklight 10
bindsym XF86MonBrightnessDown exec igotyourbacklight -10
```


## Compatibility

This currently works on my laptop where the backlight device is accessible at `/sys/class/backlight/intel_backlight/`.