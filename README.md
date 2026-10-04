![CI checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/ci.yml/badge.svg) &nbsp; ![CD checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/cd.yml/badge.svg)
# Unit Creator 🏴‍☠️

<img width="2557" height="1264" alt="Image" src="https://github.com/user-attachments/assets/8f4c463c-c477-48e1-9284-f781ba502431" />

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

Unit Creator FI is a work in progress. Today it is a Go web server that serves a static, view only website for Fieldcraft Interactive, including a Unit Portal page that previews a planned Unit Creator tool for Mil Sim and tactical gaming communities.
 
**Current limitations:** there is no database, no working login, and no API yet. The contact and login forms are layout only, and the news content is placeholder text.

### Motivation
Mil Sim enthusiasts and tactical gaming units often juggle Discord channels, Google Sheets, and forum posts to maintain their organizational identity. The goal of Unit Creator is to centralize this into a single platform with proper data persistence and a clean presentation layer. That goal is not built yet, see the roadmap below.

### What it does:
* Serves nine pages from a Go HTTP server using `net/http` and `html/template`: home, login (Unit Portal), trailer, contact, about, news, privacy policy, terms of service, and code of conduct
* Serves static assets (CSS and images) from the `static` directory
* Runs in Docker, with Docker Compose for local development
* Builds, tests, scans, and publishes automatically with GitHub Actions

### What is planned
 
* Unit pages with emblems, hierarchy, and member profiles
* Persistent storage for unit data and member rosters
* User authentication and unit administration tools
* A JSON API for unit data

## 🛠️ Built With

<p align="left">
<img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
<img src="https://github.com/user-attachments/assets/debd8e54-8d8a-4dc1-b900-0c778af06574" width="48" height="48"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=docker"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=githubactions">&nbsp;
</p>

### 🔧 Technical Stack
 
**Current**
 
* **Go HTTP server** using the standard library (`net/http`, `html/template`), with configured read, write, and idle timeouts
* **HTML and CSS** for the presentation layer, with no JavaScript yet
* **Docker and Docker Compose** for consistent builds and local development
* **GitHub Actions** for CI and CD (see below)
**Planned**
 
* **PostgreSQL** for persistent unit and roster data (Phase 2)
* **React and TypeScript** for dynamic unit page creation and editing (Phase 3)
* **GCP and Kubernetes** for hosting and scaling, to be evaluated after Phase 2
### 📍 Current Status: Phase 1 (View Only)
 
This is currently a showcase site with manually written content. Unit creation, accounts, and data storage are not implemented.
 
### 🛣️ Roadmap
 
* **Phase 1** (Current): Static pages served by a Go server, containerized, with automated CI and CD
* **Phase 2**: PostgreSQL database, user authentication, unit creation, and member management
* **Phase 3**: React frontend, search and filtering, and unit comparison features
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
 
* Optional, to build and test without Docker: Go 1.25 or newer
## 📦 Quick Start
 
### Run Pre-built Docker Image
 
Pull and run the latest production image directly from Docker Hub:
```sh
docker pull justinmnge/unit_creator_fi:latest
docker run -p 8080:8080 justinmnge/unit_creator_fi:latest
```
 
Then visit `http://localhost:8080` in your browser.
 
## 👀 What You'll See
 
After running the server, visit `http://localhost:8080` to browse:
 
* A home page, trailer page, and news page for Fieldcraft Interactive
* An about page describing the planned Unit Creator features (unit identity, admin controls, and ORBAT management)
* A Unit Portal login page and a contact page (layout only, forms do not submit)
* Privacy policy, terms of service, and code of conduct pages
The site uses a single desktop oriented stylesheet and is not yet optimized for mobile screens.
 
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
 
## 🤝 Contributing
 
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
 
## 🔁 CI and CD
 
**On every pull request to `main`** (`ci.yml`):
 
* Runs `go test`
* Scans the code with `gosec`
* Checks formatting with `go fmt` and lints with `staticcheck`
* Builds the Docker image without pushing it
**On every push to `main`** (`cd.yml`):
 
* Builds the binary and the Docker image
* Publishes the image to Docker Hub
## 🐳 Docker Hub
 
Pre-built Docker images are automatically published to Docker Hub via CD:
 
[![Docker Hub](https://img.shields.io/docker/pulls/justinmnge/unit_creator_fi)](https://hub.docker.com/r/justinmnge/unit_creator_fi)
 
* **Latest build**: `justinmnge/unit_creator_fi:latest`
* **Specific commits**: `justinmnge/unit_creator_fi:COMMIT_SHA`
## ℹ️ Contact
 
Justin Monge - hello@justin-monge.dev
 
Project Link: [https://github.com/justinmnge/unit_creator_fi](https://github.com/justinmnge/unit_creator_fi)
<p align="right">(<a href="#readme-top">back to top</a>)</p>
