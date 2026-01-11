
# Roomy Feature List

## User Stories

Entities:
- User
- Organization (the organization providing the application)

ID | Description
--- | ---
1 | As a user, I want to be able to login via GitLab to reduce my amount of credentials to remember.
2 | As the organization, I want my members to use a GitLab login, such that I dont have to manage any other authentication systems (i.e. using OAuth 2.0 flow).
3 | As the organization, I want to have a fine-grained permission/authorization system, such that I can grant/deny specific permissions to my employees and associated members (which are not users).
4 | As a user, I want to have an overview of all rooms available (that I should be able to see), such that I can find an available room/place.
5 | As a user, I want the overview to be a room plan of, for example one building floor, and as detailed as possible, such that I can find free places less time consuming. 
6 | As the organization, I want the structure to have places as the smallest entity (maybe a table/desk, a chair, seats, group of multiple seats), above these I want to be able to define an area, which groups multiple places and as the greatest entity a building, which can consists of multiple areas, such that I can manage a high amount of different places easily.
7 | As the organization, I want my members to be able to configure a location, a name, the amount of seats available, a description, whether or not it is bookable using the self-service reservation- or manual method, if a place is disabled, if a place is disabled for self-service reservation within a specific time-frame (maybe recurring weekly), of a place, such that a place can be configured to the way I want to use it for.
8 | As the organization, I want my users to be able to self-service register for specific places using a reservation overview, such that assigning a place is less work-intensive for me.
9 | As a user, I want to self-service register on a website of a place where I can select the date, time and duration of my reservation, such that I can be sure that my place is reserved.
10 | As the organization, I want my users to be required to check-in using a physical QR-code attached to a place or using a manual check-in in application for the specific time-slot. 
11 | As the organization, I want my members to be able to configure how many reservations a user can reserve in parallel (i.e. how many reservations should intersect at most for the same date and time (and duration)), such that a user is limited on how many reservation can be reserved at the same point in time.
12 | As the organization, I want my members to be able to manually reserve places for specific users (for example by their email or id), such that there is a backup mechanism if the self-service does not work for any reason.
13 | As the organization, I want my members to be able to configure "room plans" (for areas) that are shown to users, such that selecting a place is more user-friendly for users.
14 | As a user, I want to be able to see a "room plan" for an area, because that helps me identifying which place I want to reserve and where it is. 
15 | As a the organization, I want my members to be able to create a room plan using an editor, such that I can, for example upload an image of a room plan, and mark areas for places, such that a room plan can be displayed more helpful to a user.
16 | As the organization, I want my members to be able to configure a name, description, location and the areas contained of a specific building, such that I can modify a building later on.
17 | As the organization, I want my members to be able to create a building entity, such that it does not have to be configured hard-coded.
18 | As the organization, I want my members to be able to create an area of a building entity, such that it does not have to be configured hard-coded.
19 | As the organization, I want my members to be able to configure a name, a description, an location and the places of an area. 
20 | As the organization, I want reservations of places to be freed up after a configurable duration per place, such that other users can get the chance to reserve the given place as well if it is not used.
21 | As a user, I want to get notified via email if my reservation was cancelled due to decay or because of manual cancellation through a organization member, such that I know whether or not a reservation still exists.
22 | As a user, I want to be able to configure for which kind of events I receive email notifications, such that I can select over which channels I want to get notified.
23 | As the organization, I do not want user details of reservations to be shown to other users due to privacy/compliance reasons.
24 | As the organization, I want to be able to see user details of reservations to be shown to organization members, such that a booking can be verified.
25 | As an organization member, I want to be able to create a reservation QR-code for a place that shows the QR-code, the name, area name and building name, such that a place can be easily identified.
26 | As an organization member, I want to be able to select which details (name, area name, building name) should be displayed on the QR-code print-out, such that it is not fixed what is shown on the print-out.
27 | As an organization member, I want to be able to print the QR-code (including details name, area name, building name), such that I can physically attach the QR-code to places.
28 | As an organization member, I want to be able to save the QR-code (including details name, area name, building name) as .jpg, .png, .pdf, such that I can further modify the output.
29 | As an organization member, I want to be able to configure the QR-code output using a HTML-template that later can be converted to an image, PDF file or be printed, such that it is more efficient to use uniform QR-code designs.
30 | As an organization member, I want to have (anonymous) statistics about the current place usage (per place/area/building), how many reservations? how many check-ins (currently)? how many places/areas? as well as metrics over time (i.e. weekly, monthly, yearly usage statistics), such that I can detect which places/areas/buildings are used more frequently and which are not.
