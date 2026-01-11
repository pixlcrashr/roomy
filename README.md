# Roomy

A room and place reservation management system designed for organizations to manage bookable spaces with self-service user reservations and fine-grained permission controls.

## Overview

Roomy enables organizations to manage physical spaces such as desks, meeting rooms, seating areas, and other bookable units. Users can self-service reserve places through an intuitive interface, while organization members maintain control over configuration, permissions, and analytics.

## Key Features

### For Users
- Self-service reservation with calendar views
- QR code check-in for reservations
- Room plan visualization for easy navigation
- Search and filter by equipment, capacity, and availability
- Recurring reservations and favorite places
- Email notifications for confirmations, reminders, and cancellations
- Personal dashboard for managing bookings

### For Organizations
- Hierarchical structure: Buildings > Areas > Places
- Fine-grained permission system with customizable groups
- Configurable constraints (duration limits, booking windows, user whitelists)
- Room plan editor for visual space management
- QR code generation with customizable templates
- Usage statistics and analytics
- Reservation data export
- Audit logging for compliance
- API for external integrations

### Authentication
- OAuth 2.0 authentication via GitLab
- API key support for system integrations

## Project Structure

```
roomy/
├── planning/           # Feature specifications and user stories
├── CLAUDE.md           # Development guidance
├── LICENSE             # GPL-3.0 License
└── README.md           # This file
```

## Status

This project is currently in the **planning phase**. See [planning/feature-list.md](planning/feature-list.md) for the complete set of user stories defining the feature scope.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

---

> **Note:** Parts of this project, including documentation and code, are generated with the assistance of Large Language Models (LLMs). All generated content is reviewed and maintained by the project contributors.
