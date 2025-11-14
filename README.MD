![CI checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/ci.yml/badge.svg) &nbsp; ![CD checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/cd.yml/badge.svg)
# Unit Portal Creator 🏴‍☠️

<img width="1912" height="908" alt="Image" src="https://github.com/user-attachments/assets/14049d29-0d06-41f7-aae1-419d91f6fb6c" />

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

The goal of this project is to build a comprehensive CRUD application that integrates with a game API to display real-time player and unit statistics, enabling Mil-Sim enthusiasts to create custom unit pages within a tactical gaming community.

Key Objectives:
* Build a production-grade HTTP server in Go
* Integrate and transform external game API data into meaningful user experiences
* Implement secure authentication and authorization for player and unit data
* Design RESTful APIs using JSON for seamless communication between game API, backend, and frontend
* Use Postgres with type-safe SQL to store unit emblems, organizational structures, and cached player data
* Develop endpoints for unit emblem creation, customization, and hierarchy visualization
* Document APIs to enable future expansion and integration

## 🛠️ Built With

</h3>
<p align="left">
 <a href="https://go.dev/"><img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
<img src="https://github.com/user-attachments/assets/debd8e54-8d8a-4dc1-b900-0c778af06574" width="48" height="48"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=docker"></a>&nbsp;
 </p>

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

1. Clone the repo
   ```sh
   git clone https://github.com/justinmnge/unit_creator_fi.git
   ```
2. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ```

## ⚙️ Usage

* Run the server with Docker
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

## ℹ️ Contact

Justin Monge - hello@justin-monge.dev

Project Link: [https://github.com/justinmnge/unit_creator_fi](https://github.com/justinmnge/unit_creator_fi)
<p align="right">(<a href="#readme-top">back to top</a>)</p>