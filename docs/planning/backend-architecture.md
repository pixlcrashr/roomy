# Backend Architecture

This document outlines the API structure and OpenAPI definitions for Roomy based on the user stories.

## Technology Considerations

- RESTful API design
- OAuth 2.0 GitLab integration
- API key authentication for external integrations
- Fine-grained permission/authorization system
- PostgreSQL database (based on existing go.mod)
- Go backend with GORM
- Automatically generated code for frontend and backend

---

## Unified Availability Model

### Core Concept: Blocking-Based Availability

The system uses a **blocking-based availability model** where:

1. **Default state**: Everything is available 24/7
2. **Blocking periods**: Define when something is NOT available
3. **Availability** = All time − Blocking periods − Existing reservations

This unified approach simplifies availability calculations across all entity levels.

### Blocking Inheritance

Blocking periods cascade down the hierarchy:
- **Building blocking** → Applies to all areas and places in the building
- **Area blocking** → Applies to all places in the area
- **Place blocking** → Applies only to that specific place

When calculating availability for a place, the system merges blocking from:
1. The place's own blocking periods
2. The parent area's blocking periods
3. The parent building's blocking periods

### Blocking Types

All blocking periods use the same unified `Blocking` model:

| Type | Example | Recurring |
|------|---------|-----------|
| `closedHours` | Closed 18:00-09:00 daily | Yes (RRULE) |
| `weekend` | Closed Saturdays and Sundays | Yes (RRULE) |
| `holiday` | Christmas Day | Yes (yearly) or one-time |
| `maintenance` | Cleaning every Monday 07:00-09:00 | Yes (RRULE) |
| `event` | Company event on specific date | One-time |
| `disabled` | Place temporarily out of service | One-time |

### Time Slot Intervals

The `timeSlotConfig` on places defines the **booking grid** (when reservations can start/end), NOT availability. For example:
- Interval: 15 minutes
- Valid start times: 09:00, 09:15, 09:30, ...

This is independent of blocking. A slot is bookable if:
1. It falls on a valid interval time
2. It is not blocked (directly or inherited)
3. It is not already reserved

---

## Authentication & Authorization

### Authentication Methods

1. **OAuth 2.0 (GitLab)** - Primary user authentication (US-1, US-2)
2. **API Key** - For external integrations (US-69, US-70, US-71, US-72)

### Authorization

- Permission-based access control via groups (US-3, US-73-86)
- API keys inherit permissions from creating user (US-72)
- Default groups: `system` (immutable, all permissions), `default` (editable baseline)

---

## API Overview

Base path: `/api/v1`

### Authentication Endpoints

```yaml
/api/v1/auth:
  /login:
    GET:
      summary: Initiate GitLab OAuth flow
      response: Redirect to GitLab authorization URL

  /callback:
    GET:
      summary: OAuth callback handler
      parameters:
        - code: Authorization code from GitLab
        - state: CSRF protection state
      response: JWT tokens + user info

  /refresh:
    POST:
      summary: Refresh access token
      body: { refreshToken: string }
      response: New access/refresh tokens

  /logout:
    POST:
      summary: Invalidate current session
      response: 204 No Content

  /me:
    GET:
      summary: Get current user profile
      response: User object with permissions
```

---

### Buildings

