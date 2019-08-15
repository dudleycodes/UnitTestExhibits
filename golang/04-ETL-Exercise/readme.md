# GO Unit Test Exercise - Mock ETL App

In computing, **extract, transform, load (ETL)** ([wikipedia](https://en.wikipedia.org/wiki/Extract,_transform,_load)) is the general procedure of copying data from one or more sources into a destination system which represents that data differently from the source(s) or in a different context than the source(s). *Data extraction* involves extracting data from external sources; *data transformation* process data by cleansing, validating, and transforming data into new structure for storage; finally, *data loading* describes the insertion of the transformed data into the final target database.

This directory includes a mock ETL application that needs to be unit tested. Since no origin or destination database exists the exercise focuses on refactoring and then writing unit tests for the transformation processes.  The application mocks getting an appetizer, salad, and entree from an origin database, transforming them into a meal, and then sending the meal to the destination database.

```shell
MASCHEEN:01-unit user$ go run .

Getting appetizer from origin database....
Got: Cheesey Bread

Getting salad from origin database...
Got: Summer Asian Slaw

Getting entree from origin database...
Got: Butter-Roasted Rib Eye Steak

Sending courses to destination database...
Sent: "Butter-Roasted Rib Eye Steak" with "Summer Asian Slaw" salad and a starter of "Cheesey Bread".
```

## Exercise

1. Refactor the code; move any transformation logic into its own function so that it can be unit tested.

hint: Write and group tests around business logic; not around coding logic.  For example one test function could cover `hasCheese`, another `hasFruit`, another `hasFruit`, and so on.
