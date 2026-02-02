# MTG Collection Tracker - Features & UI Guide

## Overview
This application provides a comprehensive interface for managing your Magic: The Gathering card collection with a clean, modern Bootstrap-based UI.

## Features Walkthrough

### 1. Authentication

#### Registration Page (`/register`)
- Clean, centered form with gradient purple background
- Fields:
  - Username (required)
  - Password (required)
  - Confirm Password (required)
- Validates password match
- Checks for duplicate usernames
- Link to login page for existing users
- Responsive design works on mobile and desktop

#### Login Page (`/login`)
- Similar styling to registration
- Fields:
  - Username (required)
  - Password (required)
- Session-based authentication
- Secure password hashing with bcrypt
- Link to registration for new users
- Error messages displayed for invalid credentials

### 2. Card Collection Management

#### Card List Page (`/cards`)
**Layout:**
- Navigation bar with:
  - App logo and name
  - Current username display
  - Logout button
- Header section with:
  - Total card count
  - "Add Card" button (prominent, top-right)
- Search bar:
  - Full-width search input
  - Searches by card name, set code, or collector number
  - Real-time search (submit to filter)

**Card Table:**
- Responsive table with columns:
  1. **Image** - Card image thumbnail (50px height) or placeholder icon
  2. **Card Name** - Full card name
  3. **Set Code** - MTG set identifier
  4. **Collector #** - Card number in set
  5. **Language** - Card language
  6. **Quantity** - Number of copies owned
  7. **Price (THB)** - Purchase price in Thai Baht (2 decimal places)
  8. **Bought Date** - Purchase date (YYYY-MM-DD format)
  9. **Sell Date** - Sale date if sold
  10. **Actions** - Edit (yellow) and Delete (red) buttons

**Pagination:**
- 20 cards per page
- Page numbers displayed at bottom
- Previous/Next navigation buttons
- Current page highlighted
- Preserves search filter across pages

**Empty State:**
- Friendly message when no cards exist
- Direct link to "Add your first card"

#### Add Card Page (`/cards/add`)
**Form Fields:**
- **Card Name*** (required) - Text input
- **Card Image URL** - Text input for image URL
- **Set Code** - Text input (e.g., "MH2", "DMU")
- **Collector Number** - Text input (e.g., "123", "456a")
- **Language** - Text input (default: "English")
- **Quantity*** (required) - Number input (min: 1, default: 1)
- **Buying Price (THB)** - Decimal input (default: 0)
- **Bought Date** - Date picker
- **Sell Date** - Date picker

**Actions:**
- Cancel button (gray) - Returns to card list
- Add Card button (blue) - Saves card and returns to list

**Validation:**
- Required fields marked with asterisk (*)
- Client-side and server-side validation
- Error messages displayed at top if submission fails

#### Edit Card Page (`/cards/edit/:id`)
- Same form as Add Card
- Pre-populated with existing card data
- Date fields show existing dates in YYYY-MM-DD format
- Update button instead of Add button
- Can only edit cards owned by logged-in user

#### Delete Card
- Triggered from card list Actions column
- JavaScript confirmation dialog
- "Are you sure you want to delete this card?"
- Permanent deletion on confirmation
- Can only delete cards owned by logged-in user

### 3. Security Features

**Authentication:**
- Session-based login
- Bcrypt password hashing
- Secure session cookies
- Auto-redirect to login for protected routes

**Authorization:**
- Users can only view their own cards
- Users can only edit/delete their own cards
- User ID checked on all card operations
- Database-level user_id foreign key constraint

### 4. UI/UX Features

**Responsive Design:**
- Works on desktop, tablet, and mobile
- Bootstrap 5 grid system
- Responsive tables
- Mobile-friendly forms

**Visual Design:**
- Modern gradient purple navbar
- Clean white content area
- Bootstrap icons throughout
- Consistent color scheme:
  - Primary (blue) for main actions
  - Warning (yellow) for edit
  - Danger (red) for delete
  - Secondary (gray) for cancel

**User Feedback:**
- Success messages after operations
- Error messages for failures
- Confirmation dialogs for destructive actions
- Loading states (built into browser forms)

### 5. Data Management

**Search/Filter:**
- Single search box
- Searches across:
  - Card name (partial match)
  - Set code (partial match)
  - Collector number (partial match)
- Case-insensitive search
- Results update on form submission

**Pagination:**
- Fixed page size (20 cards)
- Efficient database queries (OFFSET/LIMIT)
- Total page count calculation
- Clean page number navigation

**Sorting:**
- Cards ordered by creation date (newest first)
- Consistent ordering across sessions

### 6. Database Schema

**Automatic Migration:**
- Tables created on first run
- GORM handles schema updates
- Soft deletes supported (deleted_at field)

**Data Types:**
- Decimal for prices (10,2 precision)
- Date for bought/sell dates (nullable)
- Text for card names and URLs
- Integer for quantities
- Indexed user_id for fast queries

## Technical Architecture

**Clean Architecture Layers:**
1. **Domain** - Entities and interfaces
2. **Use Cases** - Business logic
3. **Infrastructure** - Database implementations
4. **Handlers** - HTTP request handling
5. **Templates** - HTML views

**Technology Stack:**
- Go 1.21+ with Gin framework
- GORM for database operations
- Bootstrap 5 for UI
- MySQL 8.0 for data storage
- Session-based authentication

## Performance Considerations

- Indexed database queries
- Pagination to limit data transfer
- Efficient GORM queries (no N+1)
- CDN-hosted Bootstrap assets
- Minimal JavaScript (mostly vanilla browser APIs)

## Accessibility

- Semantic HTML5
- Form labels for all inputs
- ARIA labels where appropriate
- Keyboard navigation support
- Screen reader friendly
- High contrast text

## Browser Compatibility

- Modern browsers (Chrome, Firefox, Safari, Edge)
- IE11+ (with Bootstrap 5 support)
- Mobile browsers (iOS Safari, Chrome Mobile)