```yaml
/api/v1/buildings:
  GET:
    summary: List all buildings
    parameters:
      - page: int (pagination)
      - limit: int
      - search: string (name search)
    response: PaginatedList<Building>
    permissions: public (US-31)

  POST:
    summary: Create a new building
    body: CreateBuildingRequest
    response: Building
    permissions: manage:buildings (US-17)

/api/v1/buildings/{buildingId}:
  GET:
    summary: Get building details
    response: Building (with areas summary)
    permissions: public

  PUT:
    summary: Update building
    body: UpdateBuildingRequest
    response: Building
    permissions: manage:buildings (US-16)

  DELETE:
    summary: Delete building
    response: 204 No Content
    permissions: manage:buildings (US-38)

/api/v1/buildings/{buildingId}/blocking:
  # Unified blocking endpoint - replaces separate hours/holidays endpoints (US-55, US-56)
  # Blocking periods define when the building is NOT available
  # All areas and places in this building inherit these blocking periods
  GET:
    summary: Get building blocking periods
    description: |
      Returns all blocking periods for this building.
      These blocking periods cascade to all areas and places within the building.
    parameters:
      - includeRecurring: boolean (expand recurring rules for date range)
      - startDate: date (for expanding recurring)
      - endDate: date (for expanding recurring)
    response: Blocking[]
    permissions: public (US-56)

  PUT:
    summary: Replace all building blocking periods
    body: { blockings: Blocking[] }
    response: Blocking[]
    permissions: manage:buildings (US-55)

/api/v1/buildings/{buildingId}/blocking/entries:
  POST:
    summary: Add blocking periods to building
    body: { entries: CreateBlockingRequest[] }
    response: Blocking[]
    permissions: manage:buildings

  DELETE:
    summary: Remove blocking periods by IDs
    body: { blockingIds: uuid[] }
    response: 204 No Content
    permissions: manage:buildings

/api/v1/buildings/{buildingId}/areas:
  GET:
    summary: List areas in building
    response: Area[]
    permissions: public

/api/v1/buildings/{buildingId}/availability:
  # Computed availability based on blocking (inverse of blocking)
  GET:
    summary: Get available time slots for building
    description: |
      Returns time slots that are NOT blocked for the building.
      Useful for seeing when the building is open/operational.
    parameters:
      - startDate: date (required)
      - endDate: date (required)
    response: TimeRange[] (available periods)
    permissions: public

/api/v1/buildings/{buildingId}/calendar.ics:
  GET:
    summary: Public iCalendar feed for building (next 3 months)
    description: Returns .ics file with blocking periods and reservations for all places
    response: text/calendar
    permissions: public (US-89)
```

---

### Areas

```yaml
/api/v1/areas:
  GET:
    summary: List all areas (with filters)
    parameters:
      - buildingId: uuid (optional filter)
      - page: int
      - limit: int
    response: PaginatedList<Area>
    permissions: public

  POST:
    summary: Create a new area
    body: CreateAreaRequest
    response: Area
    permissions: manage:areas (US-18)

/api/v1/areas/{areaId}:
  GET:
    summary: Get area details
    response: Area (with places summary)
    permissions: public

  PUT:
    summary: Update area
    body: UpdateAreaRequest
    response: Area
    permissions: manage:areas (US-19)

  DELETE:
    summary: Delete area
    response: 204 No Content
    permissions: manage:areas (US-38)

/api/v1/areas/{areaId}/roomPlan:
  GET:
    summary: Get room plan data
    response: RoomPlan (image URL, place markers)
    permissions: public (US-14)

  PUT:
    summary: Update room plan
    body: RoomPlanData (multipart with image)
    response: RoomPlan
    permissions: manage:areas (US-13, US-15)

  DELETE:
    summary: Delete room plan
    response: 204 No Content
    permissions: manage:areas (US-91)

/api/v1/areas/{areaId}/roomPlan/markers:
  POST:
    summary: Add place markers to room plan
    body: { markers: PlaceMarker[] }
    response: PlaceMarker[]
    permissions: manage:areas

  DELETE:
    summary: Remove place markers by IDs
    body: { markerIds: uuid[] }
    response: 204 No Content
    permissions: manage:areas

/api/v1/areas/{areaId}/blocking:
  # Unified blocking endpoint for areas
  # These blocking periods cascade to all places within the area
  # Area also inherits blocking from parent building
  GET:
    summary: Get area blocking periods
    description: |
      Returns blocking periods defined at the area level.
      Use includeInherited=true to also see blocking inherited from the building.
    parameters:
      - includeInherited: boolean (include building blocking)
      - includeRecurring: boolean (expand recurring rules)
      - startDate: date
      - endDate: date
    response: Blocking[] (with source indicator if inherited)
    permissions: public (US-48)

  PUT:
    summary: Replace all area blocking periods
    body: { blockings: Blocking[] }
    response: Blocking[]
    permissions: manage:areas (US-46, US-47)

/api/v1/areas/{areaId}/blocking/entries:
  POST:
    summary: Add blocking periods to area
    body: { entries: CreateBlockingRequest[] }
    response: Blocking[]
    permissions: manage:areas

  DELETE:
    summary: Remove blocking periods by IDs
    body: { blockingIds: uuid[] }
    response: 204 No Content
    permissions: manage:areas

/api/v1/areas/{areaId}/places:
  GET:
    summary: List places in area
    response: Place[]
    permissions: public

/api/v1/areas/{areaId}/availability:
  # Aggregated availability view for all places in an area (US-4, US-32)
  GET:
    summary: Get availability summary for all places in area
    description: |
      Returns availability status for each place in the area within the given time range.
      Availability is calculated as: time slots NOT blocked (from place, area, or building)
      AND NOT reserved. Optimized for room plan visualization.
    parameters:
      - date: date (required)
      - startTime: time (optional)
      - endTime: time (optional)
    response: PlaceAvailability[]
    permissions: public

/api/v1/areas/{areaId}/calendar.ics:
  GET:
    summary: Public iCalendar feed for area (next 3 months)
    description: Returns .ics file with blocking periods and reservations for all places in area
    response: text/calendar
    permissions: public (US-89)
```

