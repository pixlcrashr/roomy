# Frontend Architecture

This document outlines the frontend routes and functionality for Roomy based on the user stories.

## Technology Considerations

- SPA (Single Page Application) with client-side routing
- OAuth 2.0 GitLab integration for authentication
- Responsive design for desktop and mobile
- Calendar/scheduling components
- Interactive room plan viewer/editor
- QR code scanning capability (for check-in)

---

## Unified Availability Model

The frontend reflects the backend's **blocking-based availability model**:

1. **Default state**: Everything is available 24/7
2. **Blocking periods**: Define when something is NOT available
3. **Availability** = All time − Blocking periods − Existing reservations

### Key UI Implications

- **Blocking editor**: Single unified component for managing closures, holidays, maintenance
- **Availability view**: Shows what's available (inverse of blocking)
- **Time slot picker**: Respects booking grid intervals and filters out blocked times
- **Inheritance indicators**: Shows where blocking comes from (building/area/place)

---

## Route Structure

### Public Routes (No Authentication Required)

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/` | Landing page with public room availability overview | US-31 |
| `/login` | GitLab OAuth login initiation | US-1, US-2 |
| `/oauth/callback` | OAuth callback handler | US-1, US-2 |
| `/buildings` | Public list of buildings | US-31 |
| `/buildings/:buildingId` | Public building detail with areas | US-31 |
| `/buildings/:buildingId/areas/:areaId` | Public area detail with places and availability | US-31 |
| `/r/:shareCode` | Shared reservation preview page (with Open Graph meta tags) | US-66, US-90 |

### Authenticated User Routes

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/dashboard` | Personal dashboard with upcoming/past reservations | US-43 |
| `/dashboard/favorites` | User's favorite places | US-45 |
| `/reservations` | List of user's reservations | US-43 |
| `/reservations/:reservationId` | Reservation detail view | US-43 |
| `/reservations/new` | Create new reservation flow (single or recurring) | US-9, US-39, US-57 |
| `/reservations/:reservationId/edit` | Edit reservation (extend/shorten) | US-44 |
| `/search` | Search places with filters (equipment, capacity, availability) | US-40, US-61, US-63, US-65 |
| `/places/:placeId` | Place detail with calendar view | US-32 |
| `/places/:placeId/reserve` | Reserve specific place | US-9 |
| `/areas/:areaId/roomPlan` | Interactive room plan view | US-5, US-14 |
| `/checkIn/:reservationId` | Manual check-in page | US-10 |
| `/checkIn/qr` | QR code scanner for check-in | US-10 |
| `/settings` | User settings | US-22, US-54 |
| `/settings/notifications` | Notification preferences | US-22, US-53, US-54 |

### Organization Member Routes (Admin Panel)

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin` | Admin dashboard with statistics overview | US-30 |
| `/admin/statistics` | Detailed usage statistics and metrics | US-30 |
| `/admin/statistics/export` | Export reservation data | US-50 |

#### Building Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/buildings` | List all buildings | US-17 |
| `/admin/buildings/new` | Create new building | US-17 |
| `/admin/buildings/:buildingId` | Building detail/edit | US-16 |
| `/admin/buildings/:buildingId/blocking` | Configure blocking periods (closures, holidays, hours) | US-55, US-56 |

#### Area Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/buildings/:buildingId/areas` | List areas in building | US-18 |
| `/admin/buildings/:buildingId/areas/new` | Create new area | US-18 |
| `/admin/areas/:areaId` | Area detail/edit | US-19 |
| `/admin/areas/:areaId/roomPlan/edit` | Room plan editor | US-13, US-15, US-91 |
| `/admin/areas/:areaId/blocking` | Configure blocking periods (inherits from building) | US-46, US-47, US-48, US-49 |

#### Place Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/areas/:areaId/places` | List places in area | US-6 |
| `/admin/areas/:areaId/places/new` | Create new place | US-6 |
| `/admin/places/:placeId` | Place detail/edit | US-7 |
| `/admin/places/:placeId/constraints` | Configure place constraints | US-11, US-20, US-33, US-34, US-36, US-58, US-59, US-60 |
| `/admin/places/:placeId/timeSlots` | Configure booking grid intervals | US-87, US-88 |
| `/admin/places/:placeId/equipment` | Configure equipment/amenities | US-62 |
| `/admin/places/:placeId/whitelist` | Configure user whitelist | US-35 |
| `/admin/places/:placeId/blocking` | Configure blocking periods (inherits from area/building) | US-46, US-47, US-48, US-49 |
| `/admin/places/:placeId/qrCode` | QR code generation and printing | US-25, US-26, US-27, US-28 |

