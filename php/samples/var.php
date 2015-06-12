#! /usr/local/php5/bin/php
<?php
$a = 'hello';
$$a = 'world';
echo "$a ${$a}".PHP_EOL;
echo "$a $hello".PHP_EOL;

$a = "foo";
$foo = "bar";
echo $$a;
?>