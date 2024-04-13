Output demo.gif

Set Shell "bash"
Set FontSize 18
Set Padding 2
Set Margin 2
Set Width 640
Set Height 480
Set Framerate 60
Set TypingSpeed 25ms
Set PlaybackSpeed 0.4
Set CursorBlink false

Require filegen

Type "filegen -count 100 -size 10K -path test-files"
Enter
Type "find test-files -type f -printf '%s %p\n' | head"
Enter
Sleep 3
Type "find test-files -type f | wc -l"
Enter
Sleep 3
Type "head -c 50 `find test-files -type f | head -1`"
Enter
Sleep 3
Hide
Type "rm -rf test-files"
Enter

Type "clear"
Enter

Show
Type "filegen -count 100 -size 1.25M -random-size-min 100K -path test-files"
Enter
Type "find test-files -type f -printf '%s %p\n' | head"
Enter
Sleep 3
Type "find test-files -type f | wc -l"
Enter
Sleep 3
Type "find test-files -type f -size +1M | wc -l"
Enter
Sleep 3
Type "du -hs test-files"
Enter
Sleep 3
Type "head -c 50 `find test-files -type f | head -1`"
Enter
Sleep 3
Hide
Type "rm -rf test-files"
Enter