---

### Places

```yaml
/api/v1/places:
  GET:
    summary: List/search places
    parameters:
      - areaId: uuid (optional)
      - buildingId: uuid (optional)
      - search: string
      - equipment: string[] (filter by equipment)
      - minCapacity: int (US-65)
      - available: boolean
      - date: date
      - startTime: time
      - endTime: time
      - page: int
      - limit: int
    response: PaginatedList<Place>
    permissions: public (US-40, US-63)

  POST:
    summary: Create a new place
    body: CreatePlaceRequest
    response: Place
    permissions: manage:places (US-6)

/api/v1/places/{placeId}:
  GET:
    summary: Get place details
    response: Place (with equipment, constraints)
    permissions: public

  PUT:
    summary: Update place
    body: UpdatePlaceRequest
    response: Place
    permissions: manage:places (US-7)

  DELETE:
    summary: Delete place
    response: 204 No Content
    permissions: manage:places (US-38)

/api/v1/places/{placeId}/constraints:
  GET:
    summary: Get place constraints
    response: PlaceConstraints
    permissions: public

  PUT:
    summary: Update place constraints
    body: PlaceConstraints
    response: PlaceConstraints
    permissions: manage:places
    # Covers: US-11, US-20, US-33, US-34, US-36, US-58, US-59, US-60

/api/v1/places/{placeId}/timeSlots:
  # Configurable time intervals for reservations (US-87, US-88)
  # This defines the BOOKING GRID (valid start/end times), not availability
  # Availability is determined by: not blocked AND not reserved
  GET:
    summary: Get configured time slot intervals
    description: |
      Returns the booking grid configuration for this place.
      Reservations must start and end on these interval boundaries.
      This does NOT define availability - blocking periods define unavailability.
    response: TimeSlotConfig
    permissions: public

  PUT:
    summary: Update time slot configuration
    body: TimeSlotConfig
    response: TimeSlotConfig
    permissions: manage:places

/api/v1/places/{placeId}/equipment:
  GET:
    summary: Get place equipment/amenities
    response: Equipment[]
    permissions: public (US-61)

  POST:
    summary: Add equipment to place
    body: { equipmentIds: uuid[] }
    response: Equipment[]
    permissions: manage:places (US-62)

  DELETE:
    summary: Remove equipment from place
    body: { equipmentIds: uuid[] }
    response: 204 No Content
    permissions: manage:places

/api/v1/places/{placeId}/whitelist:
  GET:
    summary: Get user whitelist
    response: UserReference[]
    permissions: manage:places

  POST:
    summary: Add users to whitelist
    body: { userIds: uuid[] }
    response: UserReference[]
    permissions: manage:places (US-35, US-41)

  DELETE:
    summary: Remove users from whitelist
    body: { userIds: uuid[] }
    response: 204 No Content
    permissions: manage:places

/api/v1/places/{placeId}/blocking:
  # Unified blocking endpoint for places
  # Place also inherits blocking from parent area and building
  GET:
    summary: Get place blocking periods
    description: |
      Returns blocking periods for this place.
      Use includeInherited=true to see blocking inherited from area and building.
      Blocking = times when the place is NOT available.
    parameters:
      - includeInherited: boolean (include area and building blocking)
      - includeRecurring: boolean (expand recurring rules)
      - startDate: date
      - endDate: date
    response: Blocking[] (with source indicator if inherited)
    permissions: public (US-48)

  PUT:
    summary: Replace all place blocking periods
    body: { blockings: Blocking[] }
    response: Blocking[]
    permissions: manage:places (US-46, US-47, US-49)

/api/v1/places/{placeId}/blocking/entries:
  POST:
    summary: Add blocking periods to place
    body: { entries: CreateBlockingRequest[] }
    response: Blocking[]
    permissions: manage:places

  DELETE:
    summary: Remove blocking periods by IDs
    body: { blockingIds: uuid[] }
    response: 204 No Content
    permissions: manage:places

/api/v1/places/{placeId}/availability:
  # Computed availability = NOT blocked AND NOT reserved
  GET:
    summary: Get availability calendar with time slots
    description: |
      Returns available time slots for this place.
      Availability is computed as: All valid time slot intervals
      MINUS blocking periods (own + inherited from area + building)
      MINUS existing reservations.
    parameters:
      - startDate: date
      - endDate: date
    response: AvailabilitySlot[] (status: available/blocked/reserved)
    permissions: public (US-32, US-88)

/api/v1/places/{placeId}/qrCode:
  GET:
    summary: Generate QR code
    parameters:
      - templateId: uuid (optional)
      - format: enum (png, jpg, pdf)
      - includeDetails: string[] (name, areaName, buildingName)
    response: Binary file or base64
    permissions: manage:places (US-25, US-26, US-27, US-28)

/api/v1/places/{placeId}/calendar.ics:
  GET:
    summary: Public iCalendar feed for place (next 3 months)
    description: Returns .ics file with blocking periods and reservations
    response: text/calendar
    permissions: public (US-89)
```

