#! /usr/local/php5/bin/php
<?php
// START OMIT
$juice = "apple";

echo "He drank some $juice juice.".PHP_EOL;

echo "He drank some juice made of $juices.";

$juices = array("apple", "orange", "koolaid1" => "purple");

echo "He drank some $juices[0] juice.".PHP_EOL.PHP_EOL;
echo "He drank some $juices[1] juice.".PHP_EOL.PHP_EOL;
echo "He drank some $juices[koolaid1] juice.".PHP_EOL.PHP_EOL;
// END OMIT
?>