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
 - Undo and redo? `Done`
   - Create command interface and command types `Done`
   - Update insert/delete operations to execute commands `Done`
   - Implement Undo, add tests `Done`
   - Implement Redo, add tests `Done`
   - Add keyboard shortcuts `Done`
   - Edge cases, empty undo stack, redo after new edit `Not Done. Well done all but redo after new edit`
 - Text selection `Done`
   - Add select start and end to cursor `Done`
   - Track selection in Editor `Done` 
   - Render selection in display `Done`
   - Delete selection on backspace or delete `Done`
   - Replace selection on insert `Done`
   - Edge cases `Done`
 - Copy and paste `Done`
   - System clipboard library `Done`
   - Copy selection to clipboard `Done`
   - Paste from clipboard at cursor `Done`
   - Update undo/redo to support this `Done`
   - Keyboard shortcutes `Done`
 - Line numbers `Done`
   - Add line number rendering `Done`
   - Calculate width based on line count `Done`
 - Auto indent
   - Detect indentation of current line.
   - Insert new line with same indentation
   - Figure out edge cases (empty lines, mixed tabs and spaces) 