---

### Reservations

```yaml
/api/v1/reservations:
  GET:
    summary: List reservations
    parameters:
      - userId: uuid (filter, admin only)
      - placeId: uuid (filter)
      - areaId: uuid (filter)
      - buildingId: uuid (filter)
      - status: enum (upcoming, active, past, cancelled)
      - startDate: date
      - endDate: date
      - page: int
      - limit: int
    response: PaginatedList<Reservation>
    permissions: own reservations OR manage:reservations (US-43, US-24)

  POST:
    summary: Create reservation (single or recurring)
    description: |
      Creates a single reservation or multiple recurring reservations.
      For recurring reservations, provide the recurrence field with pattern details.
      Start/end times must align with the place's configured time slot intervals.
      Reservation will fail if any time slot is blocked or already reserved.
    body: CreateReservationRequest
    response: Reservation | Reservation[] (array if recurring)
    permissions: authenticated (US-8, US-9, US-57) OR manage:reservations (US-12)

/api/v1/reservations/{reservationId}:
  GET:
    summary: Get reservation details
    response: Reservation
    permissions: owner OR manage:reservations

  PUT:
    summary: Update reservation (extend or shorten duration)
    description: |
      Allows modifying the reservation's start and/or end time.
      Extending is only allowed if the adjacent time slot is not blocked or reserved.
      Shortening is always allowed within minimum duration constraints.
      Times must align with the place's configured time slot intervals.
    body: UpdateReservationRequest { startTime?: datetime, endTime?: datetime }
    response: Reservation
    permissions: owner (US-44) OR manage:reservations

  DELETE:
    summary: Cancel reservation
    response: 204 No Content
    permissions: owner (US-42) OR manage:reservations

/api/v1/reservations/{reservationId}/checkIn:
  POST:
    summary: Check in to reservation
    body: { qrCode: string (optional) }
    response: Reservation
    permissions: owner (US-10)

/api/v1/reservations/{reservationId}/share:
  GET:
    summary: Get shareable reservation link with preview metadata
    description: |
      Returns a shareable URL and Open Graph / Twitter Card metadata
      for rich link previews in chat applications (US-90)
    response:
      shareUrl: string
      previewMetadata:
        title: string
        description: string
        image: string (optional)
        siteName: string
    permissions: owner (US-66)

/api/v1/reservations/export:
  GET:
    summary: Export reservations as CSV
    parameters:
      - startDate: date
      - endDate: date
      - buildingId: uuid (optional)
      - areaId: uuid (optional)
    response: CSV file
    permissions: view:statistics (US-50)
```

---

### Users

