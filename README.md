<h2>GO-db: Independent Address File Based Database for Go-Lang</h2>

Go-Lang is a fairly new language by Google, which translates to the best processing speeds, compared to the other languages
in the market.

However, there is lack of any good File Based Database Management System, which could be reliable enough to meet the vigorous 
computing needs, and at the same time be simple to use.

We have made an attempt to bridge this gap, by developing GO-db.

<h4>Status: Under Development (Pre-Alpha),  currently supports operations via strings only</h4>

<h3>First: Why to use GO-db?</h3>

GO-db is a simple to use file based database, which is light, simple and efficient on resources. It has been designed with 
simplicity in <strong>AIM</strong>. 

<h3>Usage</h3>

<hr>

<h5>Saving</h5>

<code>gobase.Save(path string, data string) bool</code>

Example:

<code>gobase.Save("Test", "some data")</code><br>
<code>gobase.Save("Random/Path/Test", "some data")</code>
<br>param 1: Status i.e., True if sccessful, else False

<hr>

<h5>Retriving</h5>

<code>gobase.Retrive(path string) (string, bool)</code>
<br>param 1: Available data at the specified path if successful; else ERROR_CODE: ```DOCUMENT_UNAVAILABLE``` or ```ERROR```
<br>param 2: Status i.e., True if successful, else False

Example:

<code>gobase.Retrive("Test")</code><br>
<code>gobase.Retrive("Random/Path/Test")</code>

<hr>

<h5>Get Available Collections at an Address / Path</h5>

<code>gobase.CollectionsAvaliable(path string) ([]string, bool)</code>
<br>param 1: Available Collections at the specified path
<br>param 2: Status i.e., True if successful, else False

Example:

<code>gobase.CollectionsAvaliable("/")  // output: [Test, Random]</code><br>
<code>gobase.CollectionsAvaliable("Random") // output: [Path]</code>

<hr>

<h5>Deleting</h5>

<code>gobase.Delete(path string) bool</code>

Example:

<code>gobase.Delete("/")  // deletes all collections</code><br>
<code>gobase.Delete("Random") // deletes Random collection</code>

<hr>



