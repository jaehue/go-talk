#! /usr/local/php5/bin/php
<?php

echo "Example #1 A simple array".PHP_EOL;
$array = array(
    "foo" => "bar",
    "bar" => "foo",
);

var_dump($array);

echo "<br><br>Example #2 Type Casting and Overwriting example".PHP_EOL;

$array = array(
    1    => "a",
    "1"  => "b",
    1.5  => "c",
    true => "d",
);
var_dump($array);

echo "<br><br>Example #3 Mixed integer and string keys".PHP_EOL;

$array = array(
    "foo" => "bar",
    "bar" => "foo",
    100   => -100,
    -100  => 100,
);
var_dump($array);

echo "<br><br>Example #4 Indexed arrays without key".PHP_EOL;

$array = array("foo", "bar", "hello", "world");
var_dump($array);

echo "<br><br>Example #5 Keys not on all elements".PHP_EOL;

$array = array(
         "a",
         "b",
    6 => "c",
         "d",
);
var_dump($array);

echo "<br><br>Example #6 Accessing array elements".PHP_EOL;

$array = array(
    "foo" => "bar",
    42    => 24,
    "multi" => array(
         "dimensional" => array(
             "array" => "foo"
         )
    )
);

var_dump($array["foo"]);
var_dump($array[42]);
var_dump($array["multi"]["dimensional"]["array"]);

echo "<br><br>Example #7 Array dereferencing".PHP_EOL;

function getArray() {
    return array(1, 2, 3);
}

// on PHP 5.4
$secondElement = getArray()[1];

// previously
$tmp = getArray();
$secondElement = $tmp[1];

// or
list(, $secondElement) = getArray();

?>

