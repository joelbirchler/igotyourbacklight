Simple program to adjust the backlighting on my laptop.

_Don't use this code as an example of how to write Go programs. I'm still learning the language. Feedback appreciated._

* `igotyourbacklight` outputs the current backlight stats.
* `igotyourbacklight <int>` where int is a number between -100 and 100 adjusts the brightness up or down by %.

*Example:*

`igotyourbacklight -10` makes the screen 10% dimmer.

## TODO

* write to /sys/class/backlight/intel_backlight/brightness (remember \n)
* figure out how to test
_