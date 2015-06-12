#! /usr/local/php5/bin/php
<?php
// START OMIT
$array = array(
    "foo" => "bar",
    "bar" => "foo",
    100   => -100,
    -100  => 100,
);
// END OMIT
var_dump($array);
?>

