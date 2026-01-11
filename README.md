## Go Lang Text Editor

On my journey to follow these projects and learn how to build these projects. https://austinhenley.com/blog/challengingprojects.html

I am starting out with the Text Editor.
From the information we have I am starting with a plan of how to build this.

While in the directory can build it with `go build -o texteditor`
and run it with `./texteditor testfile.txt`

Plan:
 - Research Data structures
 - Doing Piece Table `DONE`
   - Creating a Piece Table data structure `DONE`
   - Creating a Piece Table `DONE`
   - Converting away from Piece Table `DONE`
   - Inserting into Piece Table `DONE`
   - Deleting Text `DONE`
 - Figure out UX and Cursor `Done`
   - Add cursor structure `Done`
   - Create Editor and Piece Table interface `Done`
   - Add some UI (Just grab a library?) `Done`
   - Implement logic loop `Done`
   - Add tracking of location? (Up and Down) `Done`
   - Do we want to consider a scroll? Lol nope `Not Done`
 - I/O capabilities
   - File manager, with read and write file. Dirty flag?
   - Integrate file manager with Editor interface
   - Create file line args for file
   - Save keybindings
   - Save As?
   - Unsave warnings?
   - Add status bar icons for files unsaved?
   - Error handling.
 - Undo and redo? (Should be just remove from piece table and store and then readd to piece table?)
 - Yes, add more.
