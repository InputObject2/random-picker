# Random Item Selector Web App

## Overview
This is a simple web application to pick from a list of values. It allows users to input items into a list and then select one at random. 

The web server runs on port 8080.

Note that there is no session management implemented, so refreshing the page will clear all entries.

## Features
- **Add Items**: Users can add items using a text input box.
- **Random Selection**: Click a button to highlight a random item from the list.

## Running with Docker

You can run the container and it will :

```shell
docker run -p 8080:8080 inputobject2/random-picker
```