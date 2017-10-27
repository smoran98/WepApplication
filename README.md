# WepApplication in GO
# Shane Moran
# G00338607

# ====================================
Create a Web App in Go

Create a Single Git Repo as your submission
# ====================================


# 1. Guessing game
Create a web application in Go responding with text “Guessing game”, should be the response body irrespective of what request is received. Explain in your README how to examine the response, including the headers, using curl.

# 2. Make the text a H1
Change your web application to make “Guessing game” a level 1 heading (HTML). Test application to make sure HTML is rendered by your browser. If it isn’t, fix it.

# 3. Serve a page using Bootstrap
Change the web application to serve a web page rather than hard-coding the text into the web application executable. Use the Bootstrap starter template, changing the header to say “Guessing game”. Add a link on the page to the relative URL /guess with the text “New game” and this page will serve as the root resource in the web app.

# 4. Add a guess route
Create a new route in your application at /guess and have it serve a new page called guess.html. Use the same Bootstrap code for the page as in index.html but add a level 2 heading with the text “Guess a number between 1 and 20”. Add a form, with a text input with the name “guess” and a submit button with the label “Guess”. The form's action should be /guess and the method should be GET.

# 4. Turn the guess page into a template
Change the web application to use the guess.tmpl file as a template. Add a single variable to the template called Message and create a struct in Go containing a single string. Create an instance of the struct with the string set to “Guess a number between 1 and 20” and have the template render this in the H2 tag.

# 5. Check for a cookie
In the /guess handler check if the cookie called target is set. If it is not, generate a random number between 1 and 20 and set a cookie called target to that value in the response. Otherwise, leave the cookie at its current value.

# 6. Check for a variable
Have the /guess handler check if a URL encoded variable called guess has been set to an integer. If it has, have the text “You guessed {guess}.” inserted into the template where {guess} is replaced with the value of guess.

# 7. Compare the cookie and the guess
If the target cookie and the guess variable are both set, then have the /guess handler compare them. If they are equal, set the target cookie to another randomly generated integer, and have the template display a congratulations message and a link to create a new game. Otherwise, have the template display a message telling the user what their guess was and whether it was too high or too low.

# 8. Use the POST method
Change your HTML form in guess.html to use the POST method instead. Make sure your application still works, bug fixing it if necessary.