```yaml
/api/v1/users:
  GET:
    summary: List all users
    parameters:
      - search: string (email, name)
      - groupIds: uuid[] (filter by multiple groups, comma-separated)
      - status: enum (active, disabled)
      - page: int
      - limit: int
    response: PaginatedList<User>
    permissions: view:users (US-80)

/api/v1/users/{userId}:
  GET:
    summary: Get user details (OAuth info, groups)
    response: User
    permissions: view:users (US-81)

  PUT:
    summary: Update user (groups, status)
    body: UpdateUserRequest
    response: User
    permissions: manage:users (US-82)

/api/v1/users/{userId}/groups:
  GET:
    summary: Get user's groups
    response: Group[]
    permissions: view:users

  POST:
    summary: Add groups to user
    body: { groupIds: uuid[] }
    response: Group[]
    permissions: manage:users (US-75)

  DELETE:
    summary: Remove groups from user
    body: { groupIds: uuid[] }
    response: 204 No Content
    permissions: manage:users (US-76)

/api/v1/users/{userId}/disable:
  POST:
    summary: Disable user account
    response: User
    permissions: manage:users (US-67, US-68)

/api/v1/users/{userId}/enable:
  POST:
    summary: Re-enable user account
    response: User
    permissions: manage:users

/api/v1/users/me/favorites:
  GET:
    summary: Get current user's favorite places
    response: Place[]
    permissions: authenticated (US-45)

  POST:
    summary: Add places to favorites
    body: { placeIds: uuid[] }
    response: 201 Created
    permissions: authenticated

  DELETE:
    summary: Remove places from favorites
    body: { placeIds: uuid[] }
    response: 204 No Content
    permissions: authenticated

/api/v1/users/me/notifications:
  GET:
    summary: Get notification preferences
    response: NotificationPreferences
    permissions: authenticated (US-22)

  PUT:
    summary: Update notification preferences
    body: NotificationPreferences
    response: NotificationPreferences
    permissions: authenticated (US-22, US-54)
```

---

### Groups & Permissions

```yaml
/api/v1/groups:
  GET:
    summary: List all permission groups
    response: Group[]
    permissions: view:groups (US-83)

  POST:
    summary: Create new group
    body: CreateGroupRequest
    response: Group
    permissions: manage:groups (US-73)

/api/v1/groups/{groupId}:
  GET:
    summary: Get group details (permissions, members)
    response: Group (with members list)
    permissions: view:groups (US-84, US-86)

  PUT:
    summary: Update group
    body: UpdateGroupRequest
    response: Group
    permissions: manage:groups (US-74)

  DELETE:
    summary: Delete group (not system/default)
    response: 204 No Content
    permissions: manage:groups (US-85)

/api/v1/groups/{groupId}/permissions:
  GET:
    summary: Get group permissions
    response: Permission[]
    permissions: view:groups

  POST:
    summary: Add permissions to group
    body: { permissions: string[] }
    response: Permission[]
    permissions: manage:groups (US-74, US-84)

  DELETE:
    summary: Remove permissions from group
    body: { permissions: string[] }
    response: 204 No Content
    permissions: manage:groups

/api/v1/groups/{groupId}/members:
  GET:
    summary: Get group members
    response: User[]
    permissions: view:groups (US-86)

/api/v1/groups/defaultAssignment:
  GET:
    summary: Get default groups for new users
    response: Group[]
    permissions: view:groups

  PUT:
    summary: Set default groups for new users
    body: { groupIds: uuid[] }
    response: Group[]
    permissions: manage:groups (US-79)

/api/v1/permissions:
  GET:
    summary: List all available permissions
    response: Permission[]
    permissions: view:groups
```

---

### Equipment/Amenities

```yaml
/api/v1/equipment:
  GET:
    summary: List all equipment types
    response: EquipmentType[]
    permissions: public

  POST:
    summary: Create equipment type
    body: CreateEquipmentTypeRequest
    response: EquipmentType
    permissions: manage:places

/api/v1/equipment/{equipmentId}:
  GET:
    summary: Get equipment type details
    response: EquipmentType
    permissions: public

  PUT:
    summary: Update equipment type
    body: UpdateEquipmentTypeRequest
    response: EquipmentType
    permissions: manage:places

  DELETE:
    summary: Delete equipment type
    response: 204 No Content
    permissions: manage:places
```

---

### Statistics

