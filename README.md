# GO Exercise - check if older products are still being sold

## How to test the application

#### From the command line (Windows):

Inside the application folder:

<pre>
>go run main.go "2018-01-01 00:00:00" "2019-08-01 00:00:00"
</pre>
or
<pre>
>go run main.go "2018-01-01 00:00:00" "2019-08-01 00:00:00" "(1-3, 4-6, 7-12, >12)"
</pre>

#### From IntelliJ's:
* Run menu, select Edit Configurations...
* In the Program Arguments insert: "2018-01-01 00:00:00" "2019-01-01 00:00:00" "(1-3, 4-6, 7-12, >12)"
* Ok
* Run 'Main'