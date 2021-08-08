Builder pattern lets you create complex objects step-by-step. This patterns allows you to produce different types and representation of an object using same construction code.

This pattern is very useful if there are lot of things needs to be done before builging an object. On the other hand factory pattern return entire instance at one go. Simple example for builder pattern is when buidling DOM elements in HTML, there will be many nodes and each one will have different properties and all tied together at the end, this is perfect example for the builder pattern.

Also, ORM example fits well in this case, where it lets you keep adding the condition, keeps on adding before running the query to return the result.
