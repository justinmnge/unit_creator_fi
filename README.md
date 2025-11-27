![CI checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/ci.yml/badge.svg) &nbsp; ![CD checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/cd.yml/badge.svg)
# Unit Creator 🏴‍☠️

<img width="2557" height="1264" alt="Image" src="https://github.com/user-attachments/assets/8f4c463c-c477-48e1-9284-f781ba502431" />

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

Unit Creator FI is a web platform for Mil-Sim and tactical gaming communities to build authentic unit identities, manage rosters, and showcase organizational structure.

### Motivation
Mil-Sim enthusiasts and tactical gaming units currently juggle Discord channels, Google Sheets, and forum posts to maintain their organizational identity. Unit Creator centralizes this into a single, professional platform with proper data persistence and a clean presentation layer.

### What it does:
* Enables communities to create custom unit pages with emblems, hierarchy, and member profiles
* Provides a centralized platform for unit identity management across various tactical enthusiast groups
* Stores unit organizational data, member rosters, and visual assets
* Demonstrates RESTful API design patterns with clean JSON endpoints

### What This Demonstrates
This project showcases full-stack development patterns for community management platforms. It explores how tactical gaming units could centralize their organizational identity, moving from scattered Discord channels and spreadsheets to a unified system with proper data modeling and API design.

## 🛠️ Built With

<p align="left">
<img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=ts"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
<img src="https://github.com/user-attachments/assets/debd8e54-8d8a-4dc1-b900-0c778af06574" width="48" height="48"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=react"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=docker"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=postgres"></a>&nbsp;
</p>

### 🔧 Technical Stack

**Backend**
* **Go HTTP server** - Chosen for reliable concurrent request handling, simple deployment (single binary), and excellent standard library support for building RESTful APIs
* **PostgreSQL** - Persistent storage for unit data and member rosters (planned for Phase 2)
* **Docker** - Containerization for consistent deployment across environments

**Planned Frontend (Phase 2)**
* **React.js + TypeScript** - Component-based UI for dynamic unit page creation and editing
* **Modern JavaScript ecosystem** - Build tooling and state management for fluid user interactions

**Infrastructure**
* Docker containerization with planned GCP/Kubernetes scaling
* RESTful API design with comprehensive documentation for third-party integration

### 📍 Current Status: Phase 1 (View-Only)
This is currently a showcase platform with manually curated content. User-generated unit creation will be available in Phase 2.

### 🛣️ Roadmap
* **Phase 1** (Current): Static JSON data with manual curation - demonstrates platform concept and API structure
* **Phase 2**: Database-backed system with user authentication, unit creation UI, and member management
* **Phase 3**: Enhanced user experience with React frontend, advanced search/filtering, and unit comparison features

## ✅ Prerequisites

* Install Docker Desktop for your operating system
  * [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac-install/)
  * [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/)
  * [Docker Engine for Linux](https://docs.docker.com/engine/install/)

* Verify installation worked
```sh
  docker --version
  docker compose version
```

## Quick Start

### Run Pre-built Docker Image

Pull and run the latest production image directly from Docker Hub:
```sh
docker pull justinmnge/unit_creator_fi:latest
docker run -p 8080:8080 justinmnge/unit_creator_fi:latest
```

Then visit `http://localhost:8080` in your browser.

## 👀 What You'll See
After running the server, visit `http://localhost:8080` to browse:
- Example unit pages showcasing organizational structure and member profiles
- Sample unit emblems and visual identity system
- RESTful API endpoints returning JSON data for unit information
- Responsive HTML templates demonstrating the platform's presentation layer

## Usage

* Run the server
```sh
  docker compose up
```

* Run in detached mode (background)
```sh
  docker compose up -d
```

* Stop the server
```sh
  docker compose down
```

* Rebuild after code changes
```sh
  docker compose up --build
```

### Running the Pre-built Image

* Run the latest image from Docker Hub
```sh
  docker run -p 8080:8080 justinmnge/unit_creator_fi:latest
```

* Run a specific version
```sh
  docker run -p 8080:8080 justinmnge/unit_creator_fi:COMMIT_SHA
```

## Contributing

### Clone the repo
```bash
git clone https://github.com/justinmnge/unit_creator_fi
cd unit_creator_fi
```

### Build the compiled binary
```bash
go build
```

### Run the test suite
```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## 🐳 Docker Hub

Pre-built Docker images are automatically published to Docker Hub via CI/CD:

[![Docker Hub](https://img.shields.io/docker/pulls/justinmnge/unit_creator_fi)](https://hub.docker.com/r/justinmnge/unit_creator_fi)

* **Latest build**: `justinmnge/unit_creator_fi:latest`
* **Specific commits**: `justinmnge/unit_creator:COMMIT_SHA`

## ℹ️ Contact

Justin Monge - hello@justin-monge.dev

Project Link: [https://github.com/justinmnge/unit_creator_fi](https://github.com/justinmnge/unit_creator_fi)
<p align="right">(<a href="#readme-top">back to top</a>)</p>