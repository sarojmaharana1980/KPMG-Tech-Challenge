# KPMG-Tech-Challenge

## Challenge 3
We have a nested object, we would like a function that you pass in the object and a key and get back the value. How this is implemented is up to you.
Example Inputs
object = {“a”:{“b”:{“c”:”d”}}}
key = a/b/c
object = {“x”:{“y”:{“z”:”a”}}}
key = x/y/z
value = a

## Approach

Added python script which will iterate through the dictionary  value and extract the key and and values

Arguments will be passed while running the python script.

````
python json-object.py -d "{'x':{'y':{'z':'a'}}}" -k "x/y/z"
````

The output of above given input arguments is
````
Keys:['x', 'y', 'z']
Values:a
````


