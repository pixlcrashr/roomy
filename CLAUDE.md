# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Roomy** is a room/place reservation and management system designed for organizations to manage bookable spaces (desks, meeting rooms, seating areas, etc.) with self-service user reservations and fine-grained permission controls.

## Core Domain Model

The system is built around a hierarchical structure:

- **Building** (highest level) - Contains areas, has name, description, location, opening hours
- **Area** - Groups multiple places within a building, can have room plans and serve as templates for place configuration
- **Place** (lowest level) - Individual bookable units (desk, chair, meeting room, etc.)
  - Can be configured with: location, name, seat capacity, description, bookability settings
  - May have equipment/amenities (projector, whiteboard, etc.)
  - Can be disabled or have time-based restrictions

**Reservation** - Links users to places for specific time slots
  - Requires check-in (via QR code or manual) unless configured otherwise
  - Subject to decay/timeout if not checked in
  - Can be self-service or manually assigned by organization members

**User & Permission System**:
- OAuth 2.0 authentication via GitLab
- Fine-grained permission/authorization using permission groups
- Two default groups: "system" (all permissions, immutable) and "default" (baseline for new users, editable)
- Organization members have different roles (viewer, editor, admin)

## Key Features

### User-Facing
- Self-service reservation with calendar view
- Public room availability list (anonymous viewing)
- QR code check-in for reservations
- Email notifications (reservation confirmations, cancellations, reminders)
- Room plan visualization for areas
- Search and filtering by equipment, capacity, availability
- Recurring reservations
- Favorite places
- Personal reservation dashboard

### Organization Member Features
- Manual reservation assignment
- Place/area/building management (CRUD operations)
- Room plan editor
- Configurable constraints per place/area:
  - Maximum reservation duration
  - Maximum reservations per time period (day/week/month/year)
  - Required check-in settings and timeout
  - User whitelists
  - Booking windows (how far in advance)
  - Minimum booking duration
- Blocking/maintenance scheduling (one-time or recurring)
- QR code generation and printing (configurable HTML templates)
- Usage statistics and metrics (anonymous)
- Reservation data export (CSV)
- Audit logs for changes
- User and group management
- API key creation for integrations

### System Capabilities
- API for external integrations (protected by API keys)
- API keys inherit permissions from the creating user
- Privacy compliance (user details hidden from other users, visible to organization members)

## Development Notes

### Project Status
This is an early-stage project currently in the planning phase. The `planning/feature-list.md` contains 88 user stories defining the complete feature set.

### Technology Stack
Not yet determined. Implementation choices for backend, frontend, database, and deployment should align with:
- OAuth 2.0 GitLab integration requirements
- QR code generation and scanning
- Calendar/scheduling functionality
- Permission/authorization system
- API development
- Room plan image upload and editing

### License
GNU General Public License v3.0 (GPL-3.0)
