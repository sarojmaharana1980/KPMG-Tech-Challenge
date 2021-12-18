# KPMG-Tech-Challenge
(
## Challenge 2
We need to write code that will query the meta data of an instance within AWS and 
provide a json formatted output. The choice of language and implementation is up to you.

###Approach to this challenge

Introduced a 'go' script which will fetch the metadata of a EC2 instance from AWS
and returned data is presented json format

Go language is prefered for this challenge as it provides  faster execution and easy to learn

For whole meta data of EC2 instance run 
````
go run metaData.go
````
Argument to be passed (ami-id) while running the script.
````
go run metaData.go ami-id