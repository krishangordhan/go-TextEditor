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
 - I/O capabilities `Done`
   - File manager, with read and write file. Dirty flag? `Done`
   - Integrate file manager with Editor interface `Done`
   - Create file line args for file `Done` 
   - Save keybindings `Done` 
      - Add save error bubbling up.
   - Save As? `Done`
   - Unsave warnings? `Done`
   - Add status bar icons for files unsaved? `Done`
   - Error handling. `Not Done`
 - Scrolling the page `Done`
   - Implement some way to track scroll state. `Done`
   - Implement vertical scrolling logic, 3-line margin `Done`
   - Render to display only relevant text? (Hmm gonna be funky this one?) `Done`
   - Horizontal scrolling? `Done`
 - Undo and redo?
   - Create struct for edit actions
   - Add interface to editor for undo redo
   - Implement Undo, add tests
   - Implement Redo, add tests
   - Add keyboard shortcuts
   - Edge cases, empty undo stack, redo after new edit
 - Yes, add more.
