#! /usr/local/php5/bin/php
<?php
class MyClass
{
    public $var1 = 'value 1';
    public $var2 = 'value 2';
    public $var3 = 'value 3';

    protected $protected = 'protected var';
    private   $private   = 'private var';

    function iterateVisible() {
       echo "MyClass::iterateVisible:\n";
       foreach($this as $key => $value) {
           print "$key => $value".PHP_EOL;
       }
    }
}

$class = new MyClass();
foreach($class as $key => $value) {
    print "$key => $value".PHP_EOL;
}
echo "".PHP_EOL;
$class->iterateVisible();
