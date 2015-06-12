#! /usr/local/php5/bin/php
<?php
// START OMIT
class people {
    public $john = "John Smith";
    public $jane = "Jane Smith";
    public $robert = "Robert Paulsen";
    
    public $smith = "Smith";
}

$people = new people();

echo "$people->john then said hello to $people->jane.".PHP_EOL.PHP_EOL;
echo "$people->john's wife greeted $people->robert.".PHP_EOL.PHP_EOL;
echo "$people->robert greeted the two $people->smiths.";
// END OMIT
?>