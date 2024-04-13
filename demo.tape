Output demo.gif

Set Shell "bash"
Set FontSize 16
Set Padding 5
Set Margin 5
Set Width 640
Set Height 480
Set Framerate 60
Set TypingSpeed 25ms
Set PlaybackSpeed 0.4
Set CursorBlink false
Set WindowBar ColorfulRight

Require filegen

Type "filegen -count 100 -size 10K -path data"
Enter
Type "find data -type f -printf '%s %p\n' | head"
Enter
Sleep 3
Type "find data -type f | wc -l"
Enter
Sleep 3
Type "head -c 50 `find  data -type f | head -1`"
Enter
Sleep 3
Hide
Type "rm -rf data"
Enter

Type "clear"
Enter

Show
Type "filegen -count 100 -size 1.25M \"
Enter
Type "-random-size-min 100K -path data"
Enter
Sleep 3
Type "find data -type f -printf '%s %p\n' | head"
Enter
Sleep 3
Type "find data -type f | wc -l"
Enter
Sleep 3
Type "find data -type f -size +1M | wc -l"
Enter
Sleep 3
Type "du -hs data"
Enter
Sleep 3
Type "head -c 50 `find data -type f | head -1`"
Enter
Sleep 3
Hide
Type "rm -rf data"
Enter
