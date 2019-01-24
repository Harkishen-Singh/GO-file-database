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

<h3>Installation</h3>

add <code>"github.com/Harkishen-Singh/GO-db/gobase"</code> to import block of your file and open the terminal, and type <code>go get -v -t -d && go build -v</code>

<h3>Usage</h3>

<hr>

<h4>Saving</h4>

To save a data of String Type:  
<code>gobase.Save(path string, data *string) bool</code>
<br>
Similarly to save the data of different datatypes:<br>
<code>gobase.Save_datatypecode_(path string, data *_datatype_) bool</code><br>

Example:

var details = "some data"
var id int = 1234
<code>gobase.Save("Test", &details)</code><br>
<code>gobase.SaveInt("Random/Path/Test", &id)</code><br>
<br>return_param 1: Status i.e., True if successful, else False

To save an array of String Type:  
<code>gobase.SaveArr(path string, data *string) bool</code>

Similarly to save array of different datatypes:<br>
<code>gobase.SaveArr_datatypecode_(path string, data *_datatype_) bool</code><br>

Example:

var details = {"Some", "Data"}
var vals = []int{1,2,3,4,5}
<code>gobase.SaveArr("Test", &details)</code><br>
<code>gobase.SaveArrInt("Random/Path/Test", &vals)</code>
<br>return_param 1: Status i.e., True if successful, else False

**The datatype codes for different datatypes are:**
Datatype | Datatype Code
-------- | -------------
Unsigned Integer8 | uint8
Unsigned Integer16 | uint16
Unsigned Integer32 | uint32
Unsigned Integer64 | uint64
Signed Integer8 | int8
Signed Integer16 | int16
Signed Integer32 | int32
Signed Integer64 | int64
Float32 | float32
Float64 | float64

<hr>

<h5>Retriving</h5>

<code>gobase.Retrive(path string) (string, bool)</code>
<br>return_param 1: Available data at the specified path if successful; else ERROR_CODE: ```DOCUMENT_UNAVAILABLE``` or ```ERROR```
<br>return_param 2: Status i.e., True if successful, else False

Example:

<code>gobase.Retrive("Test")</code><br>
<code>gobase.Retrive("Random/Path/Test")</code>

<hr>

<h5>Get Available Collections at an Address / Path</h5>

<code>gobase.CollectionsAvaliable(path string) ([]string, bool)</code>
<br>return_param 1: Available Collections at the specified path
<br>return_param 2: Status i.e., True if successful, else False

Example:

<code>gobase.CollectionsAvaliable("/")  // output: [Test, Random]</code><br>
<code>gobase.CollectionsAvaliable("Random") // output: [Path]</code>

<hr>

<h5>Deleting</h5>

<code>gobase.Delete(path string) bool</code>
<br>return_param 1: Status i.e., True if successful, else False

Example:

<code>gobase.Delete("/")  // deletes all collections</code><br>
<code>gobase.Delete("Random") // deletes Random collection</code>

<hr>