#### QR Code Templates

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/qrTemplates` | List QR code templates | US-29 |
| `/admin/qrTemplates/new` | Create new QR template | US-29 |
| `/admin/qrTemplates/:templateId` | Edit QR template (HTML editor) | US-29 |

#### Reservation Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/reservations` | List all reservations | US-12, US-24 |
| `/admin/reservations/new` | Manually create reservation for user | US-12 |
| `/admin/reservations/:reservationId` | View/manage reservation | US-24 |

#### User Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/users` | List all users (filter by multiple groups) | US-80 |
| `/admin/users/:userId` | User detail view (OAuth details, groups) | US-81, US-82 |
| `/admin/users/:userId/groups` | Manage user group assignments | US-75, US-76, US-82 |
| `/admin/users/:userId/disable` | Disable user account | US-68 |

#### Group/Permission Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/groups` | List all permission groups | US-83 |
| `/admin/groups/new` | Create new group | US-73 |
| `/admin/groups/:groupId` | Group detail (permissions, members) | US-84, US-86 |
| `/admin/groups/:groupId/permissions` | Edit group permissions | US-74, US-84 |
| `/admin/groups/defaultAssignment` | Configure default groups for new users | US-79 |

#### Audit & Logs

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/auditLog` | View change history | US-51 |

#### API Key Management

| Route | Description | User Stories |
|-------|-------------|--------------|
| `/admin/apiKeys` | List user's API keys | US-71 |
| `/admin/apiKeys/new` | Create new API key | US-71 |

---

## Core Components

### Layout Components

- `PublicLayout` - Header with login button, navigation
- `AuthenticatedLayout` - User menu, notifications, navigation
- `AdminLayout` - Admin sidebar navigation, breadcrumbs

### Shared Components

| Component | Description | Used In |
|-----------|-------------|---------|
| `BuildingCard` | Building summary card | Building lists |
| `AreaCard` | Area summary card | Area lists |
| `PlaceCard` | Place summary with availability indicator | Place lists, search |
| `ReservationCard` | Reservation summary | Dashboard, lists |
| `CalendarView` | Weekly/monthly calendar for availability | Place detail, area detail |
| `TimeSlotPicker` | Time slot selection based on configured intervals (filters blocked) | Reservation form |
| `RoomPlanViewer` | Interactive SVG/image room plan | Area view |
| `RoomPlanEditor` | Drag-and-drop room plan editor (with delete) | Admin area edit |
| `QRScanner` | Camera-based QR code scanner | Check-in |
| `QRCodePreview` | QR code display with print options | Admin QR generation |
| `SearchFilters` | Equipment, capacity, date filters | Search page |
| `ReservationForm` | Date, time slot, duration picker (with recurring options) | Reservation creation |
| `BlockingEditor` | Unified editor for all blocking types | Building/area/place blocking |
| `BlockingCalendar` | Visual calendar showing blocked periods | Blocking management |
| `BlockingInheritanceView` | Shows inherited blocking from parent entities | Place/area blocking pages |
| `PermissionSelector` | Multi-select permission checkboxes | Group management |
| `UserSelector` | User search/select (email, ID) | Manual reservation |
| `GroupMultiSelect` | Multi-group filter selector | User list filtering |
| `StatisticsChart` | Various chart types for metrics | Admin statistics |
| `AuditLogTable` | Paginated audit log entries | Admin audit |
| `NotificationSettings` | Email notification toggles | User settings |
| `ShareButton` | Share reservation with link preview | Reservation detail |
| `IcsCalendarLink` | Subscribe to .ics calendar | Place/area detail |
| `AvailabilityIndicator` | Shows availability status with blocking source | Place cards, room plan |

### Form Components

- `DateTimePicker` - Date and time selection
- `DurationPicker` - Reservation duration selection
- `TimeSlotSelector` - Select from configured intervals (US-87, US-88)
- `RecurringScheduler` - Configure recurring reservations (US-57)
- `BlockingEntryForm` - Add/edit blocking entries (closedHours, holiday, maintenance, etc.)
- `RecurrenceRuleBuilder` - Build RRULE for recurring blocking
- `TimeRangePicker` - Time range configuration
- `HTMLEditor` - QR code template editing
- `ImageUploader` - Room plan image upload
- `WhitelistEditor` - User whitelist management (add/remove by ID)

---

## State Management

### Global State

- `AuthState` - Current user, permissions, OAuth tokens
- `NotificationState` - Unread notifications, preferences
- `FavoritesState` - User's favorite places

### Feature State

- `ReservationState` - Current reservation flow data
- `SearchState` - Active filters, results
- `AdminState` - Selected building/area/place context
- `BlockingState` - Blocking periods for current entity (with inheritance)

---

## Key User Flows

### 1. Self-Service Reservation (US-8, US-9, US-39)

```
Search/Browse → Select Place → View Calendar (shows blocked times) → Select Available Time Slot → Confirm → Reservation Created
```

### 2. QR Code Check-In (US-10)

```
Scan QR → Validate Reservation → Confirm Check-In → Success
```

### 3. Room Plan Reservation (US-5, US-14)

```
Select Area → View Room Plan (places colored by availability) → Click Place → View Availability → Reserve
```

### 4. Recurring Reservation (US-57)

```
Select Place → Configure Recurrence → Preview Slots (skips blocked) → Confirm All → Reservations Created
```

### 5. Extend/Shorten Reservation (US-44)

```
View Reservation → Edit → Adjust Start/End Time (respects blocking) → Validate Availability → Save
```

### 6. Admin Blocking Setup (US-46, US-47, US-49, US-55)

```
Select Entity (Building/Area/Place) → Open Blocking Editor → Add Entry (closedHours/holiday/maintenance/etc.) → Configure Recurrence (optional) → Save
```

### 7. Admin Place Setup (US-7, US-25, US-87)

```
Create Place → Configure Constraints → Set Time Slot Intervals → Add Equipment → Configure Blocking (or inherit) → Generate QR Code → Print
```

### 8. Share Reservation (US-66, US-90)

```
View Reservation → Click Share → Copy Link → Paste in Chat → Rich Preview Displayed
```

### 9. Subscribe to Calendar (US-89)

```
View Place/Area → Click Calendar Subscribe → Copy .ics URL → Add to Calendar App
```

---

## Blocking Editor Component

The `BlockingEditor` is the central component for managing all blocking periods. It replaces separate schedule/hours/holidays editors.

### Features

- **Unified entry types**: closedHours, weekend, holiday, maintenance, event, disabled, custom
- **Visual calendar**: Shows all blocking periods on a weekly/monthly view
- **Recurrence builder**: UI for creating RRULE patterns (daily, weekly, monthly, yearly)
- **Inheritance view**: Shows blocking inherited from parent entities (read-only, expandable)
- **Bulk actions**: Add/remove multiple entries at once
- **Preview**: Shows effective availability after blocking is applied

### Entry Form Fields

| Field | Type | Description |
|-------|------|-------------|
| `blockingType` | Select | Type of blocking (closedHours, holiday, etc.) |
| `name` | Text | Optional name (e.g., "Christmas", "Weekly Cleaning") |
| `reason` | Text | Optional reason shown to users |
| `startTime` | DateTime | Start of blocking period |
| `endTime` | DateTime | End of blocking period |
| `isRecurring` | Toggle | Enable recurring pattern |
| `recurrenceRule` | RRuleBuilder | Visual RRULE configuration |
| `recurrenceDuration` | Duration | How long each occurrence lasts |
| `recurrenceEnd` | Date | When recurring pattern ends |

### Inheritance Display

When viewing blocking for an area or place, inherited blocking is shown separately:

```
┌─────────────────────────────────────────────────┐
│ Blocking for: Meeting Room A                     │
├─────────────────────────────────────────────────┤
│ Own Blocking (editable)                          │
│ ┌─────────────────────────────────────────────┐ │
│ │ + Add Blocking                              │ │
│ │ • Event: Team Offsite (Jan 20, 14:00-18:00) │ │
│ └─────────────────────────────────────────────┘ │
├─────────────────────────────────────────────────┤
│ Inherited from Area: Floor 2 (read-only)         │
│ • Maintenance: Mon 07:00-09:00 (weekly)          │
├─────────────────────────────────────────────────┤
│ Inherited from Building: HQ (read-only)          │
│ • Closed Hours: 18:00-08:00 (daily)              │
│ • Weekend: Sat-Sun (weekly)                      │
│ • Holiday: Christmas (Dec 25, yearly)            │
└─────────────────────────────────────────────────┘
```

---

## Availability Calculation (Frontend)

When displaying availability, the frontend:

1. Fetches blocking periods with `includeInherited=true`
2. Fetches existing reservations for the time range
3. Uses the time slot configuration to generate the booking grid
4. Marks each slot as:
   - `available` - Can be booked
   - `blocked` - Blocked (with source indicator)
   - `reserved` - Already booked

### Visual Indicators

| Status | Color | Icon | Tooltip |
|--------|-------|------|---------|
| Available | Green | ✓ | "Available" |
| Blocked (own) | Gray | 🚫 | "Blocked: {reason}" |
| Blocked (inherited) | Light gray | 🔒 | "Blocked by {building/area}: {reason}" |
| Reserved (own) | Blue | 📅 | "Your reservation" |
| Reserved (other) | Orange | 👤 | "Reserved" (no user details shown) |

---

## Permissions & Access Control

### Route Guards

Routes are protected based on user permissions from assigned groups:

| Permission | Routes Accessible |
|------------|-------------------|
| `view:buildings` | `/buildings/*`, `/admin/buildings` (read-only) |
| `manage:buildings` | `/admin/buildings/*` (full CRUD, blocking) |
| `view:areas` | `/areas/*`, `/admin/areas` (read-only) |
| `manage:areas` | `/admin/areas/*` (full CRUD, blocking) |
| `view:places` | `/places/*`, `/admin/places` (read-only) |
| `manage:places` | `/admin/places/*` (full CRUD, blocking) |
| `view:reservations` | `/admin/reservations` (read-only) |
| `manage:reservations` | `/admin/reservations/*` (full CRUD) |
| `view:users` | `/admin/users` (read-only) |
| `manage:users` | `/admin/users/*` (full CRUD) |
| `view:groups` | `/admin/groups` (read-only) |
| `manage:groups` | `/admin/groups/*` (full CRUD) |
| `view:statistics` | `/admin/statistics` |
| `view:auditLog` | `/admin/auditLog` |
| `manage:apiKeys` | `/admin/apiKeys/*` |

### Component-Level Permissions

Components conditionally render based on permissions:
- Edit/Delete buttons hidden without `manage:*` permissions
- User details hidden from non-members (US-23)
- Certain statistics only visible to authorized members
- Inherited blocking is always read-only (must edit at source)

---

## Notification System (US-21, US-22, US-53, US-54)

### Notification Types

| Type | Trigger | User Configurable |
|------|---------|-------------------|
| `reservation.confirmed` | Reservation created | Yes |
| `reservation.cancelled` | Reservation cancelled (decay/manual) | Yes |
| `reservation.reminder` | X minutes before reservation | Yes (timing) |
| `reservation.expiring` | Check-in timeout warning | Yes |
| `place.blocked` | Favorite place blocked | Yes |

### Delivery Channels

- Email (primary)
- In-app notifications (optional future enhancement)

---

## Responsive Breakpoints

| Breakpoint | Target |
|------------|--------|
| `< 640px` | Mobile phones |
| `640px - 1024px` | Tablets |
| `> 1024px` | Desktop |

### Mobile Considerations

- Bottom navigation for authenticated users
- Simplified room plan view (list fallback)
- Camera access for QR scanning
- Touch-friendly calendar/date pickers
- Time slot picker optimized for touch
- Collapsible blocking inheritance sections

---

## Link Preview / Open Graph Support (US-90)

The `/r/:shareCode` route serves as a shareable reservation page with proper meta tags:

```html
<meta property="og:title" content="Reservation: Meeting Room A" />
<meta property="og:description" content="Jan 15, 2025 10:00-12:00 | Building 1, Floor 2" />
<meta property="og:image" content="https://roomy.example.com/images/room-a.jpg" />
<meta property="og:url" content="https://roomy.example.com/r/abc123" />
<meta name="twitter:card" content="summary" />
```

This enables rich previews in chat applications like WhatsApp, Slack, Teams, etc.

---

## iCalendar Integration (US-89)

Public `.ics` calendar URLs are available for:
- Individual places: `/api/v1/places/:placeId/calendar.ics`
- Areas: `/api/v1/areas/:areaId/calendar.ics`
- Buildings: `/api/v1/buildings/:buildingId/calendar.ics`

The UI provides:
- Copy-to-clipboard button for calendar URL
- "Add to Calendar" button with common calendar app options (Google Calendar, Outlook, Apple Calendar)
- Shows next 3 months of blocking periods and reservations
