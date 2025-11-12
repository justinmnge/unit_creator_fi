![CI checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/ci.yml/badge.svg)

![CD checks badge](https://github.com/justinmnge/unit_creator_fi/actions/workflows/cd.yml/badge.svg)
# Unit Portal Creator 🏴‍☠️

<img width="1912" height="908" alt="Image" src="https://github.com/user-attachments/assets/14049d29-0d06-41f7-aae1-419d91f6fb6c" />

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

The goal of this project is to build a comprehensive Unit Portal Creator that integrates with a game API to display real-time player and unit statistics.

Key Objectives:
* Understand what web servers are and how they power real-world web applications
* Understand how to integrate external game APIs and transform their data into meaningful user experiences
* Build a production-style HTTP server in Go that serves as the backbone for the unit portal system
* Implement secure user authentication to gate access to personalized player statistics and unit data
* Use RESTful API principles with JSON, headers, and status codes to communicate between the game API, backend, and frontend
* Leverage Go's performance capabilities to handle real-time game data updates efficiently
* Use type-safe SQL with a Postgres database to store and manage unit emblems, organizational structures, and cached player data
* Create a secure authorization system that controls access to unit rosters and member information
* Build API endpoints to support unit emblem creation and customization
* Design and implement a system to visualize unit organizational hierarchies and command structures
* Document the API endpoints to enable future expansion and integration possibilities

## 🛠️ Built With

</h3>
<p align="left">
 <a href="https://go.dev/"><img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
<img src="https://skillicons.dev/icons?i=css">&nbsp;
<img src="https://skillicons.dev/icons?i=docker"></a>&nbsp;
 </p>

## ✅ Prerequisites

* Install using your terminal to the latest version 
  ```sh
  curl -sS https://webi.sh/golang | sh
  ```

* Verify installation worked
  ```sh
  go version
  ```

## 📦 Installation

1. Clone the repo
   ```sh
   git clone https://github.com/justinmnge/unit_creator_fi.git
   ```
2. Download all dependencies
   ```sh
   go mod download
   ```
3. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ```

## ⚙️ Usage

* Run the server
  ```sh
  go run .
  ```

<!-- CONTACT -->
## ℹ️ Contact

Justin Monge - hello@justin-monge.dev

Project Link: [https://github.com/justinmnge/unit_creator_fi](https://github.com/justinmnge/unit_creator_fi)
<p align="right">(<a href="#readme-top">back to top</a>)</p>