```yaml
/api/v1/statistics:
  GET:
    summary: Get overall statistics
    response: Statistics
    permissions: view:statistics (US-30)

/api/v1/statistics/usage:
  GET:
    summary: Get usage statistics over time
    parameters:
      - period: enum (daily, weekly, monthly, yearly)
      - startDate: date
      - endDate: date
      - buildingId: uuid (optional)
      - areaId: uuid (optional)
      - placeId: uuid (optional)
    response: UsageStatistics[]
    permissions: view:statistics (US-30)

/api/v1/statistics/current:
  GET:
    summary: Get current occupancy
    parameters:
      - buildingId: uuid (optional)
      - areaId: uuid (optional)
    response: CurrentOccupancy
    permissions: view:statistics (US-30)
```

---

### QR Code Templates

```yaml
/api/v1/qrTemplates:
  GET:
    summary: List QR code templates
    response: QRTemplate[]
    permissions: manage:places

  POST:
    summary: Create QR template
    body: CreateQRTemplateRequest
    response: QRTemplate
    permissions: manage:places (US-29)

/api/v1/qrTemplates/{templateId}:
  GET:
    summary: Get template details
    response: QRTemplate
    permissions: manage:places

  PUT:
    summary: Update template
    body: UpdateQRTemplateRequest
    response: QRTemplate
    permissions: manage:places

  DELETE:
    summary: Delete template
    response: 204 No Content
    permissions: manage:places

/api/v1/qrTemplates/{templateId}/preview:
  GET:
    summary: Preview template with sample data
    response: Rendered HTML/image
    permissions: manage:places
```

---

### API Keys

```yaml
/api/v1/apiKeys:
  GET:
    summary: List user's API keys
    response: APIKey[] (masked)
    permissions: authenticated (US-71)

  POST:
    summary: Create new API key
    body: CreateAPIKeyRequest
    response: APIKey (with full key, only shown once)
    permissions: authenticated (US-71)

/api/v1/apiKeys/{keyId}:
  DELETE:
    summary: Revoke API key
    response: 204 No Content
    permissions: owner
```

---

### Audit Log

```yaml
/api/v1/auditLog:
  GET:
    summary: Get audit log entries
    parameters:
      - entityType: enum (building, area, place, reservation, user, group)
      - entityId: uuid (optional)
      - userId: uuid (optional, who made the change)
      - action: enum (create, update, delete)
      - startDate: datetime
      - endDate: datetime
      - page: int
      - limit: int
    response: PaginatedList<AuditLogEntry>
    permissions: view:auditLog (US-51)
```

---

## Data Models

### Core Entities

```yaml
Building:
  id: uuid
  name: string
  description: string
  location: string
  createdAt: datetime
  updatedAt: datetime
  # Blocking periods fetched via /buildings/{id}/blocking

Area:
  id: uuid
  buildingId: uuid
  name: string
  description: string
  location: string
  roomPlan: RoomPlan (nullable)
  createdAt: datetime
  updatedAt: datetime
  # Blocking periods fetched via /areas/{id}/blocking (includes inherited)

Place:
  id: uuid
  areaId: uuid
  name: string
  description: string
  location: string
  capacity: int (US-64)
  isBookable: boolean
  bookingMethod: enum (selfService, manual)
  isDisabled: boolean
  requiresCheckIn: boolean (US-33)
  equipment: Equipment[]
  constraints: PlaceConstraints
  timeSlotConfig: TimeSlotConfig
  createdAt: datetime
  updatedAt: datetime
  # Blocking periods fetched via /places/{id}/blocking (includes inherited)

Reservation:
  id: uuid
  placeId: uuid
  userId: uuid
  startTime: datetime
  endTime: datetime
  status: enum (pending, confirmed, checkedIn, completed, cancelled, expired)
  checkInTime: datetime (nullable)
  cancelReason: string (nullable)
  isRecurring: boolean
  recurringGroupId: uuid (nullable)
  createdAt: datetime
  updatedAt: datetime

User:
  id: uuid
  email: string
  username: string
  name: string
  profilePicture: string (URL)
  oauthProvider: string (gitlab)
  oauthId: string
  isActive: boolean
  groups: Group[]
  createdAt: datetime
  updatedAt: datetime

Group:
  id: uuid
  name: string
  description: string
  isSystem: boolean (true for "system" group)
  isDefault: boolean (true for "default" group)
  permissions: Permission[]
  createdAt: datetime
  updatedAt: datetime
```

