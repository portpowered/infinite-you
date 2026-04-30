# problem statement

the current website is failing to delete files, with the endpoint not being found when attempting to delete a file based on its filesystem ID and file ID. 

This is likely induced by the problem of having multiple API protocols for functionally the same thing, requesting customers to enumerate by filelabel vs file id, adn what not. 

1. Let's just consoldiate all the functionality to upsert/delete/create/list endpoints based on the same pricnipal ID + filesystem ID + file ID, or principal ID + filesystem label + file label. Remove the other elements. 

2. let's fix deletes to work with a given filesystem ID and file ID, int he backend,a nd ensure we add tests to confirm the behavior works. 