# GeoData Application

## Overview

GeoData is a web application that allows users to upload, manage, and visualize geospatial data. Users can create, update, and retrieve geospatial features on a map using a user-friendly interface. The application supports uploading GeoJSON and KML files and provides drawing tools for custom shapes.

## Features

- **User Registration and Authentication**: Secure user accounts with registration and login functionality.
- **Upload GeoJSON and KML Files**: Import geospatial data through file uploads.
- **Map Visualization**: View geospatial data on an interactive map using React Leaflet.
- **Draw Custom Shapes**: Create and edit custom shapes directly on the map.
- **Save and Retrieve Data**: Save geospatial data to the backend and retrieve it as needed.

## Tech Stack

- **Frontend**: React, React Leaflet, Axios
- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL with PostGIS extension
## Getting Started

### Prerequisites

- Node.js (v12 or later)
- Go (v1.15 or later)
- PostgreSQL with PostGIS extension

### Installation

1. **Clone the Repository**:
2. ```bash
   git clone https://github.com/deepak2875/Deepak_geo_data_app
   cd Deepak_geo_data_app

### Navigate to frontend and run npm install
-cd geo-data-frontend
-npm install
-npm start

### Navigate to backend and run main.go file 
-cd geo-backend
-go run main.go