### Unified Blocking Model

```yaml
Blocking:
  # Unified model for all types of unavailability
  # Used at building, area, and place levels
  id: uuid
  entityType: enum (building, area, place)
  entityId: uuid
  blockingType: enum (closedHours, weekend, holiday, maintenance, event, disabled, custom)
  name: string (nullable, e.g., "Christmas Day", "Weekly Cleaning")
  reason: string (nullable, shown to users) (US-47, US-48)

  # Time specification (one of the following patterns):
  # 1. One-time blocking:
  startTime: datetime
  endTime: datetime

  # 2. Recurring blocking (uses iCal RRULE format):
  isRecurring: boolean (US-49)
  recurrenceRule: string (RRULE format, nullable)
  # Examples:
  #   "FREQ=DAILY;BYHOUR=18;BYMINUTE=0" + duration = closed 18:00-09:00 daily
  #   "FREQ=WEEKLY;BYDAY=SA,SU" = weekends
  #   "FREQ=WEEKLY;BYDAY=MO;BYHOUR=7" + duration = Monday cleaning 07:00-09:00
  #   "FREQ=YEARLY;BYMONTH=12;BYMONTHDAY=25" = Christmas (yearly recurring)
  recurrenceDuration: duration (how long each occurrence lasts)
  recurrenceEnd: date (nullable, when recurring rule ends)

  # Metadata:
  source: enum (own, inheritedFromArea, inheritedFromBuilding) (read-only, for display)
  createdAt: datetime
  updatedAt: datetime

CreateBlockingRequest:
  blockingType: enum (closedHours, weekend, holiday, maintenance, event, disabled, custom)
  name: string (nullable)
  reason: string (nullable)
  startTime: datetime (required for one-time)
  endTime: datetime (required for one-time)
  isRecurring: boolean
  recurrenceRule: string (required if recurring)
  recurrenceDuration: duration (required if recurring)
  recurrenceEnd: date (nullable)
```

### Supporting Types

```yaml
PlaceConstraints:
  maxReservationDuration: duration (nullable) (US-36)
  minReservationDuration: duration (nullable) (US-59)
  maxReservationsPerDay: int (nullable) (US-37)
  maxReservationsPerWeek: int (nullable)
  maxReservationsPerMonth: int (nullable)
  maxReservationsPerYear: int (nullable)
  maxHoursPerDay: int (nullable) (US-34)
  maxHoursPerWeek: int (nullable)
  maxHoursPerMonth: int (nullable)
  maxHoursPerYear: int (nullable)
  maxConcurrentReservations: int (nullable) (US-11)
  maxAdvanceBookingDays: int (nullable) (US-58)
  checkInTimeoutMinutes: int (nullable) (US-60)
  decayTimeoutMinutes: int (nullable) (US-20)
  whitelistEnabled: boolean (US-35)

TimeSlotConfig:
  # Defines the BOOKING GRID - valid start/end times for reservations (US-87, US-88)
  # This does NOT define availability. Availability = not blocked AND not reserved.
  intervalMinutes: int (e.g., 15, 30, 60)

  # Optional: restrict booking to certain hours (still uses blocking for actual closures)
  # If not set, any time on the interval grid is valid (subject to blocking)
  earliestStartTime: time (nullable, e.g., "08:00")
  latestEndTime: time (nullable, e.g., "20:00")

  # Generated slots example for 15-min intervals, 09:00-18:00:
  # 09:00, 09:15, 09:30, 09:45, 10:00, ..., 17:45, 18:00

AvailabilitySlot:
  # Response type for /places/{placeId}/availability
  startTime: datetime
  endTime: datetime
  status: enum (available, blocked, reserved)
  blockingReason: string (nullable, if blocked)
  blockingSource: enum (place, area, building) (nullable, if blocked)
  reservationId: uuid (nullable, if reserved)

PlaceAvailability:
  # Response type for /areas/{areaId}/availability
  placeId: uuid
  placeName: string
  status: enum (available, partiallyAvailable, fullyBooked, blocked, disabled)
  nextAvailableSlot: datetime (nullable)
  blockedReason: string (nullable)
  blockedSource: enum (place, area, building) (nullable)

RoomPlan:
  imageUrl: string
  markers: PlaceMarker[]

PlaceMarker:
  id: uuid
  placeId: uuid
  x: float (percentage)
  y: float (percentage)
  width: float
  height: float
  shape: enum (rectangle, circle)

Equipment:
  id: uuid
  name: string
  icon: string (optional)
  description: string

NotificationPreferences:
  reservationConfirmed: boolean
  reservationCancelled: boolean
  reservationReminder: boolean
  reminderMinutesBefore: int (US-53, US-54)
  checkInWarning: boolean

CreateReservationRequest:
  placeId: uuid
  startTime: datetime
  endTime: datetime
  # Optional recurrence for recurring reservations (US-57)
  recurrence:
    pattern: enum (daily, weekly, biweekly, monthly)
    endDate: date
    daysOfWeek: int[] (for weekly pattern, 0-6)

UpdateReservationRequest:
  startTime: datetime (optional)
  endTime: datetime (optional)

APIKey:
  id: uuid
  name: string
  keyPrefix: string (first 8 chars for identification)
  keyHash: string (stored, not returned)
  permissions: string[] (inherited from user)
  lastUsedAt: datetime
  expiresAt: datetime (nullable)
  createdAt: datetime

AuditLogEntry:
  id: uuid
  entityType: string
  entityId: uuid
  action: enum (create, update, delete)
  userId: uuid
  changes: json (before/after)
  timestamp: datetime
```

