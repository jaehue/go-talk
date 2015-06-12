#! /usr/local/php5/bin/php
<?php
$foo = 10;   // $foo is an integer
$bar = (boolean) $foo;   // $bar is a boolean

echo $foo.PHP_EOL;
echo $bar.PHP_EOL;

$foo = 10;            // $foo is an integer
$str = "$foo";        // $str is a string
$fst = (string) $foo; // $fst is also a string

echo $foo.PHP_EOL;
echo $str.PHP_EOL;
echo $fst.PHP_EOL;
?>