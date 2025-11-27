![CI checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/ci.yml/badge.svg) &nbsp; ![CD checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/cd.yml/badge.svg)
# Unit Creator 🏴‍☠️

<img width="1912" height="908" alt="Image" src="https://github.com/user-attachments/assets/14049d29-0d06-41f7-aae1-419d91f6fb6c" />

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

Unit Creator FI is a web platform for milsim and tactical gaming communities to build authentic unit identities, manage rosters, and showcase organizational structure.

What it does:
* Enables communities to create custom unit pages with emblems, hierarchy, and member profiles
* Provides a centralized platform for unit identity management across airsoft teams and tactical gaming groups
* Stores unit organizational data, member rosters, and visual assets
* Exposes RESTful JSON APIs for community integrations and future expansion

## 🛠️ Built With

<p align="left">
 <a href="https://go.dev/"><img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
<img src="https://github.com/user-attachments/assets/debd8e54-8d8a-4dc1-b900-0c778af06574" width="48" height="48"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=docker"></a>&nbsp;
 </p>

## 🔧 Technical Stack
* Backend: Go HTTP server with PostgreSQL for persistent storage
* Authentication: Secure player and unit data access controls
* Deployment: Docker containerization with planned GCP/Kubernetes scaling
* API Design: RESTful endpoints with comprehensive documentation for third-party integration

## 🛣️ Roadmap
* Phase 1: Static JSON data with manual curation
* Phase 2: Database-backed system with user-generated content
* Phase 3: Real-time game server integration for live player statistics

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

## 📦 Installation

### Option 1: Run Pre-built Docker Image (Quickest)

Pull and run the latest production image directly from Docker Hub:
```sh
docker pull justinmnge/unit_creator_fi:latest
docker run -p 8080:8080 justinmnge/unit_creator_fi:latest
```

Then visit `http://localhost:8080` in your browser.

### Option 2: Build from Source

1. Clone the repo
```sh
   git clone https://github.com/justinmnge/unit_creator_fi.git
   cd unit_creator_fi
```
2. Change git remote url to avoid accidental pushes to base project
```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
```

## ⚙️ Usage

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

## 🐳 Docker Hub

Pre-built Docker images are automatically published to Docker Hub via CI/CD:

[![Docker Hub](https://img.shields.io/docker/pulls/justinmnge/unit_creator_fi)](https://hub.docker.com/r/justinmnge/unit_creator_fi)

* **Latest build**: `justinmnge/unit_creator_fi:latest`
* **Specific commits**: `justinmnge/unit_creator:COMMIT_SHA`

## ℹ️ Contact

Justin Monge - hello@justin-monge.dev

Project Link: [https://github.com/justinmnge/unit_creator_fi](https://github.com/justinmnge/unit_creator_fi)
<p align="right">(<a href="#readme-top">back to top</a>)</p>