---

## Availability Calculation

### Algorithm

When calculating availability for a place at a given time:

```
1. Get all blocking periods:
   - Place's own blocking
   - Area's blocking (inherited)
   - Building's blocking (inherited)

2. Merge/flatten all blocking periods into a unified timeline

3. Get all reservations for the place in the time range

4. For each time slot on the booking grid:
   - If slot overlaps with any blocking period → status: "blocked"
   - Else if slot overlaps with any reservation → status: "reserved"
   - Else → status: "available"
```

### Example

Building has blocking: `closedHours` 18:00-08:00 daily (recurring)
Area has blocking: `maintenance` every Monday 08:00-10:00 (recurring)
Place has blocking: `event` 2025-01-15 14:00-16:00 (one-time)

For 2025-01-15 (Monday):
- 00:00-08:00: blocked (building: closedHours)
- 08:00-10:00: blocked (area: maintenance)
- 10:00-14:00: available (unless reserved)
- 14:00-16:00: blocked (place: event)
- 16:00-18:00: available (unless reserved)
- 18:00-24:00: blocked (building: closedHours)

---

## Error Responses

```yaml
ErrorResponse:
  error:
    code: string (e.g., "VALIDATION_ERROR", "NOT_FOUND", "FORBIDDEN")
    message: string
    details: object (optional, field-specific errors)

Common HTTP Status Codes:
  200: Success
  201: Created
  204: No Content (successful delete)
  400: Bad Request (validation error)
  401: Unauthorized (not authenticated)
  403: Forbidden (insufficient permissions)
  404: Not Found
  409: Conflict (e.g., overlapping reservation, time slot blocked)
  422: Unprocessable Entity (business rule violation)
  429: Too Many Requests (rate limiting)
  500: Internal Server Error
```

---

## Link Preview / Open Graph Support (US-90)

The `/api/v1/reservations/{reservationId}/share` endpoint returns metadata for rich link previews:

```yaml
ShareResponse:
  shareUrl: string (e.g., "https://roomy.example.com/r/abc123")
  previewMetadata:
    title: string (e.g., "Reservation: Meeting Room A")
    description: string (e.g., "Jan 15, 2025 10:00-12:00 | Building 1, Floor 2")
    image: string (optional, room/place image URL)
    siteName: string (e.g., "Roomy")
```

The share URL should serve an HTML page with proper Open Graph and Twitter Card meta tags for preview rendering in chat applications.

---

## Webhooks (Future Enhancement)

For external system integration beyond polling:

```yaml
/api/v1/webhooks:
  POST:
    summary: Register webhook endpoint
    body:
      url: string
      events: string[] (reservation.created, reservation.cancelled, etc.)
      secret: string

Events:
  - reservation.created
  - reservation.cancelled
  - reservation.checkedIn
  - reservation.expired
  - place.blocked
  - place.updated
  - blocking.created
  - blocking.deleted
```
