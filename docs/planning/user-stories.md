
# Roomy User Stories

Entities:
- User
- Organization (the organization providing the application)
- Organization member (member with more permissions than a user that are part of the organization)

## Availability Model

The system uses a **blocking-based availability model**:
- **Default state**: Everything is available 24/7
- **Blocking periods**: Define when something is NOT available (closures, holidays, maintenance, etc.)
- **Availability** = All time − Blocking periods − Existing reservations
- **Inheritance**: Blocking periods cascade from Building → Area → Place

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
31 | As the organization, I want the list of available rooms to be public, such that users can see which places are reserved without logging in.
32 | As a user and an organization member, I want to have a calendar view for a place and an area of reserved or blocked times, such that it is easily identifyable where free slots are/can be reserved.
33 | As an organization member, I want to be able to configure whether or not a place needs a check-in by the user or not, because this maybe not needed for all places.
34 | As an organization member, I want to be able to configure a place such that a user can not reserve more than X hours per week/day/month/year, because there are shared places which are used by many people.
35 | As an organization member, I want to be able to configure a place to have an optional whitelist of users to reserve a place, because some places have the constraint that users have to fill out a formular first.
36 | As an organization member, I want to be able to configure the optional maximum duration time of a reservation of a place, such that reservation durations can be limited but not must be.
37 | As an organization member, I want to be able to configure the optional amount of how many reservations (count, not duration) can be done per week/day/month/year.
38 | As an organization member, I want to be able to delete places/areas/buildings, such that I can remove old places/areas/buildings.
39 | As a user, I want to have a reservation that only provides all necessary information and fields for, such that I do not waste much time when reserving a place.
40 | As a user, I want to have a search, such that I can find specific places/areas/buildings that I can reserve a slot in.
41 | As the organization, I want users, which are not whitelisted for a place, to not be able to register for reservations, because regulations may require it.
42 | As a user, I want to be able to cancel my own reservation, such that I can free up the place for others if my plans change.
43 | As a user, I want to see my upcoming and past reservations in a personal dashboard, such that I have an overview of my bookings.
44 | As a user, I want to be able to extend or shorten my current reservation (if the time slot allows), such that I can adjust my booking without having to cancel and re-book.
45 | As a user, I want to be able to set a favorite place, such that I can quickly reserve my preferred spot.
46 | As an organization member, I want to be able to add blocking periods to a place/area/building for a specific time range (e.g., for maintenance, events, or closures), such that users cannot reserve during that period.
47 | As an organization member, I want to be able to add a reason/note when adding a blocking period, such that users understand why the place is unavailable.
48 | As a user, I want to see why a place is blocked (if a reason was provided) and which entity (building, area, or place) the blocking comes from, such that I know if and when it might become available again.
49 | As the organization, I want to be able to configure recurring blocking periods (e.g., closed every day 18:00-08:00, weekends, every Monday for cleaning), such that regular closures are automated.
50 | As an organization member, I want to be able to export reservation data (e.g., as CSV), such that I can analyze or archive the data externally.
51 | As an organization member, I want to be able to see a log/history of changes made to places/areas/buildings, such that I can audit who changed what and when.
52 | As the organization, I want organization members to have different permission levels (e.g., viewer, editor, admin), such that not all members can modify all settings.
53 | As a user, I want to receive a reminder notification before my reservation starts (configurable time, e.g., 15 minutes before), such that I do not forget my booking.
54 | As a user, I want to be able to configure how long before a reservation I want to receive a reminder, such that I can adjust it to my preferences.
55 | As an organization member, I want to be able to add blocking periods to a building (e.g., closed hours 18:00-08:00 daily, weekends, holidays), such that these blocking periods automatically apply to all areas and places within that building.
56 | As a user, I want to see the blocking periods (closures, holidays) of a building/area/place, such that I know when I can use the place.
57 | As a user, I want to be able to book a recurring reservation (e.g., every Tuesday from 10:00-12:00 for the next 4 weeks), such that I do not have to book each slot individually.
58 | As the organization, I want to be able to limit how far in advance a user can make a reservation (e.g., max 2 weeks ahead), such that places are not blocked too far in the future.
59 | As an organization member, I want to be able to configure a minimum booking duration for a place, such that very short reservations (e.g., 5 minutes) are prevented.
60 | As the organization, I want to be able to configure whether a check-in must happen within a certain time after the reservation starts (e.g., 15 minutes), otherwise the reservation is automatically cancelled, such that no-shows free up places quickly.
61 | As a user, I want to be able to add equipment requirements to my reservation (e.g., projector, whiteboard), such that I can filter for places that have the equipment I need.
62 | As an organization member, I want to be able to define available equipment/amenities for a place, such that users can filter and search by equipment.
63 | As a user, I want to be able to filter places by available equipment/amenities, such that I can find a place that fits my needs.
64 | As an organization member, I want to be able to set a capacity (number of people) for a place, such that users know how many people can use it.
65 | As a user, I want to be able to filter places by minimum capacity, such that I can find a place suitable for my group size.
66 | As a user, I want to be able to share my reservation details (date, time, place, building) with others via a link or email, such that I can inform colleagues where to meet.
67 | As the organization, I want users to only be able to book places if their account is active/enabled, such that former employees cannot make reservations.
68 | As an organization member, I want to be able to disable a user account without deleting their reservation history, such that compliance and audit requirements are met.
69 | As the organization, I want an API to be available for integrating the reservation system with other internal systems (e.g., calendar sync, access control), such that the system can be extended.
70 | As the organization, I it necessary that the API is protected by a simple API key, such that not everyone can use the API.
71 | As an organization member, I want to be able to create an API key for myself, such that I can connect the API to other services.
72 | As the organization, I want the API keys to provide the same authorization permissions as the user that created them inherits, such that API keys are already authorized as well.
73 | As an organization member, I want to be able to create permission groups, such that I can define reusable sets of permissions for users.
74 | As an organization member, I want to be able to assign permissions to a group, such that all users in that group inherit those permissions.
75 | As an organization member, I want to be able to assign one or more groups to a user, such that the user receives the combined permissions of all assigned groups.
76 | As an organization member, I want to be able to remove groups from a user, such that I can revoke permissions when needed.
77 | As the organization, I want a "system" group to exist by default that has all available permissions and cannot be deleted or edited, such that there is always a superadmin group.
78 | As the organization, I want a "default" group to exist by default that can be edited but not deleted, such that newly registered users have a baseline set of permissions.
79 | As an organization member, I want to be able to configure which group(s) are automatically assigned to users when they first register, such that new users have appropriate initial permissions.
80 | As an organization member, I want to see an overview of all users, such that I can find and manage user accounts.
81 | As an organization member, I want to see a user detail view showing their OAuth 2.0 details (email, username, name, profile picture, etc.), such that I can identify and verify users.
82 | As an organization member, I want to be able to edit the groups assigned to a user in the user detail view, such that I can manage user permissions efficiently.
83 | As an organization member, I want to see an overview of all groups, such that I can find and manage permission groups.
84 | As an organization member, I want to see a group detail view where I can view and edit the permissions assigned to that group, such that I can configure group permissions in one place.
85 | As an organization member, I want to be able to delete custom groups (but not "system" or "default"), such that I can remove groups that are no longer needed.
86 | As an organization member, I want to see which users are assigned to a group in the group detail view, such that I understand who is affected by permission changes.
87 | As an organization member, I want to be able to configure time intervals/slots (e.g., 15-minute increments starting at 09:00, 09:15, 09:30) for a place, such that reservations can only start at defined times.
88 | As a user, I want to see available time slots based on the configured intervals when making a reservation (filtering out blocked times), such that I can only book at valid start times.
89 | As a user, I want to subscribe to a public .ics calendar URL for a place or area, such that I can see upcoming reservations and blocked times in my personal calendar application.
90 | As a user, I want shared reservation links to display a rich preview (title, date, time, location) when pasted into chat applications like WhatsApp, Slack, or Teams, such that recipients can quickly understand the reservation details.
91 | As an organization member, I want to be able to delete a room plan from an area, such that I can remove outdated or incorrect room plans.
92 | As an organization member, I want blocking periods defined at the building level to automatically cascade to all areas and places within that building, such that I don't have to configure the same closures multiple times.
93 | As an organization member, I want blocking periods defined at the area level to automatically cascade to all places within that area, such that area-wide closures are easily managed.
94 | As a user, I want to see when a blocking period is inherited from a parent building or area (vs. defined on the place itself), such that I understand the source of the unavailability